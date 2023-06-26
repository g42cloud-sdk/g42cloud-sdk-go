package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListEditingJobResponse struct {
	Total *int32 `json:"total,omitempty"`

	Jobs           *[]EditingJob `json:"jobs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListEditingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEditingJobResponse struct{}"
	}

	return strings.Join([]string{"ListEditingJobResponse", string(data)}, " ")
}
