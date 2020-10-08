// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a GraphqlApi object.
func (c *Client) UpdateGraphqlApi(ctx context.Context, params *UpdateGraphqlApiInput, optFns ...func(*Options)) (*UpdateGraphqlApiOutput, error) {
	if params == nil {
		params = &UpdateGraphqlApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGraphqlApi", params, optFns, addOperationUpdateGraphqlApiMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGraphqlApiInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The new name for the GraphqlApi object.
	//
	// This member is required.
	Name *string

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []*types.AdditionalAuthenticationProvider

	// The new authentication type for the GraphqlApi object.
	AuthenticationType types.AuthenticationType

	// The Amazon CloudWatch Logs configuration for the GraphqlApi object.
	LogConfig *types.LogConfig

	// The OpenID Connect configuration for the GraphqlApi object.
	OpenIDConnectConfig *types.OpenIDConnectConfig

	// The new Amazon Cognito user pool configuration for the GraphqlApi object.
	UserPoolConfig *types.UserPoolConfig

	// A flag indicating whether to enable X-Ray tracing for the GraphqlApi.
	XrayEnabled *bool
}

type UpdateGraphqlApiOutput struct {

	// The updated GraphqlApi object.
	GraphqlApi *types.GraphqlApi

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGraphqlApiMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGraphqlApi{}, middleware.After)
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
	addOpUpdateGraphqlApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGraphqlApi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGraphqlApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "UpdateGraphqlApi",
	}
}
