package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
