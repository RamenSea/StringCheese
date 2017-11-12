package main

import (
	"strconv"
	"strings"
	"os"
	"fmt"
	"encoding/xml"
	"regexp"
)

func CreateSwiftStringValue(value XMLStringValue, config *StringCheeseConfig) *StringValue {
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
		stringValue = reg.ReplaceAllStringFunc(stringValue, func(replaceString string) string {
			return "%" + strconv.Itoa(i + 1) + "$@"
		})
		var argNamesFromXML []string
		if len(value.ArgNames) > 0 {
			strippedArgString := strings.Replace(value.ArgNames, " ", "", -1)
			argNamesFromXML = strings.Split(strippedArgString, ",")
		}
		for i := 0; i < arguments; i++ {
			argName := "arg" + strconv.Itoa(i + 1)
			formatName := argName
			if i < len(argNamesFromXML) {
				argName = argNamesFromXML[i]
				formatName = argName
			} else if i == 0 {
				argName = "_ " + argName
			}

			if i + 1 < arguments {
				argName = argName + ": String, "
				formatName = formatName + ", "
			} else {
				argName = argName + ": String"
			}
			argumentString = argumentString + argName
			formatString = formatString + formatName
		}
	}

	s := StringValue{value.Name, value.Name, stringValue, value.Translatable != "false", arguments, argumentString, formatString}
	return &s
}
func getSwiftStringValueForLanguage(languageId string, config *StringCheeseConfig) *StringKeys {
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
		values.strings[value.Name] = CreateSwiftStringValue(value, config)
	}

	fileReader.Close()
	return &values
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