package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MasterInstance struct {
	Id string `json:"id"`

	Status string `json:"status"`

	Name string `json:"name"`

	Weight int32 `json:"weight"`

	AvailableZones []AvailableZone `json:"available_zones"`

	CpuNum int32 `json:"cpu_num"`
}

func (o MasterInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterInstance struct{}"
	}

	return strings.Join([]string{"MasterInstance", string(data)}, " ")
}
