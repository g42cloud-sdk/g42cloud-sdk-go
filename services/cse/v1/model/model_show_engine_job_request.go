package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowEngineJobRequest struct {
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	EngineId string `json:"engine_id"`

	JobId string `json:"job_id"`
}

func (o ShowEngineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineJobRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineJobRequest", string(data)}, " ")
}
