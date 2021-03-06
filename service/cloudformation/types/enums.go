// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountGateStatus string

// Enum values for AccountGateStatus
const (
	AccountGateStatusSucceeded AccountGateStatus = "SUCCEEDED"
	AccountGateStatusFailed    AccountGateStatus = "FAILED"
	AccountGateStatusSkipped   AccountGateStatus = "SKIPPED"
)

type Capability string

// Enum values for Capability
const (
	CapabilityCapability_iam         Capability = "CAPABILITY_IAM"
	CapabilityCapability_named_iam   Capability = "CAPABILITY_NAMED_IAM"
	CapabilityCapability_auto_expand Capability = "CAPABILITY_AUTO_EXPAND"
)

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionAdd    ChangeAction = "Add"
	ChangeActionModify ChangeAction = "Modify"
	ChangeActionRemove ChangeAction = "Remove"
	ChangeActionImport ChangeAction = "Import"
)

type ChangeSetStatus string

// Enum values for ChangeSetStatus
const (
	ChangeSetStatusCreate_pending     ChangeSetStatus = "CREATE_PENDING"
	ChangeSetStatusCreate_in_progress ChangeSetStatus = "CREATE_IN_PROGRESS"
	ChangeSetStatusCreate_complete    ChangeSetStatus = "CREATE_COMPLETE"
	ChangeSetStatusDelete_complete    ChangeSetStatus = "DELETE_COMPLETE"
	ChangeSetStatusFailed             ChangeSetStatus = "FAILED"
)

type ChangeSetType string

// Enum values for ChangeSetType
const (
	ChangeSetTypeCreate ChangeSetType = "CREATE"
	ChangeSetTypeUpdate ChangeSetType = "UPDATE"
	ChangeSetTypeImport ChangeSetType = "IMPORT"
)

type ChangeSource string

// Enum values for ChangeSource
const (
	ChangeSourceResourcereference  ChangeSource = "ResourceReference"
	ChangeSourceParameterreference ChangeSource = "ParameterReference"
	ChangeSourceResourceattribute  ChangeSource = "ResourceAttribute"
	ChangeSourceDirectmodification ChangeSource = "DirectModification"
	ChangeSourceAutomatic          ChangeSource = "Automatic"
)

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeResource ChangeType = "Resource"
)

type DeprecatedStatus string

// Enum values for DeprecatedStatus
const (
	DeprecatedStatusLive       DeprecatedStatus = "LIVE"
	DeprecatedStatusDeprecated DeprecatedStatus = "DEPRECATED"
)

type DifferenceType string

// Enum values for DifferenceType
const (
	DifferenceTypeAdd       DifferenceType = "ADD"
	DifferenceTypeRemove    DifferenceType = "REMOVE"
	DifferenceTypeNot_equal DifferenceType = "NOT_EQUAL"
)

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeStatic  EvaluationType = "Static"
	EvaluationTypeDynamic EvaluationType = "Dynamic"
)

type ExecutionStatus string

// Enum values for ExecutionStatus
const (
	ExecutionStatusUnavailable         ExecutionStatus = "UNAVAILABLE"
	ExecutionStatusAvailable           ExecutionStatus = "AVAILABLE"
	ExecutionStatusExecute_in_progress ExecutionStatus = "EXECUTE_IN_PROGRESS"
	ExecutionStatusExecute_complete    ExecutionStatus = "EXECUTE_COMPLETE"
	ExecutionStatusExecute_failed      ExecutionStatus = "EXECUTE_FAILED"
	ExecutionStatusObsolete            ExecutionStatus = "OBSOLETE"
)

type HandlerErrorCode string

// Enum values for HandlerErrorCode
const (
	HandlerErrorCodeNotupdatable            HandlerErrorCode = "NotUpdatable"
	HandlerErrorCodeInvalidrequest          HandlerErrorCode = "InvalidRequest"
	HandlerErrorCodeAccessdenied            HandlerErrorCode = "AccessDenied"
	HandlerErrorCodeInvalidcredentials      HandlerErrorCode = "InvalidCredentials"
	HandlerErrorCodeAlreadyexists           HandlerErrorCode = "AlreadyExists"
	HandlerErrorCodeNotfound                HandlerErrorCode = "NotFound"
	HandlerErrorCodeResourceconflict        HandlerErrorCode = "ResourceConflict"
	HandlerErrorCodeThrottling              HandlerErrorCode = "Throttling"
	HandlerErrorCodeServicelimitexceeded    HandlerErrorCode = "ServiceLimitExceeded"
	HandlerErrorCodeServicetimeout          HandlerErrorCode = "NotStabilized"
	HandlerErrorCodeGeneralserviceexception HandlerErrorCode = "GeneralServiceException"
	HandlerErrorCodeServiceinternalerror    HandlerErrorCode = "ServiceInternalError"
	HandlerErrorCodeNetworkfailure          HandlerErrorCode = "NetworkFailure"
	HandlerErrorCodeInternalfailure         HandlerErrorCode = "InternalFailure"
)

