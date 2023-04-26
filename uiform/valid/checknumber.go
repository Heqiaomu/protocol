package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"strconv"
)

// CheckNumber 校验Number类型
func CheckNumber(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vNumber := ui.ValidateNumber{}
	err := ptypes.UnmarshalAny(v, &vNumber)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateNumber, err=%v", err)
	}

	// 必填
	if vNumber.Require != nil && vNumber.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vNumber.RequireDes != "" {
				return vNumber.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}

	// 数值
	value := inputField.Value[0]
	vf, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Sprintf("cannot conv value=%s to number", value)
	}

	// 范围
	if vNumber.Min != "" {
		valueMin := vNumber.Min
		vfMin, err := strconv.ParseFloat(valueMin, 64)
		if err == nil {
			if vf < vfMin {
				return fmt.Sprintf("value=%s is less than min=%s", value, valueMin)
			}
		}
	}
	if vNumber.Max != "" {
		valueMax := vNumber.Max
		vfMax, err := strconv.ParseFloat(valueMax, 64)
		if err == nil {
			if vf > vfMax {
				return fmt.Sprintf("value=%s is greater than max=%s", value, valueMax)
			}
		}
	}

	return
}
