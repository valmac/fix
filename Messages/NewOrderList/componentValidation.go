package NewOrderList

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import ListOrdGrp_Orders "grgrbll/fix/Shared/ListOrdGrp_Orders"

import RootParties_RootPartyIDs "grgrbll/fix/Shared/RootParties_RootPartyIDs"


type newOrderList_RegexValidator struct {
    BeginString *(regexp.Regexp)
    BodyLength *(regexp.Regexp)
    MsgType *(regexp.Regexp)
    SenderCompID *(regexp.Regexp)
    TargetCompID *(regexp.Regexp)
    OnBehalfOfCompID *(regexp.Regexp)
    DeliverToCompID *(regexp.Regexp)
    SecureDataLen *(regexp.Regexp)
    SecureData *(regexp.Regexp)
    MsgSeqNum *(regexp.Regexp)
    SenderSubID *(regexp.Regexp)
    SenderLocationID *(regexp.Regexp)
    TargetSubID *(regexp.Regexp)
    TargetLocationID *(regexp.Regexp)
    OnBehalfOfSubID *(regexp.Regexp)
    OnBehalfOfLocationID *(regexp.Regexp)
    DeliverToSubID *(regexp.Regexp)
    DeliverToLocationID *(regexp.Regexp)
    PossDupFlag *(regexp.Regexp)
    PossResend *(regexp.Regexp)
    SendingTime *(regexp.Regexp)
    OrigSendingTime *(regexp.Regexp)
    XmlDataLen *(regexp.Regexp)
    XmlData *(regexp.Regexp)
    MessageEncoding *(regexp.Regexp)
    LastMsgSeqNumProcessed *(regexp.Regexp)
    ApplVerID *(regexp.Regexp)
    CstmApplVerID *(regexp.Regexp)
    ListID *(regexp.Regexp)
    BidID *(regexp.Regexp)
    ClientBidID *(regexp.Regexp)
    ProgRptReqs *(regexp.Regexp)
    BidType *(regexp.Regexp)
    ProgPeriodInterval *(regexp.Regexp)
    CancellationRights *(regexp.Regexp)
    MoneyLaunderingStatus *(regexp.Regexp)
    RegistID *(regexp.Regexp)
    ListExecInstType *(regexp.Regexp)
    ListExecInst *(regexp.Regexp)
    EncodedListExecInstLen *(regexp.Regexp)
    EncodedListExecInst *(regexp.Regexp)
    AllowableOneSidednessPct *(regexp.Regexp)
    AllowableOneSidednessValue *(regexp.Regexp)
    AllowableOneSidednessCurr *(regexp.Regexp)
    TotNoOrders *(regexp.Regexp)
    LastFragment *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myNewOrderList_RegexValidator newOrderList_RegexValidator

func init() {
    myNewOrderList_RegexValidator
    myNewOrderList_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myNewOrderList_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myNewOrderList_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myNewOrderList_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myNewOrderList_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myNewOrderList_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myNewOrderList_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.ListID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.BidID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.ClientBidID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.ProgRptReqs = regexp.MustCompile(`-?\d+`)
    myNewOrderList_RegexValidator.BidType = regexp.MustCompile(`-?\d+`)
    myNewOrderList_RegexValidator.ProgPeriodInterval = regexp.MustCompile(`-?\d+`)
    myNewOrderList_RegexValidator.CancellationRights = regexp.MustCompile(`[^|]`)
    myNewOrderList_RegexValidator.MoneyLaunderingStatus = regexp.MustCompile(`[^|]`)
    myNewOrderList_RegexValidator.RegistID = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.ListExecInstType = regexp.MustCompile(`[^|]`)
    myNewOrderList_RegexValidator.ListExecInst = regexp.MustCompile(`[^|]*`)
    myNewOrderList_RegexValidator.EncodedListExecInstLen = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.EncodedListExecInst = regexp.MustCompile(`.*`)
    myNewOrderList_RegexValidator.AllowableOneSidednessPct = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myNewOrderList_RegexValidator.AllowableOneSidednessValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myNewOrderList_RegexValidator.AllowableOneSidednessCurr = regexp.MustCompile(`[A-Z]{3}`)
    myNewOrderList_RegexValidator.TotNoOrders = regexp.MustCompile(`-?\d+`)
    myNewOrderList_RegexValidator.LastFragment = regexp.MustCompile(`[YN]`)
    myNewOrderList_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myNewOrderList_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myNewOrderList_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *NewOrderList) HasRequiredFields() bool {
    valid := true
    
    

    
    
    valid = valid && m.BeginString.HasValue()
    
    
    
    valid = valid && m.BodyLength.HasValue()
    
    
    
    valid = valid && m.MsgType.HasValue()
    
    
    
    valid = valid && m.SenderCompID.HasValue()
    
    
    
    valid = valid && m.TargetCompID.HasValue()
    
    
    
    valid = valid && m.MsgSeqNum.HasValue()
    
    
    
    valid = valid && m.SendingTime.HasValue()
    
    
    
    for _, g := range(m.Hops) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.ListID.HasValue()
    
    
    
    valid = valid && m.BidType.HasValue()
    
    
    
    valid = valid && m.TotNoOrders.HasValue()
    
    
    
    for _, g := range(m.Orders) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.RootPartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



