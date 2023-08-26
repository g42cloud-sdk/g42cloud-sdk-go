package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEditingJobRequest struct {
	JobId string `json:"job_id"`
}

func (o DeleteEditingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEditingJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteEditingJobRequest", string(data)}, " ")
}
