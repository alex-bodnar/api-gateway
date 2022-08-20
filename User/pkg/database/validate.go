package database

// Validate validates struct accordingly to fields tags
func (c Config) Validate() []string {
	var errs []string
	if c.ConnectionString == "" {
		errs = append(errs, "connection_string::is_required")
	}
	if c.Dialect == "" {
		errs = append(errs, "dialect::is_required")
	}
	if c.Driver == "" {
		errs = append(errs, "driver::is_required")
	}
	if c.MaxRetries == 0 {
		errs = append(errs, "max_retries::is_required")
	}
	if c.RetryDelay == 0 {
		errs = append(errs, "retry_delay::is_required")
	}
	if c.MigrationDirectory == "" {
		errs = append(errs, "migration_directory::is_required")
	}

	return errs
}
