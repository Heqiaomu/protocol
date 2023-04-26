package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"strconv"
)

// CheckNumbers 校验Numbers类型
func CheckNumbers(inputField *ui.InputField) (desc string) {
	v := inputField.Validate
	vNumbers := ui.ValidateNumbers{}
	err := ptypes.UnmarshalAny(v, &vNumbers)
	if err != nil {
		return fmt.Sprintf("cannot conv valid any to ui.ValidateNumbers, err=%v", err)
	}

	// 必填
	if vNumbers.Require != nil && vNumbers.Require.Value {
		if len(inputField.Value) <= 0 || inputField.Value[0] == "" {
			if vNumbers.RequireDes != "" {
				return vNumbers.RequireDes
			} else {
				return fmt.Sprintf("cannot be empty")
			}
		}
	}

	// 数量
	l := len(inputField.Value)
	if l < int(vNumbers.MinCount) || l > int(vNumbers.MaxCount) {
		return fmt.Sprintf("numbers count=%d out of range [%d, %d]", l, vNumbers.MinCount, vNumbers.MaxCount)
	}

	for _, value := range inputField.Value {
		// 数值
		vf, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Sprintf("cannot conv value=%s to number", value)
		}

		// 范围
		if vNumbers.Min != "" {
			valueMin := vNumbers.Min
			vfMin, err := strconv.ParseFloat(valueMin, 64)
			if err == nil {
				if vf < vfMin {
					return fmt.Sprintf("value=%s is less than min=%s", value, valueMin)
				}
			}
		}
		if vNumbers.Max != "" {
			valueMax := vNumbers.Max
			vfMax, err := strconv.ParseFloat(valueMax, 64)
			if err == nil {
				if vf > vfMax {
					return fmt.Sprintf("value=%s is greater than max=%s", value, valueMax)
				}
			}
		}
	}

	return
}
