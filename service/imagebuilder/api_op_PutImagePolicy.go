// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies a policy to an image. We recommend that you call the RAM API
// CreateResourceShare
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html)
// to share resources. If you call the Image Builder API PutImagePolicy, you must
// also call the RAM API PromoteResourceShareCreatedFromPolicy
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html)
// in order for the resource to be visible to all principals with whom the resource
// is shared.
func (c *Client) PutImagePolicy(ctx context.Context, params *PutImagePolicyInput, optFns ...func(*Options)) (*PutImagePolicyOutput, error) {
	stack := middleware.NewStack("PutImagePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutImagePolicyMiddlewares(stack)
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
	addOpPutImagePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutImagePolicy(options.Region), middleware.Before)

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
			OperationName: "PutImagePolicy",
			Err:           err,
		}
	}
	out := result.(*PutImagePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutImagePolicyInput struct {
	// The Amazon Resource Name (ARN) of the image that this policy should be applied
	// to.
	ImageArn *string
	// The policy to apply.
	Policy *string
}

type PutImagePolicyOutput struct {
	// The Amazon Resource Name (ARN) of the image that this policy was applied to.
	ImageArn *string
	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutImagePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutImagePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutImagePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutImagePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "PutImagePolicy",
	}
}
