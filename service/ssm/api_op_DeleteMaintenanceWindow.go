// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a maintenance window.
func (c *Client) DeleteMaintenanceWindow(ctx context.Context, params *DeleteMaintenanceWindowInput, optFns ...func(*Options)) (*DeleteMaintenanceWindowOutput, error) {
	stack := middleware.NewStack("DeleteMaintenanceWindow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteMaintenanceWindowMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteMaintenanceWindowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMaintenanceWindow(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "DeleteMaintenanceWindow",
			Err:           err,
		}
	}
	out := result.(*DeleteMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMaintenanceWindowInput struct {

	// The ID of the maintenance window to delete.
	//
	// This member is required.
	WindowId *string
}

type DeleteMaintenanceWindowOutput struct {

	// The ID of the deleted maintenance window.
	WindowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteMaintenanceWindowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMaintenanceWindow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMaintenanceWindow{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteMaintenanceWindow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeleteMaintenanceWindow",
	}
}
