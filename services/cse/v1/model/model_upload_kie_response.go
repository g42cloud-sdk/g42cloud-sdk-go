package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UploadKieResponse struct {
	Success *[]GetKieConfigs `json:"success,omitempty"`

	Failure        *[]DocFailedOfUpload `json:"failure,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UploadKieResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKieResponse struct{}"
	}

	return strings.Join([]string{"UploadKieResponse", string(data)}, " ")
}
