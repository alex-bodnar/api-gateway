package log

// Validate validates struct accordingly to fields tags
func (c Config) Validate() []string {
	var errs []string
	if c.Mode == "" {
		errs = append(errs, "mode::is_required")
	}
	if c.LogFormat == "" {
		errs = append(errs, "log_format::is_required")
	}
	if c.LogLevel == "" {
		errs = append(errs, "log_level::is_required")
	}

	return errs
}
