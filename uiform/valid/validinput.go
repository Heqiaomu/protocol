package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/Heqiaomu/protocol/uiform/walk"
)

// 检查表单的错误
func Check(input *ui.Input) (errs *ValidErrs) {
	errs = &ValidErrs{
		Errs: make([]*ValidErr, 0),
	}
	_ = walk.WalkInput(input, errs, VisitInputToValid, VisitInputFieldToValid)
	return
}
