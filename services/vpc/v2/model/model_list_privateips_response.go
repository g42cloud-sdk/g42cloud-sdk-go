package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateipsResponse struct {
	Privateips     *[]Privateip `json:"privateips,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPrivateipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateipsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateipsResponse", string(data)}, " ")
}
