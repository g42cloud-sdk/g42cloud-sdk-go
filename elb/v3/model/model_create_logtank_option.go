package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLogtankOption struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	LogGroupId string `json:"log_group_id"`

	LogTopicId string `json:"log_topic_id"`
}

func (o CreateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankOption struct{}"
	}

	return strings.Join([]string{"CreateLogtankOption", string(data)}, " ")
}
