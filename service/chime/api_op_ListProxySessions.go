// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the proxy sessions for the specified Amazon Chime Voice Connector.
func (c *Client) ListProxySessions(ctx context.Context, params *ListProxySessionsInput, optFns ...func(*Options)) (*ListProxySessionsOutput, error) {
	stack := middleware.NewStack("ListProxySessions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProxySessionsMiddlewares(stack)
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
	addOpListProxySessionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProxySessions(options.Region), middleware.Before)
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
			OperationName: "ListProxySessions",
			Err:           err,
		}
	}
	out := result.(*ListProxySessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProxySessionsInput struct {

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The proxy session status.
	Status types.ProxySessionStatus
}

type ListProxySessionsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The proxy session details.
	ProxySessions []*types.ProxySession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProxySessionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProxySessions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProxySessions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProxySessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListProxySessions",
	}
}
