package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAgenciesTaskRequest struct {
	Body *AgenciesTaskReq `json:"body,omitempty"`
}

func (o CreateAgenciesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAgenciesTaskRequest", string(data)}, " ")
}
