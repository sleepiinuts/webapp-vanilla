package forms

import (
	"net/mail"
	"net/url"
)

const (
	InvalidRequired    = "This field is required"
	InValidEmailFormat = "Wrong Format Email"
)

type Form struct {
	data   url.Values
	errors map[string][]string
}

func New(data url.Values) *Form {
	return &Form{
		data: data, errors: make(map[string][]string),
	}
}

func (f *Form) IsValid() bool {
	return len(f.errors) == 0
}

func (f *Form) AddError(field, errMsg string) {
	f.errors[field] = append(f.errors[field], errMsg)
}

func (f *Form) Require(fields ...string) {
	for _, field := range fields {
		if !f.data.Has(field) || f.data.Get(field) == "" {
			f.AddError(field, InvalidRequired)
		}
	}
}

func (f *Form) IsValidEmail(fields ...string) {
	for _, field := range fields {
		if _, err := mail.ParseAddress(f.data.Get(field)); err != nil {
			f.AddError(field, InValidEmailFormat)
		}
	}
}

func (f *Form) GetErrors() map[string][]string {
	return f.errors
}

func (f *Form) GetField(name string) string {
	return f.data.Get(name)
}
