package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckNumber(t *testing.T) {
	des := CheckNumber(&ui.InputField{})
	assert.NotEmpty(t, des)

	v := &ui.ValidateNumber{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumber(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumber{Require: &wrappers.BoolValue{Value: true}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumber(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumber{}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumber(&ui.InputField{Validate: av, Value: []string{"a"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumber{Min: "5"}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumber(&ui.InputField{Validate: av, Value: []string{"1"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumber{Max: "0"}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumber(&ui.InputField{Validate: av, Value: []string{"1"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumber{}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumber(&ui.InputField{Validate: av, Value: []string{"1"}})
	assert.Empty(t, des)
}
