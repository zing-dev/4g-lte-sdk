package lte_test

import (
	lte "github.com/zing-dev/4g-lte-sdk"
	"testing"
)

func TestCall(t *testing.T) {
	err := lte.OpenModem(8, 115200)
	if err != nil {
		t.Fatal(err)
	}
	err = lte.SendSms("this is a test sms", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCalls(t *testing.T) {
	err := lte.OpenModem(8, 115200)
	if err != nil {
		t.Fatal(err)
	}
	lte.SendMoreSms("this is a test sms", "", "", "")
}

//GOARCH=386;CGO_ENABLED=1
