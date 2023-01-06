package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateWholeImageRequest struct {
	Body *CreateWholeImageRequestBody `json:"body,omitempty"`
}

func (o CreateWholeImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWholeImageRequest struct{}"
	}

	return strings.Join([]string{"CreateWholeImageRequest", string(data)}, " ")
}
