// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const DisputeGameFactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"gameImpls\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_mapping(t_userDefinedValueType(GameType)1011,t_contract(IDisputeGame)1009)\"},{\"astId\":1006,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"initBonds\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_mapping(t_userDefinedValueType(GameType)1011,t_uint256)\"},{\"astId\":1007,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_disputeGames\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_mapping(t_userDefinedValueType(Hash)1012,t_userDefinedValueType(GameId)1010)\"},{\"astId\":1008,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_disputeGameList\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_array(t_userDefinedValueType(GameId)1010)dyn_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_array(t_userDefinedValueType(GameId)1010)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"GameId[]\",\"numberOfBytes\":\"32\",\"base\":\"t_userDefinedValueType(GameId)1010\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IDisputeGame)1009\":{\"encoding\":\"inplace\",\"label\":\"contract IDisputeGame\",\"numberOfBytes\":\"20\"},\"t_mapping(t_userDefinedValueType(GameType)1011,t_contract(IDisputeGame)1009)\":{\"encoding\":\"mapping\",\"label\":\"mapping(GameType =\u003e contract IDisputeGame)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(GameType)1011\",\"value\":\"t_contract(IDisputeGame)1009\"},\"t_mapping(t_userDefinedValueType(GameType)1011,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(GameType =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(GameType)1011\",\"value\":\"t_uint256\"},\"t_mapping(t_userDefinedValueType(Hash)1012,t_userDefinedValueType(GameId)1010)\":{\"encoding\":\"mapping\",\"label\":\"mapping(Hash =\u003e GameId)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(Hash)1012\",\"value\":\"t_userDefinedValueType(GameId)1010\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"},\"t_userDefinedValueType(GameId)1010\":{\"encoding\":\"inplace\",\"label\":\"GameId\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(GameType)1011\":{\"encoding\":\"inplace\",\"label\":\"GameType\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Hash)1012\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"}}}"

var DisputeGameFactoryStorageLayout = new(solc.StorageLayout)

