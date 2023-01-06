package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListEnginesResponse struct {
	Total *int32 `json:"total,omitempty"`

	Data           *[]EngineSimpleInfo `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesResponse struct{}"
	}

	return strings.Join([]string{"ListEnginesResponse", string(data)}, " ")
}
