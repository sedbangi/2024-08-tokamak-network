import { predeploys } from '@tokamak-network/core-utils'
import addressManagerArtifactGoerli from '@tokamak-network/thanos-contracts/deployments/goerli/AddressManager.json'
import l1CrossDomainMessengerArtifactGoerli from '@tokamak-network/thanos-contracts/deployments/goerli/L1CrossDomainMessengerProxy.json'
import l1StandardBridgeArtifactGoerli from '@tokamak-network/thanos-contracts/deployments/goerli/L1StandardBridgeProxy.json'
import l2OutputOracleArtifactsGoerli from '@tokamak-network/thanos-contracts/deployments/goerli/L2OutputOracleProxy.json'
import portalArtifactsGoerli from '@tokamak-network/thanos-contracts/deployments/goerli/OptimismPortalProxy.json'
import addressManagerArtifactMainnet from '@tokamak-network/thanos-contracts/deployments/mainnet/AddressManager.json'
import l1CrossDomainMessengerArtifactMainnet from '@tokamak-network/thanos-contracts/deployments/mainnet/L1CrossDomainMessengerProxy.json'
import l1StandardBridgeArtifactMainnet from '@tokamak-network/thanos-contracts/deployments/mainnet/L1StandardBridgeProxy.json'
import l2OutputOracleArtifactsMainnet from '@tokamak-network/thanos-contracts/deployments/mainnet/L2OutputOracleProxy.json'
import portalArtifactsMainnet from '@tokamak-network/thanos-contracts/deployments/mainnet/OptimismPortalProxy.json'
import addressManagerArtifactSepolia from '@tokamak-network/thanos-contracts/deployments/sepolia/AddressManager.json'
import l1CrossDomainMessengerArtifactSepolia from '@tokamak-network/thanos-contracts/deployments/sepolia/L1CrossDomainMessengerProxy.json'
import l1StandardBridgeArtifactSepolia from '@tokamak-network/thanos-contracts/deployments/sepolia/L1StandardBridgeProxy.json'
import l2OutputOracleArtifactsSepolia from '@tokamak-network/thanos-contracts/deployments/sepolia/L2OutputOracleProxy.json'
import portalArtifactsSepolia from '@tokamak-network/thanos-contracts/deployments/sepolia/OptimismPortalProxy.json'
import addressManagerArtifactThanosSepoliaTest from '@tokamak-network/thanos-contracts/deployments/thanos-sepolia-test/AddressManager.json'
import l1CrossDomainMessengerArtifactThanosSepoliaTest from '@tokamak-network/thanos-contracts/deployments/thanos-sepolia-test/L1CrossDomainMessengerProxy.json'
import l1StandardBridgeArtifactThanosSepoliaTest from '@tokamak-network/thanos-contracts/deployments/thanos-sepolia-test/L1StandardBridgeProxy.json'
import l2OutputOracleArtifactsThanosSepoliaTest from '@tokamak-network/thanos-contracts/deployments/thanos-sepolia-test/L2OutputOracleProxy.json'
import portalArtifactsThanosSepoliaTest from '@tokamak-network/thanos-contracts/deployments/thanos-sepolia-test/OptimismPortalProxy.json'
import { ethers } from 'ethers'

import {
  DAIBridgeAdapter,
  ECOBridgeAdapter,
  StandardBridgeAdapter,
} from '../adapters'
import {
  BridgeAdapterData,
  L1ChainID,
  L2ChainID,
  OEContractsLike,
  OEL1ContractsLike,
  OEL2ContractsLike,
} from '../interfaces'

const portalAddresses = {
  mainnet: portalArtifactsMainnet.address,
  goerli: portalArtifactsGoerli.address,
  sepolia: portalArtifactsSepolia.address,
  'thanos-sepolia-test': portalArtifactsThanosSepoliaTest.address,
}

const l2OutputOracleAddresses = {
  mainnet: l2OutputOracleArtifactsMainnet.address,
  goerli: l2OutputOracleArtifactsGoerli.address,
  sepolia: l2OutputOracleArtifactsSepolia.address,
  'thanos-sepolia-test': l2OutputOracleArtifactsThanosSepoliaTest.address,
}

