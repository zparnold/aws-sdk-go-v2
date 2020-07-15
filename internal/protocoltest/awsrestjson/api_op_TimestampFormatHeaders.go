// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how timestamp request and response headers are serialized.
func (c *Client) TimestampFormatHeaders(ctx context.Context, params *TimestampFormatHeadersInput, optFns ...func(*Options)) (*TimestampFormatHeadersOutput, error) {
	stack := middleware.NewStack("TimestampFormatHeaders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTimestampFormatHeaders(options.Region), middleware.Before)
	addawsRestjson1_serdeOpTimestampFormatHeadersMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "TimestampFormatHeaders",
			Err:           err,
		}
	}
	out := result.(*TimestampFormatHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TimestampFormatHeadersInput struct {
	MemberEpochSeconds *time.Time
	MemberHttpDate     *time.Time
	MemberDateTime     *time.Time
	DefaultFormat      *time.Time
	TargetEpochSeconds *time.Time
	TargetHttpDate     *time.Time
	TargetDateTime     *time.Time
}

type TimestampFormatHeadersOutput struct {
	MemberEpochSeconds *time.Time
	MemberHttpDate     *time.Time
	MemberDateTime     *time.Time
	DefaultFormat      *time.Time
	TargetEpochSeconds *time.Time
	TargetHttpDate     *time.Time
	TargetDateTime     *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTimestampFormatHeadersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTimestampFormatHeaders{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTimestampFormatHeaders{}, middleware.After)
}

func newServiceMetadataMiddleware_opTimestampFormatHeaders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "TimestampFormatHeaders",
	}
}
