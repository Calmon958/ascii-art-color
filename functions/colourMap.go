package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Color struct {
	R   int
	G   int
	B   int
	Hex string
}

// Maps the string of the required color to it's hexadecimal value from the commandline
func ColourMap(s []string, shade string) (string, string) {
	rgbFile, err := os.Open("resources/text/rgb_hex.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(0)
	}

	defer rgbFile.Close()

	scanner := bufio.NewScanner(rgbFile)

	colors := make(map[string]Color)

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Fields(line)

		if len(parts) < 4 {
			continue
		}

		// get the rgb values and color name
		r, _ := strconv.Atoi(parts[0])
		g, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		colorName := strings.Join(parts[3:], " ")

		// convert RGB to Hex
		hexColor := fmt.Sprintf("#%02x%02x%02x", r, g, b)

		// store the color in the map
		colors[strings.ToLower(colorName)] = Color{R: r, G: g, B: b, Hex: hexColor}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(0)
	}
	// Find the color in t}he map
	var color Color
	var Colored string
	var exists bool
	if IsValidHex(shade) {
		r, g, b := HexToRGB(shade)
		color = Color{R: r, G: g, B: b, Hex: shade}

	} else {
		color, exists = colors[strings.ToLower(shade)]
		if !exists {
			fmt.Println("Color not found")
			os.Exit(0)
		}
	}
	for _, line := range s {
		Colored = fmt.Sprintf("\033[38;2;%d;%d;%dm", color.R, color.G, color.B)
		return Colored, line
		
	}
	return "", ""
}