type OnFailure string

// Enum values for OnFailure
const (
	OnFailureDo_nothing OnFailure = "DO_NOTHING"
	OnFailureRollback   OnFailure = "ROLLBACK"
	OnFailureDelete     OnFailure = "DELETE"
)

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusPending     OperationStatus = "PENDING"
	OperationStatusIn_progress OperationStatus = "IN_PROGRESS"
	OperationStatusSuccess     OperationStatus = "SUCCESS"
	OperationStatusFailed      OperationStatus = "FAILED"
)

type PermissionModels string

// Enum values for PermissionModels
const (
	PermissionModelsService_managed PermissionModels = "SERVICE_MANAGED"
	PermissionModelsSelf_managed    PermissionModels = "SELF_MANAGED"
)

type ProvisioningType string

// Enum values for ProvisioningType
const (
	ProvisioningTypeNon_provisionable ProvisioningType = "NON_PROVISIONABLE"
	ProvisioningTypeImmutable         ProvisioningType = "IMMUTABLE"
	ProvisioningTypeFully_mutable     ProvisioningType = "FULLY_MUTABLE"
)

type RegistrationStatus string

// Enum values for RegistrationStatus
const (
	RegistrationStatusComplete    RegistrationStatus = "COMPLETE"
	RegistrationStatusIn_progress RegistrationStatus = "IN_PROGRESS"
	RegistrationStatusFailed      RegistrationStatus = "FAILED"
)

type RegistryType string

// Enum values for RegistryType
const (
	RegistryTypeResource RegistryType = "RESOURCE"
)

type Replacement string

// Enum values for Replacement
const (
	ReplacementTrue        Replacement = "True"
	ReplacementFalse       Replacement = "False"
	ReplacementConditional Replacement = "Conditional"
)

type RequiresRecreation string

// Enum values for RequiresRecreation
const (
	RequiresRecreationNever         RequiresRecreation = "Never"
	RequiresRecreationConditionally RequiresRecreation = "Conditionally"
	RequiresRecreationAlways        RequiresRecreation = "Always"
)

type ResourceAttribute string

// Enum values for ResourceAttribute
const (
	ResourceAttributeProperties     ResourceAttribute = "Properties"
	ResourceAttributeMetadata       ResourceAttribute = "Metadata"
	ResourceAttributeCreationpolicy ResourceAttribute = "CreationPolicy"
	ResourceAttributeUpdatepolicy   ResourceAttribute = "UpdatePolicy"
	ResourceAttributeDeletionpolicy ResourceAttribute = "DeletionPolicy"
	ResourceAttributeTags           ResourceAttribute = "Tags"
)

type ResourceSignalStatus string

// Enum values for ResourceSignalStatus
const (
	ResourceSignalStatusSuccess ResourceSignalStatus = "SUCCESS"
	ResourceSignalStatusFailure ResourceSignalStatus = "FAILURE"
)

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusCreate_in_progress          ResourceStatus = "CREATE_IN_PROGRESS"
	ResourceStatusCreate_failed               ResourceStatus = "CREATE_FAILED"
	ResourceStatusCreate_complete             ResourceStatus = "CREATE_COMPLETE"
	ResourceStatusDelete_in_progress          ResourceStatus = "DELETE_IN_PROGRESS"
	ResourceStatusDelete_failed               ResourceStatus = "DELETE_FAILED"
	ResourceStatusDelete_complete             ResourceStatus = "DELETE_COMPLETE"
	ResourceStatusDelete_skipped              ResourceStatus = "DELETE_SKIPPED"
	ResourceStatusUpdate_in_progress          ResourceStatus = "UPDATE_IN_PROGRESS"
	ResourceStatusUpdate_failed               ResourceStatus = "UPDATE_FAILED"
	ResourceStatusUpdate_complete             ResourceStatus = "UPDATE_COMPLETE"
	ResourceStatusImport_failed               ResourceStatus = "IMPORT_FAILED"
	ResourceStatusImport_complete             ResourceStatus = "IMPORT_COMPLETE"
	ResourceStatusImport_in_progress          ResourceStatus = "IMPORT_IN_PROGRESS"
	ResourceStatusImport_rollback_in_progress ResourceStatus = "IMPORT_ROLLBACK_IN_PROGRESS"
	ResourceStatusImport_rollback_failed      ResourceStatus = "IMPORT_ROLLBACK_FAILED"
	ResourceStatusImport_rollback_complete    ResourceStatus = "IMPORT_ROLLBACK_COMPLETE"
)

