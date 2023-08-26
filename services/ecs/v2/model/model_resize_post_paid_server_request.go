package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResizePostPaidServerRequest struct {
	ServerId string `json:"server_id"`

	Body *ResizePostPaidServerRequestBody `json:"body,omitempty"`
}

func (o ResizePostPaidServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerRequest struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerRequest", string(data)}, " ")
}
