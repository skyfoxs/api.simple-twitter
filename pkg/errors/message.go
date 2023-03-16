package errors

type Errors struct {
	Message string `json:"message"`
}

func New(m string) *Errors {
	return &Errors{Message: m}
}
