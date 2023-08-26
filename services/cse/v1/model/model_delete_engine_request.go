package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEngineRequest struct {
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	EngineId string `json:"engine_id"`
}

func (o DeleteEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEngineRequest struct{}"
	}

	return strings.Join([]string{"DeleteEngineRequest", string(data)}, " ")
}
