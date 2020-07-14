package SideCrossOrdModGrp_Sides

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import PreAllocGrp_Allocs "grgrbll/fix/Shared/PreAllocGrp_Allocs"


func (m *SideCrossOrdModGrp_Sides) MarshalFIX(res *bufio.Writer) error {
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
    if m.ClOrdLinkID.HasValue() {
        (*res).WriteString("583=")
        val, err := m.ClOrdLinkID.MarshalFIX()
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
    
    if m.TradeOriginationDate.HasValue() {
        (*res).WriteString("229=")
        val, err := m.TradeOriginationDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradeDate.HasValue() {
        (*res).WriteString("75=")
        val, err := m.TradeDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
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
    if m.DayBookingInst.HasValue() {
        (*res).WriteString("589=")
        val, err := m.DayBookingInst.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.BookingUnit.HasValue() {
        (*res).WriteString("590=")
        val, err := m.BookingUnit.MarshalFIX()
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
    
    if m.QtyType.HasValue() {
        (*res).WriteString("854=")
        val, err := m.QtyType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderQty.HasValue() {
        (*res).WriteString("38=")
        val, err := m.OrderQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CashOrderQty.HasValue() {
        (*res).WriteString("152=")
        val, err := m.CashOrderQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderPercent.HasValue() {
        (*res).WriteString("516=")
        val, err := m.OrderPercent.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RoundingDirection.HasValue() {
        (*res).WriteString("468=")
        val, err := m.RoundingDirection.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RoundingModulus.HasValue() {
        (*res).WriteString("469=")
        val, err := m.RoundingModulus.MarshalFIX()
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
    if m.ForexReq.HasValue() {
        (*res).WriteString("121=")
        val, err := m.ForexReq.MarshalFIX()
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
    if m.BookingType.HasValue() {
        (*res).WriteString("775=")
        val, err := m.BookingType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Text.HasValue() {
        (*res).WriteString("58=")
        val, err := m.Text.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedTextLen.HasValue() {
        (*res).WriteString("354=")
        val, err := m.EncodedTextLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedText.HasValue() {
        (*res).WriteString("355=")
        val, err := m.EncodedText.MarshalFIX()
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
    if m.CoveredOrUncovered.HasValue() {
        (*res).WriteString("203=")
        val, err := m.CoveredOrUncovered.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CashMargin.HasValue() {
        (*res).WriteString("544=")
        val, err := m.CashMargin.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ClearingFeeIndicator.HasValue() {
        (*res).WriteString("635=")
        val, err := m.ClearingFeeIndicator.MarshalFIX()
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
    if m.SideComplianceID.HasValue() {
        (*res).WriteString("659=")
        val, err := m.SideComplianceID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SideTimeInForce.HasValue() {
        (*res).WriteString("962=")
        val, err := m.SideTimeInForce.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PreTradeAnonymity.HasValue() {
        (*res).WriteString("1091=")
        val, err := m.PreTradeAnonymity.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *SideCrossOrdModGrp_Sides) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message SideCrossOrdModGrp_Sides (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *SideCrossOrdModGrp_Sides) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        
        
        case 583:
            if !m.ClOrdLinkID.HasValue() {
                used = true
                err = m.ClOrdLinkID.UnmarshalFIX(value)
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
        
        
        
        case 229:
            if !m.TradeOriginationDate.HasValue() {
                used = true
                err = m.TradeOriginationDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 75:
            if !m.TradeDate.HasValue() {
                used = true
                err = m.TradeDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
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
        
        
        
        case 589:
            if !m.DayBookingInst.HasValue() {
                used = true
                err = m.DayBookingInst.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 590:
            if !m.BookingUnit.HasValue() {
                used = true
                err = m.BookingUnit.UnmarshalFIX(value)
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
            m.Allocs = make([]PreAllocGrp_Allocs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Allocs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 854:
            if !m.QtyType.HasValue() {
                used = true
                err = m.QtyType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 38:
            if !m.OrderQty.HasValue() {
                used = true
                err = m.OrderQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 152:
            if !m.CashOrderQty.HasValue() {
                used = true
                err = m.CashOrderQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 516:
            if !m.OrderPercent.HasValue() {
                used = true
                err = m.OrderPercent.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 468:
            if !m.RoundingDirection.HasValue() {
                used = true
                err = m.RoundingDirection.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 469:
            if !m.RoundingModulus.HasValue() {
                used = true
                err = m.RoundingModulus.UnmarshalFIX(value)
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
        
        
        
        case 121:
            if !m.ForexReq.HasValue() {
                used = true
                err = m.ForexReq.UnmarshalFIX(value)
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
        
        
        
        case 775:
            if !m.BookingType.HasValue() {
                used = true
                err = m.BookingType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 58:
            if !m.Text.HasValue() {
                used = true
                err = m.Text.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 354:
            if !m.EncodedTextLen.HasValue() {
                used = true
                err = m.EncodedTextLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 355:
            if !m.EncodedText.HasValue() {
                used = true
                err = m.EncodedText.UnmarshalFIX(value)
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
        
        
        
        case 203:
            if !m.CoveredOrUncovered.HasValue() {
                used = true
                err = m.CoveredOrUncovered.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 544:
            if !m.CashMargin.HasValue() {
                used = true
                err = m.CashMargin.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 635:
            if !m.ClearingFeeIndicator.HasValue() {
                used = true
                err = m.ClearingFeeIndicator.UnmarshalFIX(value)
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
        
        
        
        case 659:
            if !m.SideComplianceID.HasValue() {
                used = true
                err = m.SideComplianceID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 962:
            if !m.SideTimeInForce.HasValue() {
                used = true
                err = m.SideTimeInForce.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1091:
            if !m.PreTradeAnonymity.HasValue() {
                used = true
                err = m.PreTradeAnonymity.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        default:
            used = false
        }
    }
   
    return used, err
}

