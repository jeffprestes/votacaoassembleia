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
        secretary = designateSecretary;
    }

    /** 
    @notice The designated president defines when voting ends
    @param endingVotingDate - the date and hour limit for voting
                                In Unix timestamp format - The Date and Number of seconds elapsed since January 1, 1970.
                                Format widely used in the programming environment. To convert to current date you can use:
    https://www.unixtimestamp.com/index.php
    */
    function defineEndOfVoting(uint endingVotingDate) public presidentOnly {
        require(whichEndOfVotingDate > now, "The date defined must be larger than the current date");
        endOfVotingDate = whichEndOfVotingDate;
    }

    /** 
    @notice Designated presidnt will define when voting will begin
    @param whichBegginingOfVotingDate - The date and hour voting will begin
                               In Unix timestamp format - The Date and Number of seconds elapsed since January 1, 1970.
                                Format widely used in the programming environment. To convert to current date you can use:
    https://www.unixtimestamp.com/index.php
    */
    function defineWhenVotingBegins(uint whichBegginingOfVotingDate) public presidentOnly {
        require(whichBegginingOfVotingDate > now, "The date defined must be larger than the current date");
       begginingOfVotingDate = whichBegginingOfVotingDate;
    }

    /** 
    @notice function to be used only by the Secretary to include a new eligible voter (voter)
    @param enderecoVotante - Ethereum address of the voters account
    @param quotaOfVotes - Quantity of votes (or the percentage of owenership) possesed by the voter
    */
    function includeVoter(address voterAddress, uint quotaOfVotes) public secretaryOnly {
        require(quotaOfVotess <= 99, "Quota can not be greater than 99%");
        require(voterAddress != address(0), "The voter must have a valid address");
        Voter memory newVoter = Voter(voterAddress, quotaOfVotes, false, true);
        Voters[voterAddress] = newVoter;
        numberOfVoters.push(newVoter);
    }

    /** 
    @notice only to be used by the Secretary to include a new proposal to be considered (voted on) by the Assembly
    @param whichProposalText - Text of the proposal to be considered
    @param whichProponent - Ethereum address of the member's account that brought the proposal to a vote
    @param whichQuotaIsTheMinimumForApproval -   Quantity of votes (or the percentage of owenership) minimum a proposal must receive for a statute / contract to be approved
    */
    function includeProposal(string whichProposalText, address whichProponent, uint whichQuotaIsTheMinimumForApproval) public secretaryOnly {
        Proposal memory newProposal = Proposal(whichProposalText, whichProponent, 0, whichQuotaIsTheMinimumForApproval, true);
        proposals.push(newProposal);
    }
    

    /** 
    @notice function to be used to obtain the number of votes (or percentage of shares) of the voter
    @param voterIndex - Ethereum address of the voter
    @return Ethereum address of the voter
    @return number of votes (or percentage of shares) of the voter
    */
    function searchVoter(address voterIndex) public view returns (address, uint) {
        Voter memory temporaryVoter = votantes[indiceVotante];
        if (temporaryVoter.exists == true) {
            return (temporaryVoter.address, temporaryVoter.quotaOfVotes);
        } else {
            return (0,0);
        }
    }

    /** 
    @notice function to be used to obtain the number of votes (or percentage of shares) of the voter
    @param voterIndex - Unique voter identification number
    @return Ethereum address of the voter
    @return number of votes (or percentage of shares) of the voter
    */
    function searchVoterByIndex(uint voterIndex) public view returns (address, uint) {
        require(voterIndex <= numberOfVoters.length, "Index listed is larger than number of voters");
        Voter memory temporaryVoter = numberOfVotes[voterIndex];
        if (temporaryVoter.exists == true) {
            return (temporaryVoter.address, temporaryVoter.quotaOfVotes);
        } else {
            return (0,0);
        }
    }

    /** 
    @notice function to be used to obtain details of a proposal to be considered (voted on) 
    @param proposalNumber - Unique proposal identification number
    @return Text of proposal
    @return Ethereum address of the proponent
    @return Number of votes (or percentage of shares) received supporting the proposal
    @return Number of votes (or percentage of shares)  needed for the proposal to be approved
    */
    function searchProposal(uint numeroProposta) public view returns (string, address, uint, uint)  {
        Proposal memory temporaryproposal = proposals[numberProposal];
        if (temporaryproposal.exists) {
            return (temporaryproposal.text, temporaryproposal.proponent, temporaryproposal.quotaOfVotes, temporaryproposal.quotaMinimumForApproval);
        } else {
            return ("", 0, 0, 0);
        }
    }

    /** 
    @notice Total of registered proposals
    @return Total number of registered proposals
    */
    function totalOfProposals() public view returns (uint) {
        return proposals.length;
    }

    /** 
    @notice Total participants eligible to vote (voters) registered in the Assembly
    @return Total number of participants eligible to vote (voters) registered in the Assembly
    */
    function totalOfVoters() public view returns (uint) {
        return numberOfVoters.length;
    }

    /** 
    @notice What is the reason the meeting was called
    @return Text explaining why the meeting was called
    */
    function whatIsTheReasonForMeeting() public view returns (string) {
        return reasonForMeeting;
    }

    /** 
    @notice Inform if the proposal was approved or not
    @return True if the proposal received the minimum number of votes (or percentage of shares) to be considered approved
    */
    function proposalApproved(uint proposalNumber) public view returns (bool)  {
        Proposal memory temporaryProposal = proposals[proposalNumber];
        if (propostaTemporario.existe) {
            return temporaryProposal.quotaOfVotes>=temporaryProposal.quotaMinimumForApproval;
        } else {
            return false;
        }
    }

    /** 
    @notice Inform when voting has ended
    @return Unix timestamp of when the voting ended
    */
    function endOfVotingDate() public view returns (uint) {
        return endOfVotingDate;
    }

    /** 
    @notice Summary of the Assembly
    @return Ethereum address of the president
    @return Ethereum address of the secretary
    @return Text explaining why the meeting was called
    @return When the vote opened(format Unix timestamp) - To convert: https://www.unixtimestamp.com/index.php
    @return When the vote has or will end (format Unix timestamp) - To convert: https://www.unixtimestamp.com/index.php
    @return Total number of proposals submitted for consideration
    @return Total eligible voters
    */
    function meetingDetails() public view returns(address, address, string, uint, uint, uint, uint) {
        uint tProposals = totalOfProposals();
        uint tVoters = totalNumberOfVoters();
        return (president, secretary, reasonForTheMeeting, begginingOfVotingDate, endOfVotingDate, tProposals, tVoters);
    }

    /** 
    @notice to be performed by voters to register whether or not they approve a proposal
           Voting must be open to permit
            A voter is only allowed to vote once and the vote can not be changed.
    @param proposalNumber - Unique proposal identification number
    @param inFavorOfTheProposal - True / False where the voter expresses whether or not he approves the proposal
    @return True if the vote was successfully computed
    */
    function vote(uint proposalNumber, bool inFavorOfTheProposal) public returns (bool) {
        require(endOfVotingDate>=now, "Voting ended");
        require(begginingOfVotingDate<=now, "Voting is not open yet");
        Proposal storage temporaryProposal = proposals[proposalNumber];
        if (temporaryProposal.exists) {
            Voter storage temporaryVoter = voters[msg.sender];
            if (votanteTemporario.existe == true) {
                if (votanteTemporario.votou == false) {
                    if (favoravelAProposta == true) {
                        temporaryProposal.quotaOfvotes = temporaryProposal.quotaOfVotes + temporaryVoter.quotaOfVotes;
                    }
                    emit Voted(msg.sender, proposalNumber, inFavorOfTheProposal);
                    temporaryVoter.voted = true;
                    return true;
                } 
            } 
        } 
        return false;
    }
}
