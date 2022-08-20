package status

type (
	response struct {
		Name          string `json:"name"`
		Version       string `json:"version"`
		Tag           string `json:"tag"`
		Commit        string `json:"commit"`
		Date          string `json:"date"`
		FortuneCookie string `json:"fortune_cookie"`
	}
)
