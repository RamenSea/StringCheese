package main

import "errors"

/*
	1. Get root string file
	2. Get the rest of the string files
	3. Clean this up and look for missing keys
	4. Generate Kotlin file
*/
func translateStringsToKotlin(config *StringCheeseConfig) error {

	//	1. Get root string file
	rootStringValue, err := getStringKeys(config.rootLanguageId, config, kotlinPlatformStringValueProcessor)
	if rootStringValue == nil {
		return errors.New("Error loading the root string value")
		//exit
	}
	//2. Get the rest of the string files
	ids, err := config.getAllValueFoldersLanguageIds()
	if err != nil {
		return err
		//exit
	}

	stringValues := []*StringKeys{}
	for _, id := range ids {
		sv, err := getStringKeys(id, config, kotlinPlatformStringValueProcessor)
		if err != nil {
			return err
			//exit
		}
		if sv != nil {
			stringValues = append(stringValues, sv)
		}
	}
	//3. Clean this up and look for missing keys

	//adds missing strings keys to root value
	for _, value := range stringValues {
		rootStringValue.compareAndAddValues(false, false, value, config)
	}
	//adds missing string keys to all of the string values from root value
	for _, value := range stringValues {
		value.compareAndAddValues(true, true, rootStringValue, config)
	}

	//4. Generate JavaScript file
	return writeKotlinFiles(rootStringValue, stringValues, config)
}
