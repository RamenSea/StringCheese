package main

import (
	"encoding/csv"
	"errors"
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

	var fullStringPath string
	if config.spreadSheetValuesAllInOneSheet {
		fullStringPath = config.pathToSpreadSheetFolder
	} else {
		fullStringPath = config.pathToSpreadSheetFolder + usingLanguageId + "_sheet.csv"
	}

	fileReader, err := os.Open(fullStringPath)

	if err != nil {
		return nil, err
	}
	values := StringKeys{languageId: usingLanguageId, strings: map[string]*StringValue{}}

	var reader = csv.NewReader(fileReader)
	var lineOfText = CSV_POSITION_VALUE
	var extraPaddingAfterValues = 0
	var columnNumberForArgs = CSV_POSITION_ARG_NAMES
	var columnNumberForTranslatable = CSV_POSITION_TRANSLATABLE
	var index = -1
	var headers []string = nil
	if config.spreadSheetValuesAllInOneSheet {
		extraPaddingAfterValues = config.spreadSheetNumberOfLanguages
		lineOfText = -1
		columnNumberForArgs = columnNumberForArgs + extraPaddingAfterValues - 1
		columnNumberForTranslatable = columnNumberForTranslatable + extraPaddingAfterValues - 1
	}
	for {
		index++
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if index == 0 {
			headers = line
			continue
		}

		if lineOfText <= 0 {
			if headers == nil {
				return nil, errors.New("no header in CSV")
			}
			for headerIndex, headerValue := range headers {
				if strings.ToLower(headerValue) == usingLanguageId {
					lineOfText = headerIndex
					break
				}
			}
			if lineOfText <= 0 {
				return nil, errors.New("Could not find column " + usingLanguageId + " in CSV")
			}
		}
		if len(line) < MIN_SIZE_OF_SPREAD_SHEET+extraPaddingAfterValues {
			continue
			//return nil, errors.New("Line " + strconv.Itoa(index) + " of the "+usingLanguageId+" spread sheet does not contain enough rows.")
		}
		key := line[CSV_POSITION_KEY]
		if key == "" {
			continue
		}

		value := line[lineOfText]
		translatableString := strings.ToLower(line[columnNumberForTranslatable])

		var translatable bool

		if translatableString == "false" ||
			translatableString == "0" ||
			translatableString == "no" {
			translatable = false
		} else {
			translatable = true
		}

		argNames := ""

		if len(line) > columnNumberForArgs {
			argNames = line[columnNumberForArgs]
		}

		values.strings[key] = CreateStringValue(translatable, key, value, argNames, config, platformProcessor)

	}

	fileReader.Close()
	return &values, nil
}
