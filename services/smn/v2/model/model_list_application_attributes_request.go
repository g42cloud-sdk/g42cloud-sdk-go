package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationAttributesRequest struct {
	ApplicationUrn string `json:"application_urn"`
}

func (o ListApplicationAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationAttributesRequest", string(data)}, " ")
}
