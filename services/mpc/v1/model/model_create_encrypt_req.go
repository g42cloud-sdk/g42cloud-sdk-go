package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEncryptReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	Encryption *Encryption `json:"encryption,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

func (o CreateEncryptReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptReq struct{}"
	}

	return strings.Join([]string{"CreateEncryptReq", string(data)}, " ")
}
