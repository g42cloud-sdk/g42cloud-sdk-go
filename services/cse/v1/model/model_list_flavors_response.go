package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListFlavorsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Data           *[]FlavorBrief `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
