// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is used by SaaS partners to write events to a customer's partner event bus.
// AWS customers do not use this operation.
func (c *Client) PutPartnerEvents(ctx context.Context, params *PutPartnerEventsInput, optFns ...func(*Options)) (*PutPartnerEventsOutput, error) {
	if params == nil {
		params = &PutPartnerEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPartnerEvents", params, optFns, addOperationPutPartnerEventsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPartnerEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPartnerEventsInput struct {

	// The list of events to write to the event bus.
	//
	// This member is required.
	Entries []*types.PutPartnerEventsRequestEntry
}

type PutPartnerEventsOutput struct {

	// The list of events from this operation that were successfully written to the
	// partner event bus.
	Entries []*types.PutPartnerEventsResultEntry

	// The number of events from this operation that could not be written to the
	// partner event bus.
	FailedEntryCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutPartnerEventsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPartnerEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPartnerEvents{}, middleware.After)
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
	addOpPutPartnerEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPartnerEvents(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutPartnerEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "PutPartnerEvents",
	}
}
