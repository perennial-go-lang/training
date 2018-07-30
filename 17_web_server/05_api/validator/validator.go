package validator

import (
	. "github.com/thedevsaddam/govalidator"
	"net/http"
)

func Validate(rules MapData, messages MapData, request *http.Request) (err map[string][]string) {
	opts := Options{
		Request:         request,  // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	err = New(opts).Validate()
	return
}
