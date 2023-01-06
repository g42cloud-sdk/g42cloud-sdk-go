package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEngineRequest struct {
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	Body *EngineCreateReq `json:"body,omitempty"`
}

func (o CreateEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEngineRequest struct{}"
	}

	return strings.Join([]string{"CreateEngineRequest", string(data)}, " ")
}
