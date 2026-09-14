// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package dataworkarounds

import (
	"errors"
	"fmt"
	"slices"

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
	resource, ok := input.Resources["Openapis"]
	if !ok {
		return nil, errors.New("expected a Resource named `Openapis` but  didn't get one")
	}

	resourceIdKeys := []string{
		"CassandraRoleAssignmentId",
		"GremlinRoleAssignmentId",
		"MongoMIRoleAssignmentId",
		"SqlRoleAssignmentId",
		"TableRoleAssignmentId",
	}

	for _, resourceIdKey := range resourceIdKeys {
		resourceId, ok := resource.ResourceIDs[resourceIdKey]
		if !ok {
			return nil, fmt.Errorf("couldn't find Resource ID `%s`", resourceIdKey)
		}

		index := slices.IndexFunc(resourceId.Segments, func(segment sdkModels.ResourceIDSegment) bool {
			return segment.Name == "roleAssignmentId"
		})
		resourceId.Segments[index].Name = "roleAssignmentGuid"
		resourceId.Segments[index].ExampleValue = "roleAssignmentGuid"
		resource.ResourceIDs[resourceIdKey] = resourceId
	}

	input.Resources["Openapis"] = resource

	return &input, nil
}
