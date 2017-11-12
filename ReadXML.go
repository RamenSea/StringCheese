package main

import (
	"os"
	"encoding/xml"
)

//Valid Language Codes
var validLanguageCodes = []string{"en", "aa", "ab", "af", "am", "ar", "as", "ay", "az", "ba", "be", "bg", "bh", "bi",
	"bn", "bo", "br", "ca", "co", "cs", "cy", "da", "de", "dz", "el", "eo", "es", "et", "eu", "fa", "fi", "fj", "fo", "fr",
	"fy", "ga", "gd","gl", "gn", "gu", "ha", "hi", "hr", "hu", "hy", "ia", "ie", "ik", "in", "is", "it", "iw", "ja", "ji",
	"jw", "ka", "kk","kl", "km", "kn", "ko", "ks", "ku", "ky", "la", "ln", "lo", "lt", "lv", "mg", "mi", "mk", "ml", "mn",
	"mo", "mr", "ms","mt", "my", "na", "ne", "nl", "no", "oc", "om", "pa", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru",
	"rw", "sa", "sd","sg", "sh", "si", "sk", "sl", "sm", "sn", "so", "sq", "sr", "ss", "st", "su", "sv", "sw", "ta", "te",
	"tg", "th", "ti","tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "uk", "ur", "uz", "vi", "vo", "wo", "xh", "yo", "zh", "zu"}

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

func GetStringKeysFromXML(languageId string, config *StringCheeseConfig,  platformProcessor func(value string, argNames string, sv *StringValue)) (*StringKeys, error) {
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
		//error decoding
		return nil, err
	}

	values := StringKeys{languageId: languageId, strings: map[string]*StringValue{}}
	xmlStringValue := res.Strings
	for _, value := range xmlStringValue {
		values.strings[value.Name] = CreateStringValue(value, config, platformProcessor)
	}

	fileReader.Close()
	return &values, nil
}
func CheckIfValidLanguageCode(languageCode string) bool {
	for _,id := range validLanguageCodes {
		if id == languageCode {
			return true
		}
	}
	return false
}


func CreateStringValue(value XMLStringValue, config *StringCheeseConfig, platformProcessor func(value string, argNames string, sv *StringValue)) *StringValue {
	s := StringValue{value.Translatable != "false",value.Name, value.Name, "",  0, "", ""}
	platformProcessor(value.Value, value.ArgNames, &s)
	return &s
}