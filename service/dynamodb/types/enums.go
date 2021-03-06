// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttributeAction string

// Enum values for AttributeAction
const (
	AttributeActionAdd    AttributeAction = "ADD"
	AttributeActionPut    AttributeAction = "PUT"
	AttributeActionDelete AttributeAction = "DELETE"
)

type BackupStatus string

// Enum values for BackupStatus
const (
	BackupStatusCreating  BackupStatus = "CREATING"
	BackupStatusDeleted   BackupStatus = "DELETED"
	BackupStatusAvailable BackupStatus = "AVAILABLE"
)

type BackupType string

// Enum values for BackupType
const (
	BackupTypeUser       BackupType = "USER"
	BackupTypeSystem     BackupType = "SYSTEM"
	BackupTypeAws_backup BackupType = "AWS_BACKUP"
)

type BackupTypeFilter string

// Enum values for BackupTypeFilter
const (
	BackupTypeFilterUser       BackupTypeFilter = "USER"
	BackupTypeFilterSystem     BackupTypeFilter = "SYSTEM"
	BackupTypeFilterAws_backup BackupTypeFilter = "AWS_BACKUP"
	BackupTypeFilterAll        BackupTypeFilter = "ALL"
)

type BillingMode string

// Enum values for BillingMode
const (
	BillingModeProvisioned     BillingMode = "PROVISIONED"
	BillingModePay_per_request BillingMode = "PAY_PER_REQUEST"
)

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq           ComparisonOperator = "EQ"
	ComparisonOperatorNe           ComparisonOperator = "NE"
	ComparisonOperatorIn           ComparisonOperator = "IN"
	ComparisonOperatorLe           ComparisonOperator = "LE"
	ComparisonOperatorLt           ComparisonOperator = "LT"
	ComparisonOperatorGe           ComparisonOperator = "GE"
	ComparisonOperatorGt           ComparisonOperator = "GT"
	ComparisonOperatorBetween      ComparisonOperator = "BETWEEN"
	ComparisonOperatorNot_null     ComparisonOperator = "NOT_NULL"
	ComparisonOperatorNull         ComparisonOperator = "NULL"
	ComparisonOperatorContains     ComparisonOperator = "CONTAINS"
	ComparisonOperatorNot_contains ComparisonOperator = "NOT_CONTAINS"
	ComparisonOperatorBegins_with  ComparisonOperator = "BEGINS_WITH"
)

type ConditionalOperator string

// Enum values for ConditionalOperator
const (
	ConditionalOperatorAnd ConditionalOperator = "AND"
	ConditionalOperatorOr  ConditionalOperator = "OR"
)

type ContinuousBackupsStatus string

// Enum values for ContinuousBackupsStatus
const (
	ContinuousBackupsStatusEnabled  ContinuousBackupsStatus = "ENABLED"
	ContinuousBackupsStatusDisabled ContinuousBackupsStatus = "DISABLED"
)

type ContributorInsightsAction string

// Enum values for ContributorInsightsAction
const (
	ContributorInsightsActionEnable  ContributorInsightsAction = "ENABLE"
	ContributorInsightsActionDisable ContributorInsightsAction = "DISABLE"
)

type ContributorInsightsStatus string

// Enum values for ContributorInsightsStatus
const (
	ContributorInsightsStatusEnabling  ContributorInsightsStatus = "ENABLING"
	ContributorInsightsStatusEnabled   ContributorInsightsStatus = "ENABLED"
	ContributorInsightsStatusDisabling ContributorInsightsStatus = "DISABLING"
	ContributorInsightsStatusDisabled  ContributorInsightsStatus = "DISABLED"
	ContributorInsightsStatusFailed    ContributorInsightsStatus = "FAILED"
)

type GlobalTableStatus string

// Enum values for GlobalTableStatus
const (
	GlobalTableStatusCreating GlobalTableStatus = "CREATING"
	GlobalTableStatusActive   GlobalTableStatus = "ACTIVE"
	GlobalTableStatusDeleting GlobalTableStatus = "DELETING"
	GlobalTableStatusUpdating GlobalTableStatus = "UPDATING"
)

type IndexStatus string

// Enum values for IndexStatus
const (
	IndexStatusCreating IndexStatus = "CREATING"
	IndexStatusUpdating IndexStatus = "UPDATING"
	IndexStatusDeleting IndexStatus = "DELETING"
	IndexStatusActive   IndexStatus = "ACTIVE"
)

type KeyType string

// Enum values for KeyType
const (
	KeyTypeHash  KeyType = "HASH"
	KeyTypeRange KeyType = "RANGE"
)

type PointInTimeRecoveryStatus string

// Enum values for PointInTimeRecoveryStatus
const (
	PointInTimeRecoveryStatusEnabled  PointInTimeRecoveryStatus = "ENABLED"
	PointInTimeRecoveryStatusDisabled PointInTimeRecoveryStatus = "DISABLED"
)

