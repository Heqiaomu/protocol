package walk

import "github.com/Heqiaomu/protocol/ui"

// WalkInput 遍历Input
func WalkInput(input *ui.Input, data interface{}, visitInput func(input *ui.Input, data interface{}) error, visitInputField func(inputField *ui.InputField, data interface{}) error) error {
	// 处理 Input节点 本身
	err := visitInput(input, data)
	if err != nil {
		return err
	}
	// 处理 Input节点 的 子Input
	for _, subInput := range input.SubInputs {
		err = WalkInput(subInput, data, visitInput, visitInputField)
		if err != nil {
			return err
		}
	}
	for _, inputField := range input.Fields {
		// 这里不需要访问inputField，因为walkInputField一进去会访问
		err = WalkInputField(inputField, data, visitInput, visitInputField)
		if err != nil {
			return err
		}
	}
	return nil
}

func WalkInputField(inputField *ui.InputField, data interface{}, visitInput func(input *ui.Input, data interface{}) error, visitInputField func(inputField *ui.InputField, data interface{}) error) error {
	// 处理 InputField节点 本身
	err := visitInputField(inputField, data)
	if err != nil {
		return err
	}
	for _, button := range inputField.Buttons {
		err = WalkInputField(button, data, visitInput, visitInputField)
		if err != nil {
			return err
		}
	}
	for _, react := range inputField.InputReacts {
		// 如果没有指定目标ID
		if react.TargetInputId == "" {
			for _, reactInput := range react.Inputs {
				// 进一步遍历 Input节点 的 子Input
				err = WalkInput(reactInput, data, visitInput, visitInputField)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
