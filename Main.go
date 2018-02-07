package main

import (
	"fmt"
)

func main() {
	config, err := parseAndGetConfig() //In config.go
	if err != nil {
		fmt.Println(err)
		return
	}

	err = RunStringCheese(config)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func RunStringCheese(config *StringCheeseConfig) error {
	var err error
	if config.translatingToIOS {
		fmt.Println("Starting iOS translation")
		err = translateAndroidStringsToIOS(config)
		if err == nil {
			fmt.Println("Success!")
			fmt.Println("Make sure to add all generated files to your XCode project.")
			fmt.Println("Otherwise your XCode project will not see the newly generated files.")
		} else {
			return err
		}
	}

	if config.translatingToDart {
		fmt.Println("Starting Dart translation")
		err = translateAndroidStringsToDart(config)
		if err == nil {
			fmt.Println("Success")
			fmt.Println("Make sure to add all generated files to your Dart project")
		} else {
			return err
		}
	}

	if config.translatingToJS {
		fmt.Println("Starting JS translation")
		err = translateStringsToJavaScript(config)
		if err == nil {
			fmt.Println("Success")
			fmt.Println("Make sure to add all generated files to your JS project")
		} else {
			return err
		}
	}
	return nil
}