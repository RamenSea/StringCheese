package main

import (
	"regexp"
	"strings"
	"strconv"
)

func kotlinMapPlatformStringValueProcessor(value string, argNames string, sv *StringValue)  {
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
		var argNamesFromXML []string

		if len(argNames) > 0 {
			strippedArgString := strings.Replace(argNames, " ", "", -1)
			argNamesFromXML = strings.Split(strippedArgString, ",")
		}
		formatString = reg.ReplaceAllStringFunc(usingValue, func(replaceString string) string {
			var argName string

			if i < len(argNamesFromXML) {
				argName = argNamesFromXML[i]
			} else {
				argName = "arg" + strconv.Itoa(i + 1)
			}
			if i > 0 {
				argumentString = argumentString + ": String, " + argName
			} else {
				argumentString = ": String " + argName
			}
			return "$" + argName
		})
	}

	sv.value = usingValue
	sv.numberOfArguments = arguments
	sv.argumentString = argumentString
	sv.formatString = formatString
}
