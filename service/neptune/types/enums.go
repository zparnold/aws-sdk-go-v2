// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplyMethod string

// Enum values for ApplyMethod
const (
	ApplyMethodImmediate     ApplyMethod = "immediate"
	ApplyMethodPendingReboot ApplyMethod = "pending-reboot"
)

type SourceType string

// Enum values for SourceType
const (
	SourceTypeDbInstance        SourceType = "db-instance"
	SourceTypeDbParameterGroup  SourceType = "db-parameter-group"
	SourceTypeDbSecurityGroup   SourceType = "db-security-group"
	SourceTypeDbSnapshot        SourceType = "db-snapshot"
	SourceTypeDbCluster         SourceType = "db-cluster"
	SourceTypeDbClusterSnapshot SourceType = "db-cluster-snapshot"
)
