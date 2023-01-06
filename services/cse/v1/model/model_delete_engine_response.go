package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEngineResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	JobId          *int32 `json:"job_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEngineResponse struct{}"
	}

	return strings.Join([]string{"DeleteEngineResponse", string(data)}, " ")
}
