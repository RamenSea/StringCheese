package main

import (
	"strconv"
	"fmt"
)

/*
	The StringValue struct is a single translated string
	It contains the original key and the reduced key if needed
	The value is the translated string value
 */

type StringValue struct {
	translatable bool
	originalKey string
	key string
	value string

	numberOfArguments int
	argumentString string
	formatString string
}
/*
	StringKeys are a pairing between StringValues and a language
 */

type StringKeys struct {
	languageId string
	strings map[string]*StringValue //map of key to StringValue, StringValue's contain their own key mapping in StringValue.key
}

/*
	Compares strings against another StringKeys struct,
	It adds a blank StringValue for every missing key
	If logMissingStrings is true, it logs this to console
 */
func (sk *StringKeys) CompareAndAddValues(skipUntranslatableStrings bool, other *StringKeys, config *StringCheeseConfig) {
	for otherKey, otherValue := range other.strings {
		if skipUntranslatableStrings && otherValue.translatable == false {
			continue
		}
		var found = false
		for mainKey, _ := range sk.strings {
			if mainKey == otherKey {
				found = true
				break
			}
		}

		if found == false {
			v := StringValue{}
			v = *otherValue
			v.value = ""
			sk.strings[otherKey] = &v
			if config.logMissingStrings {
				fmt.Println(sk.languageId + " String Key file is missing string with id - " + v.originalKey)
			}
		}
	}
}

/*
Reduce keys by just turning them into an int
eventually you could reduce the keys even further by using a-z A-Z 0-9, but I don't see the point
 */
func (sk *StringKeys) ReduceKeys() {
	var oldKeys = sk.strings
	sk.strings = map[string]*StringValue{}

	var index = 0
	for _, value := range oldKeys {
		newKey := strconv.Itoa(index)
		value.key = newKey
		sk.strings[newKey] = value
		index++
	}
}

/*
Coppies keys based off the original key value.
skipped non-translated strings
 */
func (sk *StringKeys) CopyKeys(other *StringKeys) {
	for key, value := range other.strings {
		if value.translatable == false {
			continue
		}
		var found = false
		for myKey, myValue := range sk.strings {
			if myKey == value.originalKey {
				myValue.key = key
				sk.strings[key] = myValue
				delete(sk.strings, myKey)
				found = true
				break
			}
		}

		if found == false {
			fmt.Println("BUG in key reducing could not find matching key " + value.originalKey)
		}
	}
}