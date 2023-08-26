package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAnimatedGraphicsTaskRequest struct {
	Body *CreateAnimatedGraphicsTaskReq `json:"body,omitempty"`
}

func (o CreateAnimatedGraphicsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskRequest", string(data)}, " ")
}
