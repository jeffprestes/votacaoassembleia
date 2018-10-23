pragma  solidity 0.4.25;

contract Assembleia {

    struct Proposta {
        string texto;
        address proponente;
        uint quotaDeVotos;
        uint quotaMinimaParaAprovacao;
        bool existe;
    }

    //Informações gerais da Assembleia
    Proposta[] propostas;
   
    function incluiProposta(string qualTextoDaProposta, address qualProponente, uint qualQuotaMinimaParaAprovacao) public somenteChairman {
        Proposta memory novaProposta = Proposta(qualTextoDaProposta, qualProponente, 0, qualQuotaMinimaParaAprovacao, true);
        propostas.push(novaProposta);
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
}