package users

import (
	"regexp"
	"unicode/utf8"
)

var (
	EmailRegexpuserRequest = regexp.MustCompile("^([a-z0-9_-]+.)*[a-z0-9_-]+@[a-z0-9_-]+(.[a-z0-9_-]+)*.[a-z]")
)

// Validate validates struct accordingly to fields tags
func (u userRequest) Validate() []string {
	var errs []string
	if u.Name == "" {
		errs = append(errs, "name::is_required")
	}
	if u.Name != "" && utf8.RuneCountInString(u.Name) < 1 {
		errs = append(errs, "name::min_length_is::1")
	}
	if u.Name != "" && utf8.RuneCountInString(u.Name) > 255 {
		errs = append(errs, "name::max_length_is::255")
	}
	if u.Phone == "" {
		errs = append(errs, "phone::is_required")
	}
	if u.Phone != "" && utf8.RuneCountInString(u.Phone) < 1 {
		errs = append(errs, "phone::min_length_is::1")
	}
	if u.Phone != "" && utf8.RuneCountInString(u.Phone) > 64 {
		errs = append(errs, "phone::max_length_is::64")
	}
	if u.Email != "" && !EmailRegexpuserRequest.Match([]byte(u.Email)) {
		errs = append(errs, "email::is_not_match_regexp")
	}

	return errs
}
