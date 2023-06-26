package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListRemuxTaskResponse struct {
	Total *int32 `json:"total,omitempty"`

	Tasks          *[]RemuxTask `json:"tasks,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"ListRemuxTaskResponse", string(data)}, " ")
}
