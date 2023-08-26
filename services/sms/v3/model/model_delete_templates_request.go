package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplatesRequest struct {
	Body *DeletetemplatesReq `json:"body,omitempty"`
}

func (o DeleteTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplatesRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesRequest", string(data)}, " ")
}
