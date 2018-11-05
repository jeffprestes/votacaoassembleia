const VotacaoAssembleiaABI = [{"constant":true,"inputs":[],"name":"totalDeVotantes","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"numeroProposta","type":"uint256"},{"name":"favoravelAProposta","type":"bool"}],"name":"votar","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"numeroProposta","type":"uint256"}],"name":"propostaAprovada","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"indiceVotante","type":"uint256"}],"name":"pesquisarVotantePorIndice","outputs":[{"name":"","type":"address"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"qualTextoDaProposta","type":"string"},{"name":"qualProponente","type":"address"},{"name":"qualQuotaMinimaParaAprovacao","type":"uint256"}],"name":"incluiProposta","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"detalhesAssembleia","outputs":[{"name":"","type":"address"},{"name":"","type":"address"},{"name":"","type":"string"},{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"qualDataInicioVotacao","type":"uint256"}],"name":"definirInicioVotacao","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"qualDataFimVotacao","type":"uint256"}],"name":"definirFimVotacao","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"quandoEncerraVotacao","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"enderecoVotante","type":"address"},{"name":"quotaDeVotos","type":"uint256"}],"name":"incluiVotante","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"numeroProposta","type":"uint256"}],"name":"pesquisarProposta","outputs":[{"name":"","type":"string"},{"name":"","type":"address"},{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"qualMotivoDaConvocatoria","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"indiceVotante","type":"address"}],"name":"pesquisarVotante","outputs":[{"name":"","type":"address"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"totalDePropostas","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"secretarioDesignado","type":"address"}],"name":"designarSecretario","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"qualMotivoConvocatoria","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}];

window.addEventListener('load', function() {

    // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (typeof window.web3 !== 'undefined') {
  
      // Use the browser's ethereum provider
      var provider = web3.currentProvider
  
    } else {
      console.log('No web3? You should consider trying MetaMask!')
    }
  
  })
  

var contract;
var conta;
/*
var web3;
if (typeof window.web3 !== 'undefined') {
    web3 = new Web3(window.web3.currentProvider);
} else {
    web3 = new Web3(new Web3.providers.HttpProvider("https://rinkeby.infura.io/QPF0qjGpH9OjFuuMrCse"))
}
*/
console.log("É o Metamask? "+ window.web3.currentProvider.isMetaMask);
contract = new web3.eth.Contract(VotacaoAssembleiaABI, "0x771FD37069b93bf0e41142712426D8046568897A"); 

window.web3.eth.getAccounts(function(err, accounts) {
    if (err != null) {
        console.error(err);
    } else {
        console.log(accounts);
    }
});
console.log("Conta metodo novo: " + web3.eth.defaultAccount);
console.log("Conta metodo velho: " + web3.eth.accounts[0]);

function waitForTxToBeMined(txHash, objStatus) {
    let txReceipt;
    web3.eth.getTransactionReceipt(txHash, function (err, result) {
        if (err) {
            console.error(err);
            return;
        }        
        txReceipt = result;
        let d = new Date();
        console.log("txHash: " + txHash + " at " + d);
        console.log(txReceipt); 
        if (txReceipt) {
            if (txReceipt.status == "0x1") {
                $(objStatus).css("background-color", "LawnGreen");
                $(objStatus).html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + " - Transaction hash: " + txHash);            
            } else {
                $(objStatus).css("background-color", "Salmon");
                $(objStatus).html("Aconteceu um erro - Transaction hash: " + txHash + " - Status final: " + txReceipt.status);            
            }
        } else {
            window.setTimeout(waitForTxToBeMined, 1500, txHash, objStatus);
        }
    });
}
