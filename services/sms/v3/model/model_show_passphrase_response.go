package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowPassphraseResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	Passphrase     *string `json:"passphrase,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPassphraseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPassphraseResponse struct{}"
	}

	return strings.Join([]string{"ShowPassphraseResponse", string(data)}, " ")
}
