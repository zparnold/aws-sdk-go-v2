// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a resource policy that is identified by its resource ARN.
func (c *Client) DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) {
	if params == nil {
		params = &DeleteResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResourcePolicy", params, optFns, addOperationDeleteResourcePolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResourcePolicyInput struct {

	// The ARN of the resource that is associated with the resource policy.
	//
	// This member is required.
	ResourceArn *string
}

type DeleteResourcePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteResourcePolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteResourcePolicy{}, middleware.After)
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
	addOpDeleteResourcePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResourcePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteResourcePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "DeleteResourcePolicy",
	}
}
