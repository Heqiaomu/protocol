package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSelect(t *testing.T) {
	des := CheckSelect(&ui.InputField{})
	assert.NotEmpty(t, des)

	v := &ui.ValidateSelect{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	av, err := ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSelect(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSelect{Require: &wrappers.BoolValue{Value: true}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSelect(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSelect{MinCount: 1, MaxCount: 2}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSelect(&ui.InputField{Validate: av})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSelect{MinCount: 1, MaxCount: 2, Options: []*ui.Name{{Id: "opt1"}}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSelect(&ui.InputField{Validate: av, Value: []string{""}})
	assert.NotEmpty(t, des)

	v = &ui.ValidateSelect{MinCount: 1, MaxCount: 2, Options: []*ui.Name{{Id: "opt1"}}}
	av, err = ptypes.MarshalAny(v)
	assert.Nil(t, err)
	des = CheckSelect(&ui.InputField{Validate: av, Value: []string{"opt1"}})
	assert.Empty(t, des)
}
