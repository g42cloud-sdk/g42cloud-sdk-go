package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListExtractTaskResponse struct {
	Total *int32 `json:"total,omitempty"`

	Tasks          *[]ExtractTask `json:"tasks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListExtractTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtractTaskResponse struct{}"
	}

	return strings.Join([]string{"ListExtractTaskResponse", string(data)}, " ")
}
