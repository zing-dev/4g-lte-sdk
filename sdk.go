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
	DefaultDllName  = "./dll/xxtSmsDll.dll"
	DefaultBaudRate = 115200

	FuncOpenModem  = "OpenModem"
	FuncSendSms    = "SendSms"
	FuncAutoDelMsg = "AutoDelMsg"
	FuncReadNewSms = "ReadNewSms"
	FuncReadSms    = "ReadSms"
	FuncCloseModem = "CloseModem"

	ContentMinLength = 5
	ContentMaxLength = 256
)

var (
	ErrIllegalMobileNumber  = errors.New("非法的手机号")
	ErrIllegalContentLength = errors.New("短信内容长度非法")
	ErrNonsupport           = errors.New("功能不支持")
)

type Proc struct {
	OpenModem  *syscall.LazyProc
	SendSms    *syscall.LazyProc
	AutoDelMsg *syscall.LazyProc
	ReadNewSms *syscall.LazyProc
	ReadSms    *syscall.LazyProc
	CloseModem *syscall.LazyProc
}

type Client struct {
	Com      int
	BaudRate int
	DllName  string
	LazyDLL  *syscall.LazyDLL
	Proc     Proc
}

func NewDefault(com int) *Client {
	return New(DefaultDllName, com)
}

func New(name string, com int) *Client {
	lazy := syscall.NewLazyDLL(name)
	return &Client{
		Com:      com,
		BaudRate: DefaultBaudRate,
		DllName:  name,
		LazyDLL:  lazy,
		Proc: Proc{
			OpenModem:  lazy.NewProc(FuncOpenModem),
			SendSms:    lazy.NewProc(FuncSendSms),
			AutoDelMsg: lazy.NewProc(FuncAutoDelMsg),
			ReadNewSms: lazy.NewProc(FuncReadNewSms),
			ReadSms:    lazy.NewProc(FuncReadSms),
			CloseModem: lazy.NewProc(FuncCloseModem),
		},
	}
}

func (client *Client) OpenModem() error {
	str, _ := syscall.BytePtrFromString("")
	ret, _, err := client.Proc.OpenModem.Call(uintptr(client.Com), uintptr(client.BaudRate), uintptr(unsafe.Pointer(str)))
	if ret == 1 {
		return nil
	}
	return err
}

func (client *Client) CloseModem() error {
	ret, _, err := client.Proc.CloseModem.Call()
	if ret == 1 {
		return nil
	}
	return err
}

func (client *Client) send(content, mobile *byte) error {
	ret, _, err := client.Proc.SendSms.Call(uintptr(4), uintptr(unsafe.Pointer(mobile)), uintptr(unsafe.Pointer(content)))
	if ret == 1 {
		return nil
	}
	return err
}

//todo
func (client *Client) ReadSms() error {
	return ErrNonsupport
}

//todo
func (client *Client) ReadNewSms() error {
	return ErrNonsupport
}

//todo
func (client *Client) AutoDelMsg() error {
	return ErrNonsupport
}

func (client *Client) SendSms(content, mobile string) error {
	if ContentMinLength > len(content) || len(content) > ContentMaxLength {
		return ErrIllegalContentLength
	}
	if !checkMobileNumber(mobile) {
		return ErrIllegalMobileNumber
	}
	m, _ := syscall.BytePtrFromString(mobile)
	c, _ := syscall.BytePtrFromString(content)
	return client.send(c, m)
}

func (client *Client) SendMoreSms(content string, mobiles ...string) {
	if ContentMinLength > len(content) || len(content) > ContentMaxLength {
		log.Println(fmt.Sprintf("Send [ %s ]:[ %s ] ", content, ErrIllegalContentLength))
	}
	c, _ := syscall.BytePtrFromString(content)
	for _, mobile := range mobiles {
		if !checkMobileNumber(mobile) {
			log.Println(fmt.Sprintf("%s:[ %s ]", ErrIllegalMobileNumber, mobile))
			continue
		}
		m, _ := syscall.BytePtrFromString(mobile)
		err := client.send(c, m)
		if err != nil {
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
