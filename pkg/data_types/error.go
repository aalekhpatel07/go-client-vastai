package data_types

import "fmt"

type VastAIError struct {
	Err     string `json:"error"`
	Message string `json:"msg"`
}

func (v VastAIError) Error() string {
	return fmt.Sprintf("VastAIError %s: %s", v.Err, v.Message)
}
