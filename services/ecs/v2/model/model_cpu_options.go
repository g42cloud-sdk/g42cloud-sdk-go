package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CpuOptions struct {
	HwcpuThreads *int32 `json:"hw:cpu_threads,omitempty"`
}

func (o CpuOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpuOptions struct{}"
	}

	return strings.Join([]string{"CpuOptions", string(data)}, " ")
}
