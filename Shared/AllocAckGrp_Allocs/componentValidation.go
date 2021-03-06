package AllocAckGrp_Allocs

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


type allocAckGrp_Allocs_RegexValidator struct {
    AllocAccount *(regexp.Regexp)
    AllocAcctIDSource *(regexp.Regexp)
    AllocPrice *(regexp.Regexp)
    IndividualAllocID *(regexp.Regexp)
    IndividualAllocRejCode *(regexp.Regexp)
    AllocText *(regexp.Regexp)
    EncodedAllocTextLen *(regexp.Regexp)
    EncodedAllocText *(regexp.Regexp)
    SecondaryIndividualAllocID *(regexp.Regexp)
    AllocCustomerCapacity *(regexp.Regexp)
    IndividualAllocType *(regexp.Regexp)
    AllocQty *(regexp.Regexp)
    AllocPositionEffect *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myAllocAckGrp_Allocs_RegexValidator allocAckGrp_Allocs_RegexValidator

func init() {
    myAllocAckGrp_Allocs_RegexValidator
    myAllocAckGrp_Allocs_RegexValidator.AllocAccount = regexp.MustCompile(`[^|]*`)
    myAllocAckGrp_Allocs_RegexValidator.AllocAcctIDSource = regexp.MustCompile(`-?\d+`)
    myAllocAckGrp_Allocs_RegexValidator.AllocPrice = regexp.MustCompile(``)
    myAllocAckGrp_Allocs_RegexValidator.IndividualAllocID = regexp.MustCompile(`[^|]*`)
    myAllocAckGrp_Allocs_RegexValidator.IndividualAllocRejCode = regexp.MustCompile(`-?\d+`)
    myAllocAckGrp_Allocs_RegexValidator.AllocText = regexp.MustCompile(`[^|]*`)
    myAllocAckGrp_Allocs_RegexValidator.EncodedAllocTextLen = regexp.MustCompile(`\d+`)
    myAllocAckGrp_Allocs_RegexValidator.EncodedAllocText = regexp.MustCompile(`.*`)
    myAllocAckGrp_Allocs_RegexValidator.SecondaryIndividualAllocID = regexp.MustCompile(`[^|]*`)
    myAllocAckGrp_Allocs_RegexValidator.AllocCustomerCapacity = regexp.MustCompile(`[^|]*`)
    myAllocAckGrp_Allocs_RegexValidator.IndividualAllocType = regexp.MustCompile(`-?\d+`)
    myAllocAckGrp_Allocs_RegexValidator.AllocQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocAckGrp_Allocs_RegexValidator.AllocPositionEffect = regexp.MustCompile(`[^|]`)
}



func (m *AllocAckGrp_Allocs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.NestedPartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}




