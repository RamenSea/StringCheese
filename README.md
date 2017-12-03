# StringCheese
### Overview:
StringCheese is a simple string translator that turns Android's strings.xml files into iOS' .strings and a Dart string file. 
It also generates a key class to help implementing the strings in your application.

### Quick use:

1. Run StringCheese.

    `./StringCheese -a <PATH_TO_ANDROID_RES_FOLDER> -i <PATH_TO_ROOT_OF_XCODE_PROJECT>`

    `./StringCheese -a <PATH_TO_ANDROID_RES_FOLDER> -dart <PATH_TO_DART_PROJECT>`

    `-a` is required and should be set to your Android's res folder

    `-i` is required if you want to output to XCode

    `-dart` is required if you want to output to Dart

2. If this is the first time running the program add the files to your XCode or Dart project.
3. Import and create the StringCheese class

### Features:

* Turns value-<language id>/strings.xml into <language id>.lproj/String.strings
* Creates turns Android strings into a Dart class structure
* Written in an extendable way to make it easy to add new targets in the future
* For iOS it creates a key class in Swift
* Both Dart and Swift key classes have string formatting methods with arguments. (ex: `stringCheese.someKeyWithFormats("arg1","arg2")`)
* Good defaults

### Todo:

* Increase tests
* Clean up code
* Clean up README
* Consider extending translation to and from XML, CSV, or a spreadsheet. This should be pretty simple

### Argument Bucket List


| Arg | Type    | Required  | Description  |
| --- | ------- | --------- | ------------ |
| a   | string  | y         |           d  |
| i   | string  | y         |           d  |
| a   | string  | y         |           d  |
| a   | string  | y         |           d  |
| a   | string  | y         |           d  |
| a   | string  | y         |           d  |


  -a string
        REQUIRED
        Path to your Android res folder.
        Ex: /Users/me/workspace/project/android/project_name/app/src/main/res
 (default "NONE")
 
  -className string
        Optional
        The name of the Class that will be generated
 (default "StringCheese")
 
  -dart string
        Optional, REQUIRED to output to Dart
        Root folder of your Dart project. This is where the Dart StringCheese classes will be generated
 (default "NONE")
 
  -dotName string
        Optional
        Name of the .strings file.
 (default "String")
 
  -i string
        Optional, REQUIRED to output to iOS
        Root folder of your iOS project. This is where the .strings files will be created.
 (default "NONE")
 
  -lang string
        Optional
        Id of language to use as the base string xml file.
        For example en if you want the file 'values-en/strings.xml' to be used as the root language.
        If left blank the default folder will be values/strings.xml
 (default "ROOT")
 
  -langId string
        Optional
        Id of language to use as the base id for -lang if none is specified
        This value will be overwritten by -lang
        For right now this is only used with Dart
 (default "en")
 
  -logMissing
        Optional
        Logs to the console if there are any missing strings in your strings.xml files
 (default true)
 
  -reduce
        Optional
        Reduces keys in outputted key classes
 (default true)
 
  -s_xml string
        Optional
        Name of your xml file containing strings.
 (default "strings")
 
  -static
        Optional
        If the generated key file is static

  -swift
        Optional
        If a swift key file should be generated.
 (default true)
 
  -swiftLoc string
        Optional
        Relative path inside the XCode project where the Swift StringCheese key class will be generated.
        If empty, this will generate it at the root of the project.

  -useArgs
        Optional
        If the key class generated generates a special argument method for strings with format points.
 (default true)
 
  -useValidLangId
        Optional
        Skips a values-LANGUAGE_ID folder if the id is not valid
