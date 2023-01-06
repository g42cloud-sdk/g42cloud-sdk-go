package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CacheConfigRequestBody struct {
	CacheConfig *CacheConfigRequest `json:"cache_config"`
}

func (o CacheConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CacheConfigRequestBody", string(data)}, " ")
}
