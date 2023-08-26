package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NetworkCheckInfoRequestBody struct {
	NetworkDelay float64 `json:"network_delay"`

	NetworkJitter float64 `json:"network_jitter"`

	MigrationSpeed float64 `json:"migration_speed"`

	LossPercentage float64 `json:"loss_percentage"`

	CpuUsage float64 `json:"cpu_usage"`

	MemUsage float64 `json:"mem_usage"`

	EvaluationResult string `json:"evaluation_result"`
}

func (o NetworkCheckInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkCheckInfoRequestBody struct{}"
	}

	return strings.Join([]string{"NetworkCheckInfoRequestBody", string(data)}, " ")
}
