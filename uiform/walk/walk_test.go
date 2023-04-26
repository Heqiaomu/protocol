package walk

import (
	"errors"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkInput(t *testing.T) {
	err := WalkInput(&ui.Input{}, nil, testVisitInputErr, testVisitField)
	assert.NotNil(t, err)

	err = WalkInput(&ui.Input{Id: "main", SubInputs: []*ui.Input{{}}}, nil, testVisitSubInputErr, testVisitField)
	assert.NotNil(t, err)

	err = WalkInput(&ui.Input{SubInputs: []*ui.Input{{}}, Fields: []*ui.InputField{{}}}, nil, testVisitInput, testVisitFieldErr)
	assert.NotNil(t, err)

	err = WalkInput(&ui.Input{SubInputs: []*ui.Input{{}}, Fields: []*ui.InputField{{}}}, nil, testVisitInput, testVisitField)
	assert.Nil(t, err)
}

func testVisitInputErr(input *ui.Input, data interface{}) error {
	return errors.New("err")
}

func testVisitSubInputErr(input *ui.Input, data interface{}) error {
	if input.Id != "main" {
		return errors.New("err")
	}
	return nil
}

func testVisitReactInputErr(input *ui.Input, data interface{}) error {
	if input.Id == "re" {
		return errors.New("err")
	}
	return nil
}

func testVisitInput(input *ui.Input, data interface{}) error {
	return nil
}

func testVisitFieldErr(inputField *ui.InputField, data interface{}) error {
	return errors.New("err")
}

func testVisitField(inputField *ui.InputField, data interface{}) error {
	return nil
}

func testVisitFieldButtonErr(inputField *ui.InputField, data interface{}) error {
	if inputField.Id == "bt" {
		return errors.New("err")
	}
	return nil
}

func TestWalkInputField(t *testing.T) {
	err := WalkInputField(&ui.InputField{}, nil, testVisitInput, testVisitFieldErr)
	assert.NotNil(t, err)

	err = WalkInputField(&ui.InputField{Buttons: []*ui.InputField{{Id: "bt"}}}, nil, testVisitInput, testVisitFieldButtonErr)
	assert.NotNil(t, err)

	err = WalkInputField(&ui.InputField{Buttons: []*ui.InputField{{}}, InputReacts: []*ui.InputReact{{Inputs: []*ui.Input{{Id: "re"}}}}}, nil, testVisitReactInputErr, testVisitField)
	assert.NotNil(t, err)

	err = WalkInputField(&ui.InputField{Buttons: []*ui.InputField{{}}, InputReacts: []*ui.InputReact{{Inputs: []*ui.Input{{}}}}}, nil, testVisitInput, testVisitField)
	assert.Nil(t, err)
}
