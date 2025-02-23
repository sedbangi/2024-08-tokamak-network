/* eslint-disable @typescript-eslint/no-unused-vars */
import {
  ethers,
  Contract,
  Overrides,
  Signer,
  BigNumber,
  CallOverrides,
} from 'ethers'
import {
  TransactionRequest,
  TransactionResponse,
  BlockTag,
} from '@ethersproject/abstract-provider'
import { hexStringEquals, predeploys } from '@tokamak-network/core-utils'
import l1StandardBridgeArtifact from '@tokamak-network/thanos-contracts/forge-artifacts/L1StandardBridge.sol/L1StandardBridge.json'
import l2StandardBridgeArtifact from '@tokamak-network/thanos-contracts/forge-artifacts/L2StandardBridge.sol/L2StandardBridge.json'
import optimismMintableERC20 from '@tokamak-network/thanos-contracts/forge-artifacts/OptimismMintableERC20.sol/OptimismMintableERC20.json'
import erc20 from '@tokamak-network/thanos-contracts/forge-artifacts/ERC20/IERC20.sol/IERC20.json'

import { CrossChainMessenger } from '../cross-chain-messenger'
import {
  IBridgeAdapter,
  NumberLike,
  AddressLike,
  TokenBridgeMessage,
  MessageDirection,
} from '../interfaces'
import { toAddress } from '../utils'

/**
 * Bridge adapter for any token bridge that uses the standard token bridge interface.
 */
export class StandardBridgeAdapter implements IBridgeAdapter {
  public messenger: CrossChainMessenger
  public l1Bridge: Contract
  public l2Bridge: Contract

  /**
   * Creates a StandardBridgeAdapter instance.
   *
   * @param opts Options for the adapter.
   * @param opts.messenger Provider used to make queries related to cross-chain interactions.
   * @param opts.l1Bridge L1 bridge contract.
   * @param opts.l2Bridge L2 bridge contract.
   */
  constructor(opts: {
    messenger: CrossChainMessenger
    l1Bridge: AddressLike
    l2Bridge: AddressLike
  }) {
    this.messenger = opts.messenger
    this.l1Bridge = new Contract(
      toAddress(opts.l1Bridge),
      l1StandardBridgeArtifact.abi,
      this.messenger.l1Provider
    )
    this.l2Bridge = new Contract(
      toAddress(opts.l2Bridge),
      l2StandardBridgeArtifact.abi,
      this.messenger.l2Provider
    )
  }

  public async getDepositsByAddress(
    address: AddressLike,
    opts?: {
      fromBlock?: BlockTag
      toBlock?: BlockTag
    }
  ): Promise<TokenBridgeMessage[]> {
    const events = await this.l1Bridge.queryFilter(
      this.l1Bridge.filters.ERC20DepositInitiated(
        undefined,
        undefined,
        address
      ),
      opts?.fromBlock,
      opts?.toBlock
    )

    return events
      .filter((event) =>
        this.supportsTokenPair(event.args.l1Token, event.args.l2Token)
      )
      .map((event) => {
        return {
          direction: MessageDirection.L1_TO_L2,
          from: event.args.from,
          to: event.args.to,
          l1Token: event.args.l1Token,
          l2Token: event.args.l2Token,
          amount: event.args.amount,
          data: event.args.extraData,
          logIndex: event.logIndex,
          blockNumber: event.blockNumber,
          transactionHash: event.transactionHash,
        }
      })
      .sort((a, b) => {
        // Sort descending by block number
        return b.blockNumber - a.blockNumber
      })
  }

  public async getWithdrawalsByAddress(
    address: AddressLike,
    opts?: {
      fromBlock?: BlockTag
      toBlock?: BlockTag
    }
  ): Promise<TokenBridgeMessage[]> {
    const events = await this.l2Bridge.queryFilter(
      this.l2Bridge.filters.WithdrawalInitiated(undefined, undefined, address),
      opts?.fromBlock,
      opts?.toBlock
    )

    return events
      .filter((event) => {
        return this.supportsTokenPair(event.args.l1Token, event.args.l2Token)
      })
      .map((event) => {
        return {
          direction: MessageDirection.L2_TO_L1,
          from: event.args.from,
          to: event.args.to,
          l1Token: event.args.l1Token,
          l2Token: event.args.l2Token,
          amount: event.args.amount,
          data: event.args.extraData,
          logIndex: event.logIndex,
          blockNumber: event.blockNumber,
          transactionHash: event.transactionHash,
        }
      })
      .sort((a, b) => {
        // Sort descending by block number
        return b.blockNumber - a.blockNumber
      })
  }

