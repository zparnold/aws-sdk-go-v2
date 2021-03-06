// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified conformance pack and all the AWS Config rules, remediation
// actions, and all evaluation results within that conformance pack. AWS Config
// sets the conformance pack to DELETE_IN_PROGRESS until the deletion is complete.
// You cannot update a conformance pack while it is in this state.
func (c *Client) DeleteConformancePack(ctx context.Context, params *DeleteConformancePackInput, optFns ...func(*Options)) (*DeleteConformancePackOutput, error) {
	stack := middleware.NewStack("DeleteConformancePack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteConformancePackMiddlewares(stack)
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
	addOpDeleteConformancePackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConformancePack(options.Region), middleware.Before)
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
			OperationName: "DeleteConformancePack",
			Err:           err,
		}
	}
	out := result.(*DeleteConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConformancePackInput struct {

	// Name of the conformance pack you want to delete.
	//
	// This member is required.
	ConformancePackName *string
}

type DeleteConformancePackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteConformancePackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteConformancePack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteConformancePack{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteConformancePack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteConformancePack",
	}
}
