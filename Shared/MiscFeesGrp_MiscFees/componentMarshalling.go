package MiscFeesGrp_MiscFees

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit




func (m *MiscFeesGrp_MiscFees) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.MiscFeeAmt.HasValue() {
        (*res).WriteString("137=")
        val, err := m.MiscFeeAmt.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MiscFeeCurr.HasValue() {
        (*res).WriteString("138=")
        val, err := m.MiscFeeCurr.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MiscFeeType.HasValue() {
        (*res).WriteString("139=")
        val, err := m.MiscFeeType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MiscFeeBasis.HasValue() {
        (*res).WriteString("891=")
        val, err := m.MiscFeeBasis.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *MiscFeesGrp_MiscFees) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message MiscFeesGrp_MiscFees (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *MiscFeesGrp_MiscFees) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 137:
            if !m.MiscFeeAmt.HasValue() {
                used = true
                err = m.MiscFeeAmt.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 138:
            if !m.MiscFeeCurr.HasValue() {
                used = true
                err = m.MiscFeeCurr.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 139:
            if !m.MiscFeeType.HasValue() {
                used = true
                err = m.MiscFeeType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 891:
            if !m.MiscFeeBasis.HasValue() {
                used = true
                err = m.MiscFeeBasis.UnmarshalFIX(value)
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

