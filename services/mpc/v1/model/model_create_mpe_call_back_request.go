package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMpeCallBackRequest struct {
	Body *MpeCallBackReq `json:"body,omitempty"`
}

func (o CreateMpeCallBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMpeCallBackRequest struct{}"
	}

	return strings.Join([]string{"CreateMpeCallBackRequest", string(data)}, " ")
}
