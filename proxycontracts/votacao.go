// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxycontracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// VotacaoAssembleiaABI is the input ABI used to generate the binding from.
const VotacaoAssembleiaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalDeVotantes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numeroProposta\",\"type\":\"uint256\"},{\"name\":\"favoravelAProposta\",\"type\":\"bool\"}],\"name\":\"votar\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numeroProposta\",\"type\":\"uint256\"}],\"name\":\"propostaAprovada\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"indiceVotante\",\"type\":\"uint256\"}],\"name\":\"pesquisarVotantePorIndice\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qualTextoDaProposta\",\"type\":\"string\"},{\"name\":\"qualProponente\",\"type\":\"address\"},{\"name\":\"qualQuotaMinimaParaAprovacao\",\"type\":\"uint256\"}],\"name\":\"incluiProposta\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"detalhesAssembleia\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qualDataInicioVotacao\",\"type\":\"uint256\"}],\"name\":\"definirInicioVotacao\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qualDataFimVotacao\",\"type\":\"uint256\"}],\"name\":\"definirFimVotacao\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quandoEncerraVotacao\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enderecoVotante\",\"type\":\"address\"},{\"name\":\"quotaDeVotos\",\"type\":\"uint256\"}],\"name\":\"incluiVotante\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numeroProposta\",\"type\":\"uint256\"}],\"name\":\"pesquisarProposta\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"qualMotivoDaConvocatoria\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"indiceVotante\",\"type\":\"address\"}],\"name\":\"pesquisarVotante\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDePropostas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secretarioDesignado\",\"type\":\"address\"}],\"name\":\"designarSecretario\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"qualMotivoConvocatoria\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"quemVotou\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"propostaVotada\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"qualVoto\",\"type\":\"bool\"}],\"name\":\"Votou\",\"type\":\"event\"}]"

