package env

import (
	"bufio"
	"errors"
	"github.com/jellycheng/gosupport"
	"io"
	"os"
	"regexp"
	"strings"
)


const (
	//特殊字符
	doubleQuoteSpecialChars = "\\\n\r\"!$`"
	//env存储方式
	StoreTypeEnv = iota
	StoreTypeDataManage
)

var (
	//export OPTION_A=2
	exportRegex = regexp.MustCompile(`^\s*(?:export\s+)?(.*?)\s*$`)
	singleQuotesRegex  = regexp.MustCompile(`\A'(.*)'\z`)
	doubleQuotesRegex  = regexp.MustCompile(`\A"(.*)"\z`)
	escapeRegex        = regexp.MustCompile(`\\.`)
	unescapeCharsRegex = regexp.MustCompile(`\\([^$])`)
	expandVarRegex = regexp.MustCompile(`(\\)?(\$)(\()?\{?([A-Z0-9_]+)?\}?`)

	currentEnv = map[string]bool{}
	currentDM = map[string]bool{}
)

//加载.env文件，支持多个env文件，但不覆盖已经在的key值
func LoadEnv(filenames ...string) (err error)  {
	return load(false, StoreTypeEnv, filenames...)
}

//加载.env文件，支持多个env文件，会覆盖已经在的key值
func Overload(filenames ...string) (err error) {
	return load(true, StoreTypeEnv, filenames...)
}

func LoadEnv2DataManage(filenames ...string) (err error)  {
	return load(false, StoreTypeDataManage, filenames...)
}

func load(isOverload bool, storeType int, filenames ...string) (err error) {
	filenames = filenamesOrDefault(filenames)
	for _, filename := range filenames {//遍历文件
		err = loadFile(filename, isOverload, storeType)
		if err != nil {
			return
		}
	}
	return
}

func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}

//加载文件
func loadFile(filename string, overload bool,storeType int) error {
	envMap, err := readFile(filename) //读文件并分析内容
	if err != nil {
		return err
	}

	if storeType == StoreTypeEnv {
		rawEnv := os.Environ()
		for _, rawEnvLine := range rawEnv {
			key := strings.Split(rawEnvLine, "=")[0]
			currentEnv[key] = true
		}
		for key, value := range envMap {
			if !currentEnv[key] || overload {
				Set(key, value)
				currentEnv[key] = true
			}
		}

	} else if storeType == StoreTypeDataManage {
		for key, value := range envMap {
			if !currentDM[key] || overload {
				globalenv := gosupport.NewGlobalEnvSingleton()
				globalenv.Set(key, value)
				currentDM[key] = true
			}
		}
	}

	return nil
}
//读文件内容
func readFile(filename string) (envMap map[string]string, err error) {
	isFile := gosupport.IsFile(filename)
	if !isFile {
		err = errors.New(filename + "文件不存在")
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	//分析文件
	return Parse(file)
}

func Parse(r io.Reader) (envMap map[string]string, err error) {
	envMap = make(map[string]string)

	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return
	}

	for _, fullLine := range lines {
		if !isIgnoredLine(fullLine) {
			var key, value string
			key, value, err = parseLine(fullLine, envMap)

			if err != nil {
				return
			}
			envMap[key] = value
		}
	}
	return
}

func parseLine(line string, envMap map[string]string) (key string, value string, err error) {
	if len(line) == 0 {
		err = errors.New("空字符串")
		return
	}

	//包含注释符号
	if strings.Contains(line, "#") {
		segmentsBetweenHashes := strings.Split(line, "#")
		quotesAreOpen := false
		var segmentsToKeep []string
		for _, segment := range segmentsBetweenHashes {
			if strings.Count(segment, "\"") == 1 || strings.Count(segment, "'") == 1 {
				if quotesAreOpen {
					quotesAreOpen = false
					segmentsToKeep = append(segmentsToKeep, segment)
				} else {
					quotesAreOpen = true
				}
			}

			if len(segmentsToKeep) == 0 || quotesAreOpen {
				segmentsToKeep = append(segmentsToKeep, segment)
			}
		}

		line = strings.Join(segmentsToKeep, "#")
	}

	firstEquals := strings.Index(line, "=")
	firstColon := strings.Index(line, ":")
	splitString := strings.SplitN(line, "=", 2)
	if firstColon != -1 && (firstColon < firstEquals || firstEquals == -1) {
		splitString = strings.SplitN(line, ":", 2)
	}

	if len(splitString) != 2 {
		err = errors.New("Can't separate key from value")
		return
	}

	key = splitString[0]
	if strings.HasPrefix(key, "export") {
		key = strings.TrimPrefix(key, "export")
	}
	key = strings.TrimSpace(key)

	key = exportRegex.ReplaceAllString(splitString[0], "$1")

	value = parseValue(splitString[1], envMap)
	return
}


func parseValue(value string, envMap map[string]string) string {

	//去掉空格
	value = strings.Trim(value, " ")
	if len(value) > 1 {
		singleQuotes := singleQuotesRegex.FindStringSubmatch(value)
		doubleQuotes := doubleQuotesRegex.FindStringSubmatch(value)

		if singleQuotes != nil || doubleQuotes != nil {
			value = value[1 : len(value)-1]
		}

		if doubleQuotes != nil {
			value = escapeRegex.ReplaceAllStringFunc(value, func(match string) string {
				c := strings.TrimPrefix(match, `\`)
				switch c {
				case "n":
					return "\n"
				case "r":
					return "\r"
				default:
					return match
				}
			})
			value = unescapeCharsRegex.ReplaceAllString(value, "$1")
		}

		if singleQuotes == nil {
			value = expandVariables(value, envMap)
		}
	}

	return value
}


func expandVariables(v string, m map[string]string) string {
	return expandVarRegex.ReplaceAllStringFunc(v, func(s string) string {
		submatch := expandVarRegex.FindStringSubmatch(s)

		if submatch == nil {
			return s
		}
		if submatch[1] == "\\" || submatch[2] == "(" {
			return submatch[0][1:]
		} else if submatch[4] != "" {
			return m[submatch[4]]
		}
		return s
	})
}


//判断是否空字符串或以#开头的字符串
func isIgnoredLine(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return len(trimmedLine) == 0 || strings.HasPrefix(trimmedLine, "#")
}

func Get(k string, defaultVal string) (string) {
	ret := os.Getenv(k)
	if ret == "" {
		return defaultVal
	}
	return ret
}

func Set(k,v string) error {
	return os.Setenv(k, v)
}

