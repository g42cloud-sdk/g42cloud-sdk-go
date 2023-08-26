package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteDbUserResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbUserResponse", string(data)}, " ")
}
