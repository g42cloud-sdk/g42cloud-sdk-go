package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListPoliciesResponse struct {
	Policies *[]Policy `json:"policies,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListPoliciesResponse", string(data)}, " ")
}
