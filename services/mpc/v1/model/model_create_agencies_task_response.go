package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAgenciesTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAgenciesTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAgenciesTaskResponse", string(data)}, " ")
}
