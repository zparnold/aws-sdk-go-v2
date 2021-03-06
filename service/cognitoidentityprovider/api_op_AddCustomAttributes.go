// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds additional user attributes to the user pool schema.
func (c *Client) AddCustomAttributes(ctx context.Context, params *AddCustomAttributesInput, optFns ...func(*Options)) (*AddCustomAttributesOutput, error) {
	stack := middleware.NewStack("AddCustomAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddCustomAttributesMiddlewares(stack)
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
	addOpAddCustomAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddCustomAttributes(options.Region), middleware.Before)
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
			OperationName: "AddCustomAttributes",
			Err:           err,
		}
	}
	out := result.(*AddCustomAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to add custom attributes.
type AddCustomAttributesInput struct {

	// An array of custom attributes, such as Mutable and Name.
	//
	// This member is required.
	CustomAttributes []*types.SchemaAttributeType

	// The user pool ID for the user pool where you want to add custom attributes.
	//
	// This member is required.
	UserPoolId *string
}

// Represents the response from the server for the request to add custom
// attributes.
type AddCustomAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddCustomAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddCustomAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddCustomAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddCustomAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AddCustomAttributes",
	}
}
