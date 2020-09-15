// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// You don't have access to this item. The provided credentials couldn't be
// validated. You might not be authorized to carry out the request. Make sure that
// your account is authorized to use the Amazon QuickSight service, that your
// policies have the correct permissions, and that you are using the correct access
// keys.
type AccessDeniedException struct {
	Message *string

	RequestId *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccessDeniedException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *AccessDeniedException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *AccessDeniedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccessDeniedException) HasMessage() bool {
	return e.Message != nil
}

// A resource is already in a state that indicates an action is happening that must
// complete before a new update can be applied.
type ConcurrentUpdatingException struct {
	Message *string

	RequestId *string
}

func (e *ConcurrentUpdatingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentUpdatingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentUpdatingException) ErrorCode() string             { return "ConcurrentUpdatingException" }
func (e *ConcurrentUpdatingException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ConcurrentUpdatingException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConcurrentUpdatingException) HasMessage() bool {
	return e.Message != nil
}
func (e *ConcurrentUpdatingException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *ConcurrentUpdatingException) HasRequestId() bool {
	return e.RequestId != nil
}

// Updating or deleting a resource can cause an inconsistent state.
type ConflictException struct {
	Message *string

	RequestId *string
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConflictException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *ConflictException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *ConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConflictException) HasMessage() bool {
	return e.Message != nil
}

// The domain specified isn't on the allow list. All domains for embedded
// dashboards must be added to the approved list by an Amazon QuickSight admin.
type DomainNotWhitelistedException struct {
	Message *string

	RequestId *string
}

func (e *DomainNotWhitelistedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DomainNotWhitelistedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DomainNotWhitelistedException) ErrorCode() string             { return "DomainNotWhitelistedException" }
func (e *DomainNotWhitelistedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DomainNotWhitelistedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DomainNotWhitelistedException) HasMessage() bool {
	return e.Message != nil
}
func (e *DomainNotWhitelistedException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *DomainNotWhitelistedException) HasRequestId() bool {
	return e.RequestId != nil
}

// The identity type specified isn't supported. Supported identity types include
// IAM and QUICKSIGHT.
type IdentityTypeNotSupportedException struct {
	Message *string

	RequestId *string
}

func (e *IdentityTypeNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdentityTypeNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdentityTypeNotSupportedException) ErrorCode() string {
	return "IdentityTypeNotSupportedException"
}
func (e *IdentityTypeNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *IdentityTypeNotSupportedException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *IdentityTypeNotSupportedException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *IdentityTypeNotSupportedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IdentityTypeNotSupportedException) HasMessage() bool {
	return e.Message != nil
}

// An internal failure occurred.
type InternalFailureException struct {
	Message *string

	RequestId *string
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string             { return "InternalFailureException" }
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalFailureException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalFailureException) HasMessage() bool {
	return e.Message != nil
}
func (e *InternalFailureException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *InternalFailureException) HasRequestId() bool {
	return e.RequestId != nil
}

// The NextToken value isn't valid.
type InvalidNextTokenException struct {
	Message *string

	RequestId *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidNextTokenException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *InvalidNextTokenException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *InvalidNextTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidNextTokenException) HasMessage() bool {
	return e.Message != nil
}

// One or more parameters has a value that isn't valid.
type InvalidParameterValueException struct {
	Message *string

	RequestId *string
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string             { return "InvalidParameterValueException" }
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterValueException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *InvalidParameterValueException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *InvalidParameterValueException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterValueException) HasMessage() bool {
	return e.Message != nil
}

