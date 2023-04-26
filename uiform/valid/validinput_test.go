package valid

import (
	"fmt"
	"github.com/Heqiaomu/protocol/ui"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

func TestCheckCreateK8s(t *testing.T) {
	inputJson, err := ioutil.ReadFile("./test_data/createk8s.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var uiInput ui.Input
	err = jsonpb.UnmarshalString(string(inputJson), &uiInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	errs := Check(&uiInput)
	fmt.Println(len(errs.Errs) == 0)
	fmt.Println(fmt.Sprintf("%v", errs))
}

func TestCheckCreateVB(t *testing.T) {
	inputJson, err := ioutil.ReadFile("./test_data/createvb.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var uiInput ui.Input
	err = jsonpb.UnmarshalString(string(inputJson), &uiInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	errs := Check(&uiInput)
	fmt.Println(len(errs.Errs) == 0)
	fmt.Println(fmt.Sprintf("%v", errs))
}

func TestCheckCreateVBErr(t *testing.T) {
	inputJson, err := ioutil.ReadFile("./test_data/createvb-err.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var uiInput ui.Input
	err = jsonpb.UnmarshalString(string(inputJson), &uiInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	errs := Check(&uiInput)
	fmt.Println(len(errs.Errs) == 1)
	fmt.Println(fmt.Sprintf("%v", errs))
}

func TestReg(t *testing.T) {
	reg := "^[a-zA-Z0-9\\u4e00-\\u9fa5]{2,20}$"
	reg = strings.Replace(reg, `\u4e00`, "\u4e00", -1)
	reg = strings.Replace(reg, `\u9fa5`, "\u9fa5", -1)
	value := "韩hpcvv1"
	matched, err := regexp.MatchString(reg, value)
	if err != nil {
		fmt.Println(fmt.Sprintf("cannot regexp pattern=%s to match value=%s. err=%v", reg, value, err))
	}
	fmt.Println(matched)
}

func TestReg2(t *testing.T) {
	reg := "^[0-9a-zA-Z\\-+=/\\s]{20,1000}[0-9a-zA-Z\\-+=/\\s]{0,1000}$"
	value := `-----BEGIN CERTIFICATE-----
MIIDtjCCAp6gAwIBAgIUJro+NQhyPh8cd6XbUGzrdQiV4w0wDQYJKoZIhvcNAQEL
BQAwYDELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JFSUpJTkcxCzAJBgNVBAcTAlhT
MQwwCgYDVQQKEwNrOHMxDzANBgNVBAsTBlN5c3RlbTETMBEGA1UEAxMKa3ViZXJu
ZXRlczAgFw0yMDA3MDEwMjQxMDBaGA8yMTIwMDYwNzAyNDEwMFowYDELMAkGA1UE
BhMCQ04xEDAOBgNVBAgTB0JFSUpJTkcxCzAJBgNVBAcTAlhTMQwwCgYDVQQKEwNr
OHMxDzANBgNVBAsTBlN5c3RlbTETMBEGA1UEAxMKa3ViZXJuZXRlczCCASIwDQYJ
KoZIhvcNAQEBBQADggEPADCCAQoCggEBAJ2RPZ8ZkgEfVOQ440yQW6eStdEnuwu7
6U5AztZAS6418W5ys+DtjO6D6nIPGWVvVCujBql6IGxpFws/AI/WfkF1TYDw49F2
uVZAgEar/1xAeX+zf5lahyjOt7jGOR/w190yqKF1BWbrgoWv3yQ1b7NRV4lfaJUJ
IrIKXKC2rC5P/XUKS2HFgC8TXpGyFfGR3BnEbH6DPCzK7OCj5owDNtPZqeHxW08u
0pEvunZNaT9hlM4T+43N76Ohi6GJOvmthvpEM7KjrWXm9gu4irr2wmsjAPe788zj
/R61+6BQazgNsPX0AFelrUpOIKGBUgftSuYXOihjyO3fCronAq+VJa8CAwEAAaNm
MGQwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8CAQIwHQYDVR0OBBYE
FJxiwQ+jzldNOvmlYHJxvkaONdVzMB8GA1UdIwQYMBaAFJxiwQ+jzldNOvmlYHJx
vkaONdVzMA0GCSqGSIb3DQEBCwUAA4IBAQBMRV1yDtRR+G6/+OrM8a7MA8USQwuy
j9Ezq4X0rJBF3RBZ0uPSNO9ykii3/NvYDIxl/i2xKqEULSBaQpD2Ua4Oextx6+2Q
q4jWH+acZ1SNSq824xCDcJNkTovlYp3yRQdzKjZ+FBiRiHKjhtfimQpAi+uyQqpH
3fQ+pTKtBCoyxbwtCzGeVFXBXAv7U8TMSzHzSI1r/R89a4s4+1WObhtLN8A/Oxtf
NATwfsjtjdhxorJ1z3wTklStESVdLRsePeoix0WK4Za0Jp/KlOowYdrlV83FrYsW
cuGL1WfrndvFgVzc88FUr104D6X7xgUq7bv6wNCHMCqcwj7Fa9EHV/EJ
-----END CERTIFICATE-----`
	matched, err := regexp.MatchString(reg, value)
	if err != nil {
		fmt.Println(fmt.Sprintf("cannot regexp pattern=%s to match value=%s. err=%v", reg, value, err))
	}
	fmt.Println(matched)
}

func TestReg3(t *testing.T) {
	reg := "^https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)$"
	value := `https://172.16.1.121:6443`
	matched, err := regexp.MatchString(reg, value)
	if err != nil {
		fmt.Println(fmt.Sprintf("cannot regexp pattern=%s to match value=%s. err=%v", reg, value, err))
	}
	fmt.Println(matched)
}

func TestReg4(t *testing.T) {
	reg := "^[a-zA-Z0-9.\\-_:]{10,2000}$"
	value := `https://172.16.1.121:6443`
	matched, err := regexp.MatchString(reg, value)
	if err != nil {
		fmt.Println(fmt.Sprintf("cannot regexp pattern=%s to match value=%s. err=%v", reg, value, err))
	}
	fmt.Println(matched)
}
