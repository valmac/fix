package AdjustedPositionReport

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import PositionQty_Positions "grgrbll/fix/Shared/PositionQty_Positions"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"


type adjustedPositionReport_RegexValidator struct {
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
    PosMaintRptID *(regexp.Regexp)
    PosReqType *(regexp.Regexp)
    ClearingBusinessDate *(regexp.Regexp)
    SettlSessID *(regexp.Regexp)
    Symbol *(regexp.Regexp)
    SymbolSfx *(regexp.Regexp)
    SecurityID *(regexp.Regexp)
    SecurityIDSource *(regexp.Regexp)
    Product *(regexp.Regexp)
    CFICode *(regexp.Regexp)
    SecurityType *(regexp.Regexp)
    SecuritySubType *(regexp.Regexp)
    MaturityMonthYear *(regexp.Regexp)
    MaturityDate *(regexp.Regexp)
    CouponPaymentDate *(regexp.Regexp)
    IssueDate *(regexp.Regexp)
    RepoCollateralSecurityType *(regexp.Regexp)
    RepurchaseTerm *(regexp.Regexp)
    RepurchaseRate *(regexp.Regexp)
    Factor *(regexp.Regexp)
    CreditRating *(regexp.Regexp)
    InstrRegistry *(regexp.Regexp)
    CountryOfIssue *(regexp.Regexp)
    StateOrProvinceOfIssue *(regexp.Regexp)
    LocaleOfIssue *(regexp.Regexp)
    RedemptionDate *(regexp.Regexp)
    StrikePrice *(regexp.Regexp)
    StrikeCurrency *(regexp.Regexp)
    OptAttribute *(regexp.Regexp)
    ContractMultiplier *(regexp.Regexp)
    CouponRate *(regexp.Regexp)
    SecurityExchange *(regexp.Regexp)
    Issuer *(regexp.Regexp)
    EncodedIssuerLen *(regexp.Regexp)
    EncodedIssuer *(regexp.Regexp)
    SecurityDesc *(regexp.Regexp)
    EncodedSecurityDescLen *(regexp.Regexp)
    EncodedSecurityDesc *(regexp.Regexp)
    Pool *(regexp.Regexp)
    ContractSettlMonth *(regexp.Regexp)
    CPProgram *(regexp.Regexp)
    CPRegType *(regexp.Regexp)
    DatedDate *(regexp.Regexp)
    InterestAccrualDate *(regexp.Regexp)
    SecurityStatus *(regexp.Regexp)
    SettleOnOpenFlag *(regexp.Regexp)
    InstrmtAssignmentMethod *(regexp.Regexp)
    StrikeMultiplier *(regexp.Regexp)
    StrikeValue *(regexp.Regexp)
    MinPriceIncrement *(regexp.Regexp)
    PositionLimit *(regexp.Regexp)
    NTPositionLimit *(regexp.Regexp)
    UnitofMeasure *(regexp.Regexp)
    TimeUnit *(regexp.Regexp)
    MaturityTime *(regexp.Regexp)
    SettlPrice *(regexp.Regexp)
    PriorSettlPrice *(regexp.Regexp)
    PosMaintRptRefID *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myAdjustedPositionReport_RegexValidator adjustedPositionReport_RegexValidator

func init() {
    myAdjustedPositionReport_RegexValidator
    myAdjustedPositionReport_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myAdjustedPositionReport_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myAdjustedPositionReport_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myAdjustedPositionReport_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAdjustedPositionReport_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAdjustedPositionReport_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myAdjustedPositionReport_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.PosMaintRptID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.PosReqType = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.ClearingBusinessDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.SettlSessID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myAdjustedPositionReport_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myAdjustedPositionReport_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myAdjustedPositionReport_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myAdjustedPositionReport_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myAdjustedPositionReport_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myAdjustedPositionReport_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myAdjustedPositionReport_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myAdjustedPositionReport_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myAdjustedPositionReport_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myAdjustedPositionReport_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myAdjustedPositionReport_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAdjustedPositionReport_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myAdjustedPositionReport_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myAdjustedPositionReport_RegexValidator.SettlPrice = regexp.MustCompile(``)
    myAdjustedPositionReport_RegexValidator.PriorSettlPrice = regexp.MustCompile(``)
    myAdjustedPositionReport_RegexValidator.PosMaintRptRefID = regexp.MustCompile(`[^|]*`)
    myAdjustedPositionReport_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myAdjustedPositionReport_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myAdjustedPositionReport_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *AdjustedPositionReport) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.PosMaintRptID.HasValue()
    
    
    
    valid = valid && m.ClearingBusinessDate.HasValue()
    
    
    
    for _, g := range(m.PartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Positions) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.SecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Events) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



