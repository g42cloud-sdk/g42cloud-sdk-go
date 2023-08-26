package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowPortResponse struct {
	Port           *Port `json:"port,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPortResponse struct{}"
	}

	return strings.Join([]string{"ShowPortResponse", string(data)}, " ")
}