  public async supportsTokenPair(
    l1Token: AddressLike,
    l2Token: AddressLike
  ): Promise<boolean> {
    try {
      const contract = new Contract(
        toAddress(l2Token),
        optimismMintableERC20.abi,
        this.messenger.l2Provider
      )
      // Specifically filter out ETH. ETH deposits and withdrawals are handled by the ETH bridge
      // adapter. Bridges that are not the ETH bridge should not be able to handle or even
      // present ETH deposits or withdrawals.
      if (this.filterEthDepositsAndWithdrawls(l1Token, l2Token)) {
        return false
      }

      // Specifically filter out the native token. L2 native token deposits and withdrawals are handled by the L2 native token bridge
      // adapter. Bridges that are not L2 the native token bridge should not be able to handle or even
      // present L2 native token deposits or withdrawals.
      if (this.filterL2NativeTokenDepositsAndWithdrawls(l1Token, l2Token)) {
        return false
      }

      // Make sure the L1 token matches.
      const remoteL1Token = await contract.l1Token()

      if (!hexStringEquals(remoteL1Token, toAddress(l1Token))) {
        return false
      }

      // Make sure the L2 bridge matches.
      const remoteL2Bridge = await contract.l2Bridge()
      if (!hexStringEquals(remoteL2Bridge, this.l2Bridge.address)) {
        return false
      }

      return true
    } catch (err) {
      // If the L2 token is not an L2StandardERC20, it may throw an error. If there's a call
      // exception then we assume that the token is not supported. Other errors are thrown. Since
      // the JSON-RPC API is not well-specified, we need to handle multiple possible error codes.
      if (
        !err?.message?.toString().includes('CALL_EXCEPTION') &&
        !err?.stack?.toString().includes('execution reverted')
      ) {
        console.error('Unexpected error when checking bridge', err)
      }
      return false
    }
  }

  public async approval(
    l1Token: AddressLike,
    l2Token: AddressLike,
    signer: ethers.Signer
  ): Promise<BigNumber> {
    if (!(await this.supportsTokenPair(l1Token, l2Token))) {
      throw new Error(`token pair not supported by bridge`)
    }

    const token = new Contract(
      toAddress(l1Token),
      erc20.abi,
      this.messenger.l1Provider
    )

    return token.allowance(await signer.getAddress(), this.l1Bridge.address)
  }

  public async approve(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    signer: Signer,
    opts?: {
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.approve(l1Token, l2Token, amount, opts)
    )
  }

  public async deposit(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    signer: Signer,
    opts?: {
      recipient?: AddressLike
      l2GasLimit?: NumberLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.deposit(l1Token, l2Token, amount, opts)
    )
  }

  public async withdraw(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    signer: Signer,
    opts?: {
      recipient?: AddressLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.withdraw(l1Token, l2Token, amount, opts)
    )
  }

  public filterEthDepositsAndWithdrawls = (
    l1Token: AddressLike,
    l2Token: AddressLike
  ): boolean => {
    return (
      hexStringEquals(toAddress(l1Token), ethers.constants.AddressZero) &&
      hexStringEquals(toAddress(l2Token), predeploys.ETH)
    )
  }

  public filterL2NativeTokenDepositsAndWithdrawls = (
    l1Token: AddressLike,
    l2Token: AddressLike
  ): boolean => {
    if (!this.messenger.nativeTokenAddress) {
      return false
    }
    return (
      hexStringEquals(toAddress(l1Token), this.messenger.nativeTokenAddress) &&
      hexStringEquals(toAddress(l2Token), predeploys.LegacyERC20NativeToken)
    )
  }

  populateTransaction = {
    approve: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (!(await this.supportsTokenPair(l1Token, l2Token))) {
        throw new Error(`token pair not supported by bridge`)
      }

      const token = new Contract(
        toAddress(l1Token),
        erc20.abi,
        this.messenger.l1Provider
      )

      return token.populateTransaction.approve(
        this.l1Bridge.address,
        amount,
        opts?.overrides || {}
      )
    },

    deposit: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (!(await this.supportsTokenPair(l1Token, l2Token))) {
        throw new Error(`token pair not supported by bridge`)
      }

      if (opts?.recipient === undefined) {
        return this.l1Bridge.populateTransaction.bridgeERC20(
          toAddress(l1Token),
          toAddress(l2Token),
          amount,
          opts?.l2GasLimit || 200_000, // Default to 200k gas limit.
          '0x', // No data.
          opts?.overrides || {}
        )
      } else {
        return this.l1Bridge.populateTransaction.bridgeERC20To(
          toAddress(l1Token),
          toAddress(l2Token),
          toAddress(opts.recipient),
          amount,
          opts?.l2GasLimit || 200_000, // Default to 200k gas limit.
          '0x', // No data.
          opts?.overrides || {}
        )
      }
    },

    withdraw: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (!(await this.supportsTokenPair(l1Token, l2Token))) {
        throw new Error(`token pair not supported by bridge`)
      }

      if (opts?.recipient === undefined) {
        return this.l2Bridge.populateTransaction.withdraw(
          toAddress(l2Token),
          amount,
          0, // L1 gas not required.
          '0x', // No data.
          opts?.overrides || {}
        )
      } else {
        return this.l2Bridge.populateTransaction.withdrawTo(
          toAddress(l2Token),
          toAddress(opts.recipient),
          amount,
          0, // L1 gas not required.
          '0x', // No data.
          opts?.overrides || {}
        )
      }
    },
  }

  estimateGas = {
    approve: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l1Provider.estimateGas(
        await this.populateTransaction.approve(l1Token, l2Token, amount, opts)
      )
    },

    deposit: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        l2GasLimit?: NumberLike
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l1Provider.estimateGas(
        await this.populateTransaction.deposit(l1Token, l2Token, amount, opts)
      )
    },

    withdraw: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l2Provider.estimateGas(
        await this.populateTransaction.withdraw(l1Token, l2Token, amount, opts)
      )
    },
  }
}
