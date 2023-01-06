package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResizeVolumeResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeResponse struct{}"
	}

	return strings.Join([]string{"ResizeVolumeResponse", string(data)}, " ")
}
