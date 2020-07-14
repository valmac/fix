package SettlInstGrp_SettlInst

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import DlvyInstGrp_DlvyInst "grgrbll/fix/Shared/DlvyInstGrp_DlvyInst"


type SettlInstGrp_SettlInst struct {
    SettlInstID Types.String `id:"162", required:"N" `
    SettlInstTransType Types.Char `id:"163", required:"N" `
    SettlInstRefID Types.String `id:"214", required:"N" `
    PartyIDs []Parties_PartyIDs.Parties_PartyIDs `counter_name:"NoPartyIDs", counter_id:"453"`
    Side Types.Char `id:"54", required:"N" `
    Product Types.Int `id:"460", required:"N" `
    SecurityType Types.String `id:"167", required:"N" `
    CFICode Types.String `id:"461", required:"N" `
    EffectiveTime Types.UTCTimestamp `id:"168", required:"N" `
    ExpireTime Types.UTCTimestamp `id:"126", required:"N" `
    LastUpdateTime Types.UTCTimestamp `id:"779", required:"N" `
    SettlDeliveryType Types.Int `id:"172", required:"N" `
    StandInstDbType Types.Int `id:"169", required:"N" `
    StandInstDbName Types.String `id:"170", required:"N" `
    StandInstDbID Types.String `id:"171", required:"N" `
    DlvyInst []DlvyInstGrp_DlvyInst.DlvyInstGrp_DlvyInst `counter_name:"NoDlvyInst", counter_id:"85"`
    PaymentMethod Types.Int `id:"492", required:"N" `
    PaymentRef Types.String `id:"476", required:"N" `
    CardHolderName Types.String `id:"488", required:"N" `
    CardNumber Types.String `id:"489", required:"N" `
    CardStartDate Types.LocalMktDate `id:"503", required:"N" `
    CardExpDate Types.LocalMktDate `id:"490", required:"N" `
    CardIssNum Types.String `id:"491", required:"N" `
    PaymentDate Types.LocalMktDate `id:"504", required:"N" `
    PaymentRemitterID Types.String `id:"505", required:"N" `
    SettlCurrency Types.Currency `id:"120", required:"N" `
    _controlBlock fixMessageControlBlock
}