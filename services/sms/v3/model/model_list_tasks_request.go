package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListTasksRequest struct {
	State *ListTasksRequestState `json:"state,omitempty"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	SourceServerId *string `json:"source_server_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestState struct {
	value string
}

type ListTasksRequestStateEnum struct {
	READY                   ListTasksRequestState
	RUNNING                 ListTasksRequestState
	SYNCING                 ListTasksRequestState
	MIGRATE_SUCCESS         ListTasksRequestState
	MIGRATE_FAIL            ListTasksRequestState
	ABORTING                ListTasksRequestState
	ABORT                   ListTasksRequestState
	DELETING                ListTasksRequestState
	SYNC_F_ROLLBACKING      ListTasksRequestState
	SYNC_F_ROLLBACK_SUCCESS ListTasksRequestState
}

func GetListTasksRequestStateEnum() ListTasksRequestStateEnum {
	return ListTasksRequestStateEnum{
		READY: ListTasksRequestState{
			value: "READY",
		},
		RUNNING: ListTasksRequestState{
			value: "RUNNING",
		},
		SYNCING: ListTasksRequestState{
			value: "SYNCING",
		},
		MIGRATE_SUCCESS: ListTasksRequestState{
			value: "MIGRATE_SUCCESS",
		},
		MIGRATE_FAIL: ListTasksRequestState{
			value: "MIGRATE_FAIL",
		},
		ABORTING: ListTasksRequestState{
			value: "ABORTING",
		},
		ABORT: ListTasksRequestState{
			value: "ABORT",
		},
		DELETING: ListTasksRequestState{
			value: "DELETING",
		},
		SYNC_F_ROLLBACKING: ListTasksRequestState{
			value: "SYNC_F_ROLLBACKING",
		},
		SYNC_F_ROLLBACK_SUCCESS: ListTasksRequestState{
			value: "SYNC_F_ROLLBACK_SUCCESS",
		},
	}
}

func (c ListTasksRequestState) Value() string {
	return c.value
}

func (c ListTasksRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
