package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func writeStringValueToDotStrings(value *StringKeys, config *StringCheeseConfig) error {
	folderPathToDotString := config.DotStringFileWithLanguageId(value.languageId)
	_ = os.MkdirAll(folderPathToDotString, os.ModePerm) //skipped err check

	pathToDotString := config.DotStringFileWithLanguageId(value.languageId)
	//remove old string file
	_ = os.Remove(pathToDotString) //skipped err check
	file, err := os.Create(pathToDotString)
	if err != nil {
		return err
	}
	valueMap := value.strings

	file.WriteString(config.timeStampString)
	for key, value := range valueMap {
		if value.translatable {
			file.WriteString("\"" + key + "\"=\"" + value.value + "\";\n")
		}
	}
	return file.Close()
}

func iOSPlatformStringValueProcessor(value string, argNames string, sv *StringValue) {
	usingValue := value
	//%1$s
	reg, _ := regexp.Compile(`(%[0-9]*\$s)`)
	numOfArg := reg.FindAllString(usingValue,-1)
	arguments := len(numOfArg)
	argumentString := ""
	formatString := ""
	if arguments > 0 {
		//find argument names here
		i := 0
		usingValue = reg.ReplaceAllStringFunc(usingValue, func(replaceString string) string {
			return "%" + strconv.Itoa(i + 1) + "$@"
		})
		var argNamesFromXML []string
		if len(argNames) > 0 {
			strippedArgString := strings.Replace(argNames, " ", "", -1)
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
	sv.value = usingValue
	sv.numberOfArguments = arguments
	sv.argumentString = argumentString
	sv.formatString = formatString
}