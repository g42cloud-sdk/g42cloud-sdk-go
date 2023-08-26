package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCacheRulesResponse struct {
	CacheConfig    *CacheConfig `json:"cache_config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowCacheRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCacheRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowCacheRulesResponse", string(data)}, " ")
}