const addressManagerAddresses = {
  mainnet: addressManagerArtifactMainnet.address,
  goerli: addressManagerArtifactGoerli.address,
  sepolia: addressManagerArtifactSepolia.address,
  'thanos-sepolia-test': addressManagerArtifactThanosSepoliaTest.address,
}

const l1StandardBridgeAddresses = {
  mainnet: l1StandardBridgeArtifactMainnet.address,
  goerli: l1StandardBridgeArtifactGoerli.address,
  sepolia: l1StandardBridgeArtifactSepolia.address,
  'thanos-sepolia-test': l1StandardBridgeArtifactThanosSepoliaTest.address,
}

const l1CrossDomainMessengerAddresses = {
  mainnet: l1CrossDomainMessengerArtifactMainnet.address,
  goerli: l1CrossDomainMessengerArtifactGoerli.address,
  sepolia: l1CrossDomainMessengerArtifactSepolia.address,
  'thanos-sepolia-test':
    l1CrossDomainMessengerArtifactThanosSepoliaTest.address,
}

const disputeGameFactoryAddresses = {
  mainnet: ethers.constants.AddressZero,
  goerli: ethers.constants.AddressZero,
  sepolia: '0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1',
}

// legacy
const stateCommitmentChainAddresses = {
  mainnet: '0xBe5dAb4A2e9cd0F27300dB4aB94BeE3A233AEB19',
  goerli: '0x9c945aC97Baf48cB784AbBB61399beB71aF7A378',
  sepolia: ethers.constants.AddressZero,
}

// legacy
const canonicalTransactionChainAddresses = {
  mainnet: '0x5E4e65926BA27467555EB562121fac00D24E9dD2',
  goerli: '0x607F755149cFEB3a14E1Dc3A4E2450Cde7dfb04D',
  sepolia: ethers.constants.AddressZero,
}

export const DEPOSIT_CONFIRMATION_BLOCKS: {
  [ChainID in L2ChainID]: number
} = {
  [L2ChainID.OPTIMISM]: 50 as const,
  [L2ChainID.OPTIMISM_GOERLI]: 12 as const,
  [L2ChainID.OPTIMISM_SEPOLIA]: 12 as const,
  [L2ChainID.OPTIMISM_HARDHAT_LOCAL]: 2 as const,
  [L2ChainID.OPTIMISM_HARDHAT_DEVNET]: 2 as const,
  [L2ChainID.OPTIMISM_LOCAL_DEVNET]: 2 as const,
  [L2ChainID.OPTIMISM_BEDROCK_ALPHA_TESTNET]: 12 as const,
  [L2ChainID.BASE_GOERLI]: 25 as const,
  [L2ChainID.BASE_SEPOLIA]: 25 as const,
  [L2ChainID.BASE_MAINNET]: 10 as const,
  [L2ChainID.ZORA_GOERLI]: 12 as const,
  [L2ChainID.ZORA_MAINNET]: 50 as const,
  [L2ChainID.MODE_SEPOLIA]: 25 as const,
  [L2ChainID.MODE_MAINNET]: 50 as const,
  [L2ChainID.THANOS_SEPOLIA]: 111551118080 as const,
}

export const CHAIN_BLOCK_TIMES: {
  [ChainID in L1ChainID]: number
} = {
  [L1ChainID.MAINNET]: 13 as const,
  [L1ChainID.GOERLI]: 15 as const,
  [L1ChainID.SEPOLIA]: 15 as const,
  [L1ChainID.HARDHAT_LOCAL]: 1 as const,
  [L1ChainID.BEDROCK_LOCAL_DEVNET]: 15 as const,
}

/**
 * Full list of default L2 contract addresses.
 */
export const DEFAULT_L2_CONTRACT_ADDRESSES: OEL2ContractsLike = {
  L2CrossDomainMessenger: predeploys.L2CrossDomainMessenger,
  L2ToL1MessagePasser: predeploys.L2ToL1MessagePasser,
  L2StandardBridge: predeploys.L2StandardBridge,
  OVM_L1BlockNumber: predeploys.L1BlockNumber,
  OVM_L2ToL1MessagePasser: predeploys.L2ToL1MessagePasser,
  OVM_DeployerWhitelist: predeploys.DeployerWhitelist,
  OVM_ETH: predeploys.LegacyERC20NativeToken,
  OVM_GasPriceOracle: predeploys.GasPriceOracle,
  OVM_SequencerFeeVault: predeploys.SequencerFeeVault,
  WETH: predeploys.WNativeToken,
  BedrockMessagePasser: predeploys.L2ToL1MessagePasser,
  L2UsdcBridge: predeploys.L2UsdcBridge,
}

