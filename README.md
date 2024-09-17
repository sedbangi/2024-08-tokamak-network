
# Tokamak Network contest details

- Join [Sherlock Discord](https://discord.gg/MABEWyASkp)
- Submit findings using the issue page in your private contest repo (label issues as med or high)
- [Read for more details](https://docs.sherlock.xyz/audits/watsons)

# Q&A

### Q: On what chains are the smart contracts going to be deployed?
 L1: Ethereum
 L2: Thanos (built by OP Stack)
___

### Q: If you are integrating tokens, are you allowing only whitelisted tokens to work with the codebase or any complying with the standard? Are they assumed to have certain properties, e.g. be non-reentrant? Are there any types of [weird tokens](https://github.com/d-xo/weird-erc20) you want to integrate?
The assets that can be bridged are categorized into ETH, native tokens, and ERC20 tokens that are not native tokens. 
Not all tokens can be chosen as the L2 native token. For a token to be used as an L2 native token, the corresponding L1 token must satisfy the ERC20 standard and the following additional requirements:

 - Has 18 decimals. The L2 native token has 18 decimals, so the corresponding L1 token must have exactly 18 decimals to ensure no loss of precision when depositing or withdrawing.
 - Must not have fees or hooks on transfer.
 - Must not have out-of-band methods for modifying balance or allowance. For example, no tokes that have rebasing logic or double entry points can be an L2 native token.
 - Any other requirements set by a standard bridge that is not mentioned here.
___

### Q: Are there any limitations on values set by admins (or other roles) in the codebase, including restrictions on array lengths?
No
___

### Q: Are there any limitations on values set by admins (or other roles) in protocols you integrate with, including restrictions on array lengths?
No
___

### Q: For permissioned functions, please list all checks and requirements that will be made before calling the function.
N/A
___

### Q: Is the codebase expected to comply with any EIPs? Can there be/are there any deviations from the specification?
N/A
___

### Q: Are there any off-chain mechanisms or off-chain procedures for the protocol (keeper bots, arbitrage bots, etc.)?
It does not require any off-chain procedures.
___

### Q: Are there any hardcoded values that you intend to change before (some) deployments?
No
___

### Q: If the codebase is to be deployed on an L2, what should be the behavior of the protocol in case of sequencer issues (if applicable)? Should Sherlock assume that the Sequencer won't misbehave, including going offline?
Yes, we can operate and control the sequencer. So, on the contract side, we don't want to consider sequencer problems. But if the problem is serious (enough to break the protocol), Watson can report that issue.
___

### Q: Should potential issues, like broken assumptions about function behavior, be reported if they could pose risks in future integrations, even if they might not be an issue in the context of the scope? If yes, can you elaborate on properties/invariants that should hold?
No.
___

### Q: Please discuss any design choices you made.
Thanos L2 Native Token Bridge has some differences compared to the Custom Gas Token of Optimism. (https://specs.optimism.io/experimental/custom-gas-token.html)

1. When Custom Gas Token uses ERC20 as a Gas Token, L1CrossDomainMessenger and L1StandardBridge cannot be used for fund transfer. We have improved the convenience of Dapp developers by making this possible, even if only to a limited extent. 

2. We use only one function for deposit assets in OptimismPortal2. Custom Gas Token uses depositTransaction and depositERC20Transaction.
OptimismPortal2.sol::depositTransaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes calldata _data)

___

### Q: Please list any known issues and explicitly state the acceptable risks for each known issue.
Users must be careful when completing withdrawals from L1. As the finalization of the withdrawal phase, the messages that can be executed in L1 are OptimismPortal2.sol::finalizeWithdrawalTransaction(Types.WithdrawalTransaction memory _tx) and L1CrossDomainMessenger::relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes calldata _message). When sending a withdrawal transaction in L2, a user may lose funds if _tx.target is Contract and _tx.data.length is not 0. Especially, in case of _tx.data.length is not 0 and _tx.data includes function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes calldata _message), user may lose funds even if _sender is EOA. We already know this issue and allow it. However, we welcome any improvements or better methods.
___

### Q: We will report issues where the core protocol functionality is inaccessible for at least 7 days. Would you like to override this value?
No
___

### Q: Please provide links to previous audits (if any).
N/A
___

### Q: Please list any relevant protocol resources.
1. Specification: https://tokamak.notion.site/L2-native-token-specification-4ba8b138747f4981a2ce4c5a0f595964
2. Test guide using Hardhat task: https://tokamak.notion.site/Bridge-Test-using-Hardhat-task-a87ad0dafd84476c8ed0e30abb4b0a27
___

### Q: Additional audit information.
1. Key Point for this audit:
The most critical issue to focus on during the audit is the potential for assets to be frozen or lost when they are deposited or withdrawn via the bridge. We ask that you prioritize identifying vulnerabilities related to this issue.

2. About Thanos L2 Native Token Bridge:
Thanos is a Layer 2 project implemented based on OP Stack v1.7.7 (https://github.com/ethereum-optimism/optimism/releases/tag/v1.7.7). The Native Token Bridge offers similar functionality to Optimism’s Custom Gas Token (https://specs.optimism.io/experimental/custom-gas-token.html), but there are clear differences. For more details, please refer to the specifications.
___



# Audit scope


[tokamak-thanos @ c3a2bd6f768f0719a63fa1e0231eaae1d42f9e24](https://github.com/tokamak-network/tokamak-thanos/tree/c3a2bd6f768f0719a63fa1e0231eaae1d42f9e24)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/L1/L1CrossDomainMessenger.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/L1/L1CrossDomainMessenger.sol)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/L1/L1StandardBridge.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/L1/L1StandardBridge.sol)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/L1/OptimismPortal2.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/L1/OptimismPortal2.sol)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/L2/L2CrossDomainMessenger.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/L2/L2CrossDomainMessenger.sol)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/L2/L2StandardBridge.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/L2/L2StandardBridge.sol)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/universal/CrossDomainMessenger.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/universal/CrossDomainMessenger.sol)
- [tokamak-thanos/packages/tokamak/contracts-bedrock/src/universal/StandardBridge.sol](tokamak-thanos/packages/tokamak/contracts-bedrock/src/universal/StandardBridge.sol)

