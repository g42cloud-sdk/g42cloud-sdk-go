package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PoliciesReqV2 struct {
	Policies []Policy `json:"policies"`
}

func (o PoliciesReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesReqV2 struct{}"
	}

	return strings.Join([]string{"PoliciesReqV2", string(data)}, " ")
}
