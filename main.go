package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var targetURL string
var showSuccessful bool
var saveResults bool
var wordlistPath string
var secure bool

var ASCII string = `
_____  _____  _____  _____  ___  _____  __ __  _____  _____ 
/   __\/  _  \/  _  \/  _  \/___\/   __\/  |  \|__   /|__   /
|  |_ ||  |  ||  _  ||   __/|   ||   __||  |  | /  _/  /  _/ 
\_____/\_____/\__|__/\__/   \___/\__/   \_____//_____|/_____|
                                                             
`

/*
Initializing required and optional flags
*/
func initializeFlags() {
	flag.StringVar(&targetURL, "url", "", "Takes the target's URL to scan.\nNOTE: Make sure to add FUZZTHIS at the end of path. This fuzz will be replaced by the characters from wordlist.\nEx: https://evil.com/v1/FUZZTHIS")
	flag.BoolVar(&showSuccessful, "show-successful", false, "Shows only the endpoints which are accessable in the target.")
	flag.StringVar(&wordlistPath, "wordlist", "", "The path of wordlist to use for brute-forcing.")
	flag.BoolVar(&secure, "secure", true, "Toggle the use of secure protocol https or not.")
	// flag.BoolVar(&saveResults, "save", false, "Saves the results in file.")

}

func main() {
	initializeFlags()
	flag.Parse()

	fmt.Println(ASCII)
	fmt.Println("=========================================================")
	PrintLog("info", "Starting GoAPIFuzz.... Let's have some fun!!")

	if wordlistPath == "" {
		PrintLog("info", "There is no wordlist path provided, so using my default wordlist.")
		PrintLog("info", "You can provide a wordlist using --wordlist flag.")

		wordlistPath = "./wordlist/list.txt"
	}

	if showSuccessful {
		PrintLog("info", "showSuccessful is set to true, so will display only successful requests.")
	}

	if !strings.Contains(targetURL, "FUZZTHIS") {
		PrintLog("error", "Target URL does not contain FUZZTHIS string to fuzz the endpoints.")
		os.Exit(1)
	}

	wordlistFile, err := os.Open(wordlistPath)
	defer wordlistFile.Close()

	if err != nil {
		PrintLog("error", "Cannot open the wordlist file at provided path: "+wordlistPath+". Please provide a valid path.")
		os.Exit(1)
	}

	Scanner(targetURL, showSuccessful, wordlistFile, secure)

	fmt.Println("=======================================================")
	fmt.Print("Thanks for using GoAPIFuzz! I hope you got what you wanted :)")
}
