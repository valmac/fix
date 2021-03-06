package ContraGrp_ContraBrokers

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type ContraGrp_ContraBrokers struct {
    ContraBroker Types.String `id:"375", required:"N" `
    ContraTrader Types.String `id:"337", required:"N" `
    ContraTradeQty Types.Qty `id:"437", required:"N" `
    ContraTradeTime Types.UTCTimestamp `id:"438", required:"N" `
    ContraLegRefID Types.String `id:"655", required:"N" `
    _controlBlock fixMessageControlBlock
}
