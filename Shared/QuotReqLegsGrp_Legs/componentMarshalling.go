package QuotReqLegsGrp_Legs

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import LegSecAltIDGrp_LegSecurityAltID "grgrbll/fix/Shared/LegSecAltIDGrp_LegSecurityAltID"

import LegStipulations_LegStipulations "grgrbll/fix/Shared/LegStipulations_LegStipulations"

import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


func (m *QuotReqLegsGrp_Legs) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.LegSymbol.HasValue() {
        (*res).WriteString("600=")
        val, err := m.LegSymbol.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSymbolSfx.HasValue() {
        (*res).WriteString("601=")
        val, err := m.LegSymbolSfx.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSecurityID.HasValue() {
        (*res).WriteString("602=")
        val, err := m.LegSecurityID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSecurityIDSource.HasValue() {
        (*res).WriteString("603=")
        val, err := m.LegSecurityIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.LegSecurityAltID) > 0 {
    
    (*res).WriteString("604=")
    (*res).WriteString(strconv.Itoa(len(m.LegSecurityAltID)))
    (*res).WriteString("\x01")
    for _, g := range m.LegSecurityAltID {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.LegProduct.HasValue() {
        (*res).WriteString("607=")
        val, err := m.LegProduct.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegCFICode.HasValue() {
        (*res).WriteString("608=")
        val, err := m.LegCFICode.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSecurityType.HasValue() {
        (*res).WriteString("609=")
        val, err := m.LegSecurityType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSecuritySubType.HasValue() {
        (*res).WriteString("764=")
        val, err := m.LegSecuritySubType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegMaturityMonthYear.HasValue() {
        (*res).WriteString("610=")
        val, err := m.LegMaturityMonthYear.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegMaturityDate.HasValue() {
        (*res).WriteString("611=")
        val, err := m.LegMaturityDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegCouponPaymentDate.HasValue() {
        (*res).WriteString("248=")
        val, err := m.LegCouponPaymentDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegIssueDate.HasValue() {
        (*res).WriteString("249=")
        val, err := m.LegIssueDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegRepoCollateralSecurityType.HasValue() {
        (*res).WriteString("250=")
        val, err := m.LegRepoCollateralSecurityType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegRepurchaseTerm.HasValue() {
        (*res).WriteString("251=")
        val, err := m.LegRepurchaseTerm.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegRepurchaseRate.HasValue() {
        (*res).WriteString("252=")
        val, err := m.LegRepurchaseRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegFactor.HasValue() {
        (*res).WriteString("253=")
        val, err := m.LegFactor.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegCreditRating.HasValue() {
        (*res).WriteString("257=")
        val, err := m.LegCreditRating.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegInstrRegistry.HasValue() {
        (*res).WriteString("599=")
        val, err := m.LegInstrRegistry.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegCountryOfIssue.HasValue() {
        (*res).WriteString("596=")
        val, err := m.LegCountryOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegStateOrProvinceOfIssue.HasValue() {
        (*res).WriteString("597=")
        val, err := m.LegStateOrProvinceOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegLocaleOfIssue.HasValue() {
        (*res).WriteString("598=")
        val, err := m.LegLocaleOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegRedemptionDate.HasValue() {
        (*res).WriteString("254=")
        val, err := m.LegRedemptionDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegStrikePrice.HasValue() {
        (*res).WriteString("612=")
        val, err := m.LegStrikePrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegStrikeCurrency.HasValue() {
        (*res).WriteString("942=")
        val, err := m.LegStrikeCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegOptAttribute.HasValue() {
        (*res).WriteString("613=")
        val, err := m.LegOptAttribute.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegContractMultiplier.HasValue() {
        (*res).WriteString("614=")
        val, err := m.LegContractMultiplier.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegCouponRate.HasValue() {
        (*res).WriteString("615=")
        val, err := m.LegCouponRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSecurityExchange.HasValue() {
        (*res).WriteString("616=")
        val, err := m.LegSecurityExchange.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegIssuer.HasValue() {
        (*res).WriteString("617=")
        val, err := m.LegIssuer.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedLegIssuerLen.HasValue() {
        (*res).WriteString("618=")
        val, err := m.EncodedLegIssuerLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedLegIssuer.HasValue() {
        (*res).WriteString("619=")
        val, err := m.EncodedLegIssuer.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSecurityDesc.HasValue() {
        (*res).WriteString("620=")
        val, err := m.LegSecurityDesc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedLegSecurityDescLen.HasValue() {
        (*res).WriteString("621=")
        val, err := m.EncodedLegSecurityDescLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedLegSecurityDesc.HasValue() {
        (*res).WriteString("622=")
        val, err := m.EncodedLegSecurityDesc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegRatioQty.HasValue() {
        (*res).WriteString("623=")
        val, err := m.LegRatioQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSide.HasValue() {
        (*res).WriteString("624=")
        val, err := m.LegSide.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegCurrency.HasValue() {
        (*res).WriteString("556=")
        val, err := m.LegCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegPool.HasValue() {
        (*res).WriteString("740=")
        val, err := m.LegPool.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegDatedDate.HasValue() {
        (*res).WriteString("739=")
        val, err := m.LegDatedDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegContractSettlMonth.HasValue() {
        (*res).WriteString("955=")
        val, err := m.LegContractSettlMonth.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegInterestAccrualDate.HasValue() {
        (*res).WriteString("956=")
        val, err := m.LegInterestAccrualDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegUnitofMeasure.HasValue() {
        (*res).WriteString("999=")
        val, err := m.LegUnitofMeasure.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegTimeUnit.HasValue() {
        (*res).WriteString("1001=")
        val, err := m.LegTimeUnit.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegOptionRatio.HasValue() {
        (*res).WriteString("1017=")
        val, err := m.LegOptionRatio.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegPrice.HasValue() {
        (*res).WriteString("566=")
        val, err := m.LegPrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegQty.HasValue() {
        (*res).WriteString("687=")
        val, err := m.LegQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSwapType.HasValue() {
        (*res).WriteString("690=")
        val, err := m.LegSwapType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSettlType.HasValue() {
        (*res).WriteString("587=")
        val, err := m.LegSettlType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSettlDate.HasValue() {
        (*res).WriteString("588=")
        val, err := m.LegSettlDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.LegStipulations) > 0 {
    
    (*res).WriteString("683=")
    (*res).WriteString(strconv.Itoa(len(m.LegStipulations)))
    (*res).WriteString("\x01")
    for _, g := range m.LegStipulations {
        err = g.MarshalFIX(res)
    }
    }
    
    if len(m.NestedPartyIDs) > 0 {
    
    (*res).WriteString("539=")
    (*res).WriteString(strconv.Itoa(len(m.NestedPartyIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.NestedPartyIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.LegBenchmarkCurveCurrency.HasValue() {
        (*res).WriteString("676=")
        val, err := m.LegBenchmarkCurveCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegBenchmarkCurveName.HasValue() {
        (*res).WriteString("677=")
        val, err := m.LegBenchmarkCurveName.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegBenchmarkCurvePoint.HasValue() {
        (*res).WriteString("678=")
        val, err := m.LegBenchmarkCurvePoint.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegBenchmarkPrice.HasValue() {
        (*res).WriteString("679=")
        val, err := m.LegBenchmarkPrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegBenchmarkPriceType.HasValue() {
        (*res).WriteString("680=")
        val, err := m.LegBenchmarkPriceType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegOrderQty.HasValue() {
        (*res).WriteString("685=")
        val, err := m.LegOrderQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegRefID.HasValue() {
        (*res).WriteString("654=")
        val, err := m.LegRefID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *QuotReqLegsGrp_Legs) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message QuotReqLegsGrp_Legs (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *QuotReqLegsGrp_Legs) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 600:
            if !m.LegSymbol.HasValue() {
                used = true
                err = m.LegSymbol.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 601:
            if !m.LegSymbolSfx.HasValue() {
                used = true
                err = m.LegSymbolSfx.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 602:
            if !m.LegSecurityID.HasValue() {
                used = true
                err = m.LegSecurityID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 603:
            if !m.LegSecurityIDSource.HasValue() {
                used = true
                err = m.LegSecurityIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 604:
            // This counter (NoLegSecurityAltID) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.LegSecurityAltID = make([]LegSecAltIDGrp_LegSecurityAltID, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.LegSecurityAltID[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 607:
            if !m.LegProduct.HasValue() {
                used = true
                err = m.LegProduct.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 608:
            if !m.LegCFICode.HasValue() {
                used = true
                err = m.LegCFICode.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 609:
            if !m.LegSecurityType.HasValue() {
                used = true
                err = m.LegSecurityType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 764:
            if !m.LegSecuritySubType.HasValue() {
                used = true
                err = m.LegSecuritySubType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 610:
            if !m.LegMaturityMonthYear.HasValue() {
                used = true
                err = m.LegMaturityMonthYear.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 611:
            if !m.LegMaturityDate.HasValue() {
                used = true
                err = m.LegMaturityDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 248:
            if !m.LegCouponPaymentDate.HasValue() {
                used = true
                err = m.LegCouponPaymentDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 249:
            if !m.LegIssueDate.HasValue() {
                used = true
                err = m.LegIssueDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 250:
            if !m.LegRepoCollateralSecurityType.HasValue() {
                used = true
                err = m.LegRepoCollateralSecurityType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 251:
            if !m.LegRepurchaseTerm.HasValue() {
                used = true
                err = m.LegRepurchaseTerm.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 252:
            if !m.LegRepurchaseRate.HasValue() {
                used = true
                err = m.LegRepurchaseRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 253:
            if !m.LegFactor.HasValue() {
                used = true
                err = m.LegFactor.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 257:
            if !m.LegCreditRating.HasValue() {
                used = true
                err = m.LegCreditRating.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 599:
            if !m.LegInstrRegistry.HasValue() {
                used = true
                err = m.LegInstrRegistry.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 596:
            if !m.LegCountryOfIssue.HasValue() {
                used = true
                err = m.LegCountryOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 597:
            if !m.LegStateOrProvinceOfIssue.HasValue() {
                used = true
                err = m.LegStateOrProvinceOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 598:
            if !m.LegLocaleOfIssue.HasValue() {
                used = true
                err = m.LegLocaleOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 254:
            if !m.LegRedemptionDate.HasValue() {
                used = true
                err = m.LegRedemptionDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 612:
            if !m.LegStrikePrice.HasValue() {
                used = true
                err = m.LegStrikePrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 942:
            if !m.LegStrikeCurrency.HasValue() {
                used = true
                err = m.LegStrikeCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 613:
            if !m.LegOptAttribute.HasValue() {
                used = true
                err = m.LegOptAttribute.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 614:
            if !m.LegContractMultiplier.HasValue() {
                used = true
                err = m.LegContractMultiplier.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 615:
            if !m.LegCouponRate.HasValue() {
                used = true
                err = m.LegCouponRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 616:
            if !m.LegSecurityExchange.HasValue() {
                used = true
                err = m.LegSecurityExchange.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 617:
            if !m.LegIssuer.HasValue() {
                used = true
                err = m.LegIssuer.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 618:
            if !m.EncodedLegIssuerLen.HasValue() {
                used = true
                err = m.EncodedLegIssuerLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 619:
            if !m.EncodedLegIssuer.HasValue() {
                used = true
                err = m.EncodedLegIssuer.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 620:
            if !m.LegSecurityDesc.HasValue() {
                used = true
                err = m.LegSecurityDesc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 621:
            if !m.EncodedLegSecurityDescLen.HasValue() {
                used = true
                err = m.EncodedLegSecurityDescLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 622:
            if !m.EncodedLegSecurityDesc.HasValue() {
                used = true
                err = m.EncodedLegSecurityDesc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 623:
            if !m.LegRatioQty.HasValue() {
                used = true
                err = m.LegRatioQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 624:
            if !m.LegSide.HasValue() {
                used = true
                err = m.LegSide.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 556:
            if !m.LegCurrency.HasValue() {
                used = true
                err = m.LegCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 740:
            if !m.LegPool.HasValue() {
                used = true
                err = m.LegPool.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 739:
            if !m.LegDatedDate.HasValue() {
                used = true
                err = m.LegDatedDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 955:
            if !m.LegContractSettlMonth.HasValue() {
                used = true
                err = m.LegContractSettlMonth.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 956:
            if !m.LegInterestAccrualDate.HasValue() {
                used = true
                err = m.LegInterestAccrualDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 999:
            if !m.LegUnitofMeasure.HasValue() {
                used = true
                err = m.LegUnitofMeasure.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1001:
            if !m.LegTimeUnit.HasValue() {
                used = true
                err = m.LegTimeUnit.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1017:
            if !m.LegOptionRatio.HasValue() {
                used = true
                err = m.LegOptionRatio.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 566:
            if !m.LegPrice.HasValue() {
                used = true
                err = m.LegPrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 687:
            if !m.LegQty.HasValue() {
                used = true
                err = m.LegQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 690:
            if !m.LegSwapType.HasValue() {
                used = true
                err = m.LegSwapType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 587:
            if !m.LegSettlType.HasValue() {
                used = true
                err = m.LegSettlType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 588:
            if !m.LegSettlDate.HasValue() {
                used = true
                err = m.LegSettlDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 683:
            // This counter (NoLegStipulations) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.LegStipulations = make([]LegStipulations_LegStipulations, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.LegStipulations[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 539:
            // This counter (NoNestedPartyIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.NestedPartyIDs = make([]NestedParties_NestedPartyIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.NestedPartyIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 676:
            if !m.LegBenchmarkCurveCurrency.HasValue() {
                used = true
                err = m.LegBenchmarkCurveCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 677:
            if !m.LegBenchmarkCurveName.HasValue() {
                used = true
                err = m.LegBenchmarkCurveName.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 678:
            if !m.LegBenchmarkCurvePoint.HasValue() {
                used = true
                err = m.LegBenchmarkCurvePoint.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 679:
            if !m.LegBenchmarkPrice.HasValue() {
                used = true
                err = m.LegBenchmarkPrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 680:
            if !m.LegBenchmarkPriceType.HasValue() {
                used = true
                err = m.LegBenchmarkPriceType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 685:
            if !m.LegOrderQty.HasValue() {
                used = true
                err = m.LegOrderQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 654:
            if !m.LegRefID.HasValue() {
                used = true
                err = m.LegRefID.UnmarshalFIX(value)
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

