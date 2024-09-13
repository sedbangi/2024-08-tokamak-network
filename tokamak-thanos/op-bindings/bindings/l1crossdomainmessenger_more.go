// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"otherMessenger\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_contract(CrossDomainMessenger)1021\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)43_storage\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_contract(SuperchainConfig)1023\"},{\"astId\":1019,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"portal\",\"offset\":0,\"slot\":\"252\",\"type\":\"t_contract(OptimismPortal)1022\"},{\"astId\":1020,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"systemConfig\",\"offset\":0,\"slot\":\"253\",\"type\":\"t_contract(SystemConfig)1024\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)43_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[43]\",\"numberOfBytes\":\"1376\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(CrossDomainMessenger)1021\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(OptimismPortal)1022\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_contract(SuperchainConfig)1023\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_contract(SystemConfig)1024\":{\"encoding\":\"inplace\",\"label\":\"contract SystemConfig\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101b75760003560e01c80635c975abb116100ec578063b1b1b2091161008a578063d764ad0b11610064578063d764ad0b1461051c578063db505d801461052f578063e0e593c51461055c578063ecc704281461057c57600080fd5b8063b1b1b209146104ac578063b28ade25146104dc578063c0c53b8b146104fc57600080fd5b806383a74074116100c657806383a740741461043a5780638cbeeef21461034c5780639fce812c14610451578063a4e7f8bd1461047c57600080fd5b80635c975abb146103e35780636425666b146103f85780636e296e451461042557600080fd5b80633dbb202b116101595780634c1d6a69116101335780634c1d6a691461034c5780634d0047ee1461036257806354fd4d50146103775780635644cfdf146103cd57600080fd5b80633dbb202b146102ef5780633f827a5a146103045780634273ca161461032c57600080fd5b80630ff754ea116101955780630ff754ea146102345780632828d7e81461028057806333d7e2bd1461029557806335e80ab3146102c257600080fd5b806301ffc9a7146101bc578063028f85f7146101f15780630c5684981461021f575b600080fd5b3480156101c857600080fd5b506101dc6101d7366004612593565b6105e1565b60405190151581526020015b60405180910390f35b3480156101fd57600080fd5b50610206601081565b60405167ffffffffffffffff90911681526020016101e8565b34801561022b57600080fd5b50610206603f81565b34801561024057600080fd5b5060fc5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e8565b34801561028c57600080fd5b50610206604081565b3480156102a157600080fd5b5060fd5461025b9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102ce57600080fd5b5060fb5461025b9073ffffffffffffffffffffffffffffffffffffffff1681565b6103026102fd36600461265c565b61067a565b005b34801561031057600080fd5b50610319600181565b60405161ffff90911681526020016101e8565b34801561033857600080fd5b506101dc6103473660046126c3565b6108d7565b34801561035857600080fd5b50610206619c4081565b34801561036e57600080fd5b5061025b6109d6565b34801561038357600080fd5b506103c06040518060400160405280600581526020017f322e342e3000000000000000000000000000000000000000000000000000000081525081565b6040516101e891906127ac565b3480156103d957600080fd5b5061020661138881565b3480156103ef57600080fd5b506101dc610a6f565b34801561040457600080fd5b5060fc5461025b9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561043157600080fd5b5061025b610b03565b34801561044657600080fd5b5061020662030d4081565b34801561045d57600080fd5b5060cf5473ffffffffffffffffffffffffffffffffffffffff1661025b565b34801561048857600080fd5b506101dc6104973660046127bf565b60ce6020526000908152604090205460ff1681565b3480156104b857600080fd5b506101dc6104c73660046127bf565b60cb6020526000908152604090205460ff1681565b3480156104e857600080fd5b506102066104f73660046127d8565b610bea565b34801561050857600080fd5b5061030261051736600461282c565b610c5a565b61030261052a366004612877565b610ed1565b34801561053b57600080fd5b5060cf5461025b9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561056857600080fd5b506103026105773660046128fd565b6119ae565b34801561058857600080fd5b506105d360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101e8565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4273ca1600000000000000000000000000000000000000000000000000000000148061067457507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60cf546107ac9073ffffffffffffffffffffffffffffffffffffffff166106a2858585610bea565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061070e60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161072a97969594939291906129b7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526119df565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561083160cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610843959493929190612a16565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60006108e16109d6565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f6f6e6c7920616363657074206e617469766520746f6b656e20617070726f766560448201527f2063616c6c6261636b000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000803660006109b08787611ae0565b93509350935093506109c68a858a868686611bb2565b5060019998505050505050505050565b60fd54604080517f4d0047ee000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691634d0047ee9160048083019260209291908290030181865afa158015610a46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6a9190612a64565b905090565b60fb54604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa158015610adf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6a9190612a81565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610bcd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f742073657400000000000000000000006064820152608401610997565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000611388619c4080603f610c06604063ffffffff8816612ad2565b610c109190612b02565b610c1b601088612ad2565b610c289062030d40612b50565b610c329190612b50565b610c3c9190612b50565b610c469190612b50565b610c509190612b50565b90505b9392505050565b6000547501000000000000000000000000000000000000000000900460ff1615808015610ca5575060005460017401000000000000000000000000000000000000000090910460ff16105b80610cd75750303b158015610cd7575060005474010000000000000000000000000000000000000000900460ff166001145b610d63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610997565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610de957600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b60fb805473ffffffffffffffffffffffffffffffffffffffff8087167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fc805486841690831617905560fd805492851692909116919091179055610e68734200000000000000000000000000000000000007611e7e565b8015610ecb57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610ed9610a6f565b15610f40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4c312043726f7373446f6d61696e4d657373656e6765723a20706175736564006044820152606401610997565b3415610fce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f0000000000000000000000000000000000000000000000006064820152608401610997565b60f087901c60028110611089576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a401610997565b8061ffff1660000361117e5760006110da878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611fba915050565b600081815260cb602052604090205490915060ff161561117c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c617965640000000000000000006064820152608401610997565b505b60006111c4898989898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611fd992505050565b905060006111d06109d6565b90506111da611ffc565b1561123157600082815260ce602052604090205460ff16156111fe576111fe612b7c565b861561122c5760fc5461122c9073ffffffffffffffffffffffffffffffffffffffff8381169116308a6120d8565b6112cf565b600082815260ce602052604090205460ff166112cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610997565b6112d88861216d565b15801561131157508073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614155b6113c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605960248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f657373206f72206e6174697665546f6b656e4164647265737300000000000000608482015260a401610997565b600082815260cb602052604090205460ff1615611462576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610997565b61148386611474611388619c40612b50565b67ffffffffffffffff166121b0565b15806114a9575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156115c357600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016115bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610997565b5050506119a5565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8b161790558615801590611627575073ffffffffffffffffffffffffffffffffffffffff881615155b156116c7576040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89811660048301526024820189905282169063095ea7b3906044016020604051808303816000875af11580156116a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c59190612a81565b505b600061171989619c405a6116db9190612bab565b600089898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121ce92505050565b9050871580159061173f575073ffffffffffffffffffffffffffffffffffffffff891615155b156117df576040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a811660048301526000602483015283169063095ea7b3906044016020604051808303816000875af11580156117b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117dd9190612a81565b505b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055801561189357600083815260cb602052604090205460ff161561183057611830612b7c565b600083815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26119a0565b600083815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016119a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610997565b505050505b50505050505050565b6119bc338686848787611bb2565b5050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b3415611a47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f44656e79206465706f736974696e6720455448000000000000000000000000006044820152606401610997565b60fc546040517fb9e5595800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063b9e5595890611aa89087908690819089906000908990600401612bc2565b600060405180830381600087803b158015611ac257600080fd5b505af1158015611ad6573d6000803e3d6000fd5b5050505050505050565b60008036816018851015611b76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f496e76616c6964206f6e417070726f7665206461746120666f72204c3143726f60448201527f7373446f6d61696e4d657373656e6765720000000000000000000000000000006064820152608401610997565b505050823560601c93601484013560e01c93601801927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8019150565b8315611c86576000611bc26109d6565b9050611be673ffffffffffffffffffffffffffffffffffffffff82168830886120d8565b60fc546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018790529082169063095ea7b3906044016020604051808303816000875af1158015611c5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c839190612a81565b50505b60cf54611d369073ffffffffffffffffffffffffffffffffffffffff16611cae848487610bea565b867fd764ad0b00000000000000000000000000000000000000000000000000000000611d1a60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8b8b8b8b8b8b60405160240161072a97969594939291906129b7565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a878484611dbb60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b88604051611dcd959493929190612a16565b60405180910390a28573ffffffffffffffffffffffffffffffffffffffff167f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d54685604051611e1d91815260200190565b60405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff00000000000000000000000000000000000000000000000000000000000090911617905550505050565b6000547501000000000000000000000000000000000000000000900460ff16611f29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610997565b60cc5473ffffffffffffffffffffffffffffffffffffffff16611f735760cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790555b60cf80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000611fc8858585856121e8565b805190602001209050949350505050565b6000611fe9878787878787612281565b8051906020012090509695505050505050565b60fc5460009073ffffffffffffffffffffffffffffffffffffffff1633148015610a6a575060cf5460fc54604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691639bf62d82916004808201926020929091908290030181865afa158015612098573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120bc9190612a64565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610ecb908590612320565b600073ffffffffffffffffffffffffffffffffffffffff821630148061067457505060fc5473ffffffffffffffffffffffffffffffffffffffff90811691161490565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016122019493929190612c21565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b606086868686868660405160240161229e96959493929190612c6b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b6000612382826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166124319092919063ffffffff16565b80519091501561242c57808060200190518101906123a09190612a81565b61242c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610997565b505050565b6060610c5084846000858573ffffffffffffffffffffffffffffffffffffffff85163b6124ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610997565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516124e39190612cb6565b60006040518083038185875af1925050503d8060008114612520576040519150601f19603f3d011682016040523d82523d6000602084013e612525565b606091505b5091509150612535828286612540565b979650505050505050565b6060831561254f575081610c53565b82511561255f5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099791906127ac565b6000602082840312156125a557600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610c5357600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146125f757600080fd5b50565b60008083601f84011261260c57600080fd5b50813567ffffffffffffffff81111561262457600080fd5b60208301915083602082850101111561263c57600080fd5b9250929050565b803563ffffffff8116811461265757600080fd5b919050565b6000806000806060858703121561267257600080fd5b843561267d816125d5565b9350602085013567ffffffffffffffff81111561269957600080fd5b6126a5878288016125fa565b90945092506126b8905060408601612643565b905092959194509250565b6000806000806000608086880312156126db57600080fd5b85356126e6816125d5565b945060208601356126f6816125d5565b935060408601359250606086013567ffffffffffffffff81111561271957600080fd5b612725888289016125fa565b969995985093965092949392505050565b60005b83811015612751578181015183820152602001612739565b83811115610ecb5750506000910152565b6000815180845261277a816020860160208601612736565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610c536020830184612762565b6000602082840312156127d157600080fd5b5035919050565b6000806000604084860312156127ed57600080fd5b833567ffffffffffffffff81111561280457600080fd5b612810868287016125fa565b9094509250612823905060208501612643565b90509250925092565b60008060006060848603121561284157600080fd5b833561284c816125d5565b9250602084013561285c816125d5565b9150604084013561286c816125d5565b809150509250925092565b600080600080600080600060c0888a03121561289257600080fd5b8735965060208801356128a4816125d5565b955060408801356128b4816125d5565b9450606088013593506080880135925060a088013567ffffffffffffffff8111156128de57600080fd5b6128ea8a828b016125fa565b989b979a50959850939692959293505050565b60008060008060006080868803121561291557600080fd5b8535612920816125d5565b945060208601359350604086013567ffffffffffffffff81111561294357600080fd5b61294f888289016125fa565b9094509250612962905060608701612643565b90509295509295909350565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152612a0960c08301848661296e565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000612a4660808301868861296e565b905083604083015263ffffffff831660608301529695505050505050565b600060208284031215612a7657600080fd5b8151610c53816125d5565b600060208284031215612a9357600080fd5b81518015158114610c5357600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615612af957612af9612aa3565b02949350505050565b600067ffffffffffffffff80841680612b44577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115612b7357612b73612aa3565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015612bbd57612bbd612aa3565b500390565b73ffffffffffffffffffffffffffffffffffffffff8716815285602082015284604082015267ffffffffffffffff84166060820152821515608082015260c060a08201526000612c1560c0830184612762565b98975050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152612c5a6080830185612762565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152612c1560c0830184612762565b60008251612cc8818460208701612736565b919091019291505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
	immutableReferences["L1CrossDomainMessenger"] = false
}
