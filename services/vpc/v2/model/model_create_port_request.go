package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreatePortRequest struct {
	Body *CreatePortRequestBody `json:"body,omitempty"`
}

func (o CreatePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortRequest struct{}"
	}

	return strings.Join([]string{"CreatePortRequest", string(data)}, " ")
}