type StackDriftDetectionStatus string

// Enum values for StackDriftDetectionStatus
const (
	StackDriftDetectionStatusDetection_in_progress StackDriftDetectionStatus = "DETECTION_IN_PROGRESS"
	StackDriftDetectionStatusDetection_failed      StackDriftDetectionStatus = "DETECTION_FAILED"
	StackDriftDetectionStatusDetection_complete    StackDriftDetectionStatus = "DETECTION_COMPLETE"
)

type StackDriftStatus string

// Enum values for StackDriftStatus
const (
	StackDriftStatusDrifted     StackDriftStatus = "DRIFTED"
	StackDriftStatusIn_sync     StackDriftStatus = "IN_SYNC"
	StackDriftStatusUnknown     StackDriftStatus = "UNKNOWN"
	StackDriftStatusNot_checked StackDriftStatus = "NOT_CHECKED"
)

type StackInstanceDetailedStatus string

// Enum values for StackInstanceDetailedStatus
const (
	StackInstanceDetailedStatusPending    StackInstanceDetailedStatus = "PENDING"
	StackInstanceDetailedStatusRunning    StackInstanceDetailedStatus = "RUNNING"
	StackInstanceDetailedStatusSucceeded  StackInstanceDetailedStatus = "SUCCEEDED"
	StackInstanceDetailedStatusFailed     StackInstanceDetailedStatus = "FAILED"
	StackInstanceDetailedStatusCancelled  StackInstanceDetailedStatus = "CANCELLED"
	StackInstanceDetailedStatusInoperable StackInstanceDetailedStatus = "INOPERABLE"
)

type StackInstanceFilterName string

// Enum values for StackInstanceFilterName
const (
	StackInstanceFilterNameDetailed_status StackInstanceFilterName = "DETAILED_STATUS"
)

type StackInstanceStatus string

// Enum values for StackInstanceStatus
const (
	StackInstanceStatusCurrent    StackInstanceStatus = "CURRENT"
	StackInstanceStatusOutdated   StackInstanceStatus = "OUTDATED"
	StackInstanceStatusInoperable StackInstanceStatus = "INOPERABLE"
)

type StackResourceDriftStatus string

// Enum values for StackResourceDriftStatus
const (
	StackResourceDriftStatusIn_sync     StackResourceDriftStatus = "IN_SYNC"
	StackResourceDriftStatusModified    StackResourceDriftStatus = "MODIFIED"
	StackResourceDriftStatusDeleted     StackResourceDriftStatus = "DELETED"
	StackResourceDriftStatusNot_checked StackResourceDriftStatus = "NOT_CHECKED"
)

type StackSetDriftDetectionStatus string

// Enum values for StackSetDriftDetectionStatus
const (
	StackSetDriftDetectionStatusCompleted       StackSetDriftDetectionStatus = "COMPLETED"
	StackSetDriftDetectionStatusFailed          StackSetDriftDetectionStatus = "FAILED"
	StackSetDriftDetectionStatusPartial_success StackSetDriftDetectionStatus = "PARTIAL_SUCCESS"
	StackSetDriftDetectionStatusIn_progress     StackSetDriftDetectionStatus = "IN_PROGRESS"
	StackSetDriftDetectionStatusStopped         StackSetDriftDetectionStatus = "STOPPED"
)

type StackSetDriftStatus string

// Enum values for StackSetDriftStatus
const (
	StackSetDriftStatusDrifted     StackSetDriftStatus = "DRIFTED"
	StackSetDriftStatusIn_sync     StackSetDriftStatus = "IN_SYNC"
	StackSetDriftStatusNot_checked StackSetDriftStatus = "NOT_CHECKED"
)

type StackSetOperationAction string

