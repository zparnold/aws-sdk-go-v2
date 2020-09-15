// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// You have exceeded the permitted request rate for the specific operation.
type CallRateLimitExceededException struct {
	Message *string
}

func (e *CallRateLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CallRateLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CallRateLimitExceededException) ErrorCode() string             { return "CallRateLimitExceededException" }
func (e *CallRateLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *CallRateLimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *CallRateLimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// These errors are usually caused by a client action, such as using an action or
// resource on behalf of a user that doesn't have permissions to use the action or
// resource, or specifying an invalid resource identifier.
type ClientException struct {
	Message *string
}

func (e *ClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientException) ErrorCode() string             { return "ClientException" }
func (e *ClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ClientException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ClientException) HasMessage() bool {
	return e.Message != nil
}

// You are not authorized to perform the requested operation.
type ForbiddenException struct {
	Message *string
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string             { return "ForbiddenException" }
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ForbiddenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ForbiddenException) HasMessage() bool {
	return e.Message != nil
}

// You have specified a client token for an operation using parameter values that
// differ from a previous request that used the same client token.
type IdempotentParameterMismatchException struct {
	Message *string
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	return "IdempotentParameterMismatchException"
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *IdempotentParameterMismatchException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IdempotentParameterMismatchException) HasMessage() bool {
	return e.Message != nil
}

// You have provided an invalid pagination token in your request.
type InvalidPaginationTokenException struct {
	Message *string
}

func (e *InvalidPaginationTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPaginationTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPaginationTokenException) ErrorCode() string {
	return "InvalidPaginationTokenException"
}
func (e *InvalidPaginationTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidPaginationTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidPaginationTokenException) HasMessage() bool {
	return e.Message != nil
}

// You have specified two or more mutually exclusive parameters. Review the error
// message for details.
type InvalidParameterCombinationException struct {
	Message *string
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	return "InvalidParameterCombinationException"
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidParameterCombinationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterCombinationException) HasMessage() bool {
	return e.Message != nil
}

// The specified parameter is invalid. Review the available parameters for the API
// request.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}

// The value that you provided for the specified parameter is invalid.
type InvalidParameterValueException struct {
	Message *string
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
func (e *InvalidParameterValueException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterValueException) HasMessage() bool {
	return e.Message != nil
}

// You have made a request for an action that is not supported by the service.
type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRequestException) HasMessage() bool {
	return e.Message != nil
}

// Your version number is out of bounds or does not follow the required syntax.
type InvalidVersionNumberException struct {
	Message *string
}

func (e *InvalidVersionNumberException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVersionNumberException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVersionNumberException) ErrorCode() string             { return "InvalidVersionNumberException" }
func (e *InvalidVersionNumberException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidVersionNumberException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidVersionNumberException) HasMessage() bool {
	return e.Message != nil
}

// The resource that you are trying to create already exists.
type ResourceAlreadyExistsException struct {
	Message *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// You have attempted to mutate or delete a resource with a dependency that
// prohibits this action. See the error message for more details.
type ResourceDependencyException struct {
	Message *string
}

func (e *ResourceDependencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceDependencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceDependencyException) ErrorCode() string             { return "ResourceDependencyException" }
func (e *ResourceDependencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceDependencyException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceDependencyException) HasMessage() bool {
	return e.Message != nil
}

// The resource that you are trying to operate on is currently in use. Review the
// message details and retry later.
type ResourceInUseException struct {
	Message *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceInUseException) HasMessage() bool {
	return e.Message != nil
}

// At least one of the resources referenced by your request does not exist.
type ResourceNotFoundException struct {
	Message *string
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
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// This exception is thrown when the service encounters an unrecoverable exception.
type ServiceException struct {
	Message *string
}

func (e *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceException) ErrorCode() string             { return "ServiceException" }
func (e *ServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ServiceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceException) HasMessage() bool {
	return e.Message != nil
}

// You have exceeded the number of permitted resources or operations for this
// service. For service quotas, see EC2 Image Builder endpoints and quotas
// (https://docs.aws.amazon.com/general/latest/gr/imagebuilder.html#limits_imagebuilder).
type ServiceQuotaExceededException struct {
	Message *string
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string             { return "ServiceQuotaExceededException" }
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ServiceQuotaExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceQuotaExceededException) HasMessage() bool {
	return e.Message != nil
}

// The service is unable to process your request at this time.
type ServiceUnavailableException struct {
	Message *string
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ServiceUnavailableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceUnavailableException) HasMessage() bool {
	return e.Message != nil
}
