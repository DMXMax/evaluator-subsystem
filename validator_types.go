package evaluator

import (
	"regexp"
	"strings"
)

/*
rules of the game

A client in the ClientAllowList is allowed unless it is covered by the ClientDenyList
A scope in the ScopeAllowList is allowed unless it is covered by the ScopeDenyList

if scope in ScopeDenylist or client in ClientDenyList
	deny

if scope in ScopeAllowList and client in ClientAllowList
	allow
if not,
	deny

*/

type Validator interface {
	Validate(string) bool
	GetFailureMessage() string
	GetSuccessMessage() string
}

type RegexValidator struct {
	Filter         string //regex string to test scope against
	FailMessage    string // message to supply on Failure
	SuccessMessage string // message to supply on success. Optional
}

func (r *RegexValidator) Validate(s string) bool {
	ok, err := regexp.MatchString(r.Filter, s)
	if err != nil {
		evlog.Error().Err(err).Send()
		return false
	}

	return ok
}
func (r *RegexValidator) GetFailureMessage() string { return r.FailMessage }
func (r *RegexValidator) GetSuccessMessage() string { return r.SuccessMessage }

func cleanScope(s string) (scope string) {
	scope = strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z',
			r >= 'a' && r <= 'z',
			r >= '0' && r <= '9',
			r == '.' || r == '-':
			return r
		}
		return -1
	}, s)
	return
}
