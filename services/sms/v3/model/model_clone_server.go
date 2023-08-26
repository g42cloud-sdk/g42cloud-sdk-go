package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CloneServer struct {
	VmId *string `json:"vm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	CloneError *string `json:"clone_error,omitempty"`

	CloneState *string `json:"clone_state,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o CloneServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneServer struct{}"
	}

	return strings.Join([]string{"CloneServer", string(data)}, " ")
}
