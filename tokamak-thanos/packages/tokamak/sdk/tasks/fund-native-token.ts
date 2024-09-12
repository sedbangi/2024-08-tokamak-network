import '@nomiclabs/hardhat-ethers'
import { BytesLike, ethers } from 'ethers'
import 'hardhat-deploy'
import { task, types } from 'hardhat/config'

import { NumberLike } from '../src'

const privateKey = process.env.PRIVATE_KEY as BytesLike

const l1Provider = new ethers.providers.StaticJsonRpcProvider(
  process.env.L1_URL
)
const l1Wallet = new ethers.Wallet(privateKey, l1Provider)

const erc20ABI = [
  {
    inputs: [
      { internalType: 'address', name: '_spender', type: 'address' },
      { internalType: 'uint256', name: '_value', type: 'uint256' },
    ],
    name: 'approve',
    outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: true,
    inputs: [{ name: '_owner', type: 'address' }],
    name: 'balanceOf',
    outputs: [{ name: 'balance', type: 'uint256' }],
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint256', name: 'amount', type: 'uint256' }],
    name: 'faucet',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const l2NativeToken = process.env.NATIVE_TOKEN || ''

const l2NativeTokenContract = new ethers.Contract(
  l2NativeToken,
  erc20ABI,
  l1Wallet
)

const fundNativeToken = async (amount: NumberLike) => {
  console.log('Faucet amount:', amount)

  console.log('Native token address:', l2NativeToken)
  const l2NativeTokenBalanceB4 = await l2NativeTokenContract.balanceOf(
    l1Wallet.address
  )
  console.log(
    'Native token balance in L1 before funding:',
    l2NativeTokenBalanceB4.toString()
  )

  const fundTx = await l2NativeTokenContract
    .connect(l1Wallet)
    .faucet(ethers.BigNumber.from('' + amount))
  console.log('Faucet transaction hash: ', fundTx.hash)
  await fundTx.wait()

  const l2NativeTokenBalance = await l2NativeTokenContract.balanceOf(
    l1Wallet.address
  )

  console.log(
    'Native token balance in L1 after funding:',
    l2NativeTokenBalance.toString()
  )
}

task('fund-native-token', 'Send L2NativeToken to L1.')
  .addParam('amount', 'Send amount', '1000000000000000000', types.string)
  .setAction(async (args) => {
    await fundNativeToken(args.amount)
  })
