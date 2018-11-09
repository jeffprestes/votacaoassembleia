var totalDePropostas = 0;

function obterTotalDePropostas() {
    if (verificaCondicoesInteragirSmartContract()) {
        contract.methods.totalDePropostas().call(trxObj, function (err, result) {
            if (err)    {
                console.error("totalDePropostas - Erro: " + err);
            } else {
                totalDePropostas = (result*1);
                pesquisarPropostas(totalDePropostas);
            }
        });
    } else {
        console.log("nao foi");
    }
}

function pesquisarPropostas(nroTotalDePropostas) {
    if (verificaCondicoesInteragirSmartContract()) {
        for (i=0; i<nroTotalDePropostas; i++) {
            contract.methods.pesquisarProposta(i).call(trxObj, function (err, proposta) {
                if (err)    {
                    console.error("pesquisarPropostas - Erro: " + err);
                } else {
                    console.log("pesquisarPropostas resultados");
                    console.log(proposta);
                    var itm = document.getElementById("template");
                    var cln = itm.cloneNode(true);
                    cln.id = "div_" + proposta[0];
                    var titProposta = cln.getElementsByClassName("panel-heading")[0];
                    titProposta.innerText = "Proposta nro: " + ((proposta[0]*1)+1) + " - " + proposta[1];
                    var cboVoto = cln.getElementsByTagName("SELECT")[0];
                    cboVoto.id = "cbovoto_" + proposta[0];
                    cboVoto.setAttribute("name", cboVoto.id); 
                    var divStatusProposta = cln.getElementsByClassName("statusvoto")[0];
                    divStatusProposta.id = "status_" + proposta[0];
                    var spanProponente = cln.getElementsByTagName("SPAN")[0];
                    spanProponente.innerText = proposta[2];
                    cln.setAttribute("style", "display: block");
                    document.getElementById("propostas").appendChild(cln);
                }
            });
        }
    }
}

function registrarVoto(event) {
    console.log("Enviando... ID: " + obtemIDProposta($(event.target).attr('id')) + " - " + $(event.target).val());
    var objStatus = "#status_" + obtemIDProposta($(event.target).attr('id'));
    $(objStatus).text("Por favor, confirme o envio da transação no Metamask");
    contract.methods.votar(obtemIDProposta($(event.target).attr('id')), $(event.target).val()).send(trxObj)
    .on('transactionHash', function(hash) {
        var objStatus = "#status_" + obtemIDProposta($(event.target).attr('id'));
        $(objStatus).css("background-color", "yellow");
        $(objStatus).html("Transação enviada a rede do Ethereum. Aguarde enquanto ela é confirmada <br>Transaction hash: " + hash);
    })
    .on('confirmation', function(confirmationNumber, txReceipt) {
        //console.log("Confirmation number: " + confirmationNumber + " - Recibo: " + txReceipt);
        if (confirmationNumber == 2) {
            if (txReceipt) {
                var objStatus = "#status_" + obtemIDProposta($(event.target).attr('id'));
                if (txReceipt.status == "0x1") {
                    $(event.target).prop("disabled", true );
                    $(objStatus).css("background-color", "LawnGreen");
                    $(objStatus).html("Voto computado com sucesso no bloco: " + txReceipt.blockNumber + "<br>Transaction hash: " + txReceipt.transactionHash + "<br>Numero confirmações: " + confirmationNumber);            
                } else {
                    $(objStatus).css("background-color", "Salmon");
                    $(objStatus).html("Aconteceu um erro - Transaction hash: " + txReceipt.transactionHash + " - Status final: " + txReceipt.status);            
                }
            }
        }
    })
    .on('error', function (err) {
        var objStatus = "#status_" + obtemIDProposta($(event.target).attr('id'));
        $(objStatus).css("background-color", "Salmon");
        $(objStatus).html("Aconteceu um erro: " + err);            
    });

}

function obtemIDProposta(objID) {
    return objID.substr(objID.indexOf("_") + 1);
}

//Area para execução das funções quando a página for carregada
$(window).on('smartcontractdisponivel', function (e) {
    console.log('smartcontract disponivel? ', e.state);
    obterTotalDePropostas();
});