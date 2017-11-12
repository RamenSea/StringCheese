package main

import (
	"errors"
	"flag"
	"os"
	"strings"
	"time"
)

/*
	Config.go handles creating the config struct that works to configure how the strings are generated

	CONFIG OPTIONS ARE:

*/

const LANGUAGE_ID_NONE = "ROOT"
const LANGUAGE_ID_NONE_NAME = "en"

//Android related
const CONFIG_ARG_ANDROID_RES = "a" //REQUIRED
const MESSAGE_ANDROID_RES = "REQUIRED\n" +
	"        Path to your Android res folder.\n" +
	"        Ex: /Users/me/workspace/project/android/project_name/app/src/main/res\n"

const CONFIG_ARG_NAME_OF_STRING_XML_FILE = "s_xml" //Optional
const MESSAGE_NAME_OF_STRING_XML_FILE = "Optional\n" +
	"        Name of your xml file containing strings.\n"
const DEFAULT_VALUE_NAME_OF_STRING_XML_FILE = "strings" //.xml

const CONFIG_ARG_ROOT_LANGUAGE = "lang" //Optional
const MESSAGE_ROOT_LANGUAGE = "Optional\n" +
	"        Id of language to use as the base string xml file.\n" +
	"        For example en if you want the file 'values-en/strings.xml' to be used as the root language.\n" +
	"        If left blank the default folder will be values/strings.xml\n"
const DEFAULT_VALUE_ROOT_LANGUAGE = LANGUAGE_ID_NONE

const CONFIG_ARG_ROOT_LANGUAGE_ID = "langId" //Optional
const MESSAGE_ROOT_LANGUAGE_ID = "Optional\n" +
	"        Id of language to use as the base id for -lang if none is specified\n" +
	"        This value will be overwritten by -lang\n" +
	"        For right now this is only used with Dart\n"
const DEFAULT_VALUE_ROOT_LANGUAGE_ID = LANGUAGE_ID_NONE_NAME

//const CONFIG_ARG_ROOT_LANGUAGE_ID = "langId" //Optional
//const MESSAGE_ROOT_LANGUAGE_ID = "Optional\n" +
//	"Id of root language. This will override -lang's value when outputting to "
//const DEFAULT_VALUE_ROOT_LANGUAGE_ID = LANGUAGE_ID_NONE_NAME

//iOS related
const CONFIG_ARG_XCODE_PROJECT = "i" //Optional REQUIRED FOR IOS
const MESSAGE_XCODE_PROJECT = "Optional, REQUIRED to output to iOS\n" +
	"        Root folder of your iOS project. This is where the .strings files will be created.\n"

const CONFIG_ARG_NAME_OF_XCODE_DOT_STRING_FILE = "dotName" //Optional
const MESSAGE_NAME_OF_XCODE_DOT_STRING_FILE = "Optional\n" +
	"        Name of the .strings file.\n"
const DEFAULT_VALUE_NAME_OF_XCODE_DOT_STRING_FILE = "String"

const CONFIG_ARG_SHOULD_CREATE_SWIFT_KEYS = "swift" //Optional
const MESSAGE_SHOULD_CREATE_SWIFT_KEYS = "Optional\n" +
	"        If a swift key file should be generated.\n"
const DEFAULT_VALUE_SHOULD_CREATE_SWIFT_KEYS = true //this is relative to XCODE_PROJECT

const CONFIG_ARG_PATH_TO_SWIFT_FILE = "swiftLoc" //Optional
const MESSAGE_PATH_TO_SWIFT_FILE = "Optional\n" +
	"        Relative path inside the XCode project where the Swift StringCheese key class will be generated.\n" +
		"        If empty, this will generate it at the root of the project.\n"
const DEFAULT_VALUE_PATH_TO_SWIFT_FILE = "" //this is relative to XCODE_PROJECT



//Dart related
const CONFIG_ARG_PATH_TO_DART_PROJECT = "dart" //Optional
const MESSAGE_PATH_TO_DART_PROJECT = "Optional, REQUIRED to output to Dart\n" +
	"        Root folder of your Dart project. This is where the Dart StringCheese classes will be generated\n"


//general
const CONFIG_ARG_LOG_MISSING_STRINGS = "logMissing" //Optional
const MESSAGE_LOG_MISSING_STRINGS = "Optional\n" +
	"        Logs to the console if there are any missing strings in your strings.xml files\n"
const DEFAULT_VALUE_LOG_MISSING_STRINGS = true