// VotacaoAssembleiaBin is the compiled bytecode used for deploying new contracts.
const VotacaoAssembleiaBin = `0x60806040523480156200001157600080fd5b50604051620015613803806200156183398101604052805160048054600160a060020a031916331790550180516200005190600790602084019062000059565b5050620000fe565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200009c57805160ff1916838001178555620000cc565b82800160010185558215620000cc579182015b82811115620000cc578251825591602001919060010190620000af565b50620000da929150620000de565b5090565b620000fb91905b80821115620000da5760008155600101620000e5565b90565b611453806200010e6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166308ae343b81146100df57806328962d4e14610106578063302dbc20146101375780633c9afde31461014f5780634080037f1461018a5780634a727aba146101f65780636de7d9c4146102d15780637fbdbe5f146102e95780638755f2ad14610301578063932bed1a14610316578063a7e421cb1461033a578063b0899b53146103f1578063ba67a8611461047b578063cf5e72a81461049c578063e030ed33146104b1575b600080fd5b3480156100eb57600080fd5b506100f46104d2565b60408051918252519081900360200190f35b34801561011257600080fd5b5061012360043560243515156104d9565b604080519115158252519081900360200190f35b34801561014357600080fd5b5061012360043561067e565b34801561015b57600080fd5b506101676004356107a2565b60408051600160a060020a03909316835260208301919091528051918290030190f35b34801561019657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101f494369492936024939284019190819084018382808284375094975050508335600160a060020a03169450505060209091013590506108bb565b005b34801561020257600080fd5b5061020b610a37565b6040518088600160a060020a0316600160a060020a0316815260200187600160a060020a0316600160a060020a0316815260200180602001868152602001858152602001848152602001838152602001828103825287818151815260200191508051906020019080838360005b83811015610290578181015183820152602001610278565b50505050905090810190601f1680156102bd5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156102dd57600080fd5b506101f4600435610b24565b3480156102f557600080fd5b506101f4600435610c2e565b34801561030d57600080fd5b506100f4610d38565b34801561032257600080fd5b506101f4600160a060020a0360043516602435610d3e565b34801561034657600080fd5b5061035260043561100e565b604051808060200185600160a060020a0316600160a060020a03168152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156103b357818101518382015260200161039b565b50505050905090810190601f1680156103e05780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156103fd57600080fd5b5061040661115d565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610440578181015183820152602001610428565b50505050905090810190601f16801561046d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561048757600080fd5b50610167600160a060020a03600435166111f3565b3480156104a857600080fd5b506100f4611270565b3480156104bd57600080fd5b506101f4600160a060020a0360043516611276565b6002545b90565b6000806000426006541015151561053a576040805160e560020a62461bcd02815260206004820152601160248201527f566f746163616f20656e63657272616461000000000000000000000000000000604482015290519081900360640190fd5b600554421015610594576040805160e560020a62461bcd02815260206004820152601f60248201527f566f7461c3a7c3a36f2061696e6461206ec3a36f20666f692061626572746100604482015290519081900360640190fd5b60008054869081106105a257fe5b60009182526020909120600590910201600481015490925060ff161561067157503360009081526001602081905260409091206002810154909161010090910460ff161515141561067157600281015460ff161515610671576001841515141561061757600181015460028301805490910190555b60408051338152602081018790528515158183015290517f62d81f5e43004f4fc2e3d45b0ff0c27b3c7bd87ee2a022db2cac7222ebf276aa9181900360600190a160028101805460ff191660019081179091559250610676565b600092505b505092915050565b600061068861132d565b600080548490811061069657fe5b600091825260209182902060408051600593909302909101805460026001821615610100026000190190911604601f8101859004909402830160c090810190925260a08301848152929390928492909184918401828280156107395780601f1061070e57610100808354040283529160200191610739565b820191906000526020600020905b81548152906001019060200180831161071c57829003601f168201915b50505091835250506001820154600160a060020a03166020820152600282015460408201526003820154606082015260049091015460ff1615156080918201528101519091501561079757806060015181604001511015915061079c565b600091505b50919050565b6000806107ad611368565b60025484111561082d576040805160e560020a62461bcd02815260206004820152603260248201527f496e6469636520696e666f726d61646f20c3a9206d61696f7220717565206f2060448201527f6e756d65726f20646520766f74616e7465730000000000000000000000000000606482015290519081900360840190fd5b600280548590811061083b57fe5b60009182526020918290206040805160808101825260039093029091018054600160a060020a03168352600180820154948401949094526002015460ff80821615159284019290925261010090041615156060820181905290925014156108ad578051602082015190935091506108b5565b600092508291505b50915091565b6108c361132d565b600354600160a060020a0316331461094b576040805160e560020a62461bcd02815260206004820152602f60248201527f53c3b3206f20736563726574c3a172696f20706f6465207265616c697a61722060448201527f65737361206f70657261c3a7c3a36f0000000000000000000000000000000000606482015290519081900360840190fd5b506040805160a081018252848152600160a060020a038416602080830191909152600092820183905260608201849052600160808301819052835490810180855593805282518051939493859360059093027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301926109ce92849291019061138f565b50602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905560408201516002820155606082015160038201556080909101516004909101805460ff19169115159190911790555050505050565b6000806060600080600080600080610a4d611270565b9150610a576104d2565b60045460035460055460065460078054604080516020601f60026000196001871615610100020190951694909404938401819004810282018101909252828152979850600160a060020a039687169795909616959194899289928791830182828015610b045780601f10610ad957610100808354040283529160200191610b04565b820191906000526020600020905b815481529060010190602001808311610ae757829003601f168201915b505050505094509850985098509850985098509850505090919293949596565b600454600160a060020a03163314610bac576040805160e560020a62461bcd02815260206004820152603260248201527f536f6d656e7465206f20707265736964656e746520706f6465207265616c697a60448201527f61722065737361206f70657261c3a7c3a36f0000000000000000000000000000606482015290519081900360840190fd5b428111610c29576040805160e560020a62461bcd02815260206004820152602b60248201527f41206461746120696e666f726d616461206465766520736572206d61696f722060448201527f717565206120617475616c000000000000000000000000000000000000000000606482015290519081900360840190fd5b600555565b600454600160a060020a03163314610cb6576040805160e560020a62461bcd02815260206004820152603260248201527f536f6d656e7465206f20707265736964656e746520706f6465207265616c697a60448201527f61722065737361206f70657261c3a7c3a36f0000000000000000000000000000606482015290519081900360840190fd5b428111610d33576040805160e560020a62461bcd02815260206004820152602b60248201527f41206461746120696e666f726d616461206465766520736572206d61696f722060448201527f717565206120617475616c000000000000000000000000000000000000000000606482015290519081900360840190fd5b600655565b60065490565b610d46611368565b600354600160a060020a03163314610dce576040805160e560020a62461bcd02815260206004820152602f60248201527f53c3b3206f20736563726574c3a172696f20706f6465207265616c697a61722060448201527f65737361206f70657261c3a7c3a36f0000000000000000000000000000000000606482015290519081900360840190fd5b6063821115610e4d576040805160e560020a62461bcd02815260206004820152602160248201527f51756f7461206e616f20706f646520736572207375706572696f72206120393960448201527f2500000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0383161515610ed3576040805160e560020a62461bcd02815260206004820152602560248201527f4f20766f74616e746520646576652074657220756d20656e64657265636f207660448201527f616c69646f000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060408051608081018252600160a060020a0393841680825260208083019485526000838501818152600160608601818152948352928390529481208451815490891673ffffffffffffffffffffffffffffffffffffffff199182161782558751828501558651600292830180548751151561010090810261ff001994151560ff19938416178516179092558454968701855593909452955160039094027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101805495909a16949091169390931790975594517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf82015592517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad0909301805491511515909402921515941693909317909216919091179055565b6060600080600061101d61132d565b600080548790811061102b57fe5b600091825260209182902060408051600593909302909101805460026001821615610100026000190190911604601f8101859004909402830160c090810190925260a08301848152929390928492909184918401828280156110ce5780601f106110a3576101008083540402835291602001916110ce565b820191906000526020600020905b8154815290600101906020018083116110b157829003601f168201915b50505091835250506001820154600160a060020a03166020820152600282015460408201526003820154606082015260049091015460ff1615156080918201528101519091501561113957805160208201516040830151606084015192975090955093509150611155565b6040805160208101909152600080825290955093508392508291505b509193509193565b60078054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156111e95780601f106111be576101008083540402835291602001916111e9565b820191906000526020600020905b8154815290600101906020018083116111cc57829003601f168201915b5050505050905090565b6000806111fe611368565b50600160a060020a0380841660009081526001602081815260409283902083516080810185528154909516855280830154918501919091526002015460ff808216151593850193909352610100900490911615156060830181905214156108ad578051602082015190935091506108b5565b60005490565b600454600160a060020a031633146112fe576040805160e560020a62461bcd02815260206004820152603260248201527f536f6d656e7465206f20707265736964656e746520706f6465207265616c697a60448201527f61722065737361206f70657261c3a7c3a36f0000000000000000000000000000606482015290519081900360840190fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60a060405190810160405280606081526020016000600160a060020a0316815260200160008152602001600081526020016000151581525090565b60408051608081018252600080825260208201819052918101829052606081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113d057805160ff19168380011785556113fd565b828001600101855582156113fd579182015b828111156113fd5782518255916020019190600101906113e2565b5061140992915061140d565b5090565b6104d691905b8082111561140957600081556001016114135600a165627a7a723058206000c0802037f156a57abe8127f0c94e65d234d49c50291ebd5786ff7f09665e0029`

