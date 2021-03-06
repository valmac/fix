package NestedParties3_Nested3PartyIDs

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import NstdPtys3SubGrp_Nested3PartySubIDs "grgrbll/fix/Shared/NstdPtys3SubGrp_Nested3PartySubIDs"


type NestedParties3_Nested3PartyIDs struct {
    Nested3PartyID Types.String `id:"949", required:"N" `
    Nested3PartyIDSource Types.Char `id:"950", required:"N" `
    Nested3PartyRole Types.Int `id:"951", required:"N" `
    Nested3PartySubIDs []NstdPtys3SubGrp_Nested3PartySubIDs.NstdPtys3SubGrp_Nested3PartySubIDs `counter_name:"NoNested3PartySubIDs", counter_id:"952"`
    _controlBlock fixMessageControlBlock
}