/**
 * Loads the L1 contracts for a given network by the network name.
 *
 * @param network The name of the network to load the contracts for.
 * @returns The L1 contracts for the given network.
 */
const getL1ContractsByNetworkName = (network: string): OEL1ContractsLike => {
  return {
    AddressManager: addressManagerAddresses[network],
    L1CrossDomainMessenger: l1CrossDomainMessengerAddresses[network],
    L1StandardBridge: l1StandardBridgeAddresses[network],
    StateCommitmentChain: stateCommitmentChainAddresses[network],
    CanonicalTransactionChain: canonicalTransactionChainAddresses[network],
    BondManager: ethers.constants.AddressZero,
    OptimismPortal: portalAddresses[network],
    L2OutputOracle: l2OutputOracleAddresses[network],
    OptimismPortal2: portalAddresses[network],
    DisputeGameFactory: disputeGameFactoryAddresses[network],
  }
}

/**
 * List of contracts that are ignorable when checking for contracts on a given network.
 */
export const IGNORABLE_CONTRACTS = ['OptimismPortal2', 'DisputeGameFactory']

/**
 * Mapping of L1 chain IDs to the appropriate contract addresses for the OE deployments to the
 * given network. Simplifies the process of getting the correct contract addresses for a given
 * contract name.
 */
