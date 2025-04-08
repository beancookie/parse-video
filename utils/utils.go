package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func RegexpMatchUrlFromString(str string) (string, error) {
	urlReg, err := regexp.Compile(`https?://[\w.-]+[\w/-]*[\w.-:]*\??[\w=&:\-+%.]*/*`)
	if err != nil {
		return "", fmt.Errorf("match url regexp compile error: %s", err.Error())
	}

	findStr := urlReg.FindString(str)
	if len(findStr) <= 0 {
		return "", fmt.Errorf("str not have url")
	}

	return findStr, nil
}

// ConvertToHttps 将URL从http转换为https
func ConvertToHttps(url string) string {
	if strings.HasPrefix(url, "http://") {
		return "https://" + strings.TrimPrefix(url, "http://")
	}
	return url
}
