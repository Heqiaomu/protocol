package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckInputs(t *testing.T) {
	des := CheckInputs(&ui.InputField{})
	assert.NotEmpty(t, des)

	v := &ui.ValidateInputs{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInputs{Require: &wrappers.BoolValue{Value: true}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInputs{MinCount: 1, MaxCount: 1}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInputs{Regex: "*.", MinCount: 1, MaxCount: 1}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInputs{Regex: "^a$", MinCount: 1, MaxCount: 1}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInputs{Regex: "^a$", RegexDes: "des", MinCount: 1, MaxCount: 1}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInputs{Regex: "^[a-z]+$", MinCount: 1, MaxCount: 1}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInputs(&ui.InputField{Validate: av, Value: []string{"a"}})
	assert.Empty(t, des)
}