// DeployVotacaoAssembleia deploys a new Ethereum contract, binding an instance of VotacaoAssembleia to it.
func DeployVotacaoAssembleia(auth *bind.TransactOpts, backend bind.ContractBackend, qualMotivoConvocatoria string) (common.Address, *types.Transaction, *VotacaoAssembleia, error) {
	parsed, err := abi.JSON(strings.NewReader(VotacaoAssembleiaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotacaoAssembleiaBin), backend, qualMotivoConvocatoria)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VotacaoAssembleia{VotacaoAssembleiaCaller: VotacaoAssembleiaCaller{contract: contract}, VotacaoAssembleiaTransactor: VotacaoAssembleiaTransactor{contract: contract}, VotacaoAssembleiaFilterer: VotacaoAssembleiaFilterer{contract: contract}}, nil
}

// VotacaoAssembleia is an auto generated Go binding around an Ethereum contract.
type VotacaoAssembleia struct {
	VotacaoAssembleiaCaller     // Read-only binding to the contract
	VotacaoAssembleiaTransactor // Write-only binding to the contract
	VotacaoAssembleiaFilterer   // Log filterer for contract events
}

// VotacaoAssembleiaCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotacaoAssembleiaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotacaoAssembleiaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotacaoAssembleiaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotacaoAssembleiaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotacaoAssembleiaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotacaoAssembleiaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotacaoAssembleiaSession struct {
	Contract     *VotacaoAssembleia // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VotacaoAssembleiaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotacaoAssembleiaCallerSession struct {
	Contract *VotacaoAssembleiaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// VotacaoAssembleiaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotacaoAssembleiaTransactorSession struct {
	Contract     *VotacaoAssembleiaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// VotacaoAssembleiaRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotacaoAssembleiaRaw struct {
	Contract *VotacaoAssembleia // Generic contract binding to access the raw methods on
}

// VotacaoAssembleiaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotacaoAssembleiaCallerRaw struct {
	Contract *VotacaoAssembleiaCaller // Generic read-only contract binding to access the raw methods on
}

// VotacaoAssembleiaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotacaoAssembleiaTransactorRaw struct {
	Contract *VotacaoAssembleiaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotacaoAssembleia creates a new instance of VotacaoAssembleia, bound to a specific deployed contract.
func NewVotacaoAssembleia(address common.Address, backend bind.ContractBackend) (*VotacaoAssembleia, error) {
	contract, err := bindVotacaoAssembleia(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotacaoAssembleia{VotacaoAssembleiaCaller: VotacaoAssembleiaCaller{contract: contract}, VotacaoAssembleiaTransactor: VotacaoAssembleiaTransactor{contract: contract}, VotacaoAssembleiaFilterer: VotacaoAssembleiaFilterer{contract: contract}}, nil
}

// NewVotacaoAssembleiaCaller creates a new read-only instance of VotacaoAssembleia, bound to a specific deployed contract.
func NewVotacaoAssembleiaCaller(address common.Address, caller bind.ContractCaller) (*VotacaoAssembleiaCaller, error) {
	contract, err := bindVotacaoAssembleia(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotacaoAssembleiaCaller{contract: contract}, nil
}

// NewVotacaoAssembleiaTransactor creates a new write-only instance of VotacaoAssembleia, bound to a specific deployed contract.
func NewVotacaoAssembleiaTransactor(address common.Address, transactor bind.ContractTransactor) (*VotacaoAssembleiaTransactor, error) {
	contract, err := bindVotacaoAssembleia(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotacaoAssembleiaTransactor{contract: contract}, nil
}

// NewVotacaoAssembleiaFilterer creates a new log filterer instance of VotacaoAssembleia, bound to a specific deployed contract.
func NewVotacaoAssembleiaFilterer(address common.Address, filterer bind.ContractFilterer) (*VotacaoAssembleiaFilterer, error) {
	contract, err := bindVotacaoAssembleia(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotacaoAssembleiaFilterer{contract: contract}, nil
}

// bindVotacaoAssembleia binds a generic wrapper to an already deployed contract.
func bindVotacaoAssembleia(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotacaoAssembleiaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotacaoAssembleia *VotacaoAssembleiaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotacaoAssembleia.Contract.VotacaoAssembleiaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotacaoAssembleia *VotacaoAssembleiaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.VotacaoAssembleiaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotacaoAssembleia *VotacaoAssembleiaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.VotacaoAssembleiaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotacaoAssembleia *VotacaoAssembleiaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotacaoAssembleia.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.contract.Transact(opts, method, params...)
}

// DetalhesAssembleia is a free data retrieval call binding the contract method 0x4a727aba.
//
// Solidity: function detalhesAssembleia() constant returns(address, address, string, uint256, uint256, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) DetalhesAssembleia(opts *bind.CallOpts) (common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _VotacaoAssembleia.contract.Call(opts, out, "detalhesAssembleia")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// DetalhesAssembleia is a free data retrieval call binding the contract method 0x4a727aba.
//
// Solidity: function detalhesAssembleia() constant returns(address, address, string, uint256, uint256, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) DetalhesAssembleia() (common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _VotacaoAssembleia.Contract.DetalhesAssembleia(&_VotacaoAssembleia.CallOpts)
}

// DetalhesAssembleia is a free data retrieval call binding the contract method 0x4a727aba.
//
// Solidity: function detalhesAssembleia() constant returns(address, address, string, uint256, uint256, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) DetalhesAssembleia() (common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _VotacaoAssembleia.Contract.DetalhesAssembleia(&_VotacaoAssembleia.CallOpts)
}

// PesquisarProposta is a free data retrieval call binding the contract method 0xa7e421cb.
//
// Solidity: function pesquisarProposta(numeroProposta uint256) constant returns(string, address, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) PesquisarProposta(opts *bind.CallOpts, numeroProposta *big.Int) (string, common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _VotacaoAssembleia.contract.Call(opts, out, "pesquisarProposta", numeroProposta)
	return *ret0, *ret1, *ret2, *ret3, err
}

// PesquisarProposta is a free data retrieval call binding the contract method 0xa7e421cb.
//
// Solidity: function pesquisarProposta(numeroProposta uint256) constant returns(string, address, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) PesquisarProposta(numeroProposta *big.Int) (string, common.Address, *big.Int, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarProposta(&_VotacaoAssembleia.CallOpts, numeroProposta)
}

