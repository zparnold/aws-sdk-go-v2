// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of all security configurations.
func (c *Client) GetSecurityConfigurations(ctx context.Context, params *GetSecurityConfigurationsInput, optFns ...func(*Options)) (*GetSecurityConfigurationsOutput, error) {
	stack := middleware.NewStack("GetSecurityConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSecurityConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSecurityConfigurations(options.Region), middleware.Before)
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
			OperationName: "GetSecurityConfigurations",
			Err:           err,
		}
	}
	out := result.(*GetSecurityConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSecurityConfigurationsInput struct {

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetSecurityConfigurationsOutput struct {

	// A continuation token, if there are more security configurations to return.
	NextToken *string

	// A list of security configurations.
	SecurityConfigurations []*types.SecurityConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSecurityConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSecurityConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSecurityConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSecurityConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetSecurityConfigurations",
	}
}
