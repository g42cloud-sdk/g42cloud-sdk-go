package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
