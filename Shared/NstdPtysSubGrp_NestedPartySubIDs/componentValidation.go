package NstdPtysSubGrp_NestedPartySubIDs

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type nstdPtysSubGrp_NestedPartySubIDs_RegexValidator struct {
    NestedPartySubID *(regexp.Regexp)
    NestedPartySubIDType *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myNstdPtysSubGrp_NestedPartySubIDs_RegexValidator nstdPtysSubGrp_NestedPartySubIDs_RegexValidator

func init() {
    myNstdPtysSubGrp_NestedPartySubIDs_RegexValidator
    myNstdPtysSubGrp_NestedPartySubIDs_RegexValidator.NestedPartySubID = regexp.MustCompile(`[^|]*`)
    myNstdPtysSubGrp_NestedPartySubIDs_RegexValidator.NestedPartySubIDType = regexp.MustCompile(`-?\d+`)
}



func (m *NstdPtysSubGrp_NestedPartySubIDs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




