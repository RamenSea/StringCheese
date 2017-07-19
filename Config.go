package main

import (
	"strings"
	"os"
	"time"
	"flag"
	"errors"
)

type StringValueConfig struct {
	/*
	NOT SET BY USER
	Adds this time stamp to the top of all generated files
	 */
	timeStampString string

	/*
	REQUIRED
	Path to the res folder in your android project

	-a STRING
	 */

	pathToAndroidRes string

	/*
	REQUIRED
	Path to the base of your iOS project

	-i STRING
	 */
	pathToIOSProject string

	/*
	OPTIONAL
	Sets what language to use as root

	-lang STRING

	DEFAULT: "NONE" -this serves to use whatever is the base language in the project
	 */
	rootLanguageId string
	/*
	OPTIONAL
	Path to where the Swift API key file will be generated.

	DEFAULT: pathToIOSProject
	 */
	pathToSwiftKey string

	/*
	OPTIONAL
	Name of XML strings file in your Android project.

	DEFAULT: "strings"
	 */
	nameOfXMLStringFile string

	/*
	OPTIONAL
	Name of .strings file that will be generated in your iOS project

	DEFAULT: "Localizable"
	 */
	nameOfDotStringFile string

	/*
	OPTIONAL
	Name of Swift API string key file that will be generated in your iOS project

	DEFAULT: "StringKeys"
	 */
	swiftClassName string

	/*
	OPTIONAL
	Whether or not the Swift API file is generated

	DEFAULT: true
	 */
	shouldCreateSwiftKey bool

	/*
	OPTIONAL
	Whether or not to generate

	DEFAULT: true
	 */
	shouldCreateArgumentsInSwiftAPI bool

	/*
	OPTIONAL
	Whether or not to generate a static

	DEFAULT: false
	 */
	useStaticForSwiftAPI bool


	/*
	OPTIONAL
	Will log if a string is missing

	DEFAULT: true
	 */
	logMissingStrings bool

	/*
	OPTIONAL
	Will reduce the keys to smallest possible

	DEFAULT: false
	 */
	reduceKeys bool
}

//Gets the path to the where the language's .strings file should be located
func (config *StringValueConfig) DotStringFileWithLanguageId(languageId string) string {
	if languageId == LANGUAGE_ID_NONE {
		return config.pathToIOSProject + "Base.lproj/" + config.nameOfDotStringFile + ".strings"
	}
	return config.pathToIOSProject + strings.Title(languageId) + ".lproj/" + config.nameOfDotStringFile + ".strings"
}
//gets all the language IDs from an Android projects res folder
func (config * StringValueConfig) GetAllValueFoldersLanguageIds() ([]string, error) {
	languageIds := []string{}
	res, err := os.Open(config.pathToAndroidRes)
	if err != nil {
		//todo
		return nil, err
	}
	fileNames, err := res.Readdirnames(0)

	for _, name := range fileNames {
		if strings.Contains(name, "values") {
			s := strings.Split(name,"-")
			if len(s) > 1 {
				//todo add valid language checks
				languageIds = append(languageIds, s[1])
			}
		}
	}
	return languageIds, nil
}

const NO_VALUE_FROM_FLAG  = "NONE"
const DEFAULT_LANGUAGE_ID = LANGUAGE_ID_NONE
const DEFAULT_DOT_STRING_FILE_NAME = "Localizable"
const DEFAULT_XML_STRING_FILE_NAME = "strings"
const DEFAULT_CREATE_SWIFT_KEY = true
/*
	Processes CLI arguments
 */
func parseAndGetConfig() (*StringValueConfig, error) {
	pathToAndroidRes := flag.String("a", NO_VALUE_FROM_FLAG, "REQUIRED\n" +
		"Path to your Android res folder.")

	pathToIOSAPP := flag.String("i", NO_VALUE_FROM_FLAG, "REQUIRED\n" +
		"Path to your iOS project.")

	defaultLang := flag.String("lang", NO_VALUE_FROM_FLAG, "OPTIONAL\n" +
		"Language to use as your default set of strings.\n" +
		"Default: base") //todo find out what defaults

	nameOfDotStringFile := flag.String("i_strings", NO_VALUE_FROM_FLAG, "OPTIONAL\n" +
		"Name of the .string file for iOS\n" +
		"Default: Localizable")

	nameOfXMLStringFile := flag.String("xml", NO_VALUE_FROM_FLAG, "OPTIONAL\n" +
		"Name of the .xml string files in your android project\n" +
		"Default: strings")

	shouldCreateSwiftKey := flag.Bool("swift", true, "OPTIONAL\n" +
		"Creates the Swift \n" +
		"Default: true")

	flag.Parse()



	if *pathToAndroidRes == NO_VALUE_FROM_FLAG {
		return nil, errors.New("Did not include path to your Android res folder.\n" +
			"Ex: ./StringValue -a /Users/me/workspace/androidApp/app/src/main/res")
	}
	if *pathToIOSAPP == NO_VALUE_FROM_FLAG {
		return nil, errors.New("Did not include path to your iOS project folder.\n" +
			"Ex: ./StringValue -a /Users/me/workspace/iOSAPP/iOSAPP")
	}

	if *defaultLang == NO_VALUE_FROM_FLAG {
		defaultLang = &DEFAULT_LANGUAGE_ID
	}
	if *nameOfDotStringFile == NO_VALUE_FROM_FLAG {
		nameOfDotStringFile = &DEFAULT_DOT_STRING_FILE_NAME
	}
	if *nameOfXMLStringFile == NO_VALUE_FROM_FLAG {
		nameOfXMLStringFile = &DEFAULT_XML_STRING_FILE_NAME
	}

	timeStamp := "// Last generated at: " + time.Now().String() + "\n"
	config := StringValueConfig{timeStampString: timeStamp,
		rootLanguageId: *defaultLang,
		pathToAndroidRes: *pathToAndroidRes,
		pathToIOSProject: *pathToIOSAPP,
		pathToSwiftKey: *pathToIOSAPP,
		nameOfDotStringFile: *nameOfDotStringFile,
		nameOfXMLStringFile: *nameOfXMLStringFile,
		shouldCreateSwiftKey: true,
		swiftClassName: "StringKeys",
		useStaticForSwiftAPI: false,
		shouldCreateArgumentsInSwiftAPI: true,
		logMissingStrings: true,
		reduceKeys: true, //todo set false
	}

	return &config, nil
}