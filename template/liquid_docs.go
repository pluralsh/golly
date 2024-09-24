package template

var functionDocs = map[string]FilterFunctionDocumentation{
	"abbrev": {
		Description: "Truncates a string with ellipses.",
		Parameters:  []string{"Max length.", "The string."},
		Example:     "`abbrev 5 \"hello world\"` returns `he...`.",
	},
	"abbrevboth": {
		Description: "Truncates both sides of a string with ellipses.",
		Parameters:  []string{"Left offset.", "Max length", "The string."},
		Example:     "`abbrevboth 5 10 \"1234 5678 9123\"` returns `...5678...`.",
	},
	"add": {
		Description: "Sums numbers. Accepts two or more inputs.",
		Parameters:  []string{"A number.", "A number.", "..."},
		Example:     "`add 1 2 3` returns `6`.",
	},
	"add1": {
		Description: "Increments by 1.",
		Parameters:  []string{"The number to increment."},
		Example:     "`add1 3` returns `4`.",
	},
	"add1f": {
		Description: "Increments float number by 1.",
		Parameters:  []string{"The float number to increment."},
		Example:     "`add1 3.0` returns `4.0`.",
	},
	"addf": {
		Description: "Sums float numbers. Accepts two or more inputs.",
		Parameters:  []string{"A float number.", "A float number.", "..."},
		Example:     "`add 1.1 2.2 3.3` returns `6.6`.",
	},
	"adler32sum": {
		Description: "Computes Adler-32 checksum of a string.",
		Parameters:  []string{"The string."},
		Example:     "`adler32sum \"Hello world!\"`.",
	},
	"ago": {
		Description: "Returns duration from current time (`time.Now`) in seconds resolution.",
		Parameters:  []string{"The string."},
		Example:     "`ago .CreatedAt` will return something like `2h34m7s`.",
	},
	"all": {
		Description: "Takes a list of values ad returns true if all values are non-empty.",
		Parameters:  []string{"The list."},
	},
	"any": {
		Description: "Takes a list of values ad returns true if any values are non-empty.",
		Parameters:  []string{"The list."},
	},
	"append": {
		Description: "Appends a new item to existing list, creating a new list.",
		Parameters:  []string{"Existing list.", "New item."},
		Example:     "`append ( list 1 2 3 ) 5` returns `[1, 2, 3, 5]`.",
	},
	"atoi": {
		Description: "Converts a string to an integer.",
		Parameters:  []string{"The string."},
	},
	"b32dec": {
		Description: "Decodes string from Base32 format.",
		Parameters:  []string{"The string to decode."},
	},
	"b32enc": {
		Description: "Encodes string with Base32 format.",
		Parameters:  []string{"The string to encode."},
	},
	"b64dec": {
		Description: "Decodes string from Base64 format.",
		Parameters:  []string{"The string to decode."},
	},
	"b64enc": {
		Description: "Encodes string with Base64 format.",
		Parameters:  []string{"The string to encode."},
	},
	"base": {
		Description: "Returns the last element of a path.",
		Parameters:  []string{"The path."},
		Example:     "`base \"foo/bar/baz\"` returns `baz`.",
	},
	"bcrypt": {
		Description: "Generates bcrypt hash of a string.",
		Parameters:  []string{"The string."},
	},
	"biggest":         {},
	"buildCustomCert": "Allows customizing the certificate. It takes Base64 encoded PEM format certificate and private key as arguments and returns object with PEM-encoded certificate and key. Note that the returned object can be passed to the `genSignedCert` function to sign a certificate using this CA.",
	"camelcase":       {},
	"cat":             {},
	"ceil": {
		Description: "Returns greatest float value greater than or equal to input value.",
		Parameters:  []string{"Input value."},
		Example:     "`ceil 123.001` will return `124.0`.",
	},
	"chunk":    {},
	"clean":    {},
	"coalesce": {},
	"compact": {
		Description: "Accepts a list and removes entries with empty values.",
		Parameters:  []string{"The list."},
	},
	"concat":         "Concatenates arbitrary number of lists into one.",
	"contains":       "Tests if one string is contained inside of another. `contains \"cat\" \"catch\"` will return `true`.",
	"date":           "Formats date.",
	"dateInZone":     "Same as `date` but with a timezone.",
	"dateModify":     {},
	"date_in_zone":   {},
	"date_modify":    {},
	"decryptAES":     "Receives a Base64 string encoded by the AES-256 CBC algorithm and returns the decoded text.",
	"deepCopy":       {},
	"deepEqual":      {},
	"default":        {},
	"derivePassword": {},
	"dict":           {},
	"dig":            {},
	"dir":            {},
	"div":            "Performs integer division.",
	"divf":           {},
	"duration":       "Formats a given amount of seconds as a `time.Duration`.",
	"durationRound":  {},
	"empty":          {},
	"encryptAES":     "Encrypts text with AES-256 CBC and returns a Base64 encoded string.",
	"env": {
		Description: "Reads environment variable.",
		Parameters:  []string{"Environment variable name."},
		Example:     "`env \"HOME\"`",
	},
	"expandenv": {
		Description: "Substitutes environment variable in a string.",
		Parameters:  []string{"String to expand."},
		Example:     "`expandenv \"Your path is set to $PATH\"`",
	},
	"ext": {
		Description: "Returns file extension.",
		Parameters:  []string{"File path."},
		Example:     "`ext \"foo.bar\"` will return `\"bar\"`",
	},
	"fail": {
		Description: "Unconditionally returns an empty string and an error with the specified text. This is useful in scenarios where other conditionals have determined that template rendering should fail.",
		Parameters:  []string{"Error message."},
		Example:     "`fail \"Please accept the end user license agreement\"`",
	},
	"first": {
		Description: "Returns head item on a list.",
		Parameters:  []string{"The list."},
	},
	"float64": {
		Description: "Converts to a `float64`.",
		Parameters:  []string{"The value to convert. It can be for example an integer or a string."},
	},
	"floor":                    "Returns the greatest float value greater than or equal to input value. `floor 123.9999` will return `123.0`.",
	"fromJson":                 {},
	"genCA":                    "Generates a new, self-signed x509 SSL Certificate Authority using 2048-bit RSA private key. It takes subject common name (CN) and cert validity duration in days as parameters. It returns object with PEM-encoded certificate and key. Note that the returned object can be passed to the `genSignedCert` function to sign a certificate using this CA.",
	"genCAWithKey":             "Generates a new, self-signed x509 SSL Certificate Authority using given private key. It takes subject common name (CN), cert validity duration in days and private key (PEM-encoded; DSA keys are not supported) as parameters. It returns object with PEM-encoded certificate and key. Note that the returned object can be passed to the `genSignedCert` function to sign a certificate using this CA.",
	"genPrivateKey":            {},
	"genSelfSignedCert":        "Generates an SSL self-signed certificate.",
	"genSelfSignedCertWithKey": {},
	"genSignedCert":            "Generates an SSL certificate and key based on a given CA.",
	"genSignedCertWithKey":     {},
	"get":                      {},
	"getHostByName":            {},
	"has": {
		Description: "Checks if a list has a particular element. It will panic if there is a problem.",
		Parameters:  []string{"Element to find.", "The list."},
		Example:     "`has 4 $myList`",
	},
	"hasKey":         {},
	"hasPrefix":      {},
	"hasSuffix":      {},
	"htmlDate":       "Formats a date for inserting into HTML date picker input field.",
	"htmlDateInZone": "Same as `htmlDate` but with a timezone.",
	"htpasswd":       {},
	"indent":         {},
	"initial":        "Compliments `last` by retuning all but the last element.",
	"initials":       {},
	"int": {
		Description: "Converts to a `int`.",
		Parameters:  []string{"The value to convert."},
	},
	"int64": {
		Description: "Converts to a `int64`.",
		Parameters:  []string{"The value to convert."},
	},
	"isAbs": {
		Description: "Checks whether a path is absolute.",
		Parameters:  []string{"The file path."},
	},
	"join":      {},
	"kebabcase": {},
	"keys": {
		Description: "Returns list of all keys from a map.",
		Parameters:  []string{"The map."},
	},
	"kindIs": {},
	"kindOf": {},
	"last":   {},
	"list":   {},
	"lower": {
		Description: "Converts the entire string to lowercase.",
		Parameters:  []string{"The string."},
		Example:     "`upper \\\"HELLO\\\"` will return `hello`.",
	},
	"max": {
		Description: "Returns the largest of a series of integers.",
		Parameters:  []string{"A number.", "A number.", "..."},
		Example:     "`max 1 2 3` will return `3`.",
	},
	"maxf": {
		Description: "Returns the largest of a series of floats.",
		Parameters:  []string{"A float number.", "A float number.", "..."},
		Example:     "`max 1 2 3.65` will return `3.65`.",
	},
	"merge":          {},
	"mergeOverwrite": {},
	"min": {
		Description: "Returns the smallest of a series of integers.",
		Parameters:  []string{"A number.", "A number.", "..."},
		Example:     "`min 1 2 3` will return `1`.",
	},
	"minf": {
		Description: "Returns the smallest of a series of floats.",
		Parameters:  []string{"A float number.", "A float number.", "..."},
		Example:     "`min 1.3 2 3` will return `1.3`.",
	},
	"mod": {},
	"mul": {
		Description: "Multiplies numbers. Accepts two or more inputs.",
		Parameters:  []string{"A number.", "A number.", "..."},
		Example:     "`mul 1 2 3` will return `6`.",
	},
	"mulf": {
		Description: "Multiplies float numbers. Accepts two or more inputs.",
		Parameters:  []string{"A float number.", "A float number.", "..."},
		Example:     "`mulf 1.5 2 2` returns `6`.",
	},
	"mustAppend":     "Appends a new item to existing list, creating a new list. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustChunk":      {},
	"mustCompact":    "Accepts a list and removes entries with empty values. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustDateModify": {},
	"mustDeepCopy":   {},
	"mustFirst":      "Returns head item on a list. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustFromJson":   {},
	"mustHas": {
		Description: "Checks if a list has a particular element. It will return an error to the template engine ifd there is a problem.",
		Parameters:  []string{"Element to find.", "The list."},
		Example:     "`mustHas 4 $myList`",
	},
	"mustInitial":                "Compliments `last` by retuning all but the last element. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustLast":                   {},
	"mustMerge":                  {},
	"mustMergeOverwrite":         {},
	"mustPrepend":                {},
	"mustPush":                   {},
	"mustRegexFind":              {},
	"mustRegexFindAll":           {},
	"mustRegexMatch":             {},
	"mustRegexReplaceAll":        {},
	"mustRegexReplaceAllLiteral": {},
	"mustRegexSplit":             {},
	"mustRest":                   "Gets tail of the list (everything but the first item). Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustReverse":                "Produces a new list with the reversed elements of the given list. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustSlice":                  {},
	"mustToDate":                 "Converts a string to a date. The first argument is the date layout and the second is the date string. If the string can’t be converted it returns the zero value. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustToJson":                 {},
	"mustToPrettyJson":           {},
	"mustToRawJson":              {},
	"mustUniq":                   "Generates a list with all of the duplicates removed. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"mustWithout":                "Filters items out of a list. Like other `must` functions instead of panicking when there is a problem it will return an error to the template engine.",
	"must_date_modify":           {},
	"nindent":                    {},
	"nospace": {
		Description: "Removes all whitespace from a string.",
		Parameters:  []string{"The string."},
		Example:     "`nospace \"hello w o r l d\"` will return `helloworld`.",
	},
	"omit":         {},
	"osBase":       "Returns the last element of a file path. `osBase \"/foo/bar/baz\"` and `osBase \"C:\\\\foo\\\\bar\\\\baz\"` will return `\"baz\"` on Linux and Windows, respectively.",
	"osClean":      "Cleans up a path. `osClean \"/foo/bar/../baz\"` and `osClean \"C:\\\\foo\\\\bar\\\\..\\\\baz\"` will resolve the `..` and return `foo/baz` on Linux and `C:\\foo\\baz` on Windows.",
	"osDir":        "Returns the directory, stripping the last part of the path. So `osDir \"/foo/bar/baz\"` returns `/foo/bar` on Linux, and `osDir \"C:\\\\foo\\\\bar\\\\baz\"` returns `C:\\\\foo\\\\bar` on Windows.",
	"osExt":        "Return the file extension. `osExt \"/foo.bar\"` and `osExt \"C:\\\\foo.bar\"` will return `.bar` on Linux and Windows, respectively.",
	"osIsAbs":      "Checks whether a file path is absolute.",
	"pick":         {},
	"pluck":        {},
	"plural":       {},
	"prepend":      {},
	"push":         {},
	"quote":        {},
	"randAlpha":    {},
	"randAlphaNum": {},
	"randAscii":    {},
	"randBytes":    "Accepts a count `N` and generates cryptographically secure random sequence of `N` bytes. The sequence is returned as a Base64 encoded string.",
	"randInt": {
		Description: "Returns a random integer value from min (inclusive) to max (exclusive).",
		Parameters:  []string{"Min value (inclusive).", "Max value (exclusive)."},
		Example:     " `randInt 12 30` will produce a random number in the range from 12 to 30.",
	},
	"randNumeric":            {},
	"regexFind":              {},
	"regexFindAll":           {},
	"regexMatch":             {},
	"regexQuoteMeta":         {},
	"regexReplaceAll":        {},
	"regexReplaceAllLiteral": {},
	"regexSplit":             {},
	"repeat":                 {},
	"replace":                {},
	"rest":                   "Gets tail of the list (everything but the first item).",
	"reverse":                "Produces a new list with the reversed elements of the given list.",
	"round":                  "Returns a float value with the remainder rounded to the given number to digits after the decimal point. `round 123.55555 3` will return `123.556`.",
	"semver":                 {},
	"semverCompare":          {},
	"seq":                    "Works like Bash `seq` command. Specify 1 parameter (`end`) to generate all counting integers between 1 and `end` inclusive. Specify 2 parameters (`start` and `end`) to generate all counting integers between `start` and `end` inclusive incrementing or decrementing by 1. Specify 3 parameters (`start`, `step` and `end`) to generate all counting integers between `start` and `end` inclusive incrementing or decrementing by `step`.",
	"set":                    {},
	"sha1sum":                {},
	"sha256sum":              {},
	"sha512sum":              {},
	"shuffle":                {},
	"slice":                  {},
	"snakecase":              {},
	"sortAlpha":              {},
	"split":                  {},
	"splitList":              {},
	"splitn":                 {},
	"squote":                 {},
	"sub":                    {},
	"subf":                   {},
	"substr":                 {},
	"swapcase":               {},
	"ternary":                "Takes two values and a test value. If the test value is true, the first value will be returned. If the test value is false, the second value will be returned. This is similar to the C ternary operator. `ternary \"foo\" \"bar\" true` or `true | \"foo\" \"bar\"` will return `\"foo\"`.",
	"title": {
		Description: "Converts a string to title case.",
		Parameters:  []string{"The string."},
		Example:     "`title \"hello world\"` returns `\"Hello World\"`.",
	},
	"toDate":       "Converts a string to a date. The first argument is the date layout and the second is the date string. If the string can’t be converted it returns the zero value.",
	"toDecimal":    "Converts a Unix octal to a `int64`.`\"0777\" | toDecimal` will convert `0777` to `511` and return the value as `int64`.",
	"toJson":       {},
	"toPrettyJson": {},
	"toRawJson":    "Encodes an item into JSON string with HTML characters unescaped. `toRawJson .Item` will return unescaped JSON string representation of `.Item`.",
	"toString":     "Converts to a string.",
	"toStrings":    "Converts a list, slice or array to a list of strings. `list 1 2 3 | toString` converts `1`, `2` and `3` to strings and then returns them as a list.",
	"trim": {
		Description: "Removes space from either side of a string.",
		Parameters:  []string{"The string."},
		Example:     "`trim \"  hello  \"` will return `hello`.",
	},
	"trimAll": {
		Description: "Removes given characters from the front or back of a string.",
		Parameters:  []string{"Character to remove", "The string."},
		Example:     "`trimAll \"$\" \"$5.00\"` will return `5.00` (as a string).",
	},
	"trimPrefix": {
		Description: "Trims just the prefix from a string.",
		Parameters:  []string{"Character to remove", "The string."},
		Example:     "`trimPrefix \"-\" \"-hello\"` will return `hello`.",
	},
	"trimSuffix": {
		Description: "Trims just the suffix from a string.",
		Parameters:  []string{"Character to remove", "The string."},
		Example:     "`trimSuffix \"-\" \"hello-\"` will return `hello`.",
	},
	"trimall":    {},
	"trunc":      {},
	"tuple":      {},
	"typeIs":     {},
	"typeIsLike": {},
	"typeOf":     {},
	"uniq": {
		Description: "Generates a list with all of the duplicates removed.",
		Parameters:  []string{"The list."},
	},
	"unixEpoch": {
		Description: "Returns the seconds since the Unix epoch for a time.",
		Parameters:  []string{"The time (`time.Time`)."},
		Example:     "`now | unixEpoch`",
	},
	"unset": {
		Description: "Given a map and a key it deletes the key from the map. It returns dictionary. Note that if the key is not found this operation will simply return. No error will be generated.",
		Parameters:  []string{"The map.", "The key of an item to delete."},
	},
	"until": {
		Description: "Builds a range of integers.",
		Parameters:  []string{"Max value (exclusive)."},
		Example:     "`until 5` will return a list `[0, 1, 2, 3, 4]`.",
	},
	"untilStep": {
		Description: "Like `until` generates a list of counting integers but it allows to define a start, stop and step.",
		Parameters:  []string{"Start value (inclusive)", "Max value (exclusive).", "Step."},
		Example:     "`untilStep 3 6 2` will return `[3, 5]` by starting with 3 and adding 2 until it is equal or greater than 6.",
	},
	"untitle": {
		Description: "Removes title casing.",
		Parameters:  []string{"The string."},
		Example:     "`untitle \"Hello World\"` returns `\"hello world\"`.",
	},
	"upper": {
		Description: "Converts the entire string to uppercase.",
		Parameters:  []string{"The string."},
		Example:     "`upper \"hello\"` returns `HELLO`",
	},
	"urlJoin": {
		Description: "Joins map produced by `urlParse` to produce URL string.",
		Parameters:  []string{"The map produced by `urlParse`."},
		Example:     "`urlJoin (dict \"fragment\" \"fragment\" \"host\" \"host:80\" \"path\" \"/path\" \"query\" \"query\" \"scheme\" \"http\")` returns `proto://host:80/path?query#fragment`.",
	},
	"urlParse": {
		Description: "Parses string for URL and produces dict with URL parts. For more info check https://golang.org/pkg/net/url/#URL.",
		Parameters:  []string{"The string with URL."},
	},
	"values": {
		Description: "Returns list of all values from a map.",
		Parameters:  []string{"The map."},
	},
	"without": {
		Description: "Filters items out of a list. It can take more than one filter.",
		Parameters:  []string{"Column count.", "The text."},
		Example:     "`without ( list 1 2 3 4 5) 1 3 5` returns `[2, 4]`.",
	},
	"wrap": {
		Description: "Wraps text at a given column count.",
		Parameters:  []string{"Column count.", "The text."},
		Example:     "`wrap 80 $text` will wrap the string in `$text` at 80 columns.",
	},
	"wrapWith": {
		Description: "Works as `wrap` but lets you specify the string to wrap with (`wrap` uses `\\n`).",
		Parameters:  []string{"Column count.", "String to wrap with.", "The text."},
		Example:     "`wrapWith 5 \"\\t\" \"Hello world\"` returns `hello world` (where the whitespace is an ASCII tab character).",
	},
}
