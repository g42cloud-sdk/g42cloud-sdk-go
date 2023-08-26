package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResizeServerRequest struct {
	ServerId string `json:"server_id"`

	Body *ResizeServerRequestBody `json:"body,omitempty"`
}

func (o ResizeServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeServerRequest struct{}"
	}

	return strings.Join([]string{"ResizeServerRequest", string(data)}, " ")
}
