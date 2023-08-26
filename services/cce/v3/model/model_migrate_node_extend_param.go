package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrateNodeExtendParam struct {
	MaxPods *int32 `json:"maxPods,omitempty"`

	DockerLVMConfigOverride *string `json:"DockerLVMConfigOverride,omitempty"`

	AlphaCcePreInstall *string `json:"alpha.cce/preInstall,omitempty"`

	AlphaCcePostInstall *string `json:"alpha.cce/postInstall,omitempty"`

	AlphaCceNodeImageID *string `json:"alpha.cce/NodeImageID,omitempty"`
}

func (o MigrateNodeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodeExtendParam struct{}"
	}

	return strings.Join([]string{"MigrateNodeExtendParam", string(data)}, " ")
}
