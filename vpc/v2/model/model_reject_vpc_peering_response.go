package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type RejectVpcPeeringResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *RejectVpcPeeringResponseStatus `json:"status,omitempty"`

	RequestVpcInfo *VpcInfo `json:"request_vpc_info,omitempty"`

	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info,omitempty"`

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RejectVpcPeeringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"RejectVpcPeeringResponse", string(data)}, " ")
}

type RejectVpcPeeringResponseStatus struct {
	value string
}

type RejectVpcPeeringResponseStatusEnum struct {
	PENDING_ACCEPTANCE RejectVpcPeeringResponseStatus
	REJECTED           RejectVpcPeeringResponseStatus
	EXPIRED            RejectVpcPeeringResponseStatus
	DELETED            RejectVpcPeeringResponseStatus
	ACTIVE             RejectVpcPeeringResponseStatus
}

func GetRejectVpcPeeringResponseStatusEnum() RejectVpcPeeringResponseStatusEnum {
	return RejectVpcPeeringResponseStatusEnum{
		PENDING_ACCEPTANCE: RejectVpcPeeringResponseStatus{
			value: "PENDING_ACCEPTANCE",
		},
		REJECTED: RejectVpcPeeringResponseStatus{
			value: "REJECTED",
		},
		EXPIRED: RejectVpcPeeringResponseStatus{
			value: "EXPIRED",
		},
		DELETED: RejectVpcPeeringResponseStatus{
			value: "DELETED",
		},
		ACTIVE: RejectVpcPeeringResponseStatus{
			value: "ACTIVE",
		},
	}
}

func (c RejectVpcPeeringResponseStatus) Value() string {
	return c.value
}

func (c RejectVpcPeeringResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RejectVpcPeeringResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
