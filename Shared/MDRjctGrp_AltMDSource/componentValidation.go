package MDRjctGrp_AltMDSource

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type mDRjctGrp_AltMDSource_RegexValidator struct {
    AltMDSourceID *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myMDRjctGrp_AltMDSource_RegexValidator mDRjctGrp_AltMDSource_RegexValidator

func init() {
    myMDRjctGrp_AltMDSource_RegexValidator
    myMDRjctGrp_AltMDSource_RegexValidator.AltMDSourceID = regexp.MustCompile(`[^|]*`)
}



func (m *MDRjctGrp_AltMDSource) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




