package SettlementInstructions

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SettlInstGrp_SettlInst "grgrbll/fix/Shared/SettlInstGrp_SettlInst"


func (m *SettlementInstructions) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.BeginString.HasValue() {
        (*res).WriteString("8=")
        val, err := m.BeginString.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.BodyLength.HasValue() {
        (*res).WriteString("9=")
        val, err := m.BodyLength.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MsgType.HasValue() {
        (*res).WriteString("35=")
        val, err := m.MsgType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SenderCompID.HasValue() {
        (*res).WriteString("49=")
        val, err := m.SenderCompID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TargetCompID.HasValue() {
        (*res).WriteString("56=")
        val, err := m.TargetCompID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OnBehalfOfCompID.HasValue() {
        (*res).WriteString("115=")
        val, err := m.OnBehalfOfCompID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.DeliverToCompID.HasValue() {
        (*res).WriteString("128=")
        val, err := m.DeliverToCompID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecureDataLen.HasValue() {
        (*res).WriteString("90=")
        val, err := m.SecureDataLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecureData.HasValue() {
        (*res).WriteString("91=")
        val, err := m.SecureData.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MsgSeqNum.HasValue() {
        (*res).WriteString("34=")
        val, err := m.MsgSeqNum.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SenderSubID.HasValue() {
        (*res).WriteString("50=")
        val, err := m.SenderSubID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SenderLocationID.HasValue() {
        (*res).WriteString("142=")
        val, err := m.SenderLocationID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TargetSubID.HasValue() {
        (*res).WriteString("57=")
        val, err := m.TargetSubID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TargetLocationID.HasValue() {
        (*res).WriteString("143=")
        val, err := m.TargetLocationID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OnBehalfOfSubID.HasValue() {
        (*res).WriteString("116=")
        val, err := m.OnBehalfOfSubID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OnBehalfOfLocationID.HasValue() {
        (*res).WriteString("144=")
        val, err := m.OnBehalfOfLocationID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.DeliverToSubID.HasValue() {
        (*res).WriteString("129=")
        val, err := m.DeliverToSubID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.DeliverToLocationID.HasValue() {
        (*res).WriteString("145=")
        val, err := m.DeliverToLocationID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PossDupFlag.HasValue() {
        (*res).WriteString("43=")
        val, err := m.PossDupFlag.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PossResend.HasValue() {
        (*res).WriteString("97=")
        val, err := m.PossResend.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SendingTime.HasValue() {
        (*res).WriteString("52=")
        val, err := m.SendingTime.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrigSendingTime.HasValue() {
        (*res).WriteString("122=")
        val, err := m.OrigSendingTime.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.XmlDataLen.HasValue() {
        (*res).WriteString("212=")
        val, err := m.XmlDataLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.XmlData.HasValue() {
        (*res).WriteString("213=")
        val, err := m.XmlData.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MessageEncoding.HasValue() {
        (*res).WriteString("347=")
        val, err := m.MessageEncoding.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LastMsgSeqNumProcessed.HasValue() {
        (*res).WriteString("369=")
        val, err := m.LastMsgSeqNumProcessed.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.Hops) > 0 {
    
    (*res).WriteString("627=")
    (*res).WriteString(strconv.Itoa(len(m.Hops)))
    (*res).WriteString("\x01")
    for _, g := range m.Hops {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.ApplVerID.HasValue() {
        (*res).WriteString("1128=")
        val, err := m.ApplVerID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CstmApplVerID.HasValue() {
        (*res).WriteString("1129=")
        val, err := m.CstmApplVerID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlInstMsgID.HasValue() {
        (*res).WriteString("777=")
        val, err := m.SettlInstMsgID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlInstReqID.HasValue() {
        (*res).WriteString("791=")
        val, err := m.SettlInstReqID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlInstMode.HasValue() {
        (*res).WriteString("160=")
        val, err := m.SettlInstMode.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlInstReqRejCode.HasValue() {
        (*res).WriteString("792=")
        val, err := m.SettlInstReqRejCode.MarshalFIX()
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
    if m.ClOrdID.HasValue() {
        (*res).WriteString("11=")
        val, err := m.ClOrdID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TransactTime.HasValue() {
        (*res).WriteString("60=")
        val, err := m.TransactTime.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.SettlInst) > 0 {
    
    (*res).WriteString("778=")
    (*res).WriteString(strconv.Itoa(len(m.SettlInst)))
    (*res).WriteString("\x01")
    for _, g := range m.SettlInst {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.SignatureLength.HasValue() {
        (*res).WriteString("93=")
        val, err := m.SignatureLength.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Signature.HasValue() {
        (*res).WriteString("89=")
        val, err := m.Signature.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CheckSum.HasValue() {
        (*res).WriteString("10=")
        val, err := m.CheckSum.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *SettlementInstructions) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message SettlementInstructions (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *SettlementInstructions) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 8:
            if !m.BeginString.HasValue() {
                used = true
                err = m.BeginString.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 9:
            if !m.BodyLength.HasValue() {
                used = true
                err = m.BodyLength.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 35:
            if !m.MsgType.HasValue() {
                used = true
                err = m.MsgType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 49:
            if !m.SenderCompID.HasValue() {
                used = true
                err = m.SenderCompID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 56:
            if !m.TargetCompID.HasValue() {
                used = true
                err = m.TargetCompID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 115:
            if !m.OnBehalfOfCompID.HasValue() {
                used = true
                err = m.OnBehalfOfCompID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 128:
            if !m.DeliverToCompID.HasValue() {
                used = true
                err = m.DeliverToCompID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 90:
            if !m.SecureDataLen.HasValue() {
                used = true
                err = m.SecureDataLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 91:
            if !m.SecureData.HasValue() {
                used = true
                err = m.SecureData.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 34:
            if !m.MsgSeqNum.HasValue() {
                used = true
                err = m.MsgSeqNum.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 50:
            if !m.SenderSubID.HasValue() {
                used = true
                err = m.SenderSubID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 142:
            if !m.SenderLocationID.HasValue() {
                used = true
                err = m.SenderLocationID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 57:
            if !m.TargetSubID.HasValue() {
                used = true
                err = m.TargetSubID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 143:
            if !m.TargetLocationID.HasValue() {
                used = true
                err = m.TargetLocationID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 116:
            if !m.OnBehalfOfSubID.HasValue() {
                used = true
                err = m.OnBehalfOfSubID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 144:
            if !m.OnBehalfOfLocationID.HasValue() {
                used = true
                err = m.OnBehalfOfLocationID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 129:
            if !m.DeliverToSubID.HasValue() {
                used = true
                err = m.DeliverToSubID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 145:
            if !m.DeliverToLocationID.HasValue() {
                used = true
                err = m.DeliverToLocationID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 43:
            if !m.PossDupFlag.HasValue() {
                used = true
                err = m.PossDupFlag.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 97:
            if !m.PossResend.HasValue() {
                used = true
                err = m.PossResend.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 52:
            if !m.SendingTime.HasValue() {
                used = true
                err = m.SendingTime.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 122:
            if !m.OrigSendingTime.HasValue() {
                used = true
                err = m.OrigSendingTime.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 212:
            if !m.XmlDataLen.HasValue() {
                used = true
                err = m.XmlDataLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 213:
            if !m.XmlData.HasValue() {
                used = true
                err = m.XmlData.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 347:
            if !m.MessageEncoding.HasValue() {
                used = true
                err = m.MessageEncoding.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 369:
            if !m.LastMsgSeqNumProcessed.HasValue() {
                used = true
                err = m.LastMsgSeqNumProcessed.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 627:
            // This counter (NoHops) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Hops = make([]HopGrp_Hops, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Hops[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 1128:
            if !m.ApplVerID.HasValue() {
                used = true
                err = m.ApplVerID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1129:
            if !m.CstmApplVerID.HasValue() {
                used = true
                err = m.CstmApplVerID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 777:
            if !m.SettlInstMsgID.HasValue() {
                used = true
                err = m.SettlInstMsgID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 791:
            if !m.SettlInstReqID.HasValue() {
                used = true
                err = m.SettlInstReqID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 160:
            if !m.SettlInstMode.HasValue() {
                used = true
                err = m.SettlInstMode.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 792:
            if !m.SettlInstReqRejCode.HasValue() {
                used = true
                err = m.SettlInstReqRejCode.UnmarshalFIX(value)
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
        
        
        
        case 11:
            if !m.ClOrdID.HasValue() {
                used = true
                err = m.ClOrdID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 60:
            if !m.TransactTime.HasValue() {
                used = true
                err = m.TransactTime.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 778:
            // This counter (NoSettlInst) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.SettlInst = make([]SettlInstGrp_SettlInst, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.SettlInst[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 93:
            if !m.SignatureLength.HasValue() {
                used = true
                err = m.SignatureLength.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 89:
            if !m.Signature.HasValue() {
                used = true
                err = m.Signature.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 10:
            if !m.CheckSum.HasValue() {
                used = true
                err = m.CheckSum.UnmarshalFIX(value)
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


