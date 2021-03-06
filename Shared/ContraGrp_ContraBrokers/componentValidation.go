package ContraGrp_ContraBrokers

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type contraGrp_ContraBrokers_RegexValidator struct {
    ContraBroker *(regexp.Regexp)
    ContraTrader *(regexp.Regexp)
    ContraTradeQty *(regexp.Regexp)
    ContraTradeTime *(regexp.Regexp)
    ContraLegRefID *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myContraGrp_ContraBrokers_RegexValidator contraGrp_ContraBrokers_RegexValidator

func init() {
    myContraGrp_ContraBrokers_RegexValidator
    myContraGrp_ContraBrokers_RegexValidator.ContraBroker = regexp.MustCompile(`[^|]*`)
    myContraGrp_ContraBrokers_RegexValidator.ContraTrader = regexp.MustCompile(`[^|]*`)
    myContraGrp_ContraBrokers_RegexValidator.ContraTradeQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myContraGrp_ContraBrokers_RegexValidator.ContraTradeTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myContraGrp_ContraBrokers_RegexValidator.ContraLegRefID = regexp.MustCompile(`[^|]*`)
}



func (m *ContraGrp_ContraBrokers) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