type ProjectionType string

// Enum values for ProjectionType
const (
	ProjectionTypeAll       ProjectionType = "ALL"
	ProjectionTypeKeys_only ProjectionType = "KEYS_ONLY"
	ProjectionTypeInclude   ProjectionType = "INCLUDE"
)

type ReplicaStatus string

// Enum values for ReplicaStatus
const (
	ReplicaStatusCreating        ReplicaStatus = "CREATING"
	ReplicaStatusCreation_failed ReplicaStatus = "CREATION_FAILED"
	ReplicaStatusUpdating        ReplicaStatus = "UPDATING"
	ReplicaStatusDeleting        ReplicaStatus = "DELETING"
	ReplicaStatusActive          ReplicaStatus = "ACTIVE"
)

type ReturnConsumedCapacity string

// Enum values for ReturnConsumedCapacity
const (
	ReturnConsumedCapacityIndexes ReturnConsumedCapacity = "INDEXES"
	ReturnConsumedCapacityTotal   ReturnConsumedCapacity = "TOTAL"
	ReturnConsumedCapacityNone    ReturnConsumedCapacity = "NONE"
)

type ReturnItemCollectionMetrics string

// Enum values for ReturnItemCollectionMetrics
const (
	ReturnItemCollectionMetricsSize ReturnItemCollectionMetrics = "SIZE"
	ReturnItemCollectionMetricsNone ReturnItemCollectionMetrics = "NONE"
)

type ReturnValue string

// Enum values for ReturnValue
const (
	ReturnValueNone        ReturnValue = "NONE"
	ReturnValueAll_old     ReturnValue = "ALL_OLD"
	ReturnValueUpdated_old ReturnValue = "UPDATED_OLD"
	ReturnValueAll_new     ReturnValue = "ALL_NEW"
	ReturnValueUpdated_new ReturnValue = "UPDATED_NEW"
)

type ReturnValuesOnConditionCheckFailure string

// Enum values for ReturnValuesOnConditionCheckFailure
const (
	ReturnValuesOnConditionCheckFailureAll_old ReturnValuesOnConditionCheckFailure = "ALL_OLD"
	ReturnValuesOnConditionCheckFailureNone    ReturnValuesOnConditionCheckFailure = "NONE"
)

type ScalarAttributeType string

// Enum values for ScalarAttributeType
const (
	ScalarAttributeTypeS ScalarAttributeType = "S"
	ScalarAttributeTypeN ScalarAttributeType = "N"
	ScalarAttributeTypeB ScalarAttributeType = "B"
)

type Select string

// Enum values for Select
const (
	SelectAll_attributes           Select = "ALL_ATTRIBUTES"
	SelectAll_projected_attributes Select = "ALL_PROJECTED_ATTRIBUTES"
	SelectSpecific_attributes      Select = "SPECIFIC_ATTRIBUTES"
	SelectCount                    Select = "COUNT"
)

type SSEStatus string

// Enum values for SSEStatus
const (
	SSEStatusEnabling  SSEStatus = "ENABLING"
	SSEStatusEnabled   SSEStatus = "ENABLED"
	SSEStatusDisabling SSEStatus = "DISABLING"
	SSEStatusDisabled  SSEStatus = "DISABLED"
	SSEStatusUpdating  SSEStatus = "UPDATING"
)

type SSEType string

// Enum values for SSEType
const (
	SSETypeAes256 SSEType = "AES256"
	SSETypeKms    SSEType = "KMS"
)

type StreamViewType string

// Enum values for StreamViewType
const (
	StreamViewTypeNew_image          StreamViewType = "NEW_IMAGE"
	StreamViewTypeOld_image          StreamViewType = "OLD_IMAGE"
	StreamViewTypeNew_and_old_images StreamViewType = "NEW_AND_OLD_IMAGES"
	StreamViewTypeKeys_only          StreamViewType = "KEYS_ONLY"
)

type TableStatus string

// Enum values for TableStatus
const (
	TableStatusCreating                            TableStatus = "CREATING"
	TableStatusUpdating                            TableStatus = "UPDATING"
	TableStatusDeleting                            TableStatus = "DELETING"
	TableStatusActive                              TableStatus = "ACTIVE"
	TableStatusInaccessible_encryption_credentials TableStatus = "INACCESSIBLE_ENCRYPTION_CREDENTIALS"
	TableStatusArchiving                           TableStatus = "ARCHIVING"
	TableStatusArchived                            TableStatus = "ARCHIVED"
)

type TimeToLiveStatus string

// Enum values for TimeToLiveStatus
const (
	TimeToLiveStatusEnabling  TimeToLiveStatus = "ENABLING"
	TimeToLiveStatusDisabling TimeToLiveStatus = "DISABLING"
	TimeToLiveStatusEnabled   TimeToLiveStatus = "ENABLED"
	TimeToLiveStatusDisabled  TimeToLiveStatus = "DISABLED"
)
