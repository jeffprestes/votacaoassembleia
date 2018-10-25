package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/jeffprestes/votacaoassembleia/proxycontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//Connect to mainnet
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	//Input address of 0x v1 contract and build event query
	contractAddress := common.HexToAddress("0x771fd37069b93bf0e41142712426d8046568897a")
	contract, err := proxycontracts.NewVotacaoAssembleia(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	filterer := new(bind.FilterOpts)
	filterer.Start = 3224024
	filterer.End = nil
	logs, err := contract.FilterVotou(filterer)
	if err != nil {
		log.Fatal(err)
	}
	var eventLog proxycontracts.VotacaoAssembleiaVotou

	for logs.Next() {
		eventLog = *logs.Event
		descricaoProposta, proponente, _, _, err := contract.PesquisarProposta(nil, eventLog.PropostaVotada)
		if err != nil {
			log.Fatal(err)
		}
		propostaAprovada, err := contract.PropostaAprovada(nil, eventLog.PropostaVotada)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(geraSinteseVoto(descricaoProposta, proponente, eventLog.QuemVotou, eventLog.QualVoto, propostaAprovada))
	}
}

func geraSinteseVoto(descricaoProposta string, proponente common.Address, quemVotou common.Address, qualVoto bool, aPropostaFoiAprovada bool) (voto string) {
	var sinteseVoto strings.Builder
	sinteseVoto.WriteString("Proposta: ")
	sinteseVoto.WriteString(descricaoProposta)
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("Proponente: ")
	sinteseVoto.WriteString(proponente.String())
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("Votante: ")
	sinteseVoto.WriteString(quemVotou.String())
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("Seu voto: ")
	sinteseVoto.WriteString(strconv.FormatBool(qualVoto))
	sinteseVoto.WriteString(" - ")
	sinteseVoto.WriteString("A proposta foi aprovada? ")
	sinteseVoto.WriteString(strconv.FormatBool(aPropostaFoiAprovada))
	voto = sinteseVoto.String()
	return
}
