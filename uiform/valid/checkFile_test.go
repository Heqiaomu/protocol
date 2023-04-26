package valid

import (
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckFile(t *testing.T) {
	des := CheckFile(&ui.InputField{})
	assert.NotEmpty(t, des)

	vf := &ui.ValidateFile{Require: &wrappers.BoolValue{Value: true}, RequireDes: "req"}
	avf, err := ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFile(&ui.InputField{Validate: avf})
	assert.NotEmpty(t, des)

	vf = &ui.ValidateFile{Require: &wrappers.BoolValue{Value: true}}
	avf, err = ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFile(&ui.InputField{Validate: avf})
	assert.NotEmpty(t, des)

	vf = &ui.ValidateFile{FileNameSuffix: ".tar.gz"}
	avf, err = ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFile(&ui.InputField{Validate: avf, Value: []string{""}})
	assert.NotEmpty(t, des)

	vf = &ui.ValidateFile{FileNameSuffix: ".tar.gz"}
	avf, err = ptypes.MarshalAny(vf)
	assert.Nil(t, err)
	des = CheckFile(&ui.InputField{Validate: avf, Value: []string{"a.tar.gz"}})
	assert.Empty(t, des)

}
