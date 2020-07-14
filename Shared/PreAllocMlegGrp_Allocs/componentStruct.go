package PreAllocMlegGrp_Allocs

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import NestedParties3_Nested3PartyIDs "grgrbll/fix/Shared/NestedParties3_Nested3PartyIDs"


type PreAllocMlegGrp_Allocs struct {
    AllocAccount Types.String `id:"79", required:"N" `
    AllocAcctIDSource Types.Int `id:"661", required:"N" `
    AllocSettlCurrency Types.Currency `id:"736", required:"N" `
    IndividualAllocID Types.String `id:"467", required:"N" `
    Nested3PartyIDs []NestedParties3_Nested3PartyIDs.NestedParties3_Nested3PartyIDs `counter_name:"NoNested3PartyIDs", counter_id:"948"`
    AllocQty Types.Qty `id:"80", required:"N" `
    _controlBlock fixMessageControlBlock
}