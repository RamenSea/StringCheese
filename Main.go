package main

import (
	"fmt"
)
/*
todo
2. Handle different types of string arguments other than just strings
 */

func main() {
	config, err := parseAndGetConfig() //In config.go

	if err != nil {
		fmt.Println(err)
		return
	}

	if config.translatingToIOS {
		err = translateAndroidStringsToIOS(config)
		if err == nil {
			fmt.Println("Success!")
			fmt.Println("Make sure to add all generated files to your XCode project.")
			fmt.Println("Otherwise your XCode project will not see the newly generated files.")
		} else {
			fmt.Println(err)
		}
	}

	if config.translatingToDart {
		err = translateAndroidStringsToDart(config)
		if err == nil {
			fmt.Println("Success")
			fmt.Println("Make sure to add all generated files to your Dart project")
		} else {
			fmt.Println(err)
		}
	}
}