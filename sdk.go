package lte

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"syscall"
	"time"
	"unsafe"
)

const (
	DllName = "./dll/xxtSmsDll.dll"

	FuncOpenModem = "OpenModem"
	FuncSendSms   = "SendSms"
)

var (
	lazy     = syscall.NewLazyDLL(DllName)
	procOpen = lazy.NewProc(FuncOpenModem)
	procSend = lazy.NewProc(FuncSendSms)

	ErrIllegalMobileNumber = errors.New("非法的手机号")
)

func OpenModem(com, baudRate int) error {
	str, _ := syscall.BytePtrFromString("")
	ret, _, err := procOpen.Call(uintptr(com), uintptr(baudRate), uintptr(unsafe.Pointer(str)))
	if ret == 1 {
		return nil
	}
	return err
}

func SendSms(content, mobile string) error {
	if !checkMobileNumber(mobile) {
		return ErrIllegalMobileNumber
	}
	m, _ := syscall.BytePtrFromString(mobile)
	c, _ := syscall.BytePtrFromString(content)
	ret, _, err := procSend.Call(uintptr(8), uintptr(unsafe.Pointer(m)), uintptr(unsafe.Pointer(c)))
	if ret == 1 {
		return nil
	}
	return err
}

func SendMoreSms(content string, mobiles ...string) {
	c, _ := syscall.BytePtrFromString(content)
	for _, mobile := range mobiles {
		if !checkMobileNumber(mobile) {
			log.Println(fmt.Sprintf("%s:[ %s ]", ErrIllegalMobileNumber, mobile))
			continue
		}
		m, _ := syscall.BytePtrFromString(mobile)
		ret, _, err := procSend.Call(uintptr(8), uintptr(unsafe.Pointer(m)), uintptr(unsafe.Pointer(c)))
		if ret != 1 {
			log.Println(fmt.Sprintf("Send [ %s ]:[ %s ] ", mobile, err))
		}
		time.Sleep(time.Second * 3)
	}
}

//验证手机号
func checkMobileNumber(mobile string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	return regexp.MustCompile(pattern).MatchString(mobile)
}
