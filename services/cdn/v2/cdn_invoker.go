package v2

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/cdn/v2/model"
)

type ShowDomainLocationStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainLocationStatsInvoker) Invoke() (*model.ShowDomainLocationStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainLocationStatsResponse), nil
	}
}

type ShowDomainStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainStatsInvoker) Invoke() (*model.ShowDomainStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainStatsResponse), nil
	}
}

type ShowTopUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopUrlInvoker) Invoke() (*model.ShowTopUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopUrlResponse), nil
	}
}
