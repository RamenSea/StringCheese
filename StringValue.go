package main

import (
	"encoding/xml"
	"os"
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

type StringKeys struct {
	languageId string
	strings map[string]*StringValue
}
type StringValue struct {
	originalKey string
	key string
	value string

	translatable bool

	numberOfArguments int
	argumentString string
	formatString string
}
type XMLStringValue struct {
	XMLName xml.Name `xml:"string"`
	Name string `xml:"name,attr"`
	Translatable string `xml:"translatable,attr"`
	ArgNames string `xml:"argNames,attr"`
	Value string `xml:",chardata"`
}
type XMLResource struct {
	XMLName xml.Name `xml:"resources"`
	Strings []XMLStringValue `xml:"string"`
}
func CreateStringValue(value XMLStringValue, config *StringValueConfig) *StringValue {
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
func getStringValueForLanguage(languageId string, config *StringValueConfig) *StringKeys {
	valueFolderName := "values-" + languageId + "/"
	if languageId == LANGUAGE_ID_NONE {
		valueFolderName = "values/"
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
		values.strings[value.Name] = CreateStringValue(value, config)
	}

	fileReader.Close()
	return &values
}

func writeStringValueToDotStrings(value *StringKeys, config *StringValueConfig) {
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

func (main *StringKeys) CompareAndAddValues(skipUntranslatableStrings bool, other *StringKeys, config *StringValueConfig) {
	for otherKey, otherValue := range other.strings {
		if skipUntranslatableStrings && otherValue.translatable == false {
			continue
		}
		var found = false
		for mainKey, _ := range main.strings {
			if mainKey == otherKey {
				found = true
				break
			}
		}

		if found == false {
			v := StringValue{}
			v = *otherValue
			v.value = ""
			main.strings[otherKey] = &v
			if config.logMissingStrings {
				fmt.Println(main.languageId + " String Key file is missing string with id - " + v.originalKey)
			}
		}
	}
}


func (main *StringKeys) ReduceKeys() {
	var oldKeys = main.strings
	main.strings = map[string]*StringValue{}

	var index = 0
	for _, value := range oldKeys {
		newKey := strconv.Itoa(index)
		value.key = newKey
		main.strings[newKey] = value
		index++
	}
}
func (main *StringKeys) CopyKeys(other *StringKeys) {
	for key, value := range other.strings {
		if value.translatable == false {
			continue
		}
		var found = false
		for myKey, myValue := range main.strings {
			if myKey == value.originalKey {
				myValue.key = key
				main.strings[key] = myValue
				delete(main.strings, myKey)
				found = true
				break
			}
		}

		if found == false {
			fmt.Println("BUG in key reducing could not find matching key " + value.originalKey)
		}
	}
}