package UnderlyingAmount_UnderlyingAmounts

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type UnderlyingAmount_UnderlyingAmounts struct {
    UnderlyingPayAmount Types.Amt `id:"985", required:"N" `
    UnderlyingCollectAmount Types.Amt `id:"986", required:"N" `
    UnderlyingSettlementDate Types.LocalMktDate `id:"987", required:"N" `
    UnderlyingSettlementStatus Types.String `id:"988", required:"N" `
    _controlBlock fixMessageControlBlock
}
