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
const VotacaoAssembleiaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalDeVotantes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numeroProposta\",\"type\":\"uint256\"},{\"name\":\"favoravelAProposta\",\"type\":\"bool\"}],\"name\":\"votar\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numeroProposta\",\"type\":\"uint256\"}],\"name\":\"propostaAprovada\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"indiceVotante\",\"type\":\"uint256\"}],\"name\":\"pesquisarVotantePorIndice\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qualTextoDaProposta\",\"type\":\"string\"},{\"name\":\"qualProponente\",\"type\":\"address\"},{\"name\":\"qualQuotaMinimaParaAprovacao\",\"type\":\"uint256\"}],\"name\":\"incluiProposta\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"detalhesAssembleia\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qualDataInicioVotacao\",\"type\":\"uint256\"}],\"name\":\"definirInicioVotacao\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qualDataFimVotacao\",\"type\":\"uint256\"}],\"name\":\"definirFimVotacao\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quandoEncerraVotacao\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numeroProposta\",\"type\":\"uint256\"}],\"name\":\"pesquisarProposta\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"qualMotivoDaConvocatoria\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"indiceVotante\",\"type\":\"address\"}],\"name\":\"pesquisarVotante\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enderecoVotante\",\"type\":\"address\"},{\"name\":\"quotaDeVotos\",\"type\":\"uint256\"},{\"name\":\"qualIDVotante\",\"type\":\"string\"}],\"name\":\"incluiVotante\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDePropostas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secretarioDesignado\",\"type\":\"address\"}],\"name\":\"designarSecretario\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"qualMotivoConvocatoria\",\"type\":\"string\"},{\"name\":\"qualNivelControle\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"quemVotou\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"propostaVotada\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"qualVoto\",\"type\":\"bool\"}],\"name\":\"Votou\",\"type\":\"event\"}]"

