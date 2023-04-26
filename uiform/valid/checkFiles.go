package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"strings"
)

// CheckFiles 校验Files类型
func CheckFiles(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vFiles := ui.ValidateFiles{}
	err := ptypes.UnmarshalAny(v, &vFiles)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateFiles, err=%v", err)
	}

	// 必填
	if vFiles.Require != nil && vFiles.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vFiles.RequireDes != "" {
				return vFiles.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}

	//l := len(inputField.Value)
	//lFN := len(vFiles.FileName)
	//if l != lFN {
	//	return fmt.Sprintf("fileName length=%d != file length=%d", lFN, l)
	//}
	//
	//if l < int(vFiles.MinCount) {
	//	return fmt.Sprintf("file count=%d is less than min=%d", l, vFiles.MinCount)
	//}
	//
	//if l > int(vFiles.MaxCount) {
	//	return fmt.Sprintf("file count=%d is greater than min=%d", l, vFiles.MaxCount)
	//}

	for _, value := range inputField.Value {
		//fileName := vFiles.FileName[i]
		//// 正则校验
		//if vFiles.Regex != "" {
		//	reg := vFiles.Regex
		//	matched, err := CheckRegex(reg, value)
		//	if err != nil {
		//		return fmt.Sprintf("cannot regexp pattern=%s to match file content which fileName=%s", reg, fileName)
		//	}
		//	if !matched {
		//		if vFiles.RegexDes != "" {
		//			return vFiles.RegexDes
		//		} else {
		//			return fmt.Sprintf("file content doesn't match regex=%s, which fileName=%s", reg, fileName)
		//		}
		//	}
		//}
		//// 文件大小
		//vb := []byte(value)
		//vMin, err := strconv.Atoi(vFiles.Min)
		//if err != nil {
		//	return fmt.Sprintf("min=%s cannot conv to int", vFiles.Min)
		//}
		//if len(vb) < vMin {
		//	fmt.Sprintf("file content len is less than min=%d, which fileName=%s", vMin, fileName)
		//}
		//vMax, err := strconv.Atoi(vFiles.Max)
		//if err != nil {
		//	return fmt.Sprintf("max=%s cannot conv to int", vFiles.Max)
		//}
		//if len(vb) > vMax {
		//	fmt.Sprintf("file content len is greater than max=%d, which fileName=%s", vMax, fileName)
		//}

		// 文件名
		if vFiles.FileNameSuffix != "" {
			// 对每个文件名检查
			//for _, fn := range vFiles.FileName {
			hasSuf := false
			fss := strings.Split(vFiles.FileNameSuffix, ",")
			for _, fs := range fss {
				if strings.HasSuffix(value, fs) {
					hasSuf = true
					break
				}
			}
			if !hasSuf {
				return fmt.Sprintf("fileName=%s doesn't has suffix=%s", value, vFiles.FileNameSuffix)
			}
			//}
		}

		//if vFiles.FileNameRegex != "" {
		//	for _, value := range vFiles.FileName {
		//		reg := vFiles.FileNameRegex
		//		matched, err := CheckRegex(reg, value)
		//		if err != nil {
		//			return fmt.Sprintf("cannot regexp pattern=%s to match fileName=%s", reg, fileName)
		//		}
		//		if !matched {
		//			if vFiles.RegexDes != "" {
		//				return vFiles.RegexDes
		//			} else {
		//				return fmt.Sprintf("fileName=%s doesn't match regex=%s", fileName, reg)
		//			}
		//		}
		//	}
		//}
	}

	return
}
