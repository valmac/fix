package TradeCaptureReport

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import PositionAmountData_PosAmt "grgrbll/fix/Shared/PositionAmountData_PosAmt"

import TrdInstrmtLegGrp_Legs "grgrbll/fix/Shared/TrdInstrmtLegGrp_Legs"

import TrdRegTimestamps_TrdRegTimestamps "grgrbll/fix/Shared/TrdRegTimestamps_TrdRegTimestamps"

import TrdCapRptSideGrp_Sides "grgrbll/fix/Shared/TrdCapRptSideGrp_Sides"

import RootParties_RootPartyIDs "grgrbll/fix/Shared/RootParties_RootPartyIDs"


type TradeCaptureReport struct {
    BeginString Types.String `id:"8", required:"Y" `
    BodyLength Types.Length `id:"9", required:"Y" `
    MsgType Types.String `id:"35", required:"Y" `
    SenderCompID Types.String `id:"49", required:"Y" `
    TargetCompID Types.String `id:"56", required:"Y" `
    OnBehalfOfCompID Types.String `id:"115", required:"N" `
    DeliverToCompID Types.String `id:"128", required:"N" `
    SecureDataLen Types.Length `id:"90", required:"N" `
    SecureData Types.Data `id:"91", required:"N" `
    MsgSeqNum Types.SeqNum `id:"34", required:"Y" `
    SenderSubID Types.String `id:"50", required:"N" `
    SenderLocationID Types.String `id:"142", required:"N" `
    TargetSubID Types.String `id:"57", required:"N" `
    TargetLocationID Types.String `id:"143", required:"N" `
    OnBehalfOfSubID Types.String `id:"116", required:"N" `
    OnBehalfOfLocationID Types.String `id:"144", required:"N" `
    DeliverToSubID Types.String `id:"129", required:"N" `
    DeliverToLocationID Types.String `id:"145", required:"N" `
    PossDupFlag Types.Boolean `id:"43", required:"N" `
    PossResend Types.Boolean `id:"97", required:"N" `
    SendingTime Types.UTCTimestamp `id:"52", required:"Y" `
    OrigSendingTime Types.UTCTimestamp `id:"122", required:"N" `
    XmlDataLen Types.Length `id:"212", required:"N" `
    XmlData Types.Data `id:"213", required:"N" `
    MessageEncoding Types.String `id:"347", required:"N" `
    LastMsgSeqNumProcessed Types.SeqNum `id:"369", required:"N" `
    Hops []HopGrp_Hops.HopGrp_Hops `counter_name:"NoHops", counter_id:"627"`
    ApplVerID Types.String `id:"1128", required:"N" `
    CstmApplVerID Types.String `id:"1129", required:"N" `
    TradeReportID Types.String `id:"571", required:"N" `
    TradeReportTransType Types.Int `id:"487", required:"N" `
    TradeReportType Types.Int `id:"856", required:"N" `
    TradeRequestID Types.String `id:"568", required:"N" `
    TrdType Types.Int `id:"828", required:"N" `
    TrdSubType Types.Int `id:"829", required:"N" `
    SecondaryTrdType Types.Int `id:"855", required:"N" `
    TransferReason Types.String `id:"830", required:"N" `
    ExecType Types.Char `id:"150", required:"N" `
    TotNumTradeReports Types.Int `id:"748", required:"N" `
    LastRptRequested Types.Boolean `id:"912", required:"N" `
    UnsolicitedIndicator Types.Boolean `id:"325", required:"N" `
    SubscriptionRequestType Types.Char `id:"263", required:"N" `
    TradeReportRefID Types.String `id:"572", required:"N" `
    SecondaryTradeReportRefID Types.String `id:"881", required:"N" `
    SecondaryTradeReportID Types.String `id:"818", required:"N" `
    TradeLinkID Types.String `id:"820", required:"N" `
    TrdMatchID Types.String `id:"880", required:"N" `
    ExecID Types.String `id:"17", required:"N" `
    OrdStatus Types.Char `id:"39", required:"N" `
    SecondaryExecID Types.String `id:"527", required:"N" `
    ExecRestatementReason Types.Int `id:"378", required:"N" `
    PreviouslyReported Types.Boolean `id:"570", required:"N" `
    PriceType Types.Int `id:"423", required:"N" `
    Symbol Types.String `id:"55", required:"N" `
    SymbolSfx Types.String `id:"65", required:"N" `
    SecurityID Types.String `id:"48", required:"N" `
    SecurityIDSource Types.String `id:"22", required:"N" `
    SecurityAltID []SecAltIDGrp_SecurityAltID.SecAltIDGrp_SecurityAltID `counter_name:"NoSecurityAltID", counter_id:"454"`
    Product Types.Int `id:"460", required:"N" `
    CFICode Types.String `id:"461", required:"N" `
    SecurityType Types.String `id:"167", required:"N" `
    SecuritySubType Types.String `id:"762", required:"N" `
    MaturityMonthYear Types.MonthYear `id:"200", required:"N" `
    MaturityDate Types.LocalMktDate `id:"541", required:"N" `
    CouponPaymentDate Types.LocalMktDate `id:"224", required:"N" `
    IssueDate Types.LocalMktDate `id:"225", required:"N" `
    RepoCollateralSecurityType Types.Int `id:"239", required:"N" `
    RepurchaseTerm Types.Int `id:"226", required:"N" `
    RepurchaseRate Types.Percentage `id:"227", required:"N" `
    Factor Types.Float `id:"228", required:"N" `
    CreditRating Types.String `id:"255", required:"N" `
    InstrRegistry Types.String `id:"543", required:"N" `
    CountryOfIssue Types.Country `id:"470", required:"N" `
    StateOrProvinceOfIssue Types.String `id:"471", required:"N" `
    LocaleOfIssue Types.String `id:"472", required:"N" `
    RedemptionDate Types.LocalMktDate `id:"240", required:"N" `
    StrikePrice Types.Price `id:"202", required:"N" `
    StrikeCurrency Types.Currency `id:"947", required:"N" `
    OptAttribute Types.Char `id:"206", required:"N" `
    ContractMultiplier Types.Float `id:"231", required:"N" `
    CouponRate Types.Percentage `id:"223", required:"N" `
    SecurityExchange Types.Exchange `id:"207", required:"N" `
    Issuer Types.String `id:"106", required:"N" `
    EncodedIssuerLen Types.Length `id:"348", required:"N" `
    EncodedIssuer Types.Data `id:"349", required:"N" `
    SecurityDesc Types.String `id:"107", required:"N" `
    EncodedSecurityDescLen Types.Length `id:"350", required:"N" `
    EncodedSecurityDesc Types.Data `id:"351", required:"N" `
    Pool Types.String `id:"691", required:"N" `
    ContractSettlMonth Types.MonthYear `id:"667", required:"N" `
    CPProgram Types.Int `id:"875", required:"N" `
    CPRegType Types.String `id:"876", required:"N" `
    Events []EvntGrp_Events.EvntGrp_Events `counter_name:"NoEvents", counter_id:"864"`
    DatedDate Types.LocalMktDate `id:"873", required:"N" `
    InterestAccrualDate Types.LocalMktDate `id:"874", required:"N" `
    SecurityStatus Types.String `id:"965", required:"N" `
    SettleOnOpenFlag Types.String `id:"966", required:"N" `
    InstrmtAssignmentMethod Types.Char `id:"1049", required:"N" `
    StrikeMultiplier Types.Float `id:"967", required:"N" `
    StrikeValue Types.Float `id:"968", required:"N" `
    MinPriceIncrement Types.Float `id:"969", required:"N" `
    PositionLimit Types.Int `id:"970", required:"N" `
    NTPositionLimit Types.Int `id:"971", required:"N" `
    InstrumentParties []InstrumentParties_InstrumentParties.InstrumentParties_InstrumentParties `counter_name:"NoInstrumentParties", counter_id:"1018"`
    UnitofMeasure Types.String `id:"996", required:"N" `
    TimeUnit Types.String `id:"997", required:"N" `
    MaturityTime Types.TZTimeOnly `id:"1079", required:"N" `
    AgreementDesc Types.String `id:"913", required:"N" `
    AgreementID Types.String `id:"914", required:"N" `
    AgreementDate Types.LocalMktDate `id:"915", required:"N" `
    AgreementCurrency Types.Currency `id:"918", required:"N" `
    TerminationType Types.Int `id:"788", required:"N" `
    StartDate Types.LocalMktDate `id:"916", required:"N" `
    EndDate Types.LocalMktDate `id:"917", required:"N" `
    DeliveryType Types.Int `id:"919", required:"N" `
    MarginRatio Types.Percentage `id:"898", required:"N" `
    OrderQty Types.Qty `id:"38", required:"N" `
    CashOrderQty Types.Qty `id:"152", required:"N" `
    OrderPercent Types.Percentage `id:"516", required:"N" `
    RoundingDirection Types.Char `id:"468", required:"N" `
    RoundingModulus Types.Float `id:"469", required:"N" `
    QtyType Types.Int `id:"854", required:"N" `
    YieldType Types.String `id:"235", required:"N" `
    Yield Types.Percentage `id:"236", required:"N" `
    YieldCalcDate Types.LocalMktDate `id:"701", required:"N" `
    YieldRedemptionDate Types.LocalMktDate `id:"696", required:"N" `
    YieldRedemptionPrice Types.Price `id:"697", required:"N" `
    YieldRedemptionPriceType Types.Int `id:"698", required:"N" `
    Underlyings []UndInstrmtGrp_Underlyings.UndInstrmtGrp_Underlyings `counter_name:"NoUnderlyings", counter_id:"711"`
    UnderlyingTradingSessionID Types.String `id:"822", required:"N" `
    UnderlyingTradingSessionSubID Types.String `id:"823", required:"N" `
    LastQty Types.Qty `id:"32", required:"Y" `
    LastPx Types.Price `id:"31", required:"Y" `
    LastParPx Types.Price `id:"669", required:"N" `
    LastSpotRate Types.Price `id:"194", required:"N" `
    LastForwardPoints Types.PriceOffset `id:"195", required:"N" `
    LastMkt Types.Exchange `id:"30", required:"N" `
    TradeDate Types.LocalMktDate `id:"75", required:"Y" `
    ClearingBusinessDate Types.LocalMktDate `id:"715", required:"N" `
    AvgPx Types.Price `id:"6", required:"N" `
    Spread Types.PriceOffset `id:"218", required:"N" `
    BenchmarkCurveCurrency Types.Currency `id:"220", required:"N" `
    BenchmarkCurveName Types.String `id:"221", required:"N" `
    BenchmarkCurvePoint Types.String `id:"222", required:"N" `
    BenchmarkPrice Types.Price `id:"662", required:"N" `
    BenchmarkPriceType Types.Int `id:"663", required:"N" `
    BenchmarkSecurityID Types.String `id:"699", required:"N" `
    BenchmarkSecurityIDSource Types.String `id:"761", required:"N" `
    AvgPxIndicator Types.Int `id:"819", required:"N" `
    PosAmt []PositionAmountData_PosAmt.PositionAmountData_PosAmt `counter_name:"NoPosAmt", counter_id:"753"`
    MultiLegReportingType Types.Char `id:"442", required:"N" `
    TradeLegRefID Types.String `id:"824", required:"N" `
    Legs []TrdInstrmtLegGrp_Legs.TrdInstrmtLegGrp_Legs `counter_name:"NoLegs", counter_id:"555"`
    TransactTime Types.UTCTimestamp `id:"60", required:"N" `
    TrdRegTimestamps []TrdRegTimestamps_TrdRegTimestamps.TrdRegTimestamps_TrdRegTimestamps `counter_name:"NoTrdRegTimestamps", counter_id:"768"`
    SettlType Types.String `id:"63", required:"N" `
    SettlDate Types.LocalMktDate `id:"64", required:"N" `
    MatchStatus Types.Char `id:"573", required:"N" `
    MatchType Types.String `id:"574", required:"N" `
    OrderCategory Types.Char `id:"1115", required:"N" `
    Sides []TrdCapRptSideGrp_Sides.TrdCapRptSideGrp_Sides `counter_name:"NoSides", counter_id:"552"`
    CopyMsgIndicator Types.Boolean `id:"797", required:"N" `
    PublishTrdIndicator Types.Boolean `id:"852", required:"N" `
    ShortSaleReason Types.Int `id:"853", required:"N" `
    TrdRptStatus Types.Int `id:"939", required:"N" `
    AsOfIndicator Types.Char `id:"1015", required:"N" `
    SettlSessID Types.String `id:"716", required:"N" `
    SettlSessSubID Types.String `id:"717", required:"N" `
    TierCode Types.String `id:"994", required:"N" `
    MessageEventSource Types.String `id:"1011", required:"N" `
    LastUpdateTime Types.UTCTimestamp `id:"779", required:"N" `
    RndPx Types.Price `id:"991", required:"N" `
    TradeID Types.String `id:"1003", required:"N" `
    SecondaryTradeID Types.String `id:"1040", required:"N" `
    FirmTradeID Types.String `id:"1041", required:"N" `
    SecondaryFirmTradeID Types.String `id:"1042", required:"N" `
    CalculatedCcyLastQty Types.Qty `id:"1056", required:"N" `
    LastSwapPoints Types.PriceOffset `id:"1071", required:"N" `
    UnderlyingSettlementDate Types.LocalMktDate `id:"987", required:"N" `
    GrossTradeAmt Types.Amt `id:"381", required:"N" `
    RootPartyIDs []RootParties_RootPartyIDs.RootParties_RootPartyIDs `counter_name:"NoRootPartyIDs", counter_id:"1116"`
    TradeHandlingInstr Types.Char `id:"1123", required:"N" `
    OrigTradeHandlingInstr Types.Char `id:"1124", required:"N" `
    OrigTradeDate Types.LocalMktDate `id:"1125", required:"N" `
    OrigTradeID Types.String `id:"1126", required:"N" `
    OrigSecondaryTradeID Types.String `id:"1127", required:"N" `
    TZTransactTime Types.TZTimestamp `id:"1132", required:"N" `
    ReportedPxDiff Types.Boolean `id:"1134", required:"N" `
    SignatureLength Types.Length `id:"93", required:"N" `
    Signature Types.Data `id:"89", required:"N" `
    CheckSum Types.String `id:"10", required:"Y" `
    _controlBlock fixMessageControlBlock
}