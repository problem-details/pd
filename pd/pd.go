// https://www.rfc-editor.org/rfc/rfc7807.txt
package pd

import (
	"encoding/json"
)

type ProblemDetails struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
	Err      error  `json:"-"`
}

func (details *ProblemDetails) Error() string {
	bs, err := json.Marshal(details)
	if err != nil {
		panic(err) // Is it possible to panic?
	}
	return string(bs)
}

func (details *ProblemDetails) Unwrap() error {
	return details.Err
}
