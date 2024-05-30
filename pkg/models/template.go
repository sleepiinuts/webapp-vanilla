package models

import "net/url"

type Template struct {
	Data       map[string]any
	Form       url.Values
	FormErrors map[string][]string
	CSRFToken  string
	Flash      Flash
}

type Flash struct {
	Body string
	Type FlashType
}

func (f *Flash) BGColor() string {
	bg := map[FlashType]string{
		FTSuccess: "#18a999",
		FTFail:    "#EC4067",
		FTInfo:    "#CFCFCD",
	}
	return bg[f.Type]
}

type FlashType int

const (
	FTSuccess FlashType = iota + 1
	FTFail
	FTInfo
)