var DisputeGameFactoryDeployedBin = "0x6080604052600436106100e85760003560e01c80636593dc6e1161008a57806396cd97201161005957806396cd972014610313578063bb8aa1fc14610333578063c4d66de814610394578063f2fde38b146103b457600080fd5b80636593dc6e14610293578063715018a6146102c057806382ecf2f6146102d55780638da5cb5b146102e857600080fd5b8063254bd683116100c6578063254bd6831461019c5780634d1975b4146101c957806354fd4d50146101e85780635f0150cb1461023e57600080fd5b806314f6b1a3146100ed5780631b685b9e1461010f5780631e3342401461017c575b600080fd5b3480156100f957600080fd5b5061010d6101083660046110c6565b6103d4565b005b34801561011b57600080fd5b5061015261012a3660046110fd565b60656020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561018857600080fd5b5061010d610197366004611118565b61045e565b3480156101a857600080fd5b506101bc6101b7366004611142565b6104aa565b60405161017391906111ef565b3480156101d557600080fd5b506068545b604051908152602001610173565b3480156101f457600080fd5b506102316040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161017391906112ac565b34801561024a57600080fd5b5061025e6102593660046112bf565b6106ee565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835267ffffffffffffffff909116602083015201610173565b34801561029f57600080fd5b506101da6102ae3660046110fd565b60666020526000908152604090205481565b3480156102cc57600080fd5b5061010d610741565b6101526102e33660046112bf565b610755565b3480156102f457600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610152565b34801561031f57600080fd5b506101da61032e3660046112bf565b6109ef565b34801561033f57600080fd5b5061035361034e366004611346565b610a28565b6040805163ffffffff909416845267ffffffffffffffff909216602084015273ffffffffffffffffffffffffffffffffffffffff1690820152606001610173565b3480156103a057600080fd5b5061010d6103af36600461135f565b610a91565b3480156103c057600080fd5b5061010d6103cf36600461135f565b610c2d565b6103dc610d00565b63ffffffff821660008181526065602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8616908117909155905190917fff513d80e2c7fa487608f70a618dfbc0cf415699dc69588c747e8c71566c88de91a35050565b610466610d00565b63ffffffff8216600081815260666020526040808220849055518392917f74d6665c4b26d5596a5aa13d3014e0c06af4d322075a797f87b03cd4c5bc91ca91a35050565b606854606090831015806104bc575081155b6106e7575060408051600583901b8101602001909152825b8381116106e5576000606882815481106104f0576104f061137c565b600091825260209091200154905060e081901c60a082901c67ffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff831663ffffffff891683036106b6576001865101865260008173ffffffffffffffffffffffffffffffffffffffff1663609d33346040518163ffffffff1660e01b8152600401600060405180830381865afa15801561058a573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526105d091908101906113da565b905060008273ffffffffffffffffffffffffffffffffffffffff1663bcef3b556040518163ffffffff1660e01b8152600401602060405180830381865afa15801561061f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064391906114a5565b90506040518060a001604052808881526020018781526020018567ffffffffffffffff168152602001828152602001838152508860018a5161068591906114be565b815181106106955761069561137c565b6020026020010181905250888851106106b3575050505050506106e5565b50505b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90920191506104d49050565b505b9392505050565b60008060006106ff878787876109ef565b60009081526067602052604090205473ffffffffffffffffffffffffffffffffffffffff81169860a09190911c67ffffffffffffffff16975095505050505050565b610749610d00565b6107536000610d81565b565b63ffffffff841660009081526065602052604081205473ffffffffffffffffffffffffffffffffffffffff16806107c5576040517f031c6de400000000000000000000000000000000000000000000000000000000815263ffffffff871660048201526024015b60405180910390fd5b63ffffffff86166000908152606660205260409020543414610813576040517f8620aa1900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006108206001436114be565b40905061088a338783888860405160200161083f9594939291906114fc565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905273ffffffffffffffffffffffffffffffffffffffff841690610df8565b92508273ffffffffffffffffffffffffffffffffffffffff16638129fc1c346040518263ffffffff1660e01b81526004016000604051808303818588803b1580156108d457600080fd5b505af11580156108e8573d6000803e3d6000fd5b505050505060006108fb888888886109ef565b60008181526067602052604090205490915015610947576040517f014f6fe5000000000000000000000000000000000000000000000000000000008152600481018290526024016107bc565b60004260a01b60e08a901b178517600083815260676020526040808220839055606880546001810182559083527fa2153420d844928b4421650203c77babc8b33d7f2e7b450e2966db0c220977530183905551919250899163ffffffff8c169173ffffffffffffffffffffffffffffffffffffffff8916917f5b565efe82411da98814f356d0e7bcb8f0219b8d970307c5afb4a6903a8b2e359190a450505050949350505050565b600084848484604051602001610a089493929190611549565b604051602081830303815290604052805190602001209050949350505050565b600080600080600080610a8160688881548110610a4757610a4761137c565b906000526020600020015460e081901c9160a082901c67ffffffffffffffff169173ffffffffffffffffffffffffffffffffffffffff1690565b9199909850909650945050505050565b600054610100900460ff1615808015610ab15750600054600160ff909116105b80610acb5750303b158015610acb575060005460ff166001145b610b57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016107bc565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610bb557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610bbd610e06565b610bc682610d81565b8015610c2957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b610c35610d00565b73ffffffffffffffffffffffffffffffffffffffff8116610cd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016107bc565b610ce181610d81565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff163314610753576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107bc565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006106e760008484610ea5565b600054610100900460ff16610e9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016107bc565b610753610feb565b600060608203516040830351602084035184518060208701018051600283016c5af43d3d93803e606057fd5bf3895289600d8a035278593da1005b363d3d373d3d3d3d610000806062363936013d738160481b1760218a03527f9e4ac34f21c619cefc926c8bd93b54bf5a39c7ab2127a895af1cc0691d7e3dff603a8a035272fd6100003d81600a3d39f336602c57343d527f6062820160781b1761ff9e82106059018a03528060f01b8352606c8101604c8a038cf097505086610f715763301164256000526004601cfd5b905285527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08501527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08401527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa09092019190915292915050565b600054610100900460ff16611082576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016107bc565b61075333610d81565b803563ffffffff8116811461109f57600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610ce157600080fd5b600080604083850312156110d957600080fd5b6110e28361108b565b915060208301356110f2816110a4565b809150509250929050565b60006020828403121561110f57600080fd5b6106e78261108b565b6000806040838503121561112b57600080fd5b6111348361108b565b946020939093013593505050565b60008060006060848603121561115757600080fd5b6111608461108b565b95602085013595506040909401359392505050565b60005b83811015611190578181015183820152602001611178565b8381111561119f576000848401525b50505050565b600081518084526111bd816020860160208601611175565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561129e578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc001855281518051845287810151888501528681015167ffffffffffffffff16878501526060808201519085015260809081015160a09185018290529061128a818601836111a5565b968901969450505090860190600101611216565b509098975050505050505050565b6020815260006106e760208301846111a5565b600080600080606085870312156112d557600080fd5b6112de8561108b565b935060208501359250604085013567ffffffffffffffff8082111561130257600080fd5b818701915087601f83011261131657600080fd5b81358181111561132557600080fd5b88602082850101111561133757600080fd5b95989497505060200194505050565b60006020828403121561135857600080fd5b5035919050565b60006020828403121561137157600080fd5b81356106e7816110a4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156113ec57600080fd5b815167ffffffffffffffff8082111561140457600080fd5b818401915084601f83011261141857600080fd5b81518181111561142a5761142a6113ab565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611470576114706113ab565b8160405282815287602084870101111561148957600080fd5b61149a836020830160208801611175565b979650505050505050565b6000602082840312156114b757600080fd5b5051919050565b6000828210156114f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008660601b1681528460148201528360348201528183605483013760009101605401908152949350505050565b63ffffffff8516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101939250505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(DisputeGameFactoryStorageLayoutJSON), DisputeGameFactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["DisputeGameFactory"] = DisputeGameFactoryStorageLayout
	deployedBytecodes["DisputeGameFactory"] = DisputeGameFactoryDeployedBin
	immutableReferences["DisputeGameFactory"] = false
}
