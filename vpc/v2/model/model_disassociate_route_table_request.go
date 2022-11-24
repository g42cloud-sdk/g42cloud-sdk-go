package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisassociateRouteTableRequest struct {
	RoutetableId string `json:"routetable_id"`

	Body *RoutetableAssociateReqbody `json:"body,omitempty"`
}

func (o DisassociateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouteTableRequest", string(data)}, " ")
}
