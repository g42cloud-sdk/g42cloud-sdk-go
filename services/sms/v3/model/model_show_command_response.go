package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCommandResponse struct {
	CommandName *string `json:"command_name,omitempty"`

	CommandParam   *ComandParam `json:"command_param,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommandResponse struct{}"
	}

	return strings.Join([]string{"ShowCommandResponse", string(data)}, " ")
}
