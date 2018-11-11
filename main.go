package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jung-kurt/gofpdf"

	"github.com/jeffprestes/votacaoassembleia/proxycontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatalf("Erro ao conectar ao nó do Ethereum: %s\n", err.Error())
	}
	contractAddress := common.HexToAddress("0xb1cdd1f13fbc5963ba5d4211300f40c10037bb03")
	contract, err := proxycontracts.NewVotacaoAssembleia(contractAddress, client)
	if err != nil {
		log.Fatalf("Erro ao gerar um proxy do Smart Contract: %s\n", err.Error())
	}
	filterer := new(bind.FilterOpts)
	filterer.Start = 3224024
	filterer.End = nil
	logs, err := contract.FilterVotou(filterer)
	if err != nil {
		log.Fatalf("Erro ao obter os logs dos eventos: %s\n", err.Error())
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 10)
	pdf.SetLeftMargin(10)

	var eventLog proxycontracts.VotacaoAssembleiaVotou
	for logs.Next() {
		eventLog = *logs.Event
		idProposta, descricaoProposta, proponente, _, _, err := contract.PesquisarProposta(nil, eventLog.PropostaVotada)
		if err != nil {
			log.Printf("Erro ao pesquisar os detalhes da proposta: %s\n", err.Error())
			continue
		}
		propostaAprovada, err := contract.PropostaAprovada(nil, eventLog.PropostaVotada)
		if err != nil {
			log.Printf("Erro ao checar se a proposta foi aprovada: %s\n", err.Error())
			continue
		}
		geraSinteseVoto(pdf, idProposta.Text(10), descricaoProposta, proponente, eventLog.QuemVotou, eventLog.QualVoto, propostaAprovada)
	}
	err = pdf.OutputFileAndClose("ata_votacao_assembleia.pdf")
	if err != nil {
		log.Fatalf("Erro gerar o arquivo da Ata: %s\n", err.Error())
	}
}

func geraSinteseVoto(pdf *gofpdf.Fpdf, idProposta string, descricaoProposta string, proponente common.Address, quemVotou common.Address, qualVoto bool, aPropostaFoiAprovada bool) {
	var sinteseVoto strings.Builder
	sinteseVoto.WriteString("Proposta: ")
	sinteseVoto.WriteString(idProposta + " - ")
	sinteseVoto.WriteString(descricaoProposta)
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("Proponente: ")
	sinteseVoto.WriteString(proponente.String())
	pdf.CellFormat(200, 10, sinteseVoto.String(), "", 1, "C", false, 0, "")
	fmt.Println(sinteseVoto.String())
	sinteseVoto.Reset()
	sinteseVoto.WriteString("Votante: ")
	sinteseVoto.WriteString(quemVotou.String())
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("Seu voto: ")
	sinteseVoto.WriteString(strconv.FormatBool(qualVoto))
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("A proposta foi aprovada? ")
	sinteseVoto.WriteString(strconv.FormatBool(aPropostaFoiAprovada))
	pdf.CellFormat(200, 10, sinteseVoto.String(), "", 1, "C", false, 0, "")
	fmt.Println(sinteseVoto.String())
	return
}
