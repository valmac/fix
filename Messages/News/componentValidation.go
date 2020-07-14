package News

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import RoutingGrp_RoutingIDs "grgrbll/fix/Shared/RoutingGrp_RoutingIDs"

import InstrmtGrp_RelatedSym "grgrbll/fix/Shared/InstrmtGrp_RelatedSym"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import LinesOfTextGrp_LinesOfText "grgrbll/fix/Shared/LinesOfTextGrp_LinesOfText"


type news_RegexValidator struct {
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
    OrigTime *(regexp.Regexp)
    Urgency *(regexp.Regexp)
    Headline *(regexp.Regexp)
    EncodedHeadlineLen *(regexp.Regexp)
    EncodedHeadline *(regexp.Regexp)
    URLLink *(regexp.Regexp)
    RawDataLength *(regexp.Regexp)
    RawData *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myNews_RegexValidator news_RegexValidator

func init() {
    myNews_RegexValidator
    myNews_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myNews_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myNews_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myNews_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myNews_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myNews_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myNews_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.OrigTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myNews_RegexValidator.Urgency = regexp.MustCompile(`[^|]`)
    myNews_RegexValidator.Headline = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.EncodedHeadlineLen = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.EncodedHeadline = regexp.MustCompile(`.*`)
    myNews_RegexValidator.URLLink = regexp.MustCompile(`[^|]*`)
    myNews_RegexValidator.RawDataLength = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.RawData = regexp.MustCompile(`.*`)
    myNews_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myNews_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myNews_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *News) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.Headline.HasValue()
    
    
    
    for _, g := range(m.RoutingIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.RelatedSym) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.LinesOfText) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



