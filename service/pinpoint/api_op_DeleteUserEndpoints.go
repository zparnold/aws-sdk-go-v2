// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes all the endpoints that are associated with a specific user ID.
func (c *Client) DeleteUserEndpoints(ctx context.Context, params *DeleteUserEndpointsInput, optFns ...func(*Options)) (*DeleteUserEndpointsOutput, error) {
	stack := middleware.NewStack("DeleteUserEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteUserEndpointsMiddlewares(stack)
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
	addOpDeleteUserEndpointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserEndpoints(options.Region), middleware.Before)
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
			OperationName: "DeleteUserEndpoints",
			Err:           err,
		}
	}
	out := result.(*DeleteUserEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserEndpointsInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the user.
	//
	// This member is required.
	UserId *string
}

type DeleteUserEndpointsOutput struct {

	// Provides information about all the endpoints that are associated with a user ID.
	//
	// This member is required.
	EndpointsResponse *types.EndpointsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteUserEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteUserEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteUserEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUserEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteUserEndpoints",
	}
}
