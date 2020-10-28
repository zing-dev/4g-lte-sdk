package lte_test

import (
	lte "github.com/zing-dev/4g-lte-sdk"
	"testing"
)

func TestCall(t *testing.T) {
	client := lte.NewDefault(8)
	err := client.OpenModem()
	if err != nil {
		t.Fatal(err)
	}
	err = client.SendSms("哈哈哈", "")
	if err != nil {
		t.Fatal(err)
	}
	err = client.CloseModem()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCalls(t *testing.T) {
	client := lte.NewDefault(8)
	err := client.OpenModem()
	if err != nil {
		t.Fatal(err)
	}
	client.SendMoreSms("this is a test sms", "", "", "")
}

//GOARCH=386
