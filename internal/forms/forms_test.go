package forms

import (
	"net/url"
	"testing"
)

type Cases struct {
	name   string
	data   url.Values
	errors map[string][]string
	fields []string
}

func TestEmptyValidForm(t *testing.T) {
	f := New(nil)
	if !f.IsValid() {
		t.Fail()
		t.Logf("expected valid empty form")
	}

	if len(f.GetErrors()) != 0 {
		t.Fail()
		t.Logf("expected empty error")
	}
}

func TestAddError(t *testing.T) {
	cases := []Cases{
		{
			name: "Add some err",
			data: url.Values{
				"firstName":   []string{"abc"},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.paw@gmail.com"},
				"phoneNumber": []string{"123-456-7890"},
			},
			errors: map[string][]string{
				"someField": {"someErrMsg"},
			},
		},
		{
			name: "Add multiple errors of the same field",
			data: url.Values{
				"firstName":   []string{"abc"},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.paw@gmail.com"},
				"phoneNumber": []string{"123-456-7890"},
			},
			errors: map[string][]string{
				"someField": {"someErrMsg1", "someErrMsg2"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := New(c.data)

			for k, errs := range c.errors {
				for _, err := range errs {
					f.AddError(k, err)

					if f.IsValid() {
						t.Fail()
						t.Log("expected error form, but got valid form")
					}

					n := len(f.errors[k])
					if f.errors[k][n-1] != err {
						t.Fail()
						t.Logf("expected error of [%s],but error not found", err)
					}
				}
			}
		})

	}
}

func TestRequireForm(t *testing.T) {
	cases := []Cases{
		{
			name: "normal require field",
			data: url.Values{
				"firstName":   []string{"abc"},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.paw@gmail.com"},
				"phoneNumber": []string{"123-456-7890"},
			},
			fields: []string{"firstName"},
			errors: map[string][]string{},
		},
		{
			name: "require not-exist field",
			data: url.Values{
				"firstName":   []string{"abc"},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.paw@gmail.com"},
				"phoneNumber": []string{"123-456-7890"},
			},
			fields: []string{"abc"},
			errors: map[string][]string{
				"abc": {InvalidRequired},
			},
		},
		{
			name: "require field but empty value",
			data: url.Values{
				"firstName":   []string{""},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.paw@gmail.com"},
				"phoneNumber": []string{"123-456-7890"},
			},
			fields: []string{"firstName"},
			errors: map[string][]string{
				"firstName": {InvalidRequired},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := New(c.data)

			for _, reqField := range c.fields {
				f.Require(reqField)

				n := len(f.errors[reqField])
				if n > 0 {
					exp := c.errors[reqField][n-1]
					got := f.errors[reqField][n-1]
					if exp != got {
						t.Fail()
						t.Logf("mismatch err: expected %s, but got %s", exp, got)
					}
				}
			}
		})
	}
}

func TestIsInvalidEmail(t *testing.T) {
	cases := []Cases{
		{
			name: "valid email",
			data: url.Values{
				"firstName":   []string{"abc"},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.paw@gmail.com"},
				"phoneNumber": []string{"123-456-7890"},
			},
			fields: []string{"email"},
			errors: map[string][]string{},
		},
		{
			name: "valid email",
			data: url.Values{
				"firstName":   []string{"abc"},
				"lastName":    []string{"abc"},
				"email":       []string{"nut.p"},
				"phoneNumber": []string{"123-456-7890"},
			},
			fields: []string{"email"},
			errors: map[string][]string{
				"email": {InValidEmailFormat},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := New(c.data)
			f.IsValidEmail("email")

			n := len(f.errors)
			if n > 0 {
				exp := c.errors["email"][n-1]
				got := f.errors["email"][n-1]

				if exp != got {
					t.Fail()
					t.Logf("expected %s, but got %s", exp, got)
				}
			}
		})
	}
}
