package config

// Validate validates struct accordingly to fields tags
func (c Config) Validate() []string {
	var errs []string
	if e := c.Logger.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if e := c.Storage.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if e := c.Delivery.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if e := c.Extra.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (s Storage) Validate() []string {
	var errs []string
	if e := s.Postgres.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (d Delivery) Validate() []string {
	var errs []string
	if e := d.HTTPServer.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (h HTTPServer) Validate() []string {
	var errs []string
	if h.ListenAddress == "" {
		errs = append(errs, "listen_address::is_required")
	}
	if h.ReadTimeout == 0 {
		errs = append(errs, "read_timeout::is_required")
	}
	if h.WriteTimeout == 0 {
		errs = append(errs, "write_timeout::is_required")
	}
	if h.BodySizeLimitBytes == 0 {
		errs = append(errs, "body_size_limit_bytes::is_required")
	}
	if h.GracefulTimeout == 0 {
		errs = append(errs, "graceful_timeout::is_required")
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (e Extra) Validate() []string {
	var errs []string
	if e := e.AuthenticatorKafka.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (a AuthenticatorKafka) Validate() []string {
	var errs []string
	if a.Brokers == nil || len(a.Brokers) == 0 {
		errs = append(errs, "brokers::is_required")
	}

	return errs
}
