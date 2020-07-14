package TrdSessLstGrp_TradingSessions

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type TrdSessLstGrp_TradingSessions struct {
    TradingSessionID Types.String `id:"336", required:"Y" `
    TradingSessionSubID Types.String `id:"625", required:"N" `
    SecurityExchange Types.Exchange `id:"207", required:"N" `
    TradSesMethod Types.Int `id:"338", required:"N" `
    TradSesMode Types.Int `id:"339", required:"N" `
    UnsolicitedIndicator Types.Boolean `id:"325", required:"N" `
    TradSesStatus Types.Int `id:"340", required:"Y" `
    TradSesStatusRejReason Types.Int `id:"567", required:"N" `
    TradSesStartTime Types.UTCTimestamp `id:"341", required:"N" `
    TradSesOpenTime Types.UTCTimestamp `id:"342", required:"N" `
    TradSesPreCloseTime Types.UTCTimestamp `id:"343", required:"N" `
    TradSesCloseTime Types.UTCTimestamp `id:"344", required:"N" `
    TradSesEndTime Types.UTCTimestamp `id:"345", required:"N" `
    TotalVolumeTraded Types.Qty `id:"387", required:"N" `
    Text Types.String `id:"58", required:"N" `
    EncodedTextLen Types.Length `id:"354", required:"N" `
    EncodedText Types.Data `id:"355", required:"N" `
    _controlBlock fixMessageControlBlock
}