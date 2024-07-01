package ascii

import (
	"strconv"
)

func IsValidHex(hex string) bool {
	if len(hex) != 7 || hex[0] != '#' {
		return false
	}
	_, err := strconv.ParseInt(hex[1:], 16, 64)
	return err == nil
}
