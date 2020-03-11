package messages

import (
	"fmt"
	"os"
)

var ngrokNotFound string = `ngrok is not installed on your system.

Please visit https://ngrok.com/download to download and setup ngrok.

ngrok executable must be accessible in order to fart your files, 
so make sure ngrok is in dir which is set in $PATH environment variable
`

func PrintNgrokNotFound(exit bool, code int) {
	fmt.Println(ngrokNotFound)
	if exit {
		os.Exit(code)
	}
}

var fileNotProvided string = "file not provided for share, use -f <filename>\n\n"

func PrintFileNotProvided(exit bool, code int) {
	fmt.Println(fileNotProvided)

	if exit {
		os.Exit(code)
	}
}

func OutputVersion(Version, BuildTarget, BuildDate string, exit bool, code int) {
	fmt.Printf("fart version %s (%s @ %s)\n", Version, BuildTarget, BuildDate)
	if exit {
		os.Exit(code)
	}
}
