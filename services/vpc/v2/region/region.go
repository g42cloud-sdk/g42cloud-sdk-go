package region

import (
	"fmt"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/region"
)

var (
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://vpc.ae-ad-1.g42cloud.com")
)

var staticFields = map[string]*region.Region{
	"ae-ad-1": AE_AD_1,
}

var provider = region.DefaultProviderChain("VPC")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
