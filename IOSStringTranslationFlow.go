package main

import (
	"errors"
	"os"
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
	rootStringValue := getSwiftStringValueForLanguage(config.rootLanguageId, config)
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
		sv := getSwiftStringValueForLanguage(id, config)
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
	writeStringValueToDotStrings(rootStringValue, config)
	for _,value := range stringValues {
		writeStringValueToDotStrings(value,config)
	}

	return nil
}


func writeStringValueToDotStrings(value *StringKeys, config *StringCheeseConfig) {
	folderPathToDotString := config.DotStringFileWithLanguageId(value.languageId)
	_ = os.MkdirAll(folderPathToDotString, os.ModePerm) //skipped err check

	pathToDotString := config.DotStringFileWithLanguageId(value.languageId)
	//remove old string file
	_ = os.Remove(pathToDotString) //skipped err check
	file, err := os.Create(pathToDotString)
	if err != nil {
		//todo log
		return
	}
	valueMap := value.strings

	file.WriteString(config.timeStampString)
	for key, value := range valueMap {
		if value.translatable {
			file.WriteString("\"" + key + "\"=\"" + value.value + "\";\n")
		}
	}

	file.Close()
}