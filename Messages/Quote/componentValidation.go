package Quote

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import QuotQualGrp_QuoteQualifiers "grgrbll/fix/Shared/QuotQualGrp_QuoteQualifiers"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import LegQuotGrp_Legs "grgrbll/fix/Shared/LegQuotGrp_Legs"


type quote_RegexValidator struct {
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
    QuoteReqID *(regexp.Regexp)
    QuoteID *(regexp.Regexp)
    QuoteRespID *(regexp.Regexp)
    QuoteType *(regexp.Regexp)
    QuoteResponseLevel *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
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
    OrderQty *(regexp.Regexp)
    CashOrderQty *(regexp.Regexp)
    OrderPercent *(regexp.Regexp)
    RoundingDirection *(regexp.Regexp)
    RoundingModulus *(regexp.Regexp)
    SettlType *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    SettlDate2 *(regexp.Regexp)
    OrderQty2 *(regexp.Regexp)
    Currency *(regexp.Regexp)
    Account *(regexp.Regexp)
    AcctIDSource *(regexp.Regexp)
    AccountType *(regexp.Regexp)
    BidPx *(regexp.Regexp)
    OfferPx *(regexp.Regexp)
    MktBidPx *(regexp.Regexp)
    MktOfferPx *(regexp.Regexp)
    MinBidSize *(regexp.Regexp)
    BidSize *(regexp.Regexp)
    MinOfferSize *(regexp.Regexp)
    OfferSize *(regexp.Regexp)
    ValidUntilTime *(regexp.Regexp)
    BidSpotRate *(regexp.Regexp)
    OfferSpotRate *(regexp.Regexp)
    BidForwardPoints *(regexp.Regexp)
    OfferForwardPoints *(regexp.Regexp)
    MidPx *(regexp.Regexp)
    BidYield *(regexp.Regexp)
    MidYield *(regexp.Regexp)
    OfferYield *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    OrdType *(regexp.Regexp)
    BidForwardPoints2 *(regexp.Regexp)
    OfferForwardPoints2 *(regexp.Regexp)
    SettlCurrBidFxRate *(regexp.Regexp)
    SettlCurrOfferFxRate *(regexp.Regexp)
    SettlCurrFxRateCalc *(regexp.Regexp)
    CommType *(regexp.Regexp)
    Commission *(regexp.Regexp)
    CustOrderCapacity *(regexp.Regexp)
    ExDestination *(regexp.Regexp)
    OrderCapacity *(regexp.Regexp)
    PriceType *(regexp.Regexp)
    Spread *(regexp.Regexp)
    BenchmarkCurveCurrency *(regexp.Regexp)
    BenchmarkCurveName *(regexp.Regexp)
    BenchmarkCurvePoint *(regexp.Regexp)
    BenchmarkPrice *(regexp.Regexp)
    BenchmarkPriceType *(regexp.Regexp)
    BenchmarkSecurityID *(regexp.Regexp)
    BenchmarkSecurityIDSource *(regexp.Regexp)
    YieldType *(regexp.Regexp)
    Yield *(regexp.Regexp)
    YieldCalcDate *(regexp.Regexp)
    YieldRedemptionDate *(regexp.Regexp)
    YieldRedemptionPrice *(regexp.Regexp)
    YieldRedemptionPriceType *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    BidSwapPoints *(regexp.Regexp)
    OfferSwapPoints *(regexp.Regexp)
    ExDestinationIDSource *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myQuote_RegexValidator quote_RegexValidator

func init() {
    myQuote_RegexValidator
    myQuote_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myQuote_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myQuote_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myQuote_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myQuote_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myQuote_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myQuote_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.QuoteReqID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.QuoteID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.QuoteRespID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.QuoteType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.QuoteResponseLevel = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myQuote_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myQuote_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myQuote_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuote_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myQuote_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myQuote_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myQuote_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myQuote_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myQuote_RegexValidator.AgreementDesc = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.AgreementID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.AgreementDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.AgreementCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuote_RegexValidator.TerminationType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.StartDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.EndDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.DeliveryType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.MarginRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.OrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.CashOrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.OrderPercent = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.RoundingDirection = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.RoundingModulus = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.SettlType = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.SettlDate2 = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.OrderQty2 = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myQuote_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.AcctIDSource = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.AccountType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.BidPx = regexp.MustCompile(``)
    myQuote_RegexValidator.OfferPx = regexp.MustCompile(``)
    myQuote_RegexValidator.MktBidPx = regexp.MustCompile(``)
    myQuote_RegexValidator.MktOfferPx = regexp.MustCompile(``)
    myQuote_RegexValidator.MinBidSize = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.BidSize = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.MinOfferSize = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.OfferSize = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.ValidUntilTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myQuote_RegexValidator.BidSpotRate = regexp.MustCompile(``)
    myQuote_RegexValidator.OfferSpotRate = regexp.MustCompile(``)
    myQuote_RegexValidator.BidForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.OfferForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.MidPx = regexp.MustCompile(``)
    myQuote_RegexValidator.BidYield = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.MidYield = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.OfferYield = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myQuote_RegexValidator.OrdType = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.BidForwardPoints2 = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.OfferForwardPoints2 = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.SettlCurrBidFxRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.SettlCurrOfferFxRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.SettlCurrFxRateCalc = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.CommType = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.Commission = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.CustOrderCapacity = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.ExDestination = regexp.MustCompile(`[A-Z0-9]{4}`)
    myQuote_RegexValidator.OrderCapacity = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.PriceType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.Spread = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.BenchmarkCurveCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuote_RegexValidator.BenchmarkCurveName = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.BenchmarkCurvePoint = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.BenchmarkPrice = regexp.MustCompile(``)
    myQuote_RegexValidator.BenchmarkPriceType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.BenchmarkSecurityID = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.BenchmarkSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.YieldType = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.Yield = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.YieldCalcDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.YieldRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myQuote_RegexValidator.YieldRedemptionPrice = regexp.MustCompile(``)
    myQuote_RegexValidator.YieldRedemptionPriceType = regexp.MustCompile(`-?\d+`)
    myQuote_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myQuote_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myQuote_RegexValidator.BidSwapPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.OfferSwapPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuote_RegexValidator.ExDestinationIDSource = regexp.MustCompile(`[^|]`)
    myQuote_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myQuote_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myQuote_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *Quote) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.QuoteID.HasValue()
    
    
    
    for _, g := range(m.QuoteQualifiers) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
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
    
    
    
    for _, g := range(m.Stipulations) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



