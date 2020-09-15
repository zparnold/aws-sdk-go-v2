// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// AWS DMS was denied access to the endpoint. Check that the role is correctly
// configured.
type AccessDeniedFault struct {
	Message *string
}

func (e *AccessDeniedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedFault) ErrorCode() string             { return "AccessDeniedFault" }
func (e *AccessDeniedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccessDeniedFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccessDeniedFault) HasMessage() bool {
	return e.Message != nil
}

// There are not enough resources allocated to the database migration.
type InsufficientResourceCapacityFault struct {
	Message *string
}

func (e *InsufficientResourceCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientResourceCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientResourceCapacityFault) ErrorCode() string {
	return "InsufficientResourceCapacityFault"
}
func (e *InsufficientResourceCapacityFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InsufficientResourceCapacityFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InsufficientResourceCapacityFault) HasMessage() bool {
	return e.Message != nil
}

// The certificate was not valid.
type InvalidCertificateFault struct {
	Message *string
}

func (e *InvalidCertificateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCertificateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCertificateFault) ErrorCode() string             { return "InvalidCertificateFault" }
func (e *InvalidCertificateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidCertificateFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidCertificateFault) HasMessage() bool {
	return e.Message != nil
}

// The resource is in a state that prevents it from being used for database
// migration.
type InvalidResourceStateFault struct {
	Message *string
}

func (e *InvalidResourceStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceStateFault) ErrorCode() string             { return "InvalidResourceStateFault" }
func (e *InvalidResourceStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidResourceStateFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidResourceStateFault) HasMessage() bool {
	return e.Message != nil
}

// The subnet provided is invalid.
type InvalidSubnet struct {
	Message *string
}

func (e *InvalidSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnet) ErrorCode() string             { return "InvalidSubnet" }
func (e *InvalidSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidSubnet) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidSubnet) HasMessage() bool {
	return e.Message != nil
}

// The ciphertext references a key that doesn't exist or that the DMS account
// doesn't have access to.
type KMSAccessDeniedFault struct {
	Message *string
}

func (e *KMSAccessDeniedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSAccessDeniedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSAccessDeniedFault) ErrorCode() string             { return "KMSAccessDeniedFault" }
func (e *KMSAccessDeniedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSAccessDeniedFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSAccessDeniedFault) HasMessage() bool {
	return e.Message != nil
}

// The specified master key (CMK) isn't enabled.
type KMSDisabledFault struct {
	Message *string
}

func (e *KMSDisabledFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSDisabledFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSDisabledFault) ErrorCode() string             { return "KMSDisabledFault" }
func (e *KMSDisabledFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSDisabledFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSDisabledFault) HasMessage() bool {
	return e.Message != nil
}

// An AWS Key Management Service (AWS KMS) error is preventing access to AWS KMS.
type KMSFault struct {
	Message *string
}

func (e *KMSFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSFault) ErrorCode() string             { return "KMSFault" }
func (e *KMSFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSFault) HasMessage() bool {
	return e.Message != nil
}

// The state of the specified AWS KMS resource isn't valid for this request.
type KMSInvalidStateFault struct {
	Message *string
}

func (e *KMSInvalidStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSInvalidStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSInvalidStateFault) ErrorCode() string             { return "KMSInvalidStateFault" }
func (e *KMSInvalidStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSInvalidStateFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSInvalidStateFault) HasMessage() bool {
	return e.Message != nil
}

// AWS DMS cannot access the AWS KMS key.
type KMSKeyNotAccessibleFault struct {
	Message *string
}

func (e *KMSKeyNotAccessibleFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSKeyNotAccessibleFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSKeyNotAccessibleFault) ErrorCode() string             { return "KMSKeyNotAccessibleFault" }
func (e *KMSKeyNotAccessibleFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSKeyNotAccessibleFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSKeyNotAccessibleFault) HasMessage() bool {
	return e.Message != nil
}

// The specified AWS KMS entity or resource can't be found.
type KMSNotFoundFault struct {
	Message *string
}

func (e *KMSNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSNotFoundFault) ErrorCode() string             { return "KMSNotFoundFault" }
func (e *KMSNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSNotFoundFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSNotFoundFault) HasMessage() bool {
	return e.Message != nil
}

// This request triggered AWS KMS request throttling.
type KMSThrottlingFault struct {
	Message *string
}

func (e *KMSThrottlingFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSThrottlingFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSThrottlingFault) ErrorCode() string             { return "KMSThrottlingFault" }
func (e *KMSThrottlingFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *KMSThrottlingFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *KMSThrottlingFault) HasMessage() bool {
	return e.Message != nil
}

// The replication subnet group does not cover enough Availability Zones (AZs).
// Edit the replication subnet group and add more AZs.
type ReplicationSubnetGroupDoesNotCoverEnoughAZs struct {
	Message *string
}

