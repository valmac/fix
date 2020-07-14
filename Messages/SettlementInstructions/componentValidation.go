package SettlementInstructions

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SettlInstGrp_SettlInst "grgrbll/fix/Shared/SettlInstGrp_SettlInst"


type settlementInstructions_RegexValidator struct {
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
    SettlInstMsgID *(regexp.Regexp)
    SettlInstReqID *(regexp.Regexp)
    SettlInstMode *(regexp.Regexp)
    SettlInstReqRejCode *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    ClOrdID *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var mySettlementInstructions_RegexValidator settlementInstructions_RegexValidator

func init() {
    mySettlementInstructions_RegexValidator
    mySettlementInstructions_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    mySettlementInstructions_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    mySettlementInstructions_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    mySettlementInstructions_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructions_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructions_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    mySettlementInstructions_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.SettlInstMsgID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.SettlInstReqID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.SettlInstMode = regexp.MustCompile(`[^|]`)
    mySettlementInstructions_RegexValidator.SettlInstReqRejCode = regexp.MustCompile(`-?\d+`)
    mySettlementInstructions_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    mySettlementInstructions_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructions_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructions_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    mySettlementInstructions_RegexValidator.Signature = regexp.MustCompile(`.*`)
    mySettlementInstructions_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *SettlementInstructions) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.SettlInstMsgID.HasValue()
    
    
    
    valid = valid && m.SettlInstMode.HasValue()
    
    
    
    valid = valid && m.TransactTime.HasValue()
    
    
    
    for _, g := range(m.SettlInst) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



