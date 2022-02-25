package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

const MIN_SIZE_OF_SPREAD_SHEET = 3
const CSV_POSITION_KEY = 0
const CSV_POSITION_VALUE = 1
const CSV_POSITION_TRANSLATABLE = 2
const CSV_POSITION_ARG_NAMES = 3

func getStringKeys(languageId string, config *StringCheeseConfig, platformProcessor func(value string, argNames string, sv *StringValue)) (*StringKeys, error) {
	if config.shouldUseSpreadSheetForStrings {
		return GetStringKeysFromCSV(languageId, config, platformProcessor)
	} else {
		return GetStringKeysFromXML(languageId, config, platformProcessor)
	}
}

func GetStringKeysFromCSV(languageId string, config *StringCheeseConfig, platformProcessor func(value string, argNames string, sv *StringValue)) (*StringKeys, error) {

	usingLanguageId := languageId
	if languageId == LANGUAGE_ID_NONE {
		usingLanguageId = config.rootLanguageIdToUse
	}

	fullStringPath := config.pathToSpreadSheetFolder + usingLanguageId + "_sheet.csv"
	fileReader, err := os.Open(fullStringPath)

	if err != nil {
		return nil, err
	}
	values := StringKeys{languageId: usingLanguageId, strings: map[string]*StringValue{}}

	var reader = csv.NewReader(fileReader)
	var index = -1
	for {
		index++
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if index == 0 {
			continue
		}

		if len(line) < MIN_SIZE_OF_SPREAD_SHEET {
			continue
			//return nil, errors.New("Line " + strconv.Itoa(index) + " of the "+usingLanguageId+" spread sheet does not contain enough rows.")
		}
		key := line[CSV_POSITION_KEY]
		if key == "" {
			continue
		}

		value := line[CSV_POSITION_VALUE]
		translatableString := strings.ToLower(line[CSV_POSITION_TRANSLATABLE])

		var translatable bool

		if translatableString == "false" ||
			translatableString == "0" ||
			translatableString == "no" {
			translatable = false
		} else {
			translatable = true
		}

		argNames := ""

		if len(line) > CSV_POSITION_ARG_NAMES {
			argNames = line[CSV_POSITION_ARG_NAMES]
		}

		values.strings[key] = CreateStringValue(translatable, key, value, argNames, config, platformProcessor)

	}

	fileReader.Close()
	return &values, nil
}
