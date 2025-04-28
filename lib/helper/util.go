package helper

import (
	"fmt"
	"io"
	"kredi-plus.com/be/config"
	"kredi-plus.com/be/lib/exception"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
	"unicode"
)

func CheckDataOnSlice[T string | int | uint](dataMustHave T, arrData []T) bool {
	for _, data := range arrData {
		if dataMustHave == data {
			return true
		}
	}
	return false
}

func CheckDataOnSliceWithFunc[T string | int | uint](dataMustHave T, arrData []T, checkFn func(T, T) bool) bool {
	for _, value := range arrData {
		if checkFn != nil {
			if checkFn(dataMustHave, value) {
				return true
			}
		} else {
			if value == dataMustHave {
				return true
			}
		}
	}
	return false
}

func ProcessFileUpload(file *multipart.FileHeader, functionName, fileName string) (pathHttps string, err error) {
	var pathServer = `./uploads/` + functionName + `/`

	hostname := getHostname()
	if hostname == "" {
		err = exception.InternalServerError
		return
	}

	var Https = hostname + "/img/" + functionName + "/"

	//fungsi untuk mengambil fileUpload
	src, err := file.Open()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return "", err
	}
	defer src.Close()

	path := pathServer + strings.ReplaceAll(fileName, " ", "_")
	https := Https + url.QueryEscape(fileName)
	//path := file.Filename
	dst, err := os.Create(path)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return "", err
	}
	defer dst.Close()

	//copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return "", err
	}

	return https, nil
}

func getHostname() string {
	// todo switch hostname berdasarkan env staging, local, atau prod
	return fmt.Sprintf("http://localhost:%s", config.Attr.App.Port)
}

func ConvertDateFormat(format string) string {
	type DataDateFormat struct {
		Old string
		New string
	}

	replacements := []DataDateFormat{
		{"YYYY", "2006"},
		{"YYY", "006"},
		{"YY", "06"},
		{"MMMMM", "J"},
		{"MMM", "Jan"},
		{"MM", "01"},
		{"M", "1"},
		{"DDDD", "Monday"},
		{"DDD", "Mon"},
		{"DD", "02"},
		{"D", "2"},
		{"hh24", "15"},
		{"hh", "03"},
		{"mm", "04"},
		{"ss", "05"},
		{"ms", ".999"},
		{"ps", ".999999"},
		{"ns", ".999999999"},
		{"AMPM", "PM"},
		{"TZ", "MST"},
		{"Z", "Z07:00"},
		{"OFF", "-07"},
		{"TT", "PM"},
	}

	for _, val := range replacements {
		format = strings.Replace(format, val.Old, val.New, -1)
	}

	return format

}

func CapitalizedEachWords(input string) string {
	var result, finalRes strings.Builder
	runesInput := []rune(input)

	for i, r := range input {
		if unicode.IsUpper(r) && i > 0 {
			prev := i - 1
			next := i + 1

			if len(runesInput) == next {
				next = i
			}
			if unicode.IsLower(runesInput[prev]) {
				result.WriteRune(' ')
			} else if unicode.IsUpper(runesInput[prev]) && (unicode.IsLower(runesInput[next]) || unicode.IsNumber(runesInput[next])) {
				result.WriteRune(' ')
			} else if unicode.IsNumber(runesInput[prev]) {
				result.WriteRune(' ')
			}
		}
		if i == 0 || (i > 0 && unicode.IsUpper(r)) {
			result.WriteRune(unicode.ToUpper(r))
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	input = result.String()
	runesInput = []rune(input)

	for i, r := range input {
		if i == 0 {
			finalRes.WriteRune(r)
		}

		if i > 0 {
			prev := i - 1

			if unicode.IsSpace(runesInput[prev]) {
				finalRes.WriteRune(unicode.ToUpper(r))
			} else {
				finalRes.WriteRune(r)
			}
		}
	}

	arrResult := strings.Split(finalRes.String(), " ")
	for i := range arrResult {
		if CheckDataOnSliceWithFunc(arrResult[i], []string{"for"}, strings.EqualFold) {
			arrResult[i] = strings.ToLower(arrResult[i])
		}
	}

	return strings.Join(arrResult, " ")
}
