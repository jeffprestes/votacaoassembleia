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

    modifier somenteSecretario() {
        require(secretario == msg.sender, "Só o Chairmain pode realizar essa operação");
        _;
    }

    //Informações gerais da Assembleia
    Proposta[] propostas;
    mapping (address =>Votante) votantes;
    Votante[] numeroVotantes;
    address secretario;
    address presidente;
    uint dataInicioVotacao;
    uint dataFimVotacao;
    string motivoConvocatoria;

    /*
    @qualDataFimVotacao - Unix timestamp - Numero de segundos decorridos desde 1 janeiro de 1970. 
    Formato usado amplamente no ambiente de programação. Para converter em data atual pode-se usar:
    https://www.unixtimestamp.com/index.php
    */
    constructor (string qualMotivoConvocatoria, uint qualDataFimVotacao) public {
        presidente = msg.sender;
        dataFimVotacao = qualDataFimVotacao;
        dataInicioVotacao = now;
        motivoConvocatoria = qualMotivoConvocatoria;
    }

    function designarSecretario(address secretarioDesignado) public {
        require(presidente == msg.sender, "Somente o presidente pode designar o secretario");
        secretario = secretarioDesignado;
    }

    function incluiVotante(address enderecoVotante, uint quotaDeVotos) public somenteSecretario {
        require(quotaDeVotos <= 99, "Quota nao pode ser superior a 99%");
        require(enderecoVotante != address(0), "O votante deve ter um endereco valido");
        Votante memory novoVotante = Votante(enderecoVotante, quotaDeVotos, false, true);
        votantes[enderecoVotante] = novoVotante;
        numeroVotantes.push(novoVotante);
    }

    function incluiProposta(string qualTextoDaProposta, address qualProponente, uint qualQuotaMinimaParaAprovacao) public somenteSecretario {
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

    function pesquisarVotantePorIndice(uint indiceVotante) public view returns (address, uint) {
        require(indiceVotante <= numeroVotantes.length, "Indice informado é maior que o numero de votantes");
        Votante memory votanteTemporario = numeroVotantes[indiceVotante];
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

    function totalDeVotantes() public view returns (uint) {
        return numeroVotantes.length;
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

    function detalhesAssembleia() public view returns(address, address, string, uint, uint, uint, uint) {
        uint tPropostas = totalDePropostas();
        uint tVotantes = totalDeVotantes();
        return (presidente, secretario, motivoConvocatoria, dataInicioVotacao, dataFimVotacao, tPropostas, tVotantes);
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