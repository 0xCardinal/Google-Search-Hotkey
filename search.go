// Objective - Google-Search-Hotkey - Read more at github.com/krAshwin/Google-Search-Hotkey
// Author - Kumar Ashwin
// Mail - connect@krash.dev

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/atotto/clipboard"
)

func openbrowser(url string) {
	var err error

	switch os := runtime.GOOS; os {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println("Unsupported Plaform!")
		os.Exit(1)
	}
}

// URL identification - basic comparision
func search(url string) {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		openbrowser(url)
	} else {
		url = "https://www.google.com/search?q=" + url
		openbrowser(url)
	}
}

func main() {

	if len(os.Args) == 1 {
		url, err := clipboard.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		search(url)
	} else {
		param := os.Args[1]
		search(param)
	}
}
