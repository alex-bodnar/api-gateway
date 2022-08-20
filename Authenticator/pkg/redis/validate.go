// Code generated by go-validator; DO NOT EDIT.
// Package redis contains models and autogenerated validation code
package redis

// Validate validates struct accordingly to fields tags
func (c Config) Validate() []string {
	var errs []string
	if c.ConnectionAddr == "" {
		errs = append(errs, "connection_addr::is_required")
	}
	if c.MaxRetries == 0 {
		errs = append(errs, "max_retries::is_required")
	}
	if c.RetryDelay == 0 {
		errs = append(errs, "retry_delay::is_required")
	}
	if c.QueryTimeout == 0 {
		errs = append(errs, "query_timeout::is_required")
	}

	return errs
}
