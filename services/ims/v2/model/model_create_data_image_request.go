package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateDataImageRequest struct {
	Body *CreateDataImageRequestBody `json:"body,omitempty"`
}

func (o CreateDataImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageRequest struct{}"
	}

	return strings.Join([]string{"CreateDataImageRequest", string(data)}, " ")
}
