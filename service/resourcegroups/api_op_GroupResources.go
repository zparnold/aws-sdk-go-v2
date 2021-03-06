// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds the specified resources to the specified group.
func (c *Client) GroupResources(ctx context.Context, params *GroupResourcesInput, optFns ...func(*Options)) (*GroupResourcesOutput, error) {
	stack := middleware.NewStack("GroupResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGroupResourcesMiddlewares(stack)
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
	addOpGroupResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGroupResources(options.Region), middleware.Before)
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
			OperationName: "GroupResources",
			Err:           err,
		}
	}
	out := result.(*GroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GroupResourcesInput struct {

	// The name or the ARN of the resource group to add resources to.
	//
	// This member is required.
	Group *string

	// The list of ARNs for resources to be added to the group.
	//
	// This member is required.
	ResourceArns []*string
}

type GroupResourcesOutput struct {

	// The ARNs of the resources that failed to be added to the group by this
	// operation.
	Failed []*types.FailedResource

	// The ARNs of the resources that were successfully added to the group by this
	// operation.
	Succeeded []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGroupResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGroupResources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGroupResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opGroupResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "GroupResources",
	}
}
