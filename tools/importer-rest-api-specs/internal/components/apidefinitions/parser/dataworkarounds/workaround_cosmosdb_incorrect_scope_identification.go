// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package dataworkarounds

import (
	sdkModels "github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
)

var _ workaround = workaroundCosmosdbIncorrectScopeIdentification{}

type workaroundCosmosdbIncorrectScopeIdentification struct {
}

func (workaroundCosmosdbIncorrectScopeIdentification) IsApplicable(serviceName string, apiVersion sdkModels.APIVersion) bool {
	return serviceName == "CosmosDB" && apiVersion.APIVersion == "2026-03-15"
}

func (workaroundCosmosdbIncorrectScopeIdentification) Name() string {
	return "CosmosDB / Fix Incorrect Identification of Role Assignment ID Segments as Scope Segments"
}

func (workaroundCosmosdbIncorrectScopeIdentification) Process(input sdkModels.APIVersion) (*sdkModels.APIVersion, error) {

	return &input, nil
}
