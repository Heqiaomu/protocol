package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
)

// CheckInputs 校验Inputs类型
func CheckInputs(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vInputs := ui.ValidateInputs{}
	err := ptypes.UnmarshalAny(v, &vInputs)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateInputs, err=%v", err)
	}
	// 必填
	if vInputs.Require != nil && vInputs.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vInputs.RequireDes != "" {
				return vInputs.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}
	// 数量
	l := len(inputField.Value)
	if l < int(vInputs.MinCount) || l > int(vInputs.MaxCount) {
		return fmt.Sprintf("inputs count=%d out of range [%d, %d]", l, vInputs.MinCount, vInputs.MaxCount)
	}
	// 正则
	if vInputs.Regex != "" {
		for _, value := range inputField.Value {
			reg := vInputs.Regex
			matched, err := CheckRegex(reg, value)
			if err != nil {
				return fmt.Sprintf("cannot regexp pattern=%s to match value=%s", reg, value)
			}
			if !matched {
				if vInputs.RegexDes != "" {
					return vInputs.RegexDes
				} else {
					return fmt.Sprintf("value=%s doesn't match regex=%s", value, reg)
				}
			}
		}
	}
	return
}
