package OrderStatusRequest

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"


type orderStatusRequest_RegexValidator struct {
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
    OrderID *(regexp.Regexp)
    ClOrdID *(regexp.Regexp)
    SecondaryClOrdID *(regexp.Regexp)
    ClOrdLinkID *(regexp.Regexp)
    OrdStatusReqID *(regexp.Regexp)
    Account *(regexp.Regexp)
    AcctIDSource *(regexp.Regexp)
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
    AgreementDesc *(regexp.Regexp)
    AgreementID *(regexp.Regexp)
    AgreementDate *(regexp.Regexp)
    AgreementCurrency *(regexp.Regexp)
    TerminationType *(regexp.Regexp)
    StartDate *(regexp.Regexp)
    EndDate *(regexp.Regexp)
    DeliveryType *(regexp.Regexp)
    MarginRatio *(regexp.Regexp)
    Side *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myOrderStatusRequest_RegexValidator orderStatusRequest_RegexValidator

func init() {
    myOrderStatusRequest_RegexValidator
    myOrderStatusRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myOrderStatusRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myOrderStatusRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myOrderStatusRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myOrderStatusRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myOrderStatusRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myOrderStatusRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SecondaryClOrdID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.ClOrdLinkID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.OrdStatusReqID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.AcctIDSource = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myOrderStatusRequest_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myOrderStatusRequest_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myOrderStatusRequest_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myOrderStatusRequest_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myOrderStatusRequest_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myOrderStatusRequest_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myOrderStatusRequest_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myOrderStatusRequest_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myOrderStatusRequest_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myOrderStatusRequest_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myOrderStatusRequest_RegexValidator.AgreementDesc = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.AgreementID = regexp.MustCompile(`[^|]*`)
    myOrderStatusRequest_RegexValidator.AgreementDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.AgreementCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myOrderStatusRequest_RegexValidator.TerminationType = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.StartDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.EndDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderStatusRequest_RegexValidator.DeliveryType = regexp.MustCompile(`-?\d+`)
    myOrderStatusRequest_RegexValidator.MarginRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrderStatusRequest_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myOrderStatusRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myOrderStatusRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myOrderStatusRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *OrderStatusRequest) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.ClOrdID.HasValue()
    
    
    
    for _, g := range(m.PartyIDs) {
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
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.Side.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




