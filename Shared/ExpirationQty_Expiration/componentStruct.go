package ExpirationQty_Expiration

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type ExpirationQty_Expiration struct {
    ExpType Types.Int `id:"982", required:"N" `
    ExpQty Types.Qty `id:"983", required:"N" `
    _controlBlock fixMessageControlBlock
}
