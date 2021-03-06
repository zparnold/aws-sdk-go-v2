// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List purchased reservations.
func (c *Client) ListReservations(ctx context.Context, params *ListReservationsInput, optFns ...func(*Options)) (*ListReservationsOutput, error) {
	stack := middleware.NewStack("ListReservations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListReservationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReservations(options.Region), middleware.Before)
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
			OperationName: "ListReservations",
			Err:           err,
		}
	}
	out := result.(*ListReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListReservationsRequest
type ListReservationsInput struct {

	// Filter by channel class, 'STANDARD' or 'SINGLE_PIPELINE'
	ChannelClass *string

	// Filter by codec, 'AVC', 'HEVC', 'MPEG2', or 'AUDIO'
	Codec *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Filter by bitrate, 'MAX_10_MBPS', 'MAX_20_MBPS', or 'MAX_50_MBPS'
	MaximumBitrate *string

	// Filter by framerate, 'MAX_30_FPS' or 'MAX_60_FPS'
	MaximumFramerate *string

	// Placeholder documentation for __string
	NextToken *string

	// Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'
	Resolution *string

	// Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'
	ResourceType *string

	// Filter by special feature, 'ADVANCED_AUDIO' or 'AUDIO_NORMALIZATION'
	SpecialFeature *string

	// Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'
	VideoQuality *string
}

// Placeholder documentation for ListReservationsResponse
type ListReservationsOutput struct {

	// Token to retrieve the next page of results
	NextToken *string

	// List of reservations
	Reservations []*types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListReservationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListReservations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListReservations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListReservations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListReservations",
	}
}
