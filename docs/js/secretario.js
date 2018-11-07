function limpaCamposProposta() {
    var frm = document.frmIncluiProposta;
    frm.textoProposta.value = "";
    frm.qualProponente.value = "";
    frm.quotaMinima.value = 0;
    $("#statusIncluiProposta").css("background-color", "white");
    $("#statusIncluiProposta").html("");  
}

function limpaCamposVotante() {
    var frm = document.frmIncluiVotante;
    frm.enderecoVotante.value = "";
    frm.quotaVotos.value = 0;
    frm.idVotante.value = "";
    $("#statusIncluiVotante").css("background-color", "white");
    $("#statusIncluiVotante").html("");  
}

function incluiVotante() {
    if (verificaCondicoesInteragirSmartContract()) {
        var frm = document.frmIncluiVotante;
        if (!web3.utils.isAddress(frm.enderecoVotante.value)) {
            $("#statusIncluiVotante").css("background-color", "Salmon");
            $("#statusIncluiVotante").html("Você deve informar um endereço Ethereum válido do votante.");
            return
        }
        let nroQuota = (frm.quotaVotos.value*1);
        if ((nroQuota<1) || (nroQuota>99)) {
            $("#statusIncluiVotante").css("background-color", "Salmon");
            $("#statusIncluiVotante").html("Você deve informar um valor de quota válida.");
            return
        }
        console.log("incluiVotante - Enviando... " + frm.enderecoVotante.value + " " + nroQuota + " " + frm.idVotante.value);
        contract.methods.incluiVotante(frm.enderecoVotante.value, nroQuota, frm.idVotante.value).send(trxObj)
            .on('transactionHash', function(hash){
                $("#statusIncluiVotante").css("background-color", "yellow");
                $("#statusIncluiVotante").html("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada.<br>Transaction hash: " + hash);
            })
            .on('confirmation', function(confirmationNumber, txReceipt) {
                console.log("Confirmation number: " + confirmationNumber + " - Recibo: " + txReceipt);
                if (confirmationNumber == 7) {
                    if (txReceipt) {
                        if (txReceipt.status == "0x1") {
                            $("#statusIncluiVotante").css("background-color", "LawnGreen");
                            $("#statusIncluiVotante").html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + "<br>Transaction hash: " + txReceipt.transactionHash + "<br>Numero confirmações: " + confirmationNumber);            
                        } else {
                            $("#statusIncluiVotante").css("background-color", "Salmon");
                            $("#statusIncluiVotante").html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                        }
                    }
                }
            })
            .on('error', function (err) {
                $("#statusIncluiVotante").css("background-color", "Salmon");
                $("#statusIncluiVotante").html("Aconteceu um erro: " + err);            
            });
    } else {
        $("#statusIncluiVotante").css("background-color", "white");
        $("#statusIncluiVotante").html("Há problemas de conexão com a rede blockchain do Ethereum");
    }
}

function incluiProposta() {
    if (verificaCondicoesInteragirSmartContract()) {
        var frm = document.frmIncluiProposta;
        if (!web3.utils.isAddress(frm.qualProponente.value)) {
            $("#statusIncluiProposta").css("background-color", "Salmon");
            $("#statusIncluiProposta").html("Você deve informar um endereço Ethereum válido do proponente da proposta.");
            return
        }
        let nroQuota = (frm.quotaMinima.value*1);
        if ((nroQuota<1) || (nroQuota>99)) {
            $("#statusIncluiProposta").css("background-color", "Salmon");
            $("#statusIncluiProposta").html("Você deve informar um valor de quota mínima para aprovação válida.");
            return
        }
        if (frm.textoProposta.value.length < 5) {
            $("#statusIncluiProposta").css("background-color", "Salmon");
            $("#statusIncluiProposta").html("Você deve informar um texto de proposta válido.");
            return
        }
        console.log("incluiProposta - Enviando... " + frm.qualProponente.value + " " + nroQuota + " " + frm.textoProposta.value);
        contract.methods.incluiProposta(frm.textoProposta.value, frm.qualProponente.value, nroQuota).send(trxObj)
            .on('transactionHash', function(hash){
                $("#statusIncluiProposta").css("background-color", "yellow");
                $("#statusIncluiProposta").html("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada <br>Transaction hash: " + hash);
            })
            .on('confirmation', function(confirmationNumber, txReceipt) {
                console.log("Confirmation number: " + confirmationNumber + " - Recibo: " + txReceipt);
                if (confirmationNumber == 7) {
                    if (txReceipt) {
                        if (txReceipt.status == "0x1") {
                            $("#statusIncluiProposta").css("background-color", "LawnGreen");
                            $("#statusIncluiProposta").html("Registro salvo com sucesso no bloco: " + txReceipt.blockNumber + "<br>Transaction hash: " + txReceipt.transactionHash + "<br>Numero confirmações: " + confirmationNumber);            
                        } else {
                            $("#statusIncluiProposta").css("background-color", "Salmon");
                            $("#statusIncluiProposta").html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                        }
                    }
                }
            })
            .on('error', function (err) {
                $("#statusIncluiProposta").css("background-color", "Salmon");
                $("#statusIncluiProposta").html("Aconteceu um erro: " + err);            
            });
    } else {
        $("#statusIncluiProposta").css("background-color", "white");
        $("#statusIncluiProposta").html("Há problemas de conexão com a rede blockchain do Ethereum");
    }
}
