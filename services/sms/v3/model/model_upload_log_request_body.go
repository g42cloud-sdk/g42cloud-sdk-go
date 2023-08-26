package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UploadLogRequestBody struct {
	LogBucket string `json:"log_bucket"`

	LogExpire int32 `json:"log_expire"`
}

func (o UploadLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadLogRequestBody struct{}"
	}

	return strings.Join([]string{"UploadLogRequestBody", string(data)}, " ")
}
