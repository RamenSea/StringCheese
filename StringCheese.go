package main

import (
	"encoding/xml"
	"strconv"
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

func (main *StringKeys) CompareAndAddValues(skipUntranslatableStrings bool, other *StringKeys, config *StringCheeseConfig) {
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