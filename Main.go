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
const LANGUAGE_ID_NONE = "ROOT"
const LANGUAGE_ID_NONE_NAME = "en"

func main() {
	config, err := parseAndGetConfig() //In config.go

	if err != nil {
		fmt.Println(err)
		return
	}

	if config.translatingToIOS {
		err = translateAndroidStringsToIOS(config)

		if err == nil {
			fmt.Println("Success")
			fmt.Println("Make sure to add all generated files to your XCode project")
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
func translateAndroidStringsToIOS(config *StringValueConfig) error {
	rootStringValue := getSwiftStringValueForLanguage(config.rootLanguageId, config)
	if rootStringValue == nil {
		return errors.New("Error loading the root string value")
		//exit
	}
	ids, err := config.GetAllValueFoldersLanguageIds()
	if err != nil {
		return err
		//exit
	}
	stringValues := []*StringKeys{}
	for _,id := range ids {
		sv := getSwiftStringValueForLanguage(id, config)
		if sv != nil {
			stringValues = append(stringValues, sv)
		}
	}

	//adds missing strings keys to root value
	for _, value := range stringValues {
		rootStringValue.CompareAndAddValues(false, value, config)
	}
	//adds missing string keys to all of the string values from root value
	for _, value := range stringValues {
		value.CompareAndAddValues(true, rootStringValue, config)
	}


	//reduce keys if option is set
	if config.reduceKeys {
		rootStringValue.ReduceKeys()
		for _, value := range stringValues {
			value.CopyKeys(rootStringValue)
		}
	}

	writeStringValueToDotStrings(rootStringValue, config)
	writeSwiftKeyFile(rootStringValue, config)
	for _,value := range stringValues {
		writeStringValueToDotStrings(value,config)
	}


	//writeStringValueToDotStrings(test, &config)
	//writeSwiftKeyFile(test, &config)

	return nil
}

func translateAndroidStringsToDart(config *StringValueConfig) error {
	rootStringValue := getDartStringValueForLanguage(config.rootLanguageId, config)
	if rootStringValue == nil {
		return errors.New("Error loading the root string value")
		//exit
	}
	ids, err := config.GetAllValueFoldersLanguageIds()
	if err != nil {
		return err
		//exit
	}
	stringValues := []*StringKeys{}
	for _,id := range ids {
		sv := getDartStringValueForLanguage(id, config)
		if sv != nil {
			stringValues = append(stringValues, sv)
		}
	}

	//adds missing strings keys to root value
	for _, value := range stringValues {
		rootStringValue.CompareAndAddValues(false, value, config)
	}
	//adds missing string keys to all of the string values from root value
	for _, value := range stringValues {
		value.CompareAndAddValues(true, rootStringValue, config)
	}


	//reduce keys if option is set
	if config.reduceKeys {
		rootStringValue.ReduceKeys()
		for _, value := range stringValues {
			value.CopyKeys(rootStringValue)
		}
	}

	writeDartFile(rootStringValue, stringValues, config)
	//writeStringValueToDotStrings(test, &config)
	//writeSwiftKeyFile(test, &config)

	return nil
}