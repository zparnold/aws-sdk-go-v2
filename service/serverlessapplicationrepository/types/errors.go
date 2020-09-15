// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// One of the parameters in the request is invalid.
type BadRequestException struct {
	Message *string

	ErrorCode_ *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *BadRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BadRequestException) HasMessage() bool {
	return e.Message != nil
}
func (e *BadRequestException) GetErrorCode_() string {
	return ptr.ToString(e.ErrorCode_)
}
func (e *BadRequestException) HasErrorCode_() bool {
	return e.ErrorCode_ != nil
}

// The resource already exists.
type ConflictException struct {
	Message *string

	ErrorCode_ *string
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
func (e *ConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConflictException) HasMessage() bool {
	return e.Message != nil
}
func (e *ConflictException) GetErrorCode_() string {
	return ptr.ToString(e.ErrorCode_)
}
func (e *ConflictException) HasErrorCode_() bool {
	return e.ErrorCode_ != nil
}

// The client is not authenticated.
type ForbiddenException struct {
	Message *string

	ErrorCode_ *string
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
func (e *ForbiddenException) GetErrorCode_() string {
	return ptr.ToString(e.ErrorCode_)
}
func (e *ForbiddenException) HasErrorCode_() bool {
	return e.ErrorCode_ != nil
}

// The AWS Serverless Application Repository service encountered an internal error.
type InternalServerErrorException struct {
	Message *string

	ErrorCode_ *string
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string             { return "InternalServerErrorException" }
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServerErrorException) GetErrorCode_() string {
	return ptr.ToString(e.ErrorCode_)
}
func (e *InternalServerErrorException) HasErrorCode_() bool {
	return e.ErrorCode_ != nil
}
func (e *InternalServerErrorException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServerErrorException) HasMessage() bool {
	return e.Message != nil
}

// The resource (for example, an access policy statement) specified in the request
// doesn't exist.
type NotFoundException struct {
	Message *string

	ErrorCode_ *string
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
func (e *NotFoundException) GetErrorCode_() string {
	return ptr.ToString(e.ErrorCode_)
}
func (e *NotFoundException) HasErrorCode_() bool {
	return e.ErrorCode_ != nil
}

// The client is sending more than the allowed number of requests per unit of time.
type TooManyRequestsException struct {
	Message *string

	ErrorCode_ *string
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TooManyRequestsException) GetErrorCode_() string {
	return ptr.ToString(e.ErrorCode_)
}
func (e *TooManyRequestsException) HasErrorCode_() bool {
	return e.ErrorCode_ != nil
}
func (e *TooManyRequestsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TooManyRequestsException) HasMessage() bool {
	return e.Message != nil
}
