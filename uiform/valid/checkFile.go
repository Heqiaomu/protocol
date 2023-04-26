package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"strings"
)

// CheckFile 校验File类型
func CheckFile(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vFile := ui.ValidateFile{}
	err := ptypes.UnmarshalAny(v, &vFile)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateFile, err=%v", err)
	}

	// 必填
	if vFile.Require != nil && vFile.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vFile.RequireDes != "" {
				return vFile.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}

	//l := len(inputField.Value)
	//lFN := len(vFile.FileName)
	//if l != lFN {
	//	return fmt.Sprintf("fileName length=%d != file length=%d", lFN, l)
	//}

	// 注意现在value就是文件名
	for _, value := range inputField.Value {
		//fileName := vFile.FileName[i]
		//// 正则校验
		//if vFile.Regex != "" {
		//	reg := vFile.Regex
		//	matched, err := regexp.MatchString(reg, value)
		//	if err != nil {
		//		return fmt.Sprintf("cannot regexp pattern=%s to match file content which fileName=%s. err=%v", reg, fileName, err)
		//	}
		//	if !matched {
		//		if vFile.RegexDes != "" {
		//			return vFile.RegexDes
		//		} else {
		//			return fmt.Sprintf("file content doesn't match regex=%s, which fileName=%s", reg, fileName)
		//		}
		//	}
		//}
		// 文件大小
		//vb := []byte(value)
		//vMin, err := strconv.Atoi(vFile.Min)
		//if err != nil {
		//	return fmt.Sprintf("min=%s cannot conv to int", vFile.Min)
		//}
		//if len(vb) < vMin {
		//	fmt.Sprintf("file content len is less than min=%d, which fileName=%s", vMin, fileName)
		//}
		//vMax, err := strconv.Atoi(vFile.Max)
		//if err != nil {
		//	return fmt.Sprintf("max=%s cannot conv to int", vFile.Max)
		//}
		//if len(vb) > vMax {
		//	fmt.Sprintf("file content len is greater than max=%d, which fileName=%s", vMax, fileName)
		//}

		// 文件名
		if vFile.FileNameSuffix != "" {
			// 对每个文件名检查
			//for _, fn := range vFile.FileName {
			hasSuf := false
			fss := strings.Split(vFile.FileNameSuffix, ",")
			for _, fs := range fss {
				if strings.HasSuffix(value, fs) {
					hasSuf = true
					break
				}
			}
			if !hasSuf {
				return fmt.Sprintf("fileName=%s doesn't has suffix=%s", value, vFile.FileNameSuffix)
			}
			//}
		}

		//if vFile.FileNameRegex != "" {
		//	for _, value := range vFile.FileName {
		//		reg := vFile.FileNameRegex
		//		matched, err := CheckRegex(reg, value)
		//		if err != nil {
		//			return fmt.Sprintf("cannot regexp pattern=%s to match fileName=%s. err=%v", reg, fileName, err)
		//		}
		//		if !matched {
		//			if vFile.RegexDes != "" {
		//				return vFile.RegexDes
		//			} else {
		//				return fmt.Sprintf("fileName=%s doesn't match regex=%s", fileName, reg)
		//			}
		//		}
		//	}
		//}
	}

	return
}
