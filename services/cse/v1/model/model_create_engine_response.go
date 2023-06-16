package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEngineResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	JobId          *int32 `json:"jobId,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEngineResponse struct{}"
	}

	return strings.Join([]string{"CreateEngineResponse", string(data)}, " ")
}
