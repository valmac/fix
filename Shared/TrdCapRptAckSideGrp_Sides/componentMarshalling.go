package TrdCapRptAckSideGrp_Sides

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import ClrInstGrp_ClearingInstructions "grgrbll/fix/Shared/ClrInstGrp_ClearingInstructions"

import ContAmtGrp_ContAmts "grgrbll/fix/Shared/ContAmtGrp_ContAmts"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import MiscFeesGrp_MiscFees "grgrbll/fix/Shared/MiscFeesGrp_MiscFees"

import TrdAllocGrp_Allocs "grgrbll/fix/Shared/TrdAllocGrp_Allocs"

import SideTrdRegTS_SideTrdRegTS "grgrbll/fix/Shared/SideTrdRegTS_SideTrdRegTS"


func (m *TrdCapRptAckSideGrp_Sides) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.Side.HasValue() {
        (*res).WriteString("54=")
        val, err := m.Side.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderID.HasValue() {
        (*res).WriteString("37=")
        val, err := m.OrderID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecondaryOrderID.HasValue() {
        (*res).WriteString("198=")
        val, err := m.SecondaryOrderID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ClOrdID.HasValue() {
        (*res).WriteString("11=")
        val, err := m.ClOrdID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecondaryClOrdID.HasValue() {
        (*res).WriteString("526=")
        val, err := m.SecondaryClOrdID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ListID.HasValue() {
        (*res).WriteString("66=")
        val, err := m.ListID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.PartyIDs) > 0 {
    
    (*res).WriteString("453=")
    (*res).WriteString(strconv.Itoa(len(m.PartyIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.PartyIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.Account.HasValue() {
        (*res).WriteString("1=")
        val, err := m.Account.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AcctIDSource.HasValue() {
        (*res).WriteString("660=")
        val, err := m.AcctIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AccountType.HasValue() {
        (*res).WriteString("581=")
        val, err := m.AccountType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ProcessCode.HasValue() {
        (*res).WriteString("81=")
        val, err := m.ProcessCode.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OddLot.HasValue() {
        (*res).WriteString("575=")
        val, err := m.OddLot.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.ClearingInstructions) > 0 {
    
    (*res).WriteString("576=")
    (*res).WriteString(strconv.Itoa(len(m.ClearingInstructions)))
    (*res).WriteString("\x01")
    for _, g := range m.ClearingInstructions {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.TradeInputSource.HasValue() {
        (*res).WriteString("578=")
        val, err := m.TradeInputSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradeInputDevice.HasValue() {
        (*res).WriteString("579=")
        val, err := m.TradeInputDevice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderInputDevice.HasValue() {
        (*res).WriteString("821=")
        val, err := m.OrderInputDevice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Currency.HasValue() {
        (*res).WriteString("15=")
        val, err := m.Currency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ComplianceID.HasValue() {
        (*res).WriteString("376=")
        val, err := m.ComplianceID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SolicitedFlag.HasValue() {
        (*res).WriteString("377=")
        val, err := m.SolicitedFlag.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderCapacity.HasValue() {
        (*res).WriteString("528=")
        val, err := m.OrderCapacity.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderRestrictions.HasValue() {
        (*res).WriteString("529=")
        val, err := m.OrderRestrictions.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CustOrderCapacity.HasValue() {
        (*res).WriteString("582=")
        val, err := m.CustOrderCapacity.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrdType.HasValue() {
        (*res).WriteString("40=")
        val, err := m.OrdType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ExecInst.HasValue() {
        (*res).WriteString("18=")
        val, err := m.ExecInst.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TransBkdTime.HasValue() {
        (*res).WriteString("483=")
        val, err := m.TransBkdTime.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradingSessionID.HasValue() {
        (*res).WriteString("336=")
        val, err := m.TradingSessionID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradingSessionSubID.HasValue() {
        (*res).WriteString("625=")
        val, err := m.TradingSessionSubID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TimeBracket.HasValue() {
        (*res).WriteString("943=")
        val, err := m.TimeBracket.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Commission.HasValue() {
        (*res).WriteString("12=")
        val, err := m.Commission.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CommType.HasValue() {
        (*res).WriteString("13=")
        val, err := m.CommType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CommCurrency.HasValue() {
        (*res).WriteString("479=")
        val, err := m.CommCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.FundRenewWaiv.HasValue() {
        (*res).WriteString("497=")
        val, err := m.FundRenewWaiv.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.NumDaysInterest.HasValue() {
        (*res).WriteString("157=")
        val, err := m.NumDaysInterest.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ExDate.HasValue() {
        (*res).WriteString("230=")
        val, err := m.ExDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AccruedInterestRate.HasValue() {
        (*res).WriteString("158=")
        val, err := m.AccruedInterestRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AccruedInterestAmt.HasValue() {
        (*res).WriteString("159=")
        val, err := m.AccruedInterestAmt.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.InterestAtMaturity.HasValue() {
        (*res).WriteString("738=")
        val, err := m.InterestAtMaturity.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EndAccruedInterestAmt.HasValue() {
        (*res).WriteString("920=")
        val, err := m.EndAccruedInterestAmt.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.StartCash.HasValue() {
        (*res).WriteString("921=")
        val, err := m.StartCash.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EndCash.HasValue() {
        (*res).WriteString("922=")
        val, err := m.EndCash.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Concession.HasValue() {
        (*res).WriteString("238=")
        val, err := m.Concession.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TotalTakedown.HasValue() {
        (*res).WriteString("237=")
        val, err := m.TotalTakedown.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.NetMoney.HasValue() {
        (*res).WriteString("118=")
        val, err := m.NetMoney.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlCurrAmt.HasValue() {
        (*res).WriteString("119=")
        val, err := m.SettlCurrAmt.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlCurrency.HasValue() {
        (*res).WriteString("120=")
        val, err := m.SettlCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlCurrFxRate.HasValue() {
        (*res).WriteString("155=")
        val, err := m.SettlCurrFxRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlCurrFxRateCalc.HasValue() {
        (*res).WriteString("156=")
        val, err := m.SettlCurrFxRateCalc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PositionEffect.HasValue() {
        (*res).WriteString("77=")
        val, err := m.PositionEffect.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideMultiLegReportingType.HasValue() {
        (*res).WriteString("752=")
        val, err := m.SideMultiLegReportingType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.ContAmts) > 0 {
    
    (*res).WriteString("518=")
    (*res).WriteString(strconv.Itoa(len(m.ContAmts)))
    (*res).WriteString("\x01")
    for _, g := range m.ContAmts {
        err = g.MarshalFIX(res)
    }
    }
    
    if len(m.Stipulations) > 0 {
    
    (*res).WriteString("232=")
    (*res).WriteString(strconv.Itoa(len(m.Stipulations)))
    (*res).WriteString("\x01")
    for _, g := range m.Stipulations {
        err = g.MarshalFIX(res)
    }
    }
    
    if len(m.MiscFees) > 0 {
    
    (*res).WriteString("136=")
    (*res).WriteString(strconv.Itoa(len(m.MiscFees)))
    (*res).WriteString("\x01")
    for _, g := range m.MiscFees {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.ExchangeRule.HasValue() {
        (*res).WriteString("825=")
        val, err := m.ExchangeRule.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradeAllocIndicator.HasValue() {
        (*res).WriteString("826=")
        val, err := m.TradeAllocIndicator.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PreallocMethod.HasValue() {
        (*res).WriteString("591=")
        val, err := m.PreallocMethod.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AllocID.HasValue() {
        (*res).WriteString("70=")
        val, err := m.AllocID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.Allocs) > 0 {
    
    (*res).WriteString("78=")
    (*res).WriteString(strconv.Itoa(len(m.Allocs)))
    (*res).WriteString("\x01")
    for _, g := range m.Allocs {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.LotType.HasValue() {
        (*res).WriteString("1093=")
        val, err := m.LotType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideGrossTradeAmt.HasValue() {
        (*res).WriteString("1072=")
        val, err := m.SideGrossTradeAmt.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AggressorIndicator.HasValue() {
        (*res).WriteString("1057=")
        val, err := m.AggressorIndicator.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideQty.HasValue() {
        (*res).WriteString("1009=")
        val, err := m.SideQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideTradeReportID.HasValue() {
        (*res).WriteString("1005=")
        val, err := m.SideTradeReportID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideFillStationCd.HasValue() {
        (*res).WriteString("1006=")
        val, err := m.SideFillStationCd.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideReasonCd.HasValue() {
        (*res).WriteString("1007=")
        val, err := m.SideReasonCd.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RptSeq.HasValue() {
        (*res).WriteString("83=")
        val, err := m.RptSeq.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideTrdSubTyp.HasValue() {
        (*res).WriteString("1008=")
        val, err := m.SideTrdSubTyp.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.SideTrdRegTS) > 0 {
    
    (*res).WriteString("1016=")
    (*res).WriteString(strconv.Itoa(len(m.SideTrdRegTS)))
    (*res).WriteString("\x01")
    for _, g := range m.SideTrdRegTS {
        err = g.MarshalFIX(res)
    }
    }
    
    return err
}

func (m *TrdCapRptAckSideGrp_Sides) UnmarshalFIX(input io.Reader) error {
    var err error
    var field []byte
    for field, err = input.ReadSlice([]byte("\x01")); err == nil {
        slices = bytes.Split(field, []byte("="))
        if len(slices) != 2 {
            // TODO handle data fields
            err = errors.New(fmt.Sprintf("Found field without seperator '=' (%s)", string(field))
        } else {
            var used bool
            used, err m.UnmarshalFieldFIX(strconv.Atoi(string(slices[0])), slices[1])
            if !used {
                err = errors.New(fmt.Sprintf("Field unused by message TrdCapRptAckSideGrp_Sides (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *TrdCapRptAckSideGrp_Sides) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
    // Check if we are currently populating a group
    used := false
    var err error = nil
    for !used && m._controlBlock.mostRecentRepeatingGroup != nil {
        used = m._controlBlock.mostRecentRepeatingGroup[m._controlBlock.mostRecentRepeatingGroupCounter].PopulateNextFieldById(id, value)
        if !used {
            // This group did not need use the latest field, and has all its mandatory fields, so consider it complete.
            m._controlBlock.mostRecentRepeatingGroupCounter++
            if m._controlBlock.mostRecentRepeatingGroupCounter >= len(m._controlBlock.mostRecentRepeatingGroup)  {
                // we have all the repeated groups we expected.
                m._controlBlock.mostRecentRepeatingGroup = nil
            }
        }
    }
    
    // If one of the groups did not consume the KV, try our own fields 
    if !used {
        switch id {
        
        case 54:
            if !m.Side.HasValue() {
                used = true
                err = m.Side.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 37:
            if !m.OrderID.HasValue() {
                used = true
                err = m.OrderID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 198:
            if !m.SecondaryOrderID.HasValue() {
                used = true
                err = m.SecondaryOrderID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 11:
            if !m.ClOrdID.HasValue() {
                used = true
                err = m.ClOrdID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 526:
            if !m.SecondaryClOrdID.HasValue() {
                used = true
                err = m.SecondaryClOrdID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 66:
            if !m.ListID.HasValue() {
                used = true
                err = m.ListID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 453:
            // This counter (NoPartyIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.PartyIDs = make([]Parties_PartyIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.PartyIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 1:
            if !m.Account.HasValue() {
                used = true
                err = m.Account.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 660:
            if !m.AcctIDSource.HasValue() {
                used = true
                err = m.AcctIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 581:
            if !m.AccountType.HasValue() {
                used = true
                err = m.AccountType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 81:
            if !m.ProcessCode.HasValue() {
                used = true
                err = m.ProcessCode.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 575:
            if !m.OddLot.HasValue() {
                used = true
                err = m.OddLot.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 576:
            // This counter (NoClearingInstructions) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.ClearingInstructions = make([]ClrInstGrp_ClearingInstructions, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.ClearingInstructions[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 578:
            if !m.TradeInputSource.HasValue() {
                used = true
                err = m.TradeInputSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 579:
            if !m.TradeInputDevice.HasValue() {
                used = true
                err = m.TradeInputDevice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 821:
            if !m.OrderInputDevice.HasValue() {
                used = true
                err = m.OrderInputDevice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 15:
            if !m.Currency.HasValue() {
                used = true
                err = m.Currency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 376:
            if !m.ComplianceID.HasValue() {
                used = true
                err = m.ComplianceID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 377:
            if !m.SolicitedFlag.HasValue() {
                used = true
                err = m.SolicitedFlag.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 528:
            if !m.OrderCapacity.HasValue() {
                used = true
                err = m.OrderCapacity.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 529:
            if !m.OrderRestrictions.HasValue() {
                used = true
                err = m.OrderRestrictions.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 582:
            if !m.CustOrderCapacity.HasValue() {
                used = true
                err = m.CustOrderCapacity.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 40:
            if !m.OrdType.HasValue() {
                used = true
                err = m.OrdType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 18:
            if !m.ExecInst.HasValue() {
                used = true
                err = m.ExecInst.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 483:
            if !m.TransBkdTime.HasValue() {
                used = true
                err = m.TransBkdTime.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 336:
            if !m.TradingSessionID.HasValue() {
                used = true
                err = m.TradingSessionID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 625:
            if !m.TradingSessionSubID.HasValue() {
                used = true
                err = m.TradingSessionSubID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 943:
            if !m.TimeBracket.HasValue() {
                used = true
                err = m.TimeBracket.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 12:
            if !m.Commission.HasValue() {
                used = true
                err = m.Commission.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 13:
            if !m.CommType.HasValue() {
                used = true
                err = m.CommType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 479:
            if !m.CommCurrency.HasValue() {
                used = true
                err = m.CommCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 497:
            if !m.FundRenewWaiv.HasValue() {
                used = true
                err = m.FundRenewWaiv.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 157:
            if !m.NumDaysInterest.HasValue() {
                used = true
                err = m.NumDaysInterest.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 230:
            if !m.ExDate.HasValue() {
                used = true
                err = m.ExDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 158:
            if !m.AccruedInterestRate.HasValue() {
                used = true
                err = m.AccruedInterestRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 159:
            if !m.AccruedInterestAmt.HasValue() {
                used = true
                err = m.AccruedInterestAmt.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 738:
            if !m.InterestAtMaturity.HasValue() {
                used = true
                err = m.InterestAtMaturity.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 920:
            if !m.EndAccruedInterestAmt.HasValue() {
                used = true
                err = m.EndAccruedInterestAmt.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 921:
            if !m.StartCash.HasValue() {
                used = true
                err = m.StartCash.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 922:
            if !m.EndCash.HasValue() {
                used = true
                err = m.EndCash.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 238:
            if !m.Concession.HasValue() {
                used = true
                err = m.Concession.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 237:
            if !m.TotalTakedown.HasValue() {
                used = true
                err = m.TotalTakedown.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 118:
            if !m.NetMoney.HasValue() {
                used = true
                err = m.NetMoney.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 119:
            if !m.SettlCurrAmt.HasValue() {
                used = true
                err = m.SettlCurrAmt.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 120:
            if !m.SettlCurrency.HasValue() {
                used = true
                err = m.SettlCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 155:
            if !m.SettlCurrFxRate.HasValue() {
                used = true
                err = m.SettlCurrFxRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 156:
            if !m.SettlCurrFxRateCalc.HasValue() {
                used = true
                err = m.SettlCurrFxRateCalc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 77:
            if !m.PositionEffect.HasValue() {
                used = true
                err = m.PositionEffect.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 752:
            if !m.SideMultiLegReportingType.HasValue() {
                used = true
                err = m.SideMultiLegReportingType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 518:
            // This counter (NoContAmts) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.ContAmts = make([]ContAmtGrp_ContAmts, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.ContAmts[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 232:
            // This counter (NoStipulations) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Stipulations = make([]Stipulations_Stipulations, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Stipulations[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 136:
            // This counter (NoMiscFees) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.MiscFees = make([]MiscFeesGrp_MiscFees, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.MiscFees[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 825:
            if !m.ExchangeRule.HasValue() {
                used = true
                err = m.ExchangeRule.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 826:
            if !m.TradeAllocIndicator.HasValue() {
                used = true
                err = m.TradeAllocIndicator.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 591:
            if !m.PreallocMethod.HasValue() {
                used = true
                err = m.PreallocMethod.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 70:
            if !m.AllocID.HasValue() {
                used = true
                err = m.AllocID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 78:
            // This counter (NoAllocs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Allocs = make([]TrdAllocGrp_Allocs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Allocs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 1093:
            if !m.LotType.HasValue() {
                used = true
                err = m.LotType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1072:
            if !m.SideGrossTradeAmt.HasValue() {
                used = true
                err = m.SideGrossTradeAmt.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1057:
            if !m.AggressorIndicator.HasValue() {
                used = true
                err = m.AggressorIndicator.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1009:
            if !m.SideQty.HasValue() {
                used = true
                err = m.SideQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1005:
            if !m.SideTradeReportID.HasValue() {
                used = true
                err = m.SideTradeReportID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1006:
            if !m.SideFillStationCd.HasValue() {
                used = true
                err = m.SideFillStationCd.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1007:
            if !m.SideReasonCd.HasValue() {
                used = true
                err = m.SideReasonCd.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 83:
            if !m.RptSeq.HasValue() {
                used = true
                err = m.RptSeq.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1008:
            if !m.SideTrdSubTyp.HasValue() {
                used = true
                err = m.SideTrdSubTyp.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1016:
            // This counter (NoSideTrdRegTS) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.SideTrdRegTS = make([]SideTrdRegTS_SideTrdRegTS, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.SideTrdRegTS[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        default:
            used = false
        }
    }
   
    return used, err
}


