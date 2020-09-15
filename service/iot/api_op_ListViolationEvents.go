// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the Device Defender security profile violations discovered during the
// given time period. You can use filters to limit the results to those alerts
// issued for a particular security profile, behavior, or thing (device).
func (c *Client) ListViolationEvents(ctx context.Context, params *ListViolationEventsInput, optFns ...func(*Options)) (*ListViolationEventsOutput, error) {
	stack := middleware.NewStack("ListViolationEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListViolationEventsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListViolationEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListViolationEvents(options.Region), middleware.Before)

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
			OperationName: "ListViolationEvents",
			Err:           err,
		}
	}
	out := result.(*ListViolationEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListViolationEventsInput struct {
	// The maximum number of results to return at one time.
	MaxResults *int32
	// The end time for the alerts to be listed.
	EndTime *time.Time
	// A filter to limit results to those alerts generated by the specified security
	// profile.
	SecurityProfileName *string
	// A filter to limit results to those alerts caused by the specified thing.
	ThingName *string
	// The start time for the alerts to be listed.
	StartTime *time.Time
	// The token for the next set of results.
	NextToken *string
}

type ListViolationEventsOutput struct {
	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string
	// The security profile violation alerts issued for this account during the given
	// time period, potentially filtered by security profile, behavior violated, or
	// thing (device) violating.
	ViolationEvents []*types.ViolationEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListViolationEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListViolationEvents{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListViolationEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opListViolationEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListViolationEvents",
	}
}