const CONFIG_ARG_REDUCE_KEYS = "reduce" //Optional
const MESSAGE_REDUCE_KEYS = "Optional\n" +
	"        Reduces keys in outputted key classes\n"
const DEFAULT_VALUE_REDUCE_KEYS = true

const CONFIG_ARG_KEY_HAS_ARGUMENTS = "useArgs" //Optional
const MESSAGE_KEY_HAS_ARGUMENTS = "Optional\n" +
	"        If the key class generated generates a special argument method for strings with format points.\n"
const DEFAULT_VALUE_KEY_HAS_ARGUMENTS = true

const CONFIG_ARG_CLASS_IS_STATIC = "static" //Optional
const MESSAGE_CLASS_IS_STATIC = "Optional\n" +
	"        If the generated key file is static\n"
const DEFAULT_VALUE_CLASS_IS_STATIC = false

const CONFIG_ARG_SKIP_NON_VALID_LANGUAGE_ID = "useValidLangId" //Optional
const MESSAGE_SKIP_NON_VALID_LANGUAGE_ID = "Optional\n" +
	"        Skips a values-LANGUAGE_ID folder if the id is not valid\n"
const DEFAULT_VALUE_SKIP_NON_VALID_LANGUAGE_ID = false

const CONFIG_ARG_CLASS_NAME = "className" //Optional
const MESSAGE_CLASS_NAME = "Optional\n" +
	"        The name of the Class that will be generated\n"
const DEFAULT_VALUE_CLASS_NAME = "StringCheese"


const NO_VALUE_FROM_FLAG = "NONE"

/*
	Processes CLI arguments
*/
func parseAndGetConfig() (*StringCheeseConfig, error) {

	//android related
	pathToAndroidRes := flag.String(CONFIG_ARG_ANDROID_RES, NO_VALUE_FROM_FLAG, MESSAGE_ANDROID_RES)
	nameOfXMLFile := flag.String(CONFIG_ARG_NAME_OF_STRING_XML_FILE, DEFAULT_VALUE_NAME_OF_STRING_XML_FILE, MESSAGE_NAME_OF_STRING_XML_FILE)
	//ios
	iOSProjectRoot := flag.String(CONFIG_ARG_XCODE_PROJECT, NO_VALUE_FROM_FLAG, MESSAGE_XCODE_PROJECT)
	nameOfDotStrings := flag.String(CONFIG_ARG_NAME_OF_XCODE_DOT_STRING_FILE, DEFAULT_VALUE_NAME_OF_XCODE_DOT_STRING_FILE, MESSAGE_NAME_OF_XCODE_DOT_STRING_FILE)
	createSwiftKey := flag.Bool(CONFIG_ARG_SHOULD_CREATE_SWIFT_KEYS, DEFAULT_VALUE_SHOULD_CREATE_SWIFT_KEYS, MESSAGE_SHOULD_CREATE_SWIFT_KEYS)
	pathToSwift := flag.String(CONFIG_ARG_PATH_TO_SWIFT_FILE, DEFAULT_VALUE_PATH_TO_SWIFT_FILE, MESSAGE_PATH_TO_SWIFT_FILE)
	className := flag.String(CONFIG_ARG_CLASS_NAME, DEFAULT_VALUE_CLASS_NAME, MESSAGE_CLASS_NAME)
	//dart
	dartProject := flag.String(CONFIG_ARG_PATH_TO_DART_PROJECT, NO_VALUE_FROM_FLAG, MESSAGE_PATH_TO_DART_PROJECT)
	//general
	rootLanguage := flag.String(CONFIG_ARG_ROOT_LANGUAGE, DEFAULT_VALUE_ROOT_LANGUAGE, MESSAGE_ROOT_LANGUAGE)
	rootLanguageIfIfNone := flag.String(CONFIG_ARG_ROOT_LANGUAGE_ID, DEFAULT_VALUE_ROOT_LANGUAGE_ID, MESSAGE_ROOT_LANGUAGE_ID)
	logMissingStrings := flag.Bool(CONFIG_ARG_LOG_MISSING_STRINGS, DEFAULT_VALUE_LOG_MISSING_STRINGS, MESSAGE_LOG_MISSING_STRINGS)
	reduceKeys := flag.Bool(CONFIG_ARG_REDUCE_KEYS, DEFAULT_VALUE_REDUCE_KEYS, MESSAGE_REDUCE_KEYS)
	keyClassesHaveArgs := flag.Bool(CONFIG_ARG_KEY_HAS_ARGUMENTS, DEFAULT_VALUE_KEY_HAS_ARGUMENTS, MESSAGE_KEY_HAS_ARGUMENTS)
	staticClassKey := flag.Bool(CONFIG_ARG_CLASS_IS_STATIC, DEFAULT_VALUE_CLASS_IS_STATIC, MESSAGE_CLASS_IS_STATIC)
	skipNonValidLanguageIds := flag.Bool(CONFIG_ARG_SKIP_NON_VALID_LANGUAGE_ID, DEFAULT_VALUE_SKIP_NON_VALID_LANGUAGE_ID, MESSAGE_SKIP_NON_VALID_LANGUAGE_ID)

	flag.Parse()

	if *pathToAndroidRes == NO_VALUE_FROM_FLAG {
		return nil, errors.New("Did not include path to your Android res folder.\n" +
			"Ex: ./StringValue -a /Users/me/workspace/androidApp/app/src/main/res")
	}
	if *iOSProjectRoot == NO_VALUE_FROM_FLAG && *dartProject == NO_VALUE_FROM_FLAG {
		return nil, errors.New("Did not include path to an iOS or Dart project folder.\n" +
			"Ex: ./StringValue -a /Users/me/workspace/iOSAPP/iOSAPP")
	}

	var usingRootLanguageId = rootLanguageIfIfNone
	if (*rootLanguage == LANGUAGE_ID_NONE) {
		usingRootLanguageId = rootLanguage
	}
	//if *defaultLang == NO_VALUE_FROM_FLAG {
	//	*defaultLang = DEFAULT_LANGUAGE_ID
	//}

	timeStamp := "// Last generated at: " + time.Now().String() + "\n"
	config := StringCheeseConfig{
		timeStampString: timeStamp,
		rootLanguageId: *rootLanguage,
		rootLanguageIdToUse: *usingRootLanguageId,

		//android
		pathToAndroidRes: *pathToAndroidRes,
		nameOfXMLStringFile: *nameOfXMLFile,

		//ios
		translatingToIOS: *iOSProjectRoot != NO_VALUE_FROM_FLAG,
		pathToIOSProject: *iOSProjectRoot,
		nameOfDotStringFile: *nameOfDotStrings,
		shouldCreateSwiftKey: *createSwiftKey,
		pathToSwiftKey: *pathToSwift,

		//dart
		translatingToDart: *dartProject == NO_VALUE_FROM_FLAG,
		pathToDartFile: *dartProject,

		//general
		className: *className,
		logMissingStrings: *logMissingStrings,
		reduceKeys: *reduceKeys,
		shouldCreateArguments: *keyClassesHaveArgs,
		createStaticKeyClass: *staticClassKey,
		skipNonValidLanguageIds: *skipNonValidLanguageIds,
	}

	return &config, nil
}

