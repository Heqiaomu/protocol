package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSwitch(t *testing.T) {
	des := CheckSwitch(&ui.InputField{})
	assert.NotEmpty(t, des)

	v := &ui.ValidateSwitch{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSwitch(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSwitch{Require: &wrappers.BoolValue{Value: true}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSwitch(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSwitch{}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSwitch(&ui.InputField{Validate: av, Value: []string{"abc"}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSwitch{}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSwitch(&ui.InputField{Validate: av, Value: []string{"true"}})
	assert.Empty(t, des)
}
