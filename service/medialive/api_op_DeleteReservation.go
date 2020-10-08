// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete an expired reservation.
func (c *Client) DeleteReservation(ctx context.Context, params *DeleteReservationInput, optFns ...func(*Options)) (*DeleteReservationOutput, error) {
	if params == nil {
		params = &DeleteReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReservation", params, optFns, addOperationDeleteReservationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteReservationRequest
type DeleteReservationInput struct {

	// Unique reservation ID, e.g. '1234567'
	//
	// This member is required.
	ReservationId *string
}

// Placeholder documentation for DeleteReservationResponse
type DeleteReservationOutput struct {

	// Unique reservation ARN, e.g.
	// 'arn:aws:medialive:us-west-2:123456789012:reservation:1234567'
	Arn *string

	// Number of reserved resources
	Count *int32

	// Currency code for usagePrice and fixedPrice in ISO-4217 format, e.g. 'USD'
	CurrencyCode *string

	// Lease duration, e.g. '12'
	Duration *int32

	// Units for duration, e.g. 'MONTHS'
	DurationUnits types.OfferingDurationUnits

	// Reservation UTC end date and time in ISO-8601 format, e.g. '2019-03-01T00:00:00'
	End *string

	// One-time charge for each reserved resource, e.g. '0.0' for a NO_UPFRONT offering
	FixedPrice *float64

	// User specified reservation name
	Name *string

	// Offering description, e.g. 'HD AVC output at 10-20 Mbps, 30 fps, and standard VQ
	// in US West (Oregon)'
	OfferingDescription *string

	// Unique offering ID, e.g. '87654321'
	OfferingId *string

	// Offering type, e.g. 'NO_UPFRONT'
	OfferingType types.OfferingType

	// AWS region, e.g. 'us-west-2'
	Region *string

	// Unique reservation ID, e.g. '1234567'
	ReservationId *string

	// Resource configuration details
	ResourceSpecification *types.ReservationResourceSpecification

	// Reservation UTC start date and time in ISO-8601 format, e.g.
	// '2018-03-01T00:00:00'
	Start *string

	// Current state of reservation, e.g. 'ACTIVE'
	State types.ReservationState

	// A collection of key-value pairs
	Tags map[string]*string

	// Recurring usage charge for each reserved resource, e.g. '157.0'
	UsagePrice *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteReservationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteReservation{}, middleware.After)
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
	addOpDeleteReservationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReservation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteReservation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteReservation",
	}
}
