package main

import (
	"fmt"
	"errors"
)
/*
TODO:
#1 Only check values folders with actually languages


Maybes:
Handle different types of string arguments other than just strings
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

	//if config.translatingToDart {
	//	err = translateAndroidStringsToDart(config)
	//	if err == nil {
	//		fmt.Println("Success")
	//		fmt.Println("Make sure to add all generated files to your Dart project")
	//	} else {
	//		fmt.Println(err)
	//	}
	//}
}
//func translateAndroidStringsToDart(config *StringCheeseConfig) error {
//	rootStringValue := getDartStringValueForLanguage(config.rootLanguageId, config)
//	if rootStringValue == nil {
//		return errors.New("Error loading the root string value")
//		//exit
//	}
//	ids, err := config.GetAllValueFoldersLanguageIds()
//	if err != nil {
//		return err
//		//exit
//	}
//	stringValues := []*StringKeys{}
//	for _,id := range ids {
//		sv := getDartStringValueForLanguage(id, config)
//		if sv != nil {
//			stringValues = append(stringValues, sv)
//		}
//	}
//
//	//adds missing strings keys to root value
//	for _, value := range stringValues {
//		rootStringValue.CompareAndAddValues(false, value, config)
//	}
//	//adds missing string keys to all of the string values from root value
//	for _, value := range stringValues {
//		value.CompareAndAddValues(true, rootStringValue, config)
//	}
//
//
//	//reduce keys if option is set
//	if config.reduceKeys {
//		rootStringValue.ReduceKeys()
//		for _, value := range stringValues {
//			value.CopyKeys(rootStringValue)
//		}
//	}
//
//	writeDartFile(rootStringValue, stringValues, config)
//	//writeStringValueToDotStrings(test, &config)
//	//writeSwiftKeyFile(test, &config)
//
//	return nil
//}