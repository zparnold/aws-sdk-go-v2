// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a target from a maintenance window.
func (c *Client) DeregisterTargetFromMaintenanceWindow(ctx context.Context, params *DeregisterTargetFromMaintenanceWindowInput, optFns ...func(*Options)) (*DeregisterTargetFromMaintenanceWindowOutput, error) {
	stack := middleware.NewStack("DeregisterTargetFromMaintenanceWindow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterTargetFromMaintenanceWindowMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeregisterTargetFromMaintenanceWindowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterTargetFromMaintenanceWindow(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeregisterTargetFromMaintenanceWindow",
			Err:           err,
		}
	}
	out := result.(*DeregisterTargetFromMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTargetFromMaintenanceWindowInput struct {
	// The ID of the target definition to remove.
	WindowTargetId *string
	// The system checks if the target is being referenced by a task. If the target is
	// being referenced, the system returns an error and does not deregister the target
	// from the maintenance window.
	Safe *bool
	// The ID of the maintenance window the target should be removed from.
	WindowId *string
}

type DeregisterTargetFromMaintenanceWindowOutput struct {
	// The ID of the maintenance window the target was removed from.
	WindowId *string
	// The ID of the removed target definition.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterTargetFromMaintenanceWindowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterTargetFromMaintenanceWindow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterTargetFromMaintenanceWindow{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterTargetFromMaintenanceWindow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeregisterTargetFromMaintenanceWindow",
	}
}
