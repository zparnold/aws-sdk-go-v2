// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a TagOption.
func (c *Client) CreateTagOption(ctx context.Context, params *CreateTagOptionInput, optFns ...func(*Options)) (*CreateTagOptionOutput, error) {
	if params == nil {
		params = &CreateTagOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTagOption", params, optFns, addOperationCreateTagOptionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTagOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTagOptionInput struct {

	// The TagOption key.
	//
	// This member is required.
	Key *string

	// The TagOption value.
	//
	// This member is required.
	Value *string
}

type CreateTagOptionOutput struct {

	// Information about the TagOption.
	TagOptionDetail *types.TagOptionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTagOptionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTagOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTagOption{}, middleware.After)
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
	addOpCreateTagOptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTagOption(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTagOption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateTagOption",
	}
}