// VotacaoAssembleiaBin is the compiled bytecode used for deploying new contracts.
const VotacaoAssembleiaBin = `0x60806040523480156200001157600080fd5b50604051620017113803806200171183398101604052805160208083015160048054600160a060020a0319163317905591909201805190926200005a9160079185019062000073565b506008805460ff19169115159190911790555062000118565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000b657805160ff1916838001178555620000e6565b82800160010185558215620000e6579182015b82811115620000e6578251825591602001919060010190620000c9565b50620000f4929150620000f8565b5090565b6200011591905b80821115620000f45760008155600101620000ff565b90565b6115e980620001286000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166308ae343b81146100df57806328962d4e14610106578063302dbc20146101375780633c9afde31461014f5780634080037f1461018a5780634a727aba146101f65780636de7d9c4146102d15780637fbdbe5f146102e95780638755f2ad14610301578063a7e421cb14610316578063b0899b53146103d4578063ba67a8611461045e578063ce64e9951461047f578063cf5e72a8146104e8578063e030ed33146104fd575b600080fd5b3480156100eb57600080fd5b506100f461051e565b60408051918252519081900360200190f35b34801561011257600080fd5b506101236004356024351515610525565b604080519115158252519081900360200190f35b34801561014357600080fd5b506101236004356106ca565b34801561015b57600080fd5b506101676004356107ee565b60408051600160a060020a03909316835260208301919091528051918290030190f35b34801561019657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101f494369492936024939284019190819084018382808284375094975050508335600160a060020a03169450505060209091013590506109a8565b005b34801561020257600080fd5b5061020b610b2f565b6040518088600160a060020a0316600160a060020a0316815260200187600160a060020a0316600160a060020a0316815260200180602001868152602001858152602001848152602001838152602001828103825287818151815260200191508051906020019080838360005b83811015610290578181015183820152602001610278565b50505050905090810190601f1680156102bd5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156102dd57600080fd5b506101f4600435610c1c565b3480156102f557600080fd5b506101f4600435610d26565b34801561030d57600080fd5b506100f4610e30565b34801561032257600080fd5b5061032e600435610e36565b604051808681526020018060200185600160a060020a0316600160a060020a03168152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561039557818101518382015260200161037d565b50505050905090810190601f1680156103c25780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156103e057600080fd5b506103e9610f8e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561042357818101518382015260200161040b565b50505050905090810190601f1680156104505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561046a57600080fd5b50610167600160a060020a0360043516611024565b34801561048b57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f4948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506110ba9650505050505050565b3480156104f457600080fd5b506100f46113fe565b34801561050957600080fd5b506101f4600160a060020a0360043516611404565b6002545b90565b60008060004260065410151515610586576040805160e560020a62461bcd02815260206004820152601160248201527f566f746163616f20656e63657272616461000000000000000000000000000000604482015290519081900360640190fd5b6005544210156105e0576040805160e560020a62461bcd02815260206004820152601f60248201527f566f7461c3a7c3a36f2061696e6461206ec3a36f20666f692061626572746100604482015290519081900360640190fd5b60008054869081106105ee57fe5b60009182526020909120600590910201600481015490925060ff16156106bd57503360009081526001602081905260409091206003810154909161010090910460ff16151514156106bd57600381015460ff1615156106bd576001841515141561066357600280820154908301805490910190555b60408051338152602081018790528515158183015290517f62d81f5e43004f4fc2e3d45b0ff0c27b3c7bd87ee2a022db2cac7222ebf276aa9181900360600190a160038101805460ff1916600190811790915592506106c2565b600092505b505092915050565b60006106d46114bb565b60008054849081106106e257fe5b600091825260209182902060408051600593909302909101805460026001821615610100026000190190911604601f8101859004909402830160c090810190925260a08301848152929390928492909184918401828280156107855780601f1061075a57610100808354040283529160200191610785565b820191906000526020600020905b81548152906001019060200180831161076857829003601f168201915b50505091835250506001820154600160a060020a03166020820152600282015460408201526003820154606082015260049091015460ff161515608091820152810151909150156107e35780606001518160400151101591506107e8565b600091505b50919050565b6000806107f96114f6565b600254841115610879576040805160e560020a62461bcd02815260206004820152603260248201527f496e6469636520696e666f726d61646f20c3a9206d61696f7220717565206f2060448201527f6e756d65726f20646520766f74616e7465730000000000000000000000000000606482015290519081900360840190fd5b600280548590811061088757fe5b60009182526020918290206040805160a08101825260049093029091018054600160a060020a03168352600180820180548451601f600260001995841615610100029590950190921693909304908101879004870283018701909452838252939491938583019391929091908301828280156109445780601f1061091957610100808354040283529160200191610944565b820191906000526020600020905b81548152906001019060200180831161092757829003601f168201915b50505091835250506002820154602082015260039091015460ff80821615156040840152610100909104161515606090910152608081015190915015156001141561099a578051604082015190935091506109a2565b600092508291505b50915091565b6109b06114bb565b60085460ff1615610a4357600354600160a060020a03163314610a43576040805160e560020a62461bcd02815260206004820152602f60248201527f53c3b3206f20736563726574c3a172696f20706f6465207265616c697a61722060448201527f65737361206f70657261c3a7c3a36f0000000000000000000000000000000000606482015290519081900360840190fd5b506040805160a081018252848152600160a060020a038416602080830191909152600092820183905260608201849052600160808301819052835490810180855593805282518051939493859360059093027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630192610ac6928492910190611525565b50602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905560408201516002820155606082015160038201556080909101516004909101805460ff19169115159190911790555050505050565b6000806060600080600080600080610b456113fe565b9150610b4f61051e565b60045460035460055460065460078054604080516020601f60026000196001871615610100020190951694909404938401819004810282018101909252828152979850600160a060020a039687169795909616959194899289928791830182828015610bfc5780601f10610bd157610100808354040283529160200191610bfc565b820191906000526020600020905b815481529060010190602001808311610bdf57829003601f168201915b505050505094509850985098509850985098509850505090919293949596565b600454600160a060020a03163314610ca4576040805160e560020a62461bcd02815260206004820152603260248201527f536f6d656e7465206f20707265736964656e746520706f6465207265616c697a60448201527f61722065737361206f70657261c3a7c3a36f0000000000000000000000000000606482015290519081900360840190fd5b428111610d21576040805160e560020a62461bcd02815260206004820152602b60248201527f41206461746120696e666f726d616461206465766520736572206d61696f722060448201527f717565206120617475616c000000000000000000000000000000000000000000606482015290519081900360840190fd5b600555565b600454600160a060020a03163314610dae576040805160e560020a62461bcd02815260206004820152603260248201527f536f6d656e7465206f20707265736964656e746520706f6465207265616c697a60448201527f61722065737361206f70657261c3a7c3a36f0000000000000000000000000000606482015290519081900360840190fd5b428111610e2b576040805160e560020a62461bcd02815260206004820152602b60248201527f41206461746120696e666f726d616461206465766520736572206d61696f722060448201527f717565206120617475616c000000000000000000000000000000000000000000606482015290519081900360840190fd5b600655565b60065490565b600060606000806000610e476114bb565b6000805488908110610e5557fe5b600091825260209182902060408051600593909302909101805460026001821615610100026000190190911604601f8101859004909402830160c090810190925260a0830184815292939092849290918491840182828015610ef85780601f10610ecd57610100808354040283529160200191610ef8565b820191906000526020600020905b815481529060010190602001808311610edb57829003601f168201915b50505091835250506001820154600160a060020a03166020820152600282015460408201526003820154606082015260049091015460ff16151560809182015281015190915015610f665780516020820151604083015160608401518a995092975090955093509150610f84565b60408051602081019091526000808252965094508593508392508291505b5091939590929450565b60078054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561101a5780601f10610fef5761010080835404028352916020019161101a565b820191906000526020600020905b815481529060010190602001808311610ffd57829003601f168201915b5050505050905090565b60008061102f6114f6565b600160a060020a03848116600090815260016020818152604092839020835160a0810185528154909516855280830180548551600261010096831615969096026000190190911694909404601f8101849004840285018401909552848452909385830193928301828280156109445780601f1061091957610100808354040283529160200191610944565b6110c26114f6565b60085460ff161561115557600354600160a060020a03163314611155576040805160e560020a62461bcd02815260206004820152602f60248201527f53c3b3206f20736563726574c3a172696f20706f6465207265616c697a61722060448201527f65737361206f70657261c3a7c3a36f0000000000000000000000000000000000606482015290519081900360840190fd5b60638311156111d4576040805160e560020a62461bcd02815260206004820152602160248201527f51756f7461206e616f20706f646520736572207375706572696f72206120393960448201527f2500000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a038416151561125a576040805160e560020a62461bcd02815260206004820152602560248201527f4f20766f74616e746520646576652074657220756d20656e64657265636f207660448201527f616c69646f000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b506040805160a081018252600160a060020a038581168083526020808401868152848601889052600060608601819052600160808701819052938152838352959095208451815473ffffffffffffffffffffffffffffffffffffffff1916941693909317835593518051939485946112d89385019290910190611525565b50604082015160028281019190915560608301516003909201805460809094015160ff199094169215159290921761ff0019166101009315159390930292909217905580546001810180835560009290925282517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace6004909202918201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091178155602080850151805186946113b9937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf909101920190611525565b506040820151600282015560608201516003909101805460809093015115156101000261ff001992151560ff1990941693909317919091169190911790555050505050565b60005490565b600454600160a060020a0316331461148c576040805160e560020a62461bcd02815260206004820152603260248201527f536f6d656e7465206f20707265736964656e746520706f6465207265616c697a60448201527f61722065737361206f70657261c3a7c3a36f0000000000000000000000000000606482015290519081900360840190fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60a060405190810160405280606081526020016000600160a060020a0316815260200160008152602001600081526020016000151581525090565b6040805160a0810182526000808252606060208301819052928201819052918101829052608081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061156657805160ff1916838001178555611593565b82800160010185558215611593579182015b82811115611593578251825591602001919060010190611578565b5061159f9291506115a3565b5090565b61052291905b8082111561159f57600081556001016115a95600a165627a7a7230582099de376ff7840ce3b678a8d9edee044406164dd057957c610cbad6e12a18e97c0029`

