package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BillbingCreateExtraInfo struct {
	CombinedOrderId *string `json:"combined_order_id,omitempty"`

	CombinedOrderEcsNum *int32 `json:"combined_order_ecs_num,omitempty"`
}

func (o BillbingCreateExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillbingCreateExtraInfo struct{}"
	}

	return strings.Join([]string{"BillbingCreateExtraInfo", string(data)}, " ")
}
