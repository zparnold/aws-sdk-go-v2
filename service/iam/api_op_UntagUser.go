// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified tags from the user. For more information about tagging,
// see Tagging IAM Identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User
// Guide.
func (c *Client) UntagUser(ctx context.Context, params *UntagUserInput, optFns ...func(*Options)) (*UntagUserOutput, error) {
	if params == nil {
		params = &UntagUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagUser", params, optFns, addOperationUntagUserMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagUserInput struct {

	// A list of key names as a simple array of strings. The tags with matching keys
	// are removed from the specified user.
	//
	// This member is required.
	TagKeys []*string

	// The name of the IAM user from which you want to remove tags. This parameter
	// accepts (through its regex pattern (http://wikipedia.org/wiki/regex)) a string
	// of characters that consist of upper and lowercase alphanumeric characters with
	// no spaces. You can also include any of the following characters: =,.@-
	//
	// This member is required.
	UserName *string
}

type UntagUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUntagUserMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUntagUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUntagUser{}, middleware.After)
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
	addOpUntagUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntagUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUntagUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UntagUser",
	}
}
