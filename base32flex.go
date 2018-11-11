package base32flex

import (
	"encoding/base32"
)

const encodeLower = "abcdefghijkmnpqrstuvwxyz23456789"
const encodeUpper = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// LowerEncoding is more readable base32 encoding as it is used with lower characters.
var LowerEncoding = base32.NewEncoding(encodeLower)

// UpperEncoding is more readable base32 encoding as it is used with upper characters.
var UpperEncoding = base32.NewEncoding(encodeUpper)