//Gets the path to the where the language's .strings file should be located
func (config *StringCheeseConfig) DotStringFileWithLanguageId(languageId string) string {
	if languageId == LANGUAGE_ID_NONE {
		return config.pathToIOSProject + "/Base.lproj/" + config.nameOfDotStringFile + ".strings"
	}
	return config.pathToIOSProject + "/" + strings.Title(languageId) + ".lproj/" + config.nameOfDotStringFile + ".strings"
}

//gets all the language IDs from an Android projects res folder
func (config *StringCheeseConfig) GetAllValueFoldersLanguageIds() ([]string, error) {
	languageIds := []string{}
	res, err := os.Open(config.pathToAndroidRes)
	if err != nil {
		return nil, err
	}
	fileNames, err := res.Readdirnames(0)

	for _, name := range fileNames {
		if strings.Contains(name, "values") {
			s := strings.Split(name, "-")
			if len(s) > 1 {
				langId := s[1]
				if CheckIfValidLanguageCode(langId) || !config.skipNonValidLanguageIds {
					languageIds = append(languageIds, s[1])
				}
			}
		}
	}
	return languageIds, nil
}

/*
	Config Struct
*/
type StringCheeseConfig struct {
	//base
	timeStampString string
	rootLanguageId string
	rootLanguageIdToUse string

	//android
	pathToAndroidRes string
	nameOfXMLStringFile string

	//ios
	translatingToIOS bool
	pathToIOSProject string
	nameOfDotStringFile string
	shouldCreateSwiftKey bool
	pathToSwiftKey string

	//dart
	translatingToDart bool
	pathToDartFile string

	//general
	className string
	shouldCreateArguments bool
	createStaticKeyClass bool
	reduceKeys bool
	logMissingStrings bool
	skipNonValidLanguageIds bool
}
