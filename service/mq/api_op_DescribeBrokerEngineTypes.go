// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe available engine types and versions.
func (c *Client) DescribeBrokerEngineTypes(ctx context.Context, params *DescribeBrokerEngineTypesInput, optFns ...func(*Options)) (*DescribeBrokerEngineTypesOutput, error) {
	if params == nil {
		params = &DescribeBrokerEngineTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBrokerEngineTypes", params, optFns, addOperationDescribeBrokerEngineTypesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBrokerEngineTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBrokerEngineTypesInput struct {

	// Filter response by engine type.
	EngineType *string

	// The maximum number of engine types that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults *int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string
}

type DescribeBrokerEngineTypesOutput struct {

	// List of available engine types and versions.
	BrokerEngineTypes []*types.BrokerEngineType

	// Required. The maximum number of engine types that can be returned per page (20
	// by default). This value must be an integer from 5 to 100.
	MaxResults *int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBrokerEngineTypesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBrokerEngineTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBrokerEngineTypes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBrokerEngineTypes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBrokerEngineTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeBrokerEngineTypes",
	}
}
