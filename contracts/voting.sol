pragma  solidity 0.4.25;

/**
@title Voting
@notice Smart contract to store voting results
@author Jeff Prestes
*/
contract Voting {

    struct Proposal {
        string text;
        address proponent;
        uint quotaOfVotes;
        uint quotaMinimumForApproval;
        bool exists;
    }

    struct Voter {
        address account;
        uint quotaOfVotes;
        bool voted;
        bool exists;
    }

    modifier secretaryOnly() {
        require(secretary == msg.sender, "Only the secretary can carry out this operation");
        _;
    }

    modifier presidentOnly() {
        require(president == msg.sender, "Only the president can carry out this operation");
        _;
    }

    //Event to be discharged when a voter has defined his or her vote
    event Voted(address whoVoted, uint proposalVoted, bool whichVote);

    //General information
    Proposal[] proposals;
    mapping (address =>Voter) voters;
    Voter[] numberOFVoters;
    address secretary;
    address president;
    uint startinglVotingDate;
    uint endingVotingDate;
    string reasonForInvitation;

    /**
    @notice The person who will preside over the voting must register the Smart Contract
    @dev Creates a specific contract for a vote. The designated president is the one who registered the Smart Contract on the network
    @param reasonForInvitation - The designated president writes the reason for the vote
    */
    constructor (string reasonForInvitation) public {
        president = msg.sender;
        reasonForInvitation = whatIsTheReason;
    }

    /**
    @notice The President of the Assembly appoints a secretary who will register the proposals and the eligible voters
    @param Designatedsecretary - Ethereum address of the Secretary 
    */
    function Designatesecretary(address Designatedsecretary) public somentePresidente {
        secretario = secretarioDesignado;
    }

    /** 
    @notice Presidente da Assembleia define quando a votação se encerrará
    @param qualDataFimVotacao - A data e hora limite para votação
                                Em formato Unix timestamp - A data Numero de segundos decorridos desde 1 janeiro de 1970. 
                                Formato usado amplamente no ambiente de programação. Para converter em data atual pode-se usar:
    https://www.unixtimestamp.com/index.php
    */
    function definirFimVotacao(uint qualDataFimVotacao) public somentePresidente {
        require(qualDataFimVotacao > now, "A data informada deve ser maior que a atual");
        dataFimVotacao = qualDataFimVotacao;
    }

    /** 
    @notice Presidente da Assembleia define quando a votação se iniciará
    @param qualDataInicioVotacao - A data e hora para inicio para votação
                                Em formato Unix timestamp - A data Numero de segundos decorridos desde 1 janeiro de 1970. 
                                Formato usado amplamente no ambiente de programação. Para converter em data atual pode-se usar:
    https://www.unixtimestamp.com/index.php
    */
    function definirInicioVotacao(uint qualDataInicioVotacao) public somentePresidente {
        require(qualDataInicioVotacao > now, "A data informada deve ser maior que a atual");
        dataInicioVotacao = qualDataInicioVotacao;
    }

    /** 
    @notice função a ser utilizado somente pelo Secretário para incluir um novo participante apto a votar (votante)
    @param enderecoVotante - Endereço Ethereum da conta do votante
    @param quotaDeVotos - Quantidade de votos (ou percentual de ações/participações) que possuí o votante
    */
    function incluiVotante(address enderecoVotante, uint quotaDeVotos) public somenteSecretario {
        require(quotaDeVotos <= 99, "Quota nao pode ser superior a 99%");
        require(enderecoVotante != address(0), "O votante deve ter um endereco valido");
        Votante memory novoVotante = Votante(enderecoVotante, quotaDeVotos, false, true);
        votantes[enderecoVotante] = novoVotante;
        numeroVotantes.push(novoVotante);
    }

    /** 
    @notice função a ser utilizado somente pelo Secretário para incluir uma nova proposta a ser apreciada (votada) pela Assembleia
    @param qualTextoDaProposta - Texto da proposta a ser apreciada
    @param qualProponente - Endereço Ethereum da conta do membro da Assembleia que trouxe a proposta a votação
    @param qualQuotaMinimaParaAprovacao -   Quantidade de votos (ou percentual de ações/participações) mínima que por estatuto/contrato a 
                                            proposta deve receber para ser aprovada
    */
    function incluiProposta(string qualTextoDaProposta, address qualProponente, uint qualQuotaMinimaParaAprovacao) public somenteSecretario {
        Proposta memory novaProposta = Proposta(qualTextoDaProposta, qualProponente, 0, qualQuotaMinimaParaAprovacao, true);
        propostas.push(novaProposta);
    }
    

    /** 
    @notice função a ser usada para obter na prática o número de votos (ou percentual de ações/participações) de um votante
    @param indiceVotante - Endereço Ethereum da conta do votante
    @return Endereço Ethereum da conta do votante
    @return Número de votos (ou percentual de ações/participações) do votante
    */
    function pesquisarVotante(address indiceVotante) public view returns (address, uint) {
        Votante memory votanteTemporario = votantes[indiceVotante];
        if (votanteTemporario.existe == true) {
            return (votanteTemporario.conta, votanteTemporario.quotaDeVotos);
        } else {
            return (0,0);
        }
    }

    /** 
    @notice função a ser usada para obter na prática o número de votos (ou percentual de ações/participações) de um votante
    @param indiceVotante - Número inteiro e único que identifica o votante
    @return Endereço Ethereum da conta do votante
    @return Número de votos (ou percentual de ações/participações) do votante
    */
    function pesquisarVotantePorIndice(uint indiceVotante) public view returns (address, uint) {
        require(indiceVotante <= numeroVotantes.length, "Indice informado é maior que o numero de votantes");
        Votante memory votanteTemporario = numeroVotantes[indiceVotante];
        if (votanteTemporario.existe == true) {
            return (votanteTemporario.conta, votanteTemporario.quotaDeVotos);
        } else {
            return (0,0);
        }
    }

    /** 
    @notice função a ser usada para obter detalhes de uma proposta a ser apreciada (votada) na Assembleia
    @param numeroProposta - Número inteiro e único que identifica a proposta
    @return Texto da proposta
    @return Endereço Ethereum da conta do proponente
    @return Número de votos (ou percentual de ações/participações) de apoio recebido
    @return Número de votos (ou percentual de ações/participações) de apoio mínimo que a proposta necessita para ser aprovada
    */
    function pesquisarProposta(uint numeroProposta) public view returns (string, address, uint, uint)  {
        Proposta memory propostaTemporario = propostas[numeroProposta];
        if (propostaTemporario.existe) {
            return (propostaTemporario.texto, propostaTemporario.proponente, propostaTemporario.quotaDeVotos, propostaTemporario.quotaMinimaParaAprovacao);
        } else {
            return ("", 0, 0, 0);
        }
    }

    /** 
    @notice Total de propostas registradas na Assembléia
    @return Numero total de propostas registradas na Assembléia
    */
    function totalDePropostas() public view returns (uint) {
        return propostas.length;
    }

    /** 
    @notice Total de participantes aptos a votar (votantes) registradas na Assembléia
    @return Numero total de participantes aptos a votar (votantes) registradas na Assembléia
    */
    function totalDeVotantes() public view returns (uint) {
        return numeroVotantes.length;
    }

    /** 
    @notice Qual o motivo razão da Assembléia ter sido convocada
    @return Texto que explica o motivo razão da Assembléia ter sido convocada
    */
    function qualMotivoDaConvocatoria() public view returns (string) {
        return motivoConvocatoria;
    }

    /** 
    @notice Informa se a proposta foi aprovada ou não
    @return Verdadeiro (true) caso a proposta tenha recebidos votos (ou percentual de ações/participações) de apoio mínimo 
            para ser considerada aprovada
    */
    function propostaAprovada(uint numeroProposta) public view returns (bool)  {
        Proposta memory propostaTemporario = propostas[numeroProposta];
        if (propostaTemporario.existe) {
            return propostaTemporario.quotaDeVotos>=propostaTemporario.quotaMinimaParaAprovacao;
        } else {
            return false;
        }
    }

    /** 
    @notice Informa quando a votação se encerrará
    @return Unix timestamp de quando a votação será encerrada
    */
    function quandoEncerraVotacao() public view returns (uint) {
        return dataFimVotacao;
    }

    /** 
    @notice Traz um resumo da Assembleia
    @return Endereço Ethereum da conta do presidente da Assembleia
    @return Endereço Ethereum da conta do secretário da Assembleia
    @return Texto que explica o motivo razão da Assembléia ter sido convocada
    @return Quando a votação abriu (formato Unix timestamp) - para converter acesse: https://www.unixtimestamp.com/index.php
    @return Quando a votação será ou foi encerrada (formato Unix timestamp) - para converter acesse: https://www.unixtimestamp.com/index.php
    @return Total de propostas enviadas para apreciação
    @return Total de votantes aptos a votar
    */
    function detalhesAssembleia() public view returns(address, address, string, uint, uint, uint, uint) {
        uint tPropostas = totalDePropostas();
        uint tVotantes = totalDeVotantes();
        return (presidente, secretario, motivoConvocatoria, dataInicioVotacao, dataFimVotacao, tPropostas, tVotantes);
    }

    /** 
    @notice função a ser executada pelos votantes para registrar se aprovam ou não uma proposta
            A votação deve estar aberta para ser permitida
            Um votante só tem permissão de votar uma vez e o voto não pode ser alterado.
    @param numeroProposta - Numero identificador da proposta
    @param favoravelAProposta - Verdadeiro/Falso onde o votante expressa se aprova ou não a proposta
    @return Verdadeiro caso o voto tenha sido computado com sucesso
    */
    function votar(uint numeroProposta, bool favoravelAProposta) public returns (bool) {
        require(dataFimVotacao>=now, "Votacao encerrada");
        require(dataInicioVotacao<=now, "Votação ainda não foi aberta");
        Proposta storage propostaTemporario = propostas[numeroProposta];
        if (propostaTemporario.existe) {
            Votante storage votanteTemporario = votantes[msg.sender];
            if (votanteTemporario.existe == true) {
                if (votanteTemporario.votou == false) {
                    if (favoravelAProposta == true) {
                        propostaTemporario.quotaDeVotos = propostaTemporario.quotaDeVotos + votanteTemporario.quotaDeVotos;
                    }
                    emit Votou(msg.sender, numeroProposta, favoravelAProposta);
                    votanteTemporario.votou = true;
                    return true;
                } 
            } 
        } 
        return false;
    }
}
