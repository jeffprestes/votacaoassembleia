//Inicializa os seletores de data
$('#dtPickerInicioVotacao').datetimepicker();
$('#dtPickerFimVotacao').datetimepicker();

function registraSecretario() {
    if (verificaCondicoesInteragirSmartContract()) {
        var frm = document.frmSecretario;
        if (!web3.utils.isAddress(frm.secretario.value)) {
            $("#statusSecretario").css("background-color", "Salmon");
            $("#statusSecretario").html("Você deve informar um endereço Ethereum válido.");
            return
        }
        console.log("registraSecretario - Enviando..." + frm.secretario.value);
        contract.methods.designarSecretario(frm.secretario.value).send({from: conta, gas: 3000000, value: 0})
            .on('transactionHash', function(hash){
                $("#statusSecretario").css("background-color", "yellow");
                $("#statusSecretario").text("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada. Transaction hash: " + hash);
            })
            .on('confirmation', function(confirmationNumber, txReceipt){
                if (confirmationNumber > 23) {
                    if (txReceipt) {
                        if (txReceipt.status == "0x1") {
                            $("#statusSecretario").css("background-color", "LawnGreen");
                            $("#statusSecretario").html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + " - Transaction hash: " + txReceipt.transactionHash + " - Numero confirmações: " + confirmationNumber);            
                        } else {
                            $("#statusSecretario").css("background-color", "Salmon");
                            $("#statusSecretario").html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                        }
                    }
                }
            })
            .on('error', function (err) {
                $("#statusSecretario").css("background-color", "Salmon");
                $("#statusSecretario").html("Aconteceu um erro: " + err);            
            });
    }
}

function registraInicioVotacao() {
    if (verificaCondicoesInteragirSmartContract()) {
        var frm = document.frmInicioVotacao;
        try {
            let tempo = converteDataParaTimestamp(frm.inicioVotacao.value);
            console.log("registraInicioVotacao - Enviando..." + tempo);
            contract.methods.definirInicioVotacao(tempo).send({from: conta, gas: 3000000, value: 0})
            .on('transactionHash', function(hash){
                $("#statusInicioVotacao").css("background-color", "yellow");
                $("#statusInicioVotacao").text("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada. Transaction hash: " + hash);
            })
            .on('confirmation', function(confirmationNumber, txReceipt){
                if (confirmationNumber == 7) {
                    if (txReceipt) {
                        if (txReceipt.status == "0x1") {
                            $("#statusInicioVotacao").css("background-color", "LawnGreen");
                            $("#statusInicioVotacao").html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + " - Transaction hash: " + txReceipt.transactionHash + " - Numero confirmações: " + confirmationNumber);            
                        } else {
                            $("#statusInicioVotacao").css("background-color", "Salmon");
                            $("#statusInicioVotacao").html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                        }
                    }
                }
            })
            .on('error', function (err) {
                $("#statusInicioVotacao").css("background-color", "Salmon");
                $("#statusInicioVotacao").html("Aconteceu um erro: " + err);            
            });
        } catch (err) {
            $("#statusInicioVotacao").css("background-color", "Salmon");
            $("#statusInicioVotacao").html("Aconteceu um erro: " + err);  
        }
    }
}

function registraFimVotacao() {
    if (verificaCondicoesInteragirSmartContract()) {
        var frm = document.frmFimVotacao;
        try {
            let tempo = converteDataParaTimestamp(frm.fimVotacao.value);
            console.log("registraFimVotacao - Enviando..." + tempo);
            contract.methods.definirFimVotacao(tempo).send({from: conta, gas: 3000000, value: 0})
            .on('transactionHash', function(hash){
                $("#statusFimVotacao").css("background-color", "yellow");
                $("#statusFimVotacao").text("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada. Transaction hash: " + hash);
            })
            .on('confirmation', function(confirmationNumber, txReceipt){
                if (confirmationNumber == 7) {
                    if (txReceipt) {
                        if (txReceipt.status == "0x1") {
                            $("#statusFimVotacao").css("background-color", "LawnGreen");
                            $("#statusFimVotacao").html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + " - Transaction hash: " + txReceipt.transactionHash + " - Numero confirmações: " + confirmationNumber);            
                        } else {
                            $("#statusFimVotacao").css("background-color", "Salmon");
                            $("#statusFimVotacao").html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                        }
                    }
                }
            })
            .on('error', function (err) {
                $("#statusFimVotacao").css("background-color", "Salmon");
                $("#statusFimVotacao").html("Aconteceu um erro: " + err);            
            });
        } catch (err) {
            $("#statusFimVotacao").css("background-color", "Salmon");
            $("#statusFimVotacao").html("Aconteceu um erro: " + err);  
        }
    }
}

function converteDataParaTimestamp(valor) {
    try {
        var dateParts = valor.match(/(\d+)\/(\d+)\/(\d+) (\d+):(\d+)/);
        var dataX = new Date(dateParts[3], ((dateParts[2])-1), dateParts[1], dateParts[4], dateParts[5], 0, 0);
        return (dataX.getTime()/1000);
    } catch (err) {
        console.error("Erro ao converter a data: " + err.message);
        throw err;
    }
}
