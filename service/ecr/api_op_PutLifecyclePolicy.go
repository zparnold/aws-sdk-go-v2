// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates the lifecycle policy for the specified repository. For more
// information, see Lifecycle Policy Template
// (https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).
func (c *Client) PutLifecyclePolicy(ctx context.Context, params *PutLifecyclePolicyInput, optFns ...func(*Options)) (*PutLifecyclePolicyOutput, error) {
	if params == nil {
		params = &PutLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecyclePolicy", params, optFns, addOperationPutLifecyclePolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecyclePolicyInput struct {

	// The JSON repository policy text to apply to the repository.
	//
	// This member is required.
	LifecyclePolicyText *string

	// The name of the repository to receive the policy.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the repository. If
	// you do  not specify a registry, the default registry is assumed.
	RegistryId *string
}

type PutLifecyclePolicyOutput struct {

	// The JSON repository policy text.
	LifecyclePolicyText *string

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutLifecyclePolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLifecyclePolicy{}, middleware.After)
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
	addOpPutLifecyclePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecyclePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutLifecyclePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "PutLifecyclePolicy",
	}
}
