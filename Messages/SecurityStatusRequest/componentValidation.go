package SecurityStatusRequest

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import AttrbGrp_InstrAttrib "grgrbll/fix/Shared/AttrbGrp_InstrAttrib"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"


type securityStatusRequest_RegexValidator struct {
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
    SecurityStatusReqID *(regexp.Regexp)
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
    DeliveryForm *(regexp.Regexp)
    PctAtRisk *(regexp.Regexp)
    Currency *(regexp.Regexp)
    SubscriptionRequestType *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var mySecurityStatusRequest_RegexValidator securityStatusRequest_RegexValidator

func init() {
    mySecurityStatusRequest_RegexValidator
    mySecurityStatusRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    mySecurityStatusRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    mySecurityStatusRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    mySecurityStatusRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityStatusRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityStatusRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    mySecurityStatusRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SecurityStatusReqID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    mySecurityStatusRequest_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatusRequest_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatusRequest_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatusRequest_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    mySecurityStatusRequest_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatusRequest_RegexValidator.StrikePrice = regexp.MustCompile(``)
    mySecurityStatusRequest_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    mySecurityStatusRequest_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    mySecurityStatusRequest_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    mySecurityStatusRequest_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    mySecurityStatusRequest_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    mySecurityStatusRequest_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    mySecurityStatusRequest_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatusRequest_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatusRequest_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    mySecurityStatusRequest_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    mySecurityStatusRequest_RegexValidator.DeliveryForm = regexp.MustCompile(`-?\d+`)
    mySecurityStatusRequest_RegexValidator.PctAtRisk = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatusRequest_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    mySecurityStatusRequest_RegexValidator.SubscriptionRequestType = regexp.MustCompile(`[^|]`)
    mySecurityStatusRequest_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatusRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    mySecurityStatusRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    mySecurityStatusRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *SecurityStatusRequest) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.SecurityStatusReqID.HasValue()
    
    
    
    for _, g := range(m.SecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Events) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrAttrib) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.SubscriptionRequestType.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