// PesquisarProposta is a free data retrieval call binding the contract method 0xa7e421cb.
//
// Solidity: function pesquisarProposta(numeroProposta uint256) constant returns(string, address, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) PesquisarProposta(numeroProposta *big.Int) (string, common.Address, *big.Int, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarProposta(&_VotacaoAssembleia.CallOpts, numeroProposta)
}

// PesquisarVotante is a free data retrieval call binding the contract method 0xba67a861.
//
// Solidity: function pesquisarVotante(indiceVotante address) constant returns(address, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) PesquisarVotante(opts *bind.CallOpts, indiceVotante common.Address) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _VotacaoAssembleia.contract.Call(opts, out, "pesquisarVotante", indiceVotante)
	return *ret0, *ret1, err
}

// PesquisarVotante is a free data retrieval call binding the contract method 0xba67a861.
//
// Solidity: function pesquisarVotante(indiceVotante address) constant returns(address, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) PesquisarVotante(indiceVotante common.Address) (common.Address, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarVotante(&_VotacaoAssembleia.CallOpts, indiceVotante)
}

// PesquisarVotante is a free data retrieval call binding the contract method 0xba67a861.
//
// Solidity: function pesquisarVotante(indiceVotante address) constant returns(address, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) PesquisarVotante(indiceVotante common.Address) (common.Address, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarVotante(&_VotacaoAssembleia.CallOpts, indiceVotante)
}

// PesquisarVotantePorIndice is a free data retrieval call binding the contract method 0x3c9afde3.
//
// Solidity: function pesquisarVotantePorIndice(indiceVotante uint256) constant returns(address, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) PesquisarVotantePorIndice(opts *bind.CallOpts, indiceVotante *big.Int) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _VotacaoAssembleia.contract.Call(opts, out, "pesquisarVotantePorIndice", indiceVotante)
	return *ret0, *ret1, err
}

// PesquisarVotantePorIndice is a free data retrieval call binding the contract method 0x3c9afde3.
//
// Solidity: function pesquisarVotantePorIndice(indiceVotante uint256) constant returns(address, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) PesquisarVotantePorIndice(indiceVotante *big.Int) (common.Address, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarVotantePorIndice(&_VotacaoAssembleia.CallOpts, indiceVotante)
}

// PesquisarVotantePorIndice is a free data retrieval call binding the contract method 0x3c9afde3.
//
// Solidity: function pesquisarVotantePorIndice(indiceVotante uint256) constant returns(address, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) PesquisarVotantePorIndice(indiceVotante *big.Int) (common.Address, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarVotantePorIndice(&_VotacaoAssembleia.CallOpts, indiceVotante)
}

// PropostaAprovada is a free data retrieval call binding the contract method 0x302dbc20.
//
// Solidity: function propostaAprovada(numeroProposta uint256) constant returns(bool)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) PropostaAprovada(opts *bind.CallOpts, numeroProposta *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotacaoAssembleia.contract.Call(opts, out, "propostaAprovada", numeroProposta)
	return *ret0, err
}

// PropostaAprovada is a free data retrieval call binding the contract method 0x302dbc20.
//
// Solidity: function propostaAprovada(numeroProposta uint256) constant returns(bool)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) PropostaAprovada(numeroProposta *big.Int) (bool, error) {
	return _VotacaoAssembleia.Contract.PropostaAprovada(&_VotacaoAssembleia.CallOpts, numeroProposta)
}

// PropostaAprovada is a free data retrieval call binding the contract method 0x302dbc20.
//
// Solidity: function propostaAprovada(numeroProposta uint256) constant returns(bool)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) PropostaAprovada(numeroProposta *big.Int) (bool, error) {
	return _VotacaoAssembleia.Contract.PropostaAprovada(&_VotacaoAssembleia.CallOpts, numeroProposta)
}

