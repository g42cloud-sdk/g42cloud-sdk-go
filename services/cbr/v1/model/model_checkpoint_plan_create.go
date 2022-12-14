package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CheckpointPlanCreate struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Resources *[]CheckpointResourceResp `json:"resources,omitempty"`

	SkippedResources *[]CheckpointCreateSkippedResource `json:"skipped_resources,omitempty"`
}

func (o CheckpointPlanCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointPlanCreate struct{}"
	}

	return strings.Join([]string{"CheckpointPlanCreate", string(data)}, " ")
}
