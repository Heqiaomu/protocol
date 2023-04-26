package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckNumbers(t *testing.T) {
	des := CheckNumbers(&ui.InputField{})
	assert.NotEmpty(t, des)

	v := &ui.ValidateNumbers{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumbers{Require: &wrappers.BoolValue{Value: true}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumbers{MinCount: 2, MaxCount: 10}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av, Value: []string{"a"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumbers{MinCount: 0, MaxCount: 10}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av, Value: []string{"a"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumbers{Min: "5", MinCount: 0, MaxCount: 10}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av, Value: []string{"1"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumbers{Max: "0", MinCount: 0, MaxCount: 10}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av, Value: []string{"1"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateNumbers{MinCount: 0, MaxCount: 10}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckNumbers(&ui.InputField{Validate: av, Value: []string{"1"}})
	assert.Empty(t, des)
}
