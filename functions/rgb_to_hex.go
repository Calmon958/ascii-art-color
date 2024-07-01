package ascii

import "strconv"



func HexToRGB(hex string) (int, int, int) {
	r, _ := strconv.ParseInt(hex[1:3], 16, 0)
	g, _ := strconv.ParseInt(hex[3:5], 16, 0)
	b, _ := strconv.ParseInt(hex[5:], 16, 0)
	return int(r), int(g), int(b)
}
