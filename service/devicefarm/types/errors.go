// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// An invalid argument was specified.
type ArgumentException struct {
	Message *string
}

func (e *ArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ArgumentException) ErrorCode() string             { return "ArgumentException" }
func (e *ArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ArgumentException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ArgumentException) HasMessage() bool {
	return e.Message != nil
}

// The requested object could not be deleted.
type CannotDeleteException struct {
	Message *string
}

func (e *CannotDeleteException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CannotDeleteException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CannotDeleteException) ErrorCode() string             { return "CannotDeleteException" }
func (e *CannotDeleteException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *CannotDeleteException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *CannotDeleteException) HasMessage() bool {
	return e.Message != nil
}

// An entity with the same name already exists.
type IdempotencyException struct {
	Message *string
}

func (e *IdempotencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotencyException) ErrorCode() string             { return "IdempotencyException" }
func (e *IdempotencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *IdempotencyException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IdempotencyException) HasMessage() bool {
	return e.Message != nil
}

// An internal exception was raised in the service. Contact
// aws-devicefarm-support@amazon.com (mailto:aws-devicefarm-support@amazon.com) if
// you see this error.
type InternalServiceException struct {
	Message *string
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServiceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServiceException) HasMessage() bool {
	return e.Message != nil
}

// There was an error with the update request, or you do not have sufficient
// permissions to update this VPC endpoint configuration.
type InvalidOperationException struct {
	Message *string
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidOperationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidOperationException) HasMessage() bool {
	return e.Message != nil
}

// A limit was exceeded.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// Exception gets thrown when a user is not eligible to perform the specified
// transaction.
type NotEligibleException struct {
	Message *string
}

func (e *NotEligibleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotEligibleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotEligibleException) ErrorCode() string             { return "NotEligibleException" }
func (e *NotEligibleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotEligibleException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotEligibleException) HasMessage() bool {
	return e.Message != nil
}

// The specified entity was not found.
type NotFoundException struct {
	Message *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotFoundException) HasMessage() bool {
	return e.Message != nil
}

// There was a problem with the service account.
type ServiceAccountException struct {
	Message *string
}

func (e *ServiceAccountException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceAccountException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceAccountException) ErrorCode() string             { return "ServiceAccountException" }
func (e *ServiceAccountException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ServiceAccountException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceAccountException) HasMessage() bool {
	return e.Message != nil
}

// The operation was not successful. Try again.
type TagOperationException struct {
	Message *string

	ResourceName *string
}

func (e *TagOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagOperationException) ErrorCode() string             { return "TagOperationException" }
func (e *TagOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TagOperationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TagOperationException) HasMessage() bool {
	return e.Message != nil
}
func (e *TagOperationException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *TagOperationException) HasResourceName() bool {
	return e.ResourceName != nil
}

// The request doesn't comply with the AWS Identity and Access Management (IAM) tag
// policy. Correct your request and then retry it.
type TagPolicyException struct {
	Message *string

	ResourceName *string
}

func (e *TagPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagPolicyException) ErrorCode() string             { return "TagPolicyException" }
func (e *TagPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TagPolicyException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TagPolicyException) HasMessage() bool {
	return e.Message != nil
}
func (e *TagPolicyException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *TagPolicyException) HasResourceName() bool {
	return e.ResourceName != nil
}

// The list of tags on the repository is over the limit. The maximum number of tags
// that can be applied to a repository is 50.
type TooManyTagsException struct {
	Message *string

	ResourceName *string
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TooManyTagsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TooManyTagsException) HasMessage() bool {
	return e.Message != nil
}
func (e *TooManyTagsException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *TooManyTagsException) HasResourceName() bool {
	return e.ResourceName != nil
}
