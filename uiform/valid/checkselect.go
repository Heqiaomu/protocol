package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
)

// CheckSelect 校验Select类型
func CheckSelect(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vSelect := ui.ValidateSelect{}
	err := ptypes.UnmarshalAny(v, &vSelect)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateSelect, err=%v", err)
	}

	// 必填
	if vSelect.Require != nil && vSelect.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vSelect.RequireDes != "" {
				return vSelect.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}

	// 数量
	l := len(inputField.Value)
	if l < int(vSelect.MinCount) || l > int(vSelect.MaxCount) {
		return fmt.Sprintf("inputs count=%d out of range [%d, %d]", l, vSelect.MinCount, vSelect.MaxCount)
	}

	// 选择框中的一项
	if l > 0 && len(vSelect.Options) > 0 {
		for _, value := range inputField.Value {
			isFind := false
			for _, name := range vSelect.Options {
				if name.Id == value {
					isFind = true
					break
				}
			}
			if !isFind {
				return fmt.Sprintf("value=%s is not in option", value)
			}
		}
	}
	return
}
