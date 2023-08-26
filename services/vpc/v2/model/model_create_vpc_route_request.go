package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcRouteRequest struct {
	Body *CreateVpcRouteRequestBody `json:"body,omitempty"`
}

func (o CreateVpcRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteRequest", string(data)}, " ")
}
