package SecurityListUpdateReport

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecLstUpdRelSymGrp_RelatedSym "grgrbll/fix/Shared/SecLstUpdRelSymGrp_RelatedSym"


type securityListUpdateReport_RegexValidator struct {
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
    SecurityReportID *(regexp.Regexp)
    SecurityReqID *(regexp.Regexp)
    SecurityResponseID *(regexp.Regexp)
    SecurityRequestResult *(regexp.Regexp)
    TotNoRelatedSym *(regexp.Regexp)
    ClearingBusinessDate *(regexp.Regexp)
    SecurityUpdateAction *(regexp.Regexp)
    CorporateAction *(regexp.Regexp)
    LastFragment *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var mySecurityListUpdateReport_RegexValidator securityListUpdateReport_RegexValidator

func init() {
    mySecurityListUpdateReport_RegexValidator
    mySecurityListUpdateReport_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    mySecurityListUpdateReport_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    mySecurityListUpdateReport_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    mySecurityListUpdateReport_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    mySecurityListUpdateReport_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    mySecurityListUpdateReport_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    mySecurityListUpdateReport_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityListUpdateReport_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityListUpdateReport_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    mySecurityListUpdateReport_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    mySecurityListUpdateReport_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    mySecurityListUpdateReport_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.SecurityReportID = regexp.MustCompile(`-?\d+`)
    mySecurityListUpdateReport_RegexValidator.SecurityReqID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.SecurityResponseID = regexp.MustCompile(`[^|]*`)
    mySecurityListUpdateReport_RegexValidator.SecurityRequestResult = regexp.MustCompile(`-?\d+`)
    mySecurityListUpdateReport_RegexValidator.TotNoRelatedSym = regexp.MustCompile(`-?\d+`)
    mySecurityListUpdateReport_RegexValidator.ClearingBusinessDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityListUpdateReport_RegexValidator.SecurityUpdateAction = regexp.MustCompile(`[^|]`)
    mySecurityListUpdateReport_RegexValidator.CorporateAction = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    mySecurityListUpdateReport_RegexValidator.LastFragment = regexp.MustCompile(`[YN]`)
    mySecurityListUpdateReport_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    mySecurityListUpdateReport_RegexValidator.Signature = regexp.MustCompile(`.*`)
    mySecurityListUpdateReport_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *SecurityListUpdateReport) HasRequiredFields() bool {
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
    
    
    
    for _, g := range(m.RelatedSym) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



