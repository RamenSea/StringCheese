package main

import (
	"errors"
)

/*
	1. Get root string file
	2. Get the rest of the string files
	3. Clean this up and look for missing keys
	4. Reduce if set
	5. Generate Dart file
 */
func translateAndroidStringsToDart(config *StringCheeseConfig) error {

	//	1. Get root string file
	rootStringValue, err := GetStringKeysFromXML(config.rootLanguageId, config, DartPlatformStringValueProcessor)
	if rootStringValue == nil {
		return errors.New("Error loading the root string value")
		//exit
	}
	//2. Get the rest of the string files
	ids, err := config.GetAllValueFoldersLanguageIds()
	if err != nil {
		return err
		//exit
	}
	stringValues := []*StringKeys{}
	for _,id := range ids {
		sv, err := GetStringKeysFromXML(id, config, DartPlatformStringValueProcessor)
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
		rootStringValue.CompareAndAddValues(false, value, config)
	}
	//adds missing string keys to all of the string values from root value
	for _, value := range stringValues {
		value.CompareAndAddValues(true, rootStringValue, config)
	}


	//4. Reduce if set
	//reduce keys if option is set
	if config.reduceKeys {
		rootStringValue.ReduceKeys()
		for _, value := range stringValues {
			value.CopyKeys(rootStringValue)
		}
	}

	//5. Generate Dart file
	return WriteDartFiles(rootStringValue, stringValues, config)
}