// Enum values for StackSetOperationAction
const (
	StackSetOperationActionCreate       StackSetOperationAction = "CREATE"
	StackSetOperationActionUpdate       StackSetOperationAction = "UPDATE"
	StackSetOperationActionDelete       StackSetOperationAction = "DELETE"
	StackSetOperationActionDetect_drift StackSetOperationAction = "DETECT_DRIFT"
)

type StackSetOperationResultStatus string

// Enum values for StackSetOperationResultStatus
const (
	StackSetOperationResultStatusPending   StackSetOperationResultStatus = "PENDING"
	StackSetOperationResultStatusRunning   StackSetOperationResultStatus = "RUNNING"
	StackSetOperationResultStatusSucceeded StackSetOperationResultStatus = "SUCCEEDED"
	StackSetOperationResultStatusFailed    StackSetOperationResultStatus = "FAILED"
	StackSetOperationResultStatusCancelled StackSetOperationResultStatus = "CANCELLED"
)

type StackSetOperationStatus string

// Enum values for StackSetOperationStatus
const (
	StackSetOperationStatusRunning   StackSetOperationStatus = "RUNNING"
	StackSetOperationStatusSucceeded StackSetOperationStatus = "SUCCEEDED"
	StackSetOperationStatusFailed    StackSetOperationStatus = "FAILED"
	StackSetOperationStatusStopping  StackSetOperationStatus = "STOPPING"
	StackSetOperationStatusStopped   StackSetOperationStatus = "STOPPED"
	StackSetOperationStatusQueued    StackSetOperationStatus = "QUEUED"
)

type StackSetStatus string

// Enum values for StackSetStatus
const (
	StackSetStatusActive  StackSetStatus = "ACTIVE"
	StackSetStatusDeleted StackSetStatus = "DELETED"
)

type StackStatus string

// Enum values for StackStatus
const (
	StackStatusCreate_in_progress                           StackStatus = "CREATE_IN_PROGRESS"
	StackStatusCreate_failed                                StackStatus = "CREATE_FAILED"
	StackStatusCreate_complete                              StackStatus = "CREATE_COMPLETE"
	StackStatusRollback_in_progress                         StackStatus = "ROLLBACK_IN_PROGRESS"
	StackStatusRollback_failed                              StackStatus = "ROLLBACK_FAILED"
	StackStatusRollback_complete                            StackStatus = "ROLLBACK_COMPLETE"
	StackStatusDelete_in_progress                           StackStatus = "DELETE_IN_PROGRESS"
	StackStatusDelete_failed                                StackStatus = "DELETE_FAILED"
	StackStatusDelete_complete                              StackStatus = "DELETE_COMPLETE"
	StackStatusUpdate_in_progress                           StackStatus = "UPDATE_IN_PROGRESS"
	StackStatusUpdate_complete_cleanup_in_progress          StackStatus = "UPDATE_COMPLETE_CLEANUP_IN_PROGRESS"
	StackStatusUpdate_complete                              StackStatus = "UPDATE_COMPLETE"
	StackStatusUpdate_rollback_in_progress                  StackStatus = "UPDATE_ROLLBACK_IN_PROGRESS"
	StackStatusUpdate_rollback_failed                       StackStatus = "UPDATE_ROLLBACK_FAILED"
	StackStatusUpdate_rollback_complete_cleanup_in_progress StackStatus = "UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS"
	StackStatusUpdate_rollback_complete                     StackStatus = "UPDATE_ROLLBACK_COMPLETE"
	StackStatusReview_in_progress                           StackStatus = "REVIEW_IN_PROGRESS"
	StackStatusImport_in_progress                           StackStatus = "IMPORT_IN_PROGRESS"
	StackStatusImport_complete                              StackStatus = "IMPORT_COMPLETE"
	StackStatusImport_rollback_in_progress                  StackStatus = "IMPORT_ROLLBACK_IN_PROGRESS"
	StackStatusImport_rollback_failed                       StackStatus = "IMPORT_ROLLBACK_FAILED"
	StackStatusImport_rollback_complete                     StackStatus = "IMPORT_ROLLBACK_COMPLETE"
)

type TemplateStage string

// Enum values for TemplateStage
const (
	TemplateStageOriginal  TemplateStage = "Original"
	TemplateStageProcessed TemplateStage = "Processed"
)

type Visibility string

// Enum values for Visibility
const (
	VisibilityPublic  Visibility = "PUBLIC"
	VisibilityPrivate Visibility = "PRIVATE"
)
