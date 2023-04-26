package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"strconv"
)

// CheckSwitch 校验Switch类型
func CheckSwitch(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vSwitch := ui.ValidateSwitch{}
	err := ptypes.UnmarshalAny(v, &vSwitch)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateSwitch, err=%v", err)
	}

	// 必填
	if vSwitch.Require != nil && vSwitch.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vSwitch.RequireDes != "" {
				return vSwitch.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}

	// 判断是否是false或者true
	if len(inputField.Value) > 0 {
		value := inputField.Value[0]
		_, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Sprintf("cannot conv value=%s to bool. err=%v", value, err)
		}
	}

	return
}