// QualMotivoDaConvocatoria is a free data retrieval call binding the contract method 0xb0899b53.
//
// Solidity: function qualMotivoDaConvocatoria() constant returns(string)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) QualMotivoDaConvocatoria(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VotacaoAssembleia.contract.Call(opts, out, "qualMotivoDaConvocatoria")
	return *ret0, err
}

// QualMotivoDaConvocatoria is a free data retrieval call binding the contract method 0xb0899b53.
//
// Solidity: function qualMotivoDaConvocatoria() constant returns(string)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) QualMotivoDaConvocatoria() (string, error) {
	return _VotacaoAssembleia.Contract.QualMotivoDaConvocatoria(&_VotacaoAssembleia.CallOpts)
}

// QualMotivoDaConvocatoria is a free data retrieval call binding the contract method 0xb0899b53.
//
// Solidity: function qualMotivoDaConvocatoria() constant returns(string)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) QualMotivoDaConvocatoria() (string, error) {
	return _VotacaoAssembleia.Contract.QualMotivoDaConvocatoria(&_VotacaoAssembleia.CallOpts)
}

// QuandoEncerraVotacao is a free data retrieval call binding the contract method 0x8755f2ad.
//
// Solidity: function quandoEncerraVotacao() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) QuandoEncerraVotacao(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotacaoAssembleia.contract.Call(opts, out, "quandoEncerraVotacao")
	return *ret0, err
}

// QuandoEncerraVotacao is a free data retrieval call binding the contract method 0x8755f2ad.
//
// Solidity: function quandoEncerraVotacao() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) QuandoEncerraVotacao() (*big.Int, error) {
	return _VotacaoAssembleia.Contract.QuandoEncerraVotacao(&_VotacaoAssembleia.CallOpts)
}

// QuandoEncerraVotacao is a free data retrieval call binding the contract method 0x8755f2ad.
//
// Solidity: function quandoEncerraVotacao() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) QuandoEncerraVotacao() (*big.Int, error) {
	return _VotacaoAssembleia.Contract.QuandoEncerraVotacao(&_VotacaoAssembleia.CallOpts)
}

// TotalDePropostas is a free data retrieval call binding the contract method 0xcf5e72a8.
//
// Solidity: function totalDePropostas() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) TotalDePropostas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotacaoAssembleia.contract.Call(opts, out, "totalDePropostas")
	return *ret0, err
}

// TotalDePropostas is a free data retrieval call binding the contract method 0xcf5e72a8.
//
// Solidity: function totalDePropostas() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) TotalDePropostas() (*big.Int, error) {
	return _VotacaoAssembleia.Contract.TotalDePropostas(&_VotacaoAssembleia.CallOpts)
}

// TotalDePropostas is a free data retrieval call binding the contract method 0xcf5e72a8.
//
// Solidity: function totalDePropostas() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) TotalDePropostas() (*big.Int, error) {
	return _VotacaoAssembleia.Contract.TotalDePropostas(&_VotacaoAssembleia.CallOpts)
}

// TotalDeVotantes is a free data retrieval call binding the contract method 0x08ae343b.
//
// Solidity: function totalDeVotantes() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) TotalDeVotantes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotacaoAssembleia.contract.Call(opts, out, "totalDeVotantes")
	return *ret0, err
}

// TotalDeVotantes is a free data retrieval call binding the contract method 0x08ae343b.
//
// Solidity: function totalDeVotantes() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) TotalDeVotantes() (*big.Int, error) {
	return _VotacaoAssembleia.Contract.TotalDeVotantes(&_VotacaoAssembleia.CallOpts)
}

// TotalDeVotantes is a free data retrieval call binding the contract method 0x08ae343b.
//
// Solidity: function totalDeVotantes() constant returns(uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) TotalDeVotantes() (*big.Int, error) {
	return _VotacaoAssembleia.Contract.TotalDeVotantes(&_VotacaoAssembleia.CallOpts)
}

// DefinirFimVotacao is a paid mutator transaction binding the contract method 0x7fbdbe5f.
//
// Solidity: function definirFimVotacao(qualDataFimVotacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) DefinirFimVotacao(opts *bind.TransactOpts, qualDataFimVotacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "definirFimVotacao", qualDataFimVotacao)
}