func (e *ReplicationSubnetGroupDoesNotCoverEnoughAZs) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationSubnetGroupDoesNotCoverEnoughAZs) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationSubnetGroupDoesNotCoverEnoughAZs) ErrorCode() string {
	return "ReplicationSubnetGroupDoesNotCoverEnoughAZs"
}
func (e *ReplicationSubnetGroupDoesNotCoverEnoughAZs) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *ReplicationSubnetGroupDoesNotCoverEnoughAZs) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ReplicationSubnetGroupDoesNotCoverEnoughAZs) HasMessage() bool {
	return e.Message != nil
}

// The resource you are attempting to create already exists.
type ResourceAlreadyExistsFault struct {
	Message *string

	ResourceArn *string
}

func (e *ResourceAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsFault) ErrorCode() string             { return "ResourceAlreadyExistsFault" }
func (e *ResourceAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAlreadyExistsFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAlreadyExistsFault) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceAlreadyExistsFault) GetResourceArn() string {
	return ptr.ToString(e.ResourceArn)
}
func (e *ResourceAlreadyExistsFault) HasResourceArn() bool {
	return e.ResourceArn != nil
}

// The resource could not be found.
type ResourceNotFoundFault struct {
	Message *string
}

func (e *ResourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundFault) ErrorCode() string             { return "ResourceNotFoundFault" }
func (e *ResourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundFault) HasMessage() bool {
	return e.Message != nil
}

// The quota for this resource quota has been exceeded.
type ResourceQuotaExceededFault struct {
	Message *string
}

func (e *ResourceQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceQuotaExceededFault) ErrorCode() string             { return "ResourceQuotaExceededFault" }
func (e *ResourceQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceQuotaExceededFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceQuotaExceededFault) HasMessage() bool {
	return e.Message != nil
}

// Insufficient privileges are preventing access to an Amazon S3 object.
type S3AccessDeniedFault struct {
	Message *string
}

func (e *S3AccessDeniedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *S3AccessDeniedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *S3AccessDeniedFault) ErrorCode() string             { return "S3AccessDeniedFault" }
func (e *S3AccessDeniedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *S3AccessDeniedFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *S3AccessDeniedFault) HasMessage() bool {
	return e.Message != nil
}

// A specified Amazon S3 bucket, bucket folder, or other object can't be found.
type S3ResourceNotFoundFault struct {
	Message *string
}

func (e *S3ResourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *S3ResourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *S3ResourceNotFoundFault) ErrorCode() string             { return "S3ResourceNotFoundFault" }
func (e *S3ResourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *S3ResourceNotFoundFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *S3ResourceNotFoundFault) HasMessage() bool {
	return e.Message != nil
}

// The SNS topic is invalid.
type SNSInvalidTopicFault struct {
	Message *string
}

func (e *SNSInvalidTopicFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSInvalidTopicFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSInvalidTopicFault) ErrorCode() string             { return "SNSInvalidTopicFault" }
func (e *SNSInvalidTopicFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SNSInvalidTopicFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SNSInvalidTopicFault) HasMessage() bool {
	return e.Message != nil
}

// You are not authorized for the SNS subscription.
type SNSNoAuthorizationFault struct {
	Message *string
}

func (e *SNSNoAuthorizationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSNoAuthorizationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSNoAuthorizationFault) ErrorCode() string             { return "SNSNoAuthorizationFault" }
func (e *SNSNoAuthorizationFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SNSNoAuthorizationFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SNSNoAuthorizationFault) HasMessage() bool {
	return e.Message != nil
}

// The storage quota has been exceeded.
type StorageQuotaExceededFault struct {
	Message *string
}

func (e *StorageQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageQuotaExceededFault) ErrorCode() string             { return "StorageQuotaExceededFault" }
func (e *StorageQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *StorageQuotaExceededFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *StorageQuotaExceededFault) HasMessage() bool {
	return e.Message != nil
}

// The specified subnet is already in use.
type SubnetAlreadyInUse struct {
	Message *string
}

func (e *SubnetAlreadyInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetAlreadyInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetAlreadyInUse) ErrorCode() string             { return "SubnetAlreadyInUse" }
func (e *SubnetAlreadyInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SubnetAlreadyInUse) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SubnetAlreadyInUse) HasMessage() bool {
	return e.Message != nil
}

// An upgrade dependency is preventing the database migration.
type UpgradeDependencyFailureFault struct {
	Message *string
}

func (e *UpgradeDependencyFailureFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UpgradeDependencyFailureFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UpgradeDependencyFailureFault) ErrorCode() string             { return "UpgradeDependencyFailureFault" }
func (e *UpgradeDependencyFailureFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UpgradeDependencyFailureFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UpgradeDependencyFailureFault) HasMessage() bool {
	return e.Message != nil
}
