package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	A "ascii/functions"
)

func main() {
	// Get command-line arguments
	arg := os.Args[1:]
	var colorName string
	var bannerFile string
	flag.StringVar(&colorName, "color", "", "color str")
	flag.StringVar(&bannerFile, "banner", "", "banner file type")
	flag.Parse()

	// Check if there are any arguments
	if len(arg) < 1 {
		fmt.Println("Enter message to be printed")
		return
	}
	// If only one argument is provided, use the default ASCII art

	if len(arg) == 1 {

		data, err := os.ReadFile("resources/text/standard.txt")
		// confirms if the file used is available
		if err != nil {
			fmt.Println("File not found")
			return
		}

		mapped := A.AsciiArt(string(data))
		tab := A.Tab(os.Args[1])
		word := A.Paragraph(tab, mapped)
		fmt.Print(strings.Join(word, ""))
		return
	}

	if len(arg) == 2 {
		if colorName != "" {
			A.Checkfiles("resources/text/standard.txt")
			data, err := os.ReadFile("resources/text/standard.txt")
			if err != nil {
				fmt.Println("File not found")
				return
			}

			mapped := A.AsciiArt(string(data))
			tab := A.Tab(arg[1])
			shade := strings.ToLower(colorName)
			colored := A.Paragraph(tab, mapped)
			Colored, line := A.ColourMap(colored, shade)
			fmt.Printf("%s%s\033[0m", Colored, line)
		} else {
			words := []string{"shadow", "standard", "thinkertoy"}
			var file string
			for _, char := range words {
				file = arg[1]
				if char == file {
					continue
				}
			}
			switch file {
			case "standard":
				file = "resources/text/standard.txt"
			case "shadow":
				file = "resources/text/shadow.txt"
			case "thinkertoy":
				file = "resources/text/thinkertoy.txt"
			default:
				file = "resources/text/standard.txt"
			}
			A.Checkfiles(file)
			data, err := os.ReadFile(file)
			// confirms if the file used is available
			if err != nil {
				// confirms if all the lines on the ascii art file are there

				fmt.Println("File not found")
				return
			}

			mapped := A.AsciiArt(string(data))
			tab := A.Tab(os.Args[1])
			word := A.Paragraph(tab, mapped)
			for _, char := range word {
				fmt.Print(char)
			}
			return
		}
	}

	if len(arg) == 3 {
		if colorName != "" {
			if bannerFile != "" {
				banners := []string{"standard", "thinkertoy", "shadow"}
				for i := range banners {
					if bannerFile != banners[i] && i == len(banners)-1 {
						fmt.Println("format for banner is : --banner=<shadow/standard/thinkertoy>")
						return
					} else if bannerFile == banners[i] {
						break
					}
				}
				switch bannerFile {

				case "standard":
					bannerFile = "resources/text/standard.txt"
				case "shadow":
					bannerFile = "resources/text/shadow.txt"
				case "thinkertoy":
					bannerFile = "resources/text/thinkertoy.txt"
				default:
					bannerFile = "resources/text/standard.txt"

				}

				A.Checkfiles(bannerFile)
				data, err := os.ReadFile(bannerFile)
				if err != nil {
					fmt.Println("File not found")
					return
				}

				mapped := A.AsciiArt(string(data))
				tab := A.Tab(arg[2])
				shade := strings.ToLower(colorName)
				colored := A.Paragraph(tab, mapped)
				Colored, line := A.ColourMap(colored, shade)
				fmt.Printf("%s%s\033[0m", Colored, line)
				return
			} else if bannerFile == "" {
				A.Checkfiles("resources/text/standard.txt")
				data, err := os.ReadFile("resources/text/standard.txt")
				if err != nil {
					fmt.Println("File not found")
					return
				}
				sub := A.Tab(arg[1])
				str := A.Tab(arg[2])
				

				shade := strings.ToLower(colorName)
				mapped := A.AsciiArt(string(data))
				index := A.IndexOf(str, sub)
				uncolored := A.Paragraph(str[:index], mapped)
				
				
				uncolored2 := A.Paragraph(str, mapped)
				colored3 := A.Paragraph(str[index:index+len(sub)], mapped)
				Colored, line := A.ColourMap(colored3, shade)
				colored33:=fmt.Sprintf("%s%s\033[0m",Colored,line)
				
				fmt.Printf("%s%s%s",  strings.Join(uncolored, "\\n"),colored33,strings.Join(uncolored2, ""))
			}
		}
	}
}
