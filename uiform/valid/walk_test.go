package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/Heqiaomu/protocol/uiform/cons"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitInputToValid(t *testing.T) {
	err := VisitInputToValid(nil, nil)
	assert.Nil(t, err)
}

func TestVisitInputFieldToValid(t *testing.T) {
	err := VisitInputFieldToValid(&ui.InputField{}, nil)
	assert.Nil(t, err)

	v := &ui.ValidateInput{}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	ve := &ValidErrs{}

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, Title: &ui.Name{Text: "t"}, InputType: cons.Input}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.Inputs}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.NumberInput}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.NumberInputs}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.Switch}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.Select}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.File}, ve)
	assert.Nil(t, err)

	err = VisitInputFieldToValid(&ui.InputField{Validate: av, InputType: cons.Files}, ve)
	assert.Nil(t, err)
}