// DefinirFimVotacao is a paid mutator transaction binding the contract method 0x7fbdbe5f.
//
// Solidity: function definirFimVotacao(qualDataFimVotacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaSession) DefinirFimVotacao(qualDataFimVotacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.DefinirFimVotacao(&_VotacaoAssembleia.TransactOpts, qualDataFimVotacao)
}

// DefinirFimVotacao is a paid mutator transaction binding the contract method 0x7fbdbe5f.
//
// Solidity: function definirFimVotacao(qualDataFimVotacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) DefinirFimVotacao(qualDataFimVotacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.DefinirFimVotacao(&_VotacaoAssembleia.TransactOpts, qualDataFimVotacao)
}

// DefinirInicioVotacao is a paid mutator transaction binding the contract method 0x6de7d9c4.
//
// Solidity: function definirInicioVotacao(qualDataInicioVotacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) DefinirInicioVotacao(opts *bind.TransactOpts, qualDataInicioVotacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "definirInicioVotacao", qualDataInicioVotacao)
}

// DefinirInicioVotacao is a paid mutator transaction binding the contract method 0x6de7d9c4.
//
// Solidity: function definirInicioVotacao(qualDataInicioVotacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaSession) DefinirInicioVotacao(qualDataInicioVotacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.DefinirInicioVotacao(&_VotacaoAssembleia.TransactOpts, qualDataInicioVotacao)
}

// DefinirInicioVotacao is a paid mutator transaction binding the contract method 0x6de7d9c4.
//
// Solidity: function definirInicioVotacao(qualDataInicioVotacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) DefinirInicioVotacao(qualDataInicioVotacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.DefinirInicioVotacao(&_VotacaoAssembleia.TransactOpts, qualDataInicioVotacao)
}

// DesignarSecretario is a paid mutator transaction binding the contract method 0xe030ed33.
//
// Solidity: function designarSecretario(secretarioDesignado address) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) DesignarSecretario(opts *bind.TransactOpts, secretarioDesignado common.Address) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "designarSecretario", secretarioDesignado)
}

// DesignarSecretario is a paid mutator transaction binding the contract method 0xe030ed33.
//
// Solidity: function designarSecretario(secretarioDesignado address) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaSession) DesignarSecretario(secretarioDesignado common.Address) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.DesignarSecretario(&_VotacaoAssembleia.TransactOpts, secretarioDesignado)
}

// DesignarSecretario is a paid mutator transaction binding the contract method 0xe030ed33.
//
// Solidity: function designarSecretario(secretarioDesignado address) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) DesignarSecretario(secretarioDesignado common.Address) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.DesignarSecretario(&_VotacaoAssembleia.TransactOpts, secretarioDesignado)
}

// IncluiProposta is a paid mutator transaction binding the contract method 0x4080037f.
//
// Solidity: function incluiProposta(qualTextoDaProposta string, qualProponente address, qualQuotaMinimaParaAprovacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) IncluiProposta(opts *bind.TransactOpts, qualTextoDaProposta string, qualProponente common.Address, qualQuotaMinimaParaAprovacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "incluiProposta", qualTextoDaProposta, qualProponente, qualQuotaMinimaParaAprovacao)
}

// IncluiProposta is a paid mutator transaction binding the contract method 0x4080037f.
//
// Solidity: function incluiProposta(qualTextoDaProposta string, qualProponente address, qualQuotaMinimaParaAprovacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaSession) IncluiProposta(qualTextoDaProposta string, qualProponente common.Address, qualQuotaMinimaParaAprovacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.IncluiProposta(&_VotacaoAssembleia.TransactOpts, qualTextoDaProposta, qualProponente, qualQuotaMinimaParaAprovacao)
}

// IncluiProposta is a paid mutator transaction binding the contract method 0x4080037f.
//
// Solidity: function incluiProposta(qualTextoDaProposta string, qualProponente address, qualQuotaMinimaParaAprovacao uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) IncluiProposta(qualTextoDaProposta string, qualProponente common.Address, qualQuotaMinimaParaAprovacao *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.IncluiProposta(&_VotacaoAssembleia.TransactOpts, qualTextoDaProposta, qualProponente, qualQuotaMinimaParaAprovacao)
}