// DeployVotacaoAssembleia deploys a new Ethereum contract, binding an instance of VotacaoAssembleia to it.
func DeployVotacaoAssembleia(auth *bind.TransactOpts, backend bind.ContractBackend, qualMotivoConvocatoria string, qualNivelControle bool) (common.Address, *types.Transaction, *VotacaoAssembleia, error) {
	parsed, err := abi.JSON(strings.NewReader(VotacaoAssembleiaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotacaoAssembleiaBin), backend, qualMotivoConvocatoria, qualNivelControle)
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
// Solidity: function pesquisarProposta(numeroProposta uint256) constant returns(uint256, string, address, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCaller) PesquisarProposta(opts *bind.CallOpts, numeroProposta *big.Int) (*big.Int, string, common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _VotacaoAssembleia.contract.Call(opts, out, "pesquisarProposta", numeroProposta)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// PesquisarProposta is a free data retrieval call binding the contract method 0xa7e421cb.
//
// Solidity: function pesquisarProposta(numeroProposta uint256) constant returns(uint256, string, address, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaSession) PesquisarProposta(numeroProposta *big.Int) (*big.Int, string, common.Address, *big.Int, *big.Int, error) {
	return _VotacaoAssembleia.Contract.PesquisarProposta(&_VotacaoAssembleia.CallOpts, numeroProposta)
}

