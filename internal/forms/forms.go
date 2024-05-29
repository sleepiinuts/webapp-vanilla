package forms

import "net/url"

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
			f.AddError(field, "This field is required")
		}
	}
}
