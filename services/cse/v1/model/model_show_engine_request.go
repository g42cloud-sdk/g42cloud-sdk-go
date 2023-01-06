package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowEngineRequest struct {
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	EngineId string `json:"engine_id"`
}

func (o ShowEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineRequest", string(data)}, " ")
}
