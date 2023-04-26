package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckFiles(t *testing.T) {
	des := CheckFiles(&ui.InputField{})
	assert.NotEmpty(t, des)

	vf := &ui.ValidateFiles{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	avf, err := ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFiles(&ui.InputField{Validate: avf})
	assert.NotEmpty(t, des)

	vf = &ui.ValidateFiles{Require: &wrappers.BoolValue{Value: true}}
	avf, err = ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFiles(&ui.InputField{Validate: avf})
	assert.NotEmpty(t, des)

	vf = &ui.ValidateFiles{FileNameSuffix: ".tar.gz"}
	avf, err = ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFiles(&ui.InputField{Validate: avf, Value: []string{""}})
	assert.NotEmpty(t, des)

	vf = &ui.ValidateFiles{FileNameSuffix: ".tar.gz"}
	avf, err = ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFiles(&ui.InputField{Validate: avf, Value: []string{"a.tar.gz"}})
	assert.Empty(t, des)
}
