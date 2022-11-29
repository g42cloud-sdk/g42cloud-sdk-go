package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Flavor struct {
	Id string `json:"id"`

	Info *FlavorInfo `json:"info"`

	Name string `json:"name"`

	Shared bool `json:"shared"`

	ProjectId string `json:"project_id"`

	Type string `json:"type"`

	FlavorSoldOut bool `json:"flavor_sold_out"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
