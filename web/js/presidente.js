function registraSecretario() {
    var frm = document.frmSecretario;
    if (frm.secretario.value.length < 40 || !web3.utils.isAddress(frm.secretario.value)) {
        $("#statusSecretario").css("background-color", "Salmon");
        $("#statusSecretario").html("Você deve informar um endereço Ethereum válido.");
        return
    }
    console.log("Enviando..." + frm.secretario.value);
    contract.methods.designarSecretario(frm.secretario.value).send({from: web3.eth.defaultAccount, gas: 3000000, value: 0})
        .on('transactionHash', function(hash){
            $("#statusSecretario").css("background-color", "yellow");
            $("#statusSecretario").text("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada. Transaction hash: " + result);
        })
        .on('receipt', function(receipt){
            console.log("Recibo: " + receipt)
        })
        .on('confirmation', function(confirmationNumber, txReceipt){
            console.log("Numero de confirmações: " + confirmationNumber);
            console.log(txReceipt); 
            if (txReceipt) {
                if (txReceipt.status == "0x1") {
                    $("#statusSecretario").css("background-color", "LawnGreen");
                    $("#statusSecretario").html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + " - Transaction hash: " + txReceipt.transactionHash);            
                } else {
                    $("#statusSecretario").css("background-color", "Salmon");
                    $("#statusSecretario").html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                }
            }
        })
        .on('error', function (err) {
            $("#statusSecretario").css("background-color", "Salmon");
            $("#statusSecretario").html("Aconteceu um erro: " + err);            
        });
}