// IncluiVotante is a paid mutator transaction binding the contract method 0x932bed1a.
//
// Solidity: function incluiVotante(enderecoVotante address, quotaDeVotos uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) IncluiVotante(opts *bind.TransactOpts, enderecoVotante common.Address, quotaDeVotos *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "incluiVotante", enderecoVotante, quotaDeVotos)
}

// IncluiVotante is a paid mutator transaction binding the contract method 0x932bed1a.
//
// Solidity: function incluiVotante(enderecoVotante address, quotaDeVotos uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaSession) IncluiVotante(enderecoVotante common.Address, quotaDeVotos *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.IncluiVotante(&_VotacaoAssembleia.TransactOpts, enderecoVotante, quotaDeVotos)
}

// IncluiVotante is a paid mutator transaction binding the contract method 0x932bed1a.
//
// Solidity: function incluiVotante(enderecoVotante address, quotaDeVotos uint256) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) IncluiVotante(enderecoVotante common.Address, quotaDeVotos *big.Int) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.IncluiVotante(&_VotacaoAssembleia.TransactOpts, enderecoVotante, quotaDeVotos)
}

// Votar is a paid mutator transaction binding the contract method 0x28962d4e.
//
// Solidity: function votar(numeroProposta uint256, favoravelAProposta bool) returns(bool)
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) Votar(opts *bind.TransactOpts, numeroProposta *big.Int, favoravelAProposta bool) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "votar", numeroProposta, favoravelAProposta)
}

// Votar is a paid mutator transaction binding the contract method 0x28962d4e.
//
// Solidity: function votar(numeroProposta uint256, favoravelAProposta bool) returns(bool)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) Votar(numeroProposta *big.Int, favoravelAProposta bool) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.Votar(&_VotacaoAssembleia.TransactOpts, numeroProposta, favoravelAProposta)
}

// Votar is a paid mutator transaction binding the contract method 0x28962d4e.
//
// Solidity: function votar(numeroProposta uint256, favoravelAProposta bool) returns(bool)
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) Votar(numeroProposta *big.Int, favoravelAProposta bool) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.Votar(&_VotacaoAssembleia.TransactOpts, numeroProposta, favoravelAProposta)
}

// VotacaoAssembleiaVotouIterator is returned from FilterVotou and is used to iterate over the raw logs and unpacked data for Votou events raised by the VotacaoAssembleia contract.
type VotacaoAssembleiaVotouIterator struct {
	Event *VotacaoAssembleiaVotou // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotacaoAssembleiaVotouIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotacaoAssembleiaVotou)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotacaoAssembleiaVotou)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotacaoAssembleiaVotouIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotacaoAssembleiaVotouIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotacaoAssembleiaVotou represents a Votou event raised by the VotacaoAssembleia contract.
type VotacaoAssembleiaVotou struct {
	QuemVotou      common.Address
	PropostaVotada *big.Int
	QualVoto       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotou is a free log retrieval operation binding the contract event 0x62d81f5e43004f4fc2e3d45b0ff0c27b3c7bd87ee2a022db2cac7222ebf276aa.
//
// Solidity: e Votou(quemVotou address, propostaVotada uint256, qualVoto bool)
func (_VotacaoAssembleia *VotacaoAssembleiaFilterer) FilterVotou(opts *bind.FilterOpts) (*VotacaoAssembleiaVotouIterator, error) {

	logs, sub, err := _VotacaoAssembleia.contract.FilterLogs(opts, "Votou")
	if err != nil {
		return nil, err
	}
	return &VotacaoAssembleiaVotouIterator{contract: _VotacaoAssembleia.contract, event: "Votou", logs: logs, sub: sub}, nil
}

// WatchVotou is a free log subscription operation binding the contract event 0x62d81f5e43004f4fc2e3d45b0ff0c27b3c7bd87ee2a022db2cac7222ebf276aa.
//
// Solidity: e Votou(quemVotou address, propostaVotada uint256, qualVoto bool)
func (_VotacaoAssembleia *VotacaoAssembleiaFilterer) WatchVotou(opts *bind.WatchOpts, sink chan<- *VotacaoAssembleiaVotou) (event.Subscription, error) {

	logs, sub, err := _VotacaoAssembleia.contract.WatchLogs(opts, "Votou")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotacaoAssembleiaVotou)
				if err := _VotacaoAssembleia.contract.UnpackLog(event, "Votou", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
