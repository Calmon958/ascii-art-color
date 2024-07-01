# Ascii_Art-Colour

Ascii_Art-Colour is a command-line Go application that renders input text strings into ASCII art representations. Transform your simple text into an exciting visual that stands out in any plain text environment.

## Objectives

Ascii-art aims to:

- Receive input strings and convert them into graphic representations using ASCII characters.
- Support varied input including alphanumerics, spaces, special characters, and line breaks.
- Offer multiple ASCII art fonts such as shadow, standard, and thinkertoy.

## Technical Specifications

- **Language**: The entire project is implemented in Go (Golang).
- **Coding Standards**: Follow Go's best practices for a clean, maintainable codebase.
- **Unit Testing**: Includes a test to ensure the reliability and stability of the application.

## Banner Format

Ascii-art characters adhere to the following specifications:

- Fixed height of 8 lines.
- New line separation for each character representation.

## Installation
    

This project is written in Go. If you don't have Go installed, you can follow these steps:
### step 1
1. Download the Go installer from the official [Go website](https://go.dev/doc/install).
2. Follow the installation instructions for your specific operating system.
3. Verify your installation by opening a command prompt and typing `go version`. This should print the installed version of Go.

### step 2

clone into this repo by entering this into the terminal

    git clone https://learn.zone01kisumu.ke/git/cliffootieno/ascii-art.git

### step 3
get into to the directory where the program is contained with this.

```
cd ascii-art
```

## test files
    main_test.go

## running the program
use the Go command-line interface with the following syntax:

```shell
go run . "Your desired string"
```

or 
```
go run main.go "Your desired string"
```

## Example
command:

```go
go run . "Hello\n" | cat -e
```
output:
```

 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```

if you use characters which are out of range in ascii between 32 and 126 it print's out an error message displaying it's not an ascii character
```shell
go run . "さようなら" 
```
the output will be: "Not an ascii character"


If you want to change the color of the output the number or aurgument on the command line will change. There are two instances;

* #### if you want to change the whole word:
    Here you will specify the color you want to change to in the following format:

    ```go
      go run . --color=<color> "something"
    ```

    This will colour the whole word in a particular color.
    ``` go 
    go run . --color=green "Apocalypse"
    ```

* #### If you wnat to color a specific or set of letter in a word
    Here you have to specify the letters that you want to color along with the arguments. The syntax should be as follows:

    ```go 
    go run . --color=<color> <letters to be colored> "something"
    ```

    Here is an example
    ```go
    go run . ---color=yellow b "bubble"
    ```

If you happen to use a wrong syntax or arrange them in a different way you will get an error which directs you on the the proper way to use it.

```bash
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> 'something'

```

## things to note
all the possible colors to use are in the rgb_hex.txt file in the ascii-art-color/resources/text folder.

## Contribution

The project was worked on by 3 software developers:

* Wilfred Njuguna
* Clifford Otieno
* Arnold Adero