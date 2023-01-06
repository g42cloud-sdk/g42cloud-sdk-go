package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVolumeResponse struct {
	JobId *string `json:"job_id,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	VolumeIds      *[]string `json:"volume_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeResponse struct{}"
	}

	return strings.Join([]string{"CreateVolumeResponse", string(data)}, " ")
}
