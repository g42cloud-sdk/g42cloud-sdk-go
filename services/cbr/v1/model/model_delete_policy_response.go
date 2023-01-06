package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeletePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyResponse", string(data)}, " ")
}
