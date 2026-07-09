// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	sdkModels "github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
)

var _ commonIdMatcher = commonIdCassandraRoleAssignment{}

type commonIdCassandraRoleAssignment struct{}

func (c commonIdCassandraRoleAssignment) ID() sdkModels.ResourceID {
	name := "CassandraRoleAssignment"
	return sdkModels.ResourceID{
		CommonIDAlias: &name,
		ConstantNames: []string{},
		Segments: []sdkModels.ResourceIDSegment{
			sdkModels.NewStaticValueResourceIDSegment("subscriptions", "subscriptions"),
			sdkModels.NewSubscriptionIDResourceIDSegment("subscriptionId"),
			sdkModels.NewStaticValueResourceIDSegment("resourceGroups", "resourceGroups"),
			sdkModels.NewResourceGroupNameResourceIDSegment("resourceGroupName"),
			sdkModels.NewStaticValueResourceIDSegment("providers", "providers"),
			sdkModels.NewResourceProviderResourceIDSegment("resourceProvider", "Microsoft.DocumentDB"),
			sdkModels.NewStaticValueResourceIDSegment("databaseAccounts", "databaseAccounts"),
			sdkModels.NewUserSpecifiedResourceIDSegment("accountName", "accountName"),
			sdkModels.NewStaticValueResourceIDSegment("cassandraRoleAssignments", "cassandraRoleAssignments"),
			sdkModels.NewScopeResourceIDSegment("roleAssignmentId"),
		},
	}
}