// PesquisarProposta is a free data retrieval call binding the contract method 0xa7e421cb.
//
// Solidity: function pesquisarProposta(numeroProposta uint256) constant returns(uint256, string, address, uint256, uint256)
func (_VotacaoAssembleia *VotacaoAssembleiaCallerSession) PesquisarProposta(numeroProposta *big.Int) (*big.Int, string, common.Address, *big.Int, *big.Int, error) {
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

// IncluiVotante is a paid mutator transaction binding the contract method 0xce64e995.
//
// Solidity: function incluiVotante(enderecoVotante address, quotaDeVotos uint256, qualIDVotante string) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactor) IncluiVotante(opts *bind.TransactOpts, enderecoVotante common.Address, quotaDeVotos *big.Int, qualIDVotante string) (*types.Transaction, error) {
	return _VotacaoAssembleia.contract.Transact(opts, "incluiVotante", enderecoVotante, quotaDeVotos, qualIDVotante)
}

// IncluiVotante is a paid mutator transaction binding the contract method 0xce64e995.
//
// Solidity: function incluiVotante(enderecoVotante address, quotaDeVotos uint256, qualIDVotante string) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaSession) IncluiVotante(enderecoVotante common.Address, quotaDeVotos *big.Int, qualIDVotante string) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.IncluiVotante(&_VotacaoAssembleia.TransactOpts, enderecoVotante, quotaDeVotos, qualIDVotante)
}

// IncluiVotante is a paid mutator transaction binding the contract method 0xce64e995.
//
// Solidity: function incluiVotante(enderecoVotante address, quotaDeVotos uint256, qualIDVotante string) returns()
func (_VotacaoAssembleia *VotacaoAssembleiaTransactorSession) IncluiVotante(enderecoVotante common.Address, quotaDeVotos *big.Int, qualIDVotante string) (*types.Transaction, error) {
	return _VotacaoAssembleia.Contract.IncluiVotante(&_VotacaoAssembleia.TransactOpts, enderecoVotante, quotaDeVotos, qualIDVotante)
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
