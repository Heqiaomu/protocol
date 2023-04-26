package valid

import (
	"regexp"
	"strings"
)

// CheckRegex 检验正则
func CheckRegex(reg string, value string) (matched bool, err error) {
	reg = strings.Replace(reg, `\u4e00`, "\u4e00", -1)
	reg = strings.Replace(reg, `\u9fa5`, "\u9fa5", -1)

	return regexp.MatchString(reg, value)
}
