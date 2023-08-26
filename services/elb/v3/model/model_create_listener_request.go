package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerRequest struct {
	Body *CreateListenerRequestBody `json:"body,omitempty"`
}

func (o CreateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerRequest", string(data)}, " ")
}
