package main

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
	1. Get root string file
	2. Get the rest of the string files
	3. Clean this up and look for missing keys
	4. Generate Swift file if needed
	5. Reduce keys if need, but only if a Swift file is needed
	6. Generate .strings files
 */
func translateAndroidStringsToIOS(config *StringCheeseConfig) error {
	//	1. Get root string file
	rootStringValue, err := GetStringKeysFromXML(config.rootLanguageId, config, iOSPlatformStringValueProcessor)
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
		sv, err := GetStringKeysFromXML(id, config, iOSPlatformStringValueProcessor)
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

	//	4. Generate Swift file if needed
	if config.shouldCreateSwiftKey {
		//	5. Reduce keys if need, but only if a Swift file is needed
		if config.reduceKeys{
			rootStringValue.ReduceKeys()
			for _, value := range stringValues {
				value.CopyKeys(rootStringValue)
			}
		}
		WriteSwiftKeyFile(rootStringValue, config)
	}

	//	6. Generate .strings files
	err = writeStringValueToDotStrings(rootStringValue, config)
	if err != nil {
		return err
	}
	for _,value := range stringValues {
		err = writeStringValueToDotStrings(value,config)
		if err != nil {
			return err
		}
	}

	return nil
}