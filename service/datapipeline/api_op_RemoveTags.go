// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes existing tags from the specified pipeline.
func (c *Client) RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error) {
	if params == nil {
		params = &RemoveTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTags", params, optFns, addOperationRemoveTagsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RemoveTags.
type RemoveTagsInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// The keys of the tags to remove.
	//
	// This member is required.
	TagKeys []*string
}

// Contains the output of RemoveTags.
type RemoveTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveTagsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTags{}, middleware.After)
	if err != nil {
		return err
	}
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
	addOpRemoveTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "RemoveTags",
	}
}
