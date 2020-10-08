// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Integration.
func (c *Client) DeleteIntegration(ctx context.Context, params *DeleteIntegrationInput, optFns ...func(*Options)) (*DeleteIntegrationOutput, error) {
	if params == nil {
		params = &DeleteIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIntegration", params, optFns, addOperationDeleteIntegrationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIntegrationInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The integration ID.
	//
	// This member is required.
	IntegrationId *string
}

type DeleteIntegrationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteIntegrationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIntegration{}, middleware.After)
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
	addOpDeleteIntegrationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIntegration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteIntegration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteIntegration",
	}
}
