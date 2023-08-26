package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteApplicationResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationResponse", string(data)}, " ")
}
