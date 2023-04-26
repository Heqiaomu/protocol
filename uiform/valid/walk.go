package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/Heqiaomu/protocol/uiform/cons"
)

type ValidErrs struct {
	Errs []*ValidErr
}

type ValidErr struct {
	ID    string
	Title string
	Desc  string
}

func VisitInputToValid(input *ui.Input, data interface{}) error {
	return nil
}

func VisitInputFieldToValid(inputField *ui.InputField, data interface{}) (err error) {
	// 不需要格式校验
	if inputField.Validate == nil {
		return
	}

	// 标题
	var title string
	if inputField.Title != nil && inputField.Title.Text != "" {
		title = inputField.Title.Text
	}
	// ID
	ID := inputField.Id

	// 错误返回结构体
	verrs := data.(*ValidErrs)
	ve := &ValidErr{
		ID:    ID,
		Title: title,
	}

	t := inputField.InputType
	// input
	if t == cons.Input || t == cons.Data || t == cons.Time || t == cons.DateTime || t == cons.IP || t == cons.TextArea || t == cons.Password {
		desc := CheckInput(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// inputs
	if t == cons.Inputs {
		desc := CheckInputs(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// ValidateNumber
	if t == cons.NumberInput || t == cons.NumberRange || t == cons.NumberSlide {
		desc := CheckNumber(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// ValidateNumbers
	if t == cons.NumberInputs {
		desc := CheckNumbers(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// ValidateSwitch
	if t == cons.Switch {
		desc := CheckSwitch(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// ValidateSelect
	if t == cons.Select || t == cons.Radio || t == cons.RadioButton {
		desc := CheckSelect(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// ValidateFile
	if t == cons.File {
		desc := CheckFile(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	// ValidateFiles
	if t == cons.Files {
		desc := CheckFiles(inputField)
		if desc != "" {
			ve.Desc = desc
			verrs.Errs = append(verrs.Errs, ve)
		}
		return
	}

	return nil
}
