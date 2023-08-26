package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteVolumeResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeResponse struct{}"
	}

	return strings.Join([]string{"DeleteVolumeResponse", string(data)}, " ")
}