export const CONTRACT_ADDRESSES: {
  [ChainID in L2ChainID]: OEContractsLike
} = {
  [L2ChainID.OPTIMISM]: {
    l1: getL1ContractsByNetworkName('mainnet'),
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.OPTIMISM_GOERLI]: {
    l1: getL1ContractsByNetworkName('goerli'),
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.OPTIMISM_SEPOLIA]: {
    l1: getL1ContractsByNetworkName('sepolia'),
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.OPTIMISM_HARDHAT_LOCAL]: {
    l1: {
      AddressManager: '0x5FbDB2315678afecb367f032d93F642f64180aa3' as const,
      L1CrossDomainMessenger:
        '0x8A791620dd6260079BF849Dc5567aDC3F2FdC318' as const,
      L1StandardBridge: '0x610178dA211FEF7D417bC0e6FeD39F05609AD788' as const,
      StateCommitmentChain:
        '0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9' as const,
      CanonicalTransactionChain:
        '0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9' as const,
      BondManager: '0x5FC8d32690cc91D4c39d9d3abcBD16989F875707' as const,
      // FIXME
      OptimismPortal: '0x0000000000000000000000000000000000000000' as const,
      L2OutputOracle: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.OPTIMISM_HARDHAT_DEVNET]: {
    l1: {
      AddressManager: '0x5FbDB2315678afecb367f032d93F642f64180aa3' as const,
      L1CrossDomainMessenger:
        '0x8A791620dd6260079BF849Dc5567aDC3F2FdC318' as const,
      L1StandardBridge: '0x610178dA211FEF7D417bC0e6FeD39F05609AD788' as const,
      StateCommitmentChain:
        '0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9' as const,
      CanonicalTransactionChain:
        '0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9' as const,
      BondManager: '0x5FC8d32690cc91D4c39d9d3abcBD16989F875707' as const,
      OptimismPortal: '0x0000000000000000000000000000000000000000' as const,
      L2OutputOracle: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.OPTIMISM_BEDROCK_ALPHA_TESTNET]: {
    l1: {
      AddressManager: '0xb4e08DcE1F323608229265c9d4125E22a4B9dbAF' as const,
      L1CrossDomainMessenger:
        '0x838a6DC4E37CA45D4Ef05bb776bf05eEf50798De' as const,
      L1StandardBridge: '0xFf94B6C486350aD92561Ba09bad3a59df764Da92' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0xA581Ca3353DB73115C4625FFC7aDF5dB379434A8' as const,
      L2OutputOracle: '0x3A234299a14De50027eA65dCdf1c0DaC729e04A6' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.BASE_GOERLI]: {
    l1: {
      AddressManager: '0x4Cf6b56b14c6CFcB72A75611080514F94624c54e' as const,
      L1CrossDomainMessenger:
        '0x8e5693140eA606bcEB98761d9beB1BC87383706D' as const,
      L1StandardBridge: '0xfA6D8Ee5BE770F84FC001D098C4bD604Fe01284a' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0xe93c8cD0D409341205A592f8c4Ac1A5fe5585cfA' as const,
      L2OutputOracle: '0x2A35891ff30313CcFa6CE88dcf3858bb075A2298' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.BASE_SEPOLIA]: {
    l1: {
      AddressManager: '0x709c2B8ef4A9feFc629A8a2C1AF424Dc5BD6ad1B' as const,
      L1CrossDomainMessenger:
        '0xC34855F4De64F1840e5686e64278da901e261f20' as const,
      L1StandardBridge: '0xfd0Bf71F60660E2f608ed56e1659C450eB113120' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0x49f53e41452C74589E85cA1677426Ba426459e85' as const,
      L2OutputOracle: '0x84457ca9D0163FbC4bbfe4Dfbb20ba46e48DF254' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.BASE_MAINNET]: {
    l1: {
      AddressManager: '0x8EfB6B5c4767B09Dc9AA6Af4eAA89F749522BaE2' as const,
      L1CrossDomainMessenger:
        '0x866E82a600A1414e583f7F13623F1aC5d58b0Afa' as const,
      L1StandardBridge: '0x3154Cf16ccdb4C6d922629664174b904d80F2C35' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0x49048044D57e1C92A77f79988d21Fa8fAF74E97e' as const,
      L2OutputOracle: '0x56315b90c40730925ec5485cf004d835058518A0' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  // Zora Goerli
  [L2ChainID.ZORA_GOERLI]: {
    l1: {
      AddressManager: '0x54f4676203dEDA6C08E0D40557A119c602bFA246' as const,
      L1CrossDomainMessenger:
        '0xD87342e16352D33170557A7dA1e5fB966a60FafC' as const,
      L1StandardBridge: '0x7CC09AC2452D6555d5e0C213Ab9E2d44eFbFc956' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0xDb9F51790365e7dc196e7D072728df39Be958ACe' as const,
      L2OutputOracle: '0xdD292C9eEd00f6A32Ff5245d0BCd7f2a15f24e00' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.ZORA_MAINNET]: {
    l1: {
      AddressManager: '0xEF8115F2733fb2033a7c756402Fc1deaa56550Ef' as const,
      L1CrossDomainMessenger:
        '0xdC40a14d9abd6F410226f1E6de71aE03441ca506' as const,
      L1StandardBridge: '0x3e2Ea9B92B7E48A52296fD261dc26fd995284631' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0x1a0ad011913A150f69f6A19DF447A0CfD9551054' as const,
      L2OutputOracle: '0x9E6204F750cD866b299594e2aC9eA824E2e5f95c' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MODE_SEPOLIA]: {
    l1: {
      AddressManager: '0x83D45725d6562d8CD717673D6bb4c67C07dC1905' as const,
      L1CrossDomainMessenger:
        '0xc19a60d9E8C27B9A43527c3283B4dd8eDC8bE15C' as const,
      L1StandardBridge: '0xbC5C679879B2965296756CD959C3C739769995E2' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0x320e1580effF37E008F1C92700d1eBa47c1B23fD' as const,
      L2OutputOracle: '0x2634BD65ba27AB63811c74A63118ACb312701Bfa' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MODE_MAINNET]: {
    l1: {
      AddressManager: '0x50eF494573f28Cad6B64C31b7a00Cdaa48306e15' as const,
      L1CrossDomainMessenger:
        '0x95bDCA6c8EdEB69C98Bd5bd17660BaCef1298A6f' as const,
      L1StandardBridge: '0x735aDBbE72226BD52e818E7181953f42E3b0FF21' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0x8B34b14c7c7123459Cf3076b8Cb929BE097d0C07' as const,
      L2OutputOracle: '0x4317ba146D4933D889518a3e5E11Fe7a53199b04' as const,
      OptimismPortal2: '0x0000000000000000000000000000000000000000' as const,
      DisputeGameFactory: '0x0000000000000000000000000000000000000000' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  // internal testnet
  [L2ChainID.THANOS_SEPOLIA]: {
    l1: {
      AddressManager: '0xd9Ada4AE23bFA230351D1D14a561525D77Eb15Eb' as const,
      L1CrossDomainMessenger:
        '0x9D28a920206281B4a56AEf8bC1c515Cc4C656d3f' as const,
      L1StandardBridge: '0x385076516318551d566CAaE5EC59c23fe09cbF65' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0x7b6db1316e22167b56211cDDC33431098BaBC3c2' as const,
      L2OutputOracle: '0xaB8A5Ac696675D65D09E24C2876Aa8a7e1Af5640' as const,
      L1UsdcBridge: '0xE390EE020Afb7F8e4A2Dc44a71088db2acd72CF3' as const,
    },
    // TODO: change predeploys in L2
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.OPTIMISM_LOCAL_DEVNET]: {
    l1: {
      AddressManager: '0xe4EB561155AFCe723bB1fF8606Fbfe9b28d5d38D' as const,
      L1CrossDomainMessenger:
        '0x4d8eC2972eb0bC4210c64E651638D4a00ad3B400' as const,
      L1StandardBridge: '0x072B5bdBFC5e66B55317Ef4B4d1AE7d61592ebB2' as const,
      StateCommitmentChain:
        '0x0000000000000000000000000000000000000000' as const,
      CanonicalTransactionChain:
        '0x0000000000000000000000000000000000000000' as const,
      BondManager: '0x0000000000000000000000000000000000000000' as const,
      OptimismPortal: '0xfe36E31dFE8Cb3A3Aa0CB9f35B191DdB5451b090' as const,
      OptimismPortal2: '0xfe36E31dFE8Cb3A3Aa0CB9f35B191DdB5451b090' as const,
      L2OutputOracle: '0x754A91555a8dd5037315ABFd3702ED49d92887b7' as const,
      DisputeGameFactory: '0x11c81c1A7979cdd309096D1ea53F887EA9f8D14d' as const,
      L1UsdcBridge: '0x201B36B26b816D061fC552B679f8279Db0Fbbc6A' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
}

/**
 * Mapping of L1 chain IDs to the list of custom bridge addresses for each chain.
 */
export const BRIDGE_ADAPTER_DATA: {
  [ChainID in L2ChainID]?: BridgeAdapterData
} = {
  [L2ChainID.OPTIMISM]: {
    wstETH: {
      Adapter: DAIBridgeAdapter,
      l1Bridge: '0x76943C0D61395d8F2edF9060e1533529cAe05dE6' as const,
      l2Bridge: '0x8E01013243a96601a86eb3153F0d9Fa4fbFb6957' as const,
    },
    BitBTC: {
      Adapter: StandardBridgeAdapter,
      l1Bridge: '0xaBA2c5F108F7E820C049D5Af70B16ac266c8f128' as const,
      l2Bridge: '0x158F513096923fF2d3aab2BcF4478536de6725e2' as const,
    },
    DAI: {
      Adapter: DAIBridgeAdapter,
      l1Bridge: '0x10E6593CDda8c58a1d0f14C5164B376352a55f2F' as const,
      l2Bridge: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65' as const,
    },
    ECO: {
      Adapter: ECOBridgeAdapter,
      l1Bridge: '0xAa029BbdC947F5205fBa0F3C11b592420B58f824' as const,
      l2Bridge: '0xAa029BbdC947F5205fBa0F3C11b592420B58f824' as const,
    },
  },
  [L2ChainID.OPTIMISM_GOERLI]: {
    DAI: {
      Adapter: DAIBridgeAdapter,
      l1Bridge: '0x05a388Db09C2D44ec0b00Ee188cD42365c42Df23' as const,
      l2Bridge: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65' as const,
    },
    ECO: {
      Adapter: ECOBridgeAdapter,
      l1Bridge: '0x9A4464D6bFE006715382D39D183AAf66c952a3e0' as const,
      l2Bridge: '0x6aA809bAeA2e4C057b3994127cB165119c6fc3B2' as const,
    },
  },
}
