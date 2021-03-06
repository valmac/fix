package BidCompReqGrp_BidComponents

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type bidCompReqGrp_BidComponents_RegexValidator struct {
    ListID *(regexp.Regexp)
    Side *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    NetGrossInd *(regexp.Regexp)
    SettlType *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    Account *(regexp.Regexp)
    AcctIDSource *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myBidCompReqGrp_BidComponents_RegexValidator bidCompReqGrp_BidComponents_RegexValidator

func init() {
    myBidCompReqGrp_BidComponents_RegexValidator
    myBidCompReqGrp_BidComponents_RegexValidator.ListID = regexp.MustCompile(`[^|]*`)
    myBidCompReqGrp_BidComponents_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myBidCompReqGrp_BidComponents_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myBidCompReqGrp_BidComponents_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myBidCompReqGrp_BidComponents_RegexValidator.NetGrossInd = regexp.MustCompile(`-?\d+`)
    myBidCompReqGrp_BidComponents_RegexValidator.SettlType = regexp.MustCompile(`[^|]*`)
    myBidCompReqGrp_BidComponents_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myBidCompReqGrp_BidComponents_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myBidCompReqGrp_BidComponents_RegexValidator.AcctIDSource = regexp.MustCompile(`-?\d+`)
}



func (m *BidCompReqGrp_BidComponents) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