// A limit is exceeded.
type LimitExceededException struct {
	Message *string

	RequestId    *string
	ResourceType ExceptionResourceType
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
func (e *LimitExceededException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *LimitExceededException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}
func (e *LimitExceededException) GetResourceType() ExceptionResourceType {
	return e.ResourceType
}

// One or more preconditions aren't met.
type PreconditionNotMetException struct {
	Message *string

	RequestId *string
}

func (e *PreconditionNotMetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionNotMetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionNotMetException) ErrorCode() string             { return "PreconditionNotMetException" }
func (e *PreconditionNotMetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PreconditionNotMetException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PreconditionNotMetException) HasMessage() bool {
	return e.Message != nil
}
func (e *PreconditionNotMetException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *PreconditionNotMetException) HasRequestId() bool {
	return e.RequestId != nil
}

// The user with the provided name isn't found. This error can happen in any
// operation that requires finding a user based on a provided user name, such as
// DeleteUser, DescribeUser, and so on.
type QuickSightUserNotFoundException struct {
	Message *string

	RequestId *string
}

func (e *QuickSightUserNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QuickSightUserNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QuickSightUserNotFoundException) ErrorCode() string {
	return "QuickSightUserNotFoundException"
}
func (e *QuickSightUserNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *QuickSightUserNotFoundException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *QuickSightUserNotFoundException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *QuickSightUserNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *QuickSightUserNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The resource specified already exists.
type ResourceExistsException struct {
	Message *string

	ResourceType ExceptionResourceType
	RequestId    *string
}

func (e *ResourceExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceExistsException) ErrorCode() string             { return "ResourceExistsException" }
func (e *ResourceExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceExistsException) GetResourceType() ExceptionResourceType {
	return e.ResourceType
}
func (e *ResourceExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceExistsException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceExistsException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *ResourceExistsException) HasRequestId() bool {
	return e.RequestId != nil
}

// One or more resources can't be found.
type ResourceNotFoundException struct {
	Message *string

	RequestId    *string
	ResourceType ExceptionResourceType
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *ResourceNotFoundException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceNotFoundException) GetResourceType() ExceptionResourceType {
	return e.ResourceType
}

// This resource is currently unavailable.
type ResourceUnavailableException struct {
	Message *string

	ResourceType ExceptionResourceType
	RequestId    *string
}

func (e *ResourceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceUnavailableException) ErrorCode() string             { return "ResourceUnavailableException" }
func (e *ResourceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ResourceUnavailableException) GetResourceType() ExceptionResourceType {
	return e.ResourceType
}
func (e *ResourceUnavailableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceUnavailableException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceUnavailableException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *ResourceUnavailableException) HasRequestId() bool {
	return e.RequestId != nil
}

// The number of minutes specified for the lifetime of a session isn't valid. The
// session lifetime must be 15-600 minutes.
type SessionLifetimeInMinutesInvalidException struct {
	Message *string

	RequestId *string
}

func (e *SessionLifetimeInMinutesInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorCode() string {
	return "SessionLifetimeInMinutesInvalidException"
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *SessionLifetimeInMinutesInvalidException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SessionLifetimeInMinutesInvalidException) HasMessage() bool {
	return e.Message != nil
}
func (e *SessionLifetimeInMinutesInvalidException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *SessionLifetimeInMinutesInvalidException) HasRequestId() bool {
	return e.RequestId != nil
}

// Access is throttled.
type ThrottlingException struct {
	Message *string

	RequestId *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ThrottlingException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *ThrottlingException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *ThrottlingException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ThrottlingException) HasMessage() bool {
	return e.Message != nil
}

// This error indicates that you are calling an operation on an Amazon QuickSight
// subscription where the edition doesn't include support for that operation.
// Amazon QuickSight currently has Standard Edition and Enterprise Edition. Not
// every operation and capability is available in every edition.
type UnsupportedUserEditionException struct {
	Message *string

	RequestId *string
}

func (e *UnsupportedUserEditionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedUserEditionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedUserEditionException) ErrorCode() string {
	return "UnsupportedUserEditionException"
}
func (e *UnsupportedUserEditionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnsupportedUserEditionException) GetRequestId() string {
	return ptr.ToString(e.RequestId)
}
func (e *UnsupportedUserEditionException) HasRequestId() bool {
	return e.RequestId != nil
}
func (e *UnsupportedUserEditionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnsupportedUserEditionException) HasMessage() bool {
	return e.Message != nil
}
