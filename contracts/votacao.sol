pragma  solidity 0.4.25;

contract Assembleia {

    struct Proposta {
        string texto;
        address proponente;
        uint quotaDeVotos;
        uint quotaMinimaParaAprovacao;
        bool existe;
    }

    struct Votante {
        address conta;
        uint quotaDeVotos;
        bool votou;
        bool existe;
    }

    modifier somenteChairman() {
        require(chairman == msg.sender, "Só o Chairmain pode realizar essa operação");
        _;
    }

    //Informações gerais da Assembleia
    Proposta[] propostas;
    mapping (address =>Votante) votantes;
    address chairman;
    uint dataInicioVotacao;
    uint dataFimVotacao;
    string motivoConvocatoria;

    constructor (string qualMotivoConvocatoria, uint qualDataFimVotacao) public {
        chairman = msg.sender;
        dataFimVotacao = qualDataFimVotacao;
        motivoConvocatoria = qualMotivoConvocatoria;
    }

    function incluiVotante(address enderecoVotante, uint quotaDeVotos) public somenteChairman {
        require(quotaDeVotos <= 99, "Quota nao pode ser superior a 99%");
        require(enderecoVotante != address(0), "O votante deve ter um endereco valido");
        Votante memory novoVotante = Votante(enderecoVotante, quotaDeVotos, false, true);
        votantes[enderecoVotante] = novoVotante;
    }

    function incluiProposta(string qualTextoDaProposta, address qualProponente, uint qualQuotaMinimaParaAprovacao) public somenteChairman {
        Proposta memory novaProposta = Proposta(qualTextoDaProposta, qualProponente, 0, qualQuotaMinimaParaAprovacao, true);
        propostas.push(novaProposta);
    }
    
    function pesquisarVotante(address indiceVotante) public view returns (address, uint) {
        Votante memory votanteTemporario = votantes[indiceVotante];
        if (votanteTemporario.existe == true) {
            return (votanteTemporario.conta, votanteTemporario.quotaDeVotos);
        } else {
            return (0,0);
        }
    }

    function pesquisarProposta(uint numeroProposta) public view returns (string, address, uint, uint)  {
        Proposta memory propostaTemporario = propostas[numeroProposta];
        if (propostaTemporario.existe) {
            return (propostaTemporario.texto, propostaTemporario.proponente, propostaTemporario.quotaDeVotos, propostaTemporario.quotaMinimaParaAprovacao);
        } else {
            return ("", 0, 0, 0);
        }
    }

    function totalDePropostas() public view returns (uint) {
        return propostas.length;
    }

    function qualMotivoDaConvocatoria() public view returns (string) {
        return motivoConvocatoria;
    }

    function propostaAprovada(uint numeroProposta) public view returns (bool)  {
        Proposta memory propostaTemporario = propostas[numeroProposta];
        if (propostaTemporario.existe) {
            return propostaTemporario.quotaDeVotos>=propostaTemporario.quotaMinimaParaAprovacao;
        } else {
            return false;
        }
    }

    function quandoEncerraVotacao() public view returns (uint) {
        return dataFimVotacao;
    }

    function votar(uint numeroProposta, bool favoravelAProposta) public returns (bool) {
        require(dataFimVotacao>=now, "Votacao encerrada");
        Proposta storage propostaTemporario = propostas[numeroProposta];
        if (propostaTemporario.existe) {
            Votante storage votanteTemporario = votantes[msg.sender];
            if (votanteTemporario.existe == true) {
                if (votanteTemporario.votou == false) {
                    if (favoravelAProposta == true) {
                        propostaTemporario.quotaDeVotos = propostaTemporario.quotaDeVotos + votanteTemporario.quotaDeVotos;
                    }
                    votanteTemporario.votou = true;
                    return true;
                } 
            } 
        } 
        return false;
    }
}