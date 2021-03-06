package Parties_PartyIDs

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import PtysSubGrp_PartySubIDs "grgrbll/fix/Shared/PtysSubGrp_PartySubIDs"


func (m *Parties_PartyIDs) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.PartyID.HasValue() {
        (*res).WriteString("448=")
        val, err := m.PartyID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PartyIDSource.HasValue() {
        (*res).WriteString("447=")
        val, err := m.PartyIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PartyRole.HasValue() {
        (*res).WriteString("452=")
        val, err := m.PartyRole.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.PartySubIDs) > 0 {
    
    (*res).WriteString("802=")
    (*res).WriteString(strconv.Itoa(len(m.PartySubIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.PartySubIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    return err
}

func (m *Parties_PartyIDs) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message Parties_PartyIDs (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *Parties_PartyIDs) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 448:
            if !m.PartyID.HasValue() {
                used = true
                err = m.PartyID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 447:
            if !m.PartyIDSource.HasValue() {
                used = true
                err = m.PartyIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 452:
            if !m.PartyRole.HasValue() {
                used = true
                err = m.PartyRole.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 802:
            // This counter (NoPartySubIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.PartySubIDs = make([]PtysSubGrp_PartySubIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.PartySubIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        default:
            used = false
        }
    }
   
    return used, err
}


