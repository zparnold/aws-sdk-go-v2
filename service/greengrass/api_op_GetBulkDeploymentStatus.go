// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the status of a bulk deployment.
func (c *Client) GetBulkDeploymentStatus(ctx context.Context, params *GetBulkDeploymentStatusInput, optFns ...func(*Options)) (*GetBulkDeploymentStatusOutput, error) {
	stack := middleware.NewStack("GetBulkDeploymentStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBulkDeploymentStatusMiddlewares(stack)
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
	addOpGetBulkDeploymentStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBulkDeploymentStatus(options.Region), middleware.Before)
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
			OperationName: "GetBulkDeploymentStatus",
			Err:           err,
		}
	}
	out := result.(*GetBulkDeploymentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBulkDeploymentStatusInput struct {

	// The ID of the bulk deployment.
	//
	// This member is required.
	BulkDeploymentId *string
}

type GetBulkDeploymentStatusOutput struct {

	// Relevant metrics on input records processed during bulk deployment.
	BulkDeploymentMetrics *types.BulkDeploymentMetrics

	// The status of the bulk deployment.
	BulkDeploymentStatus types.BulkDeploymentStatus

	// The time, in ISO format, when the deployment was created.
	CreatedAt *string

	// Error details
	ErrorDetails []*types.ErrorDetail

	// Error message
	ErrorMessage *string

	// Tag(s) attached to the resource arn.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBulkDeploymentStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBulkDeploymentStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBulkDeploymentStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBulkDeploymentStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetBulkDeploymentStatus",
	}
}
