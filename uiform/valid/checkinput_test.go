package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckInput(t *testing.T) {
	des := CheckInput(&ui.InputField{})
	assert.NotEmpty(t, des)

	v := &ui.ValidateInput{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInput(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInput{Require: &wrappers.BoolValue{Value: true}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInput(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInput{Regex: "*."}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInput(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInput{Regex: "^a$"}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInput(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInput{Regex: "^a$", RegexDes: "des"}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInput(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateInput{Regex: "^[a-z]+$"}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckInput(&ui.InputField{Validate: av, Value: []string{"a"}})
	assert.Empty(t, des)
}
