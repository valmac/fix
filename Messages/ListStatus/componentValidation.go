package ListStatus

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import OrdListStatGrp_Orders "grgrbll/fix/Shared/OrdListStatGrp_Orders"


type listStatus_RegexValidator struct {
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
    ListStatusType *(regexp.Regexp)
    NoRpts *(regexp.Regexp)
    ListOrderStatus *(regexp.Regexp)
    RptSeq *(regexp.Regexp)
    ListStatusText *(regexp.Regexp)
    EncodedListStatusTextLen *(regexp.Regexp)
    EncodedListStatusText *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    TotNoOrders *(regexp.Regexp)
    LastFragment *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myListStatus_RegexValidator listStatus_RegexValidator

func init() {
    myListStatus_RegexValidator
    myListStatus_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myListStatus_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myListStatus_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myListStatus_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myListStatus_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myListStatus_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myListStatus_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.ListID = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.ListStatusType = regexp.MustCompile(`-?\d+`)
    myListStatus_RegexValidator.NoRpts = regexp.MustCompile(`-?\d+`)
    myListStatus_RegexValidator.ListOrderStatus = regexp.MustCompile(`-?\d+`)
    myListStatus_RegexValidator.RptSeq = regexp.MustCompile(`-?\d+`)
    myListStatus_RegexValidator.ListStatusText = regexp.MustCompile(`[^|]*`)
    myListStatus_RegexValidator.EncodedListStatusTextLen = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.EncodedListStatusText = regexp.MustCompile(`.*`)
    myListStatus_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myListStatus_RegexValidator.TotNoOrders = regexp.MustCompile(`-?\d+`)
    myListStatus_RegexValidator.LastFragment = regexp.MustCompile(`[YN]`)
    myListStatus_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myListStatus_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myListStatus_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *ListStatus) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.ListStatusType.HasValue()
    
    
    
    valid = valid && m.NoRpts.HasValue()
    
    
    
    valid = valid && m.ListOrderStatus.HasValue()
    
    
    
    valid = valid && m.RptSeq.HasValue()
    
    
    
    valid = valid && m.TotNoOrders.HasValue()
    
    
    
    for _, g := range(m.Orders) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



