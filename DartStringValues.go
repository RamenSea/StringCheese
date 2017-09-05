package main

import (
	"strconv"
	"strings"
	"os"
	"fmt"
	"encoding/xml"
	"regexp"
)

func CreateDartStringValue(value XMLStringValue, config *StringValueConfig) *StringValue {
	stringValue := value.Value

	//%1$s
	reg, _ := regexp.Compile(`(%[0-9]*\$s)`)
	numOfArg := reg.FindAllString(stringValue,-1)
	arguments := len(numOfArg)
	argumentString := ""
	formatString := ""
	if arguments > 0 {
		//find argument names here
		i := 0
		var argNamesFromXML []string

		if len(value.ArgNames) > 0 {
			strippedArgString := strings.Replace(value.ArgNames, " ", "", -1)
			argNamesFromXML = strings.Split(strippedArgString, ",")
		}
		formatString = reg.ReplaceAllStringFunc(stringValue, func(replaceString string) string {
			var argName string

			if i < len(argNamesFromXML) {
				argName = argNamesFromXML[i]
			} else {
				argName = "arg" + strconv.Itoa(i + 1)
			}
			if i > 0 {
				argumentString = argumentString + ", String " + argName
			} else {
				argumentString = "String " + argName
			}
			return "$" + argName
		})
	}

	s := StringValue{value.Name, value.Name, stringValue, value.Translatable != "false", arguments, argumentString, formatString}
	return &s
}
func getDartStringValueForLanguage(languageId string, config *StringValueConfig) *StringKeys {
	valueFolderName := "/values-" + languageId + "/"
	if languageId == LANGUAGE_ID_NONE {
		valueFolderName = "/values/"
	}
	fullStringPath := config.pathToAndroidRes + valueFolderName + config.nameOfXMLStringFile + ".xml"
	fileReader, err := os.Open(fullStringPath)
	if err != nil {
		fmt.Println("Error loading xml string at: " + fullStringPath)
		//couldn't find strings.xml
		return nil
	}

	values := StringKeys{languageId: languageId, strings: map[string]*StringValue{}}
	var res XMLResource
	if err := xml.NewDecoder(fileReader).Decode(&res); err != nil {
		return nil
	}

	xmlStringValue := res.Strings
	for _, value := range xmlStringValue {
		values.strings[value.Name] = CreateDartStringValue(value, config)
	}

	fileReader.Close()
	return &values
}