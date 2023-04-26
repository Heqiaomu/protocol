package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
)

// CheckInput 校验Input类型
func CheckInput(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vInput := ui.ValidateInput{}
	err := ptypes.UnmarshalAny(v, &vInput)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateInput, err=%v", err)
	}
	// 必填
	if vInput.Require != nil && vInput.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vInput.RequireDes != "" {
				return vInput.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}
	// 正则校验
	if vInput.Regex != "" {
		if len(inputField.Value) > 0 {
			value := inputField.Value[0]
			reg := vInput.Regex
			matched, err := CheckRegex(reg, value)
			if err != nil {
				return fmt.Sprintf("cannot regexp pattern=%s to match value=%s. err=%v", reg, value, err)
			}
			if !matched {
				if vInput.RegexDes != "" {
					return vInput.RegexDes
				} else {
					return fmt.Sprintf("value=%s doesn't match regex=%s", value, reg)
				}
			}
		}
	}
	return
}
