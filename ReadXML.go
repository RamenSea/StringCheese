package main

import (
	"os"
	"fmt"
	"encoding/xml"
	"errors"
)

/*
	Go XML structure of android xml string files
 */

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

func GetStringValueFromXML(languageId string, config *StringCheeseConfig, platformStringValueCreator func(value XMLStringValue, config *StringCheeseConfig) *StringValue) (*StringKeys, error) {
	valueFolderName := "/values-" + languageId + "/"
	if languageId == LANGUAGE_ID_NONE {
		valueFolderName = "/values/"
	}

	fullStringPath := config.pathToAndroidRes + valueFolderName + config.nameOfXMLStringFile + ".xml"
	fileReader, err := os.Open(fullStringPath)
	if err != nil {
		//couldn't find strings.xml
		return nil, err
	}

	var res XMLResource
	if err := xml.NewDecoder(fileReader).Decode(&res); err != nil {
		return nil, err
	}

	values := StringKeys{languageId: languageId, strings: map[string]*StringValue{}}
	xmlStringValue := res.Strings
	for _, value := range xmlStringValue {
		values.strings[value.Name] = platformStringValueCreator(value, config)
	}

	fileReader.Close()
	return &values, nil
}