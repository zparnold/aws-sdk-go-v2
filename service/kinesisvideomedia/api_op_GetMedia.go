// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideomedia

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Use this API to retrieve media content from a Kinesis video stream. In the
// request, you identify the stream name or stream Amazon Resource Name (ARN), and
// the starting chunk. Kinesis Video Streams then returns a stream of chunks in
// order by fragment number. You must first call the GetDataEndpoint API to get an
// endpoint. Then send the GetMedia requests to this endpoint using the
// --endpoint-url parameter (https://docs.aws.amazon.com/cli/latest/reference/).
// When you put media data (fragments) on a stream, Kinesis Video Streams stores
// each incoming fragment and related metadata in what is called a "chunk." For
// more information, see PutMedia
// (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_dataplane_PutMedia.html).
// The GetMedia API returns a stream of these chunks starting from the chunk that
// you specify in the request. The following limits apply when using the GetMedia
// API:
//
//     * A client can call GetMedia up to five times per second per stream.
//
//
// * Kinesis Video Streams sends media data at a rate of up to 25 megabytes per
// second (or 200 megabits per second) during a GetMedia session.
//
//     <note> <p>If
// an error is thrown after invoking a Kinesis Video Streams media API, in addition
// to the HTTP status code and the response body, it includes the following pieces
// of information: </p> <ul> <li> <p> <code>x-amz-ErrorType</code> HTTP header –
// contains a more specific error type in addition to what the HTTP status code
// provides. </p> </li> <li> <p> <code>x-amz-RequestId</code> HTTP header – if you
// want to report an issue to AWS, the support team can better diagnose the problem
// if given the Request Id.</p> </li> </ul> <p>Both the HTTP status code and the
// ErrorType header can be utilized to make programmatic decisions about whether
// errors are retry-able and under what conditions, as well as provide information
// on what actions the client programmer might need to take in order to
// successfully try again.</p> <p>For more information, see the <b>Errors</b>
// section at the bottom of this topic, as well as <a
// href="https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html">Common
// Errors</a>. </p> </note>
func (c *Client) GetMedia(ctx context.Context, params *GetMediaInput, optFns ...func(*Options)) (*GetMediaOutput, error) {
	stack := middleware.NewStack("GetMedia", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetMediaMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpGetMediaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMedia(options.Region), middleware.Before)
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
			OperationName: "GetMedia",
			Err:           err,
		}
	}
	out := result.(*GetMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMediaInput struct {

	// Identifies the starting chunk to get from the specified stream.
	//
	// This member is required.
	StartSelector *types.StartSelector

	// The ARN of the stream from where you want to get the media content. If you don't
	// specify the streamARN, you must specify the streamName.
	StreamARN *string

	// The Kinesis video stream name from where you want to get the media content. If
	// you don't specify the streamName, you must specify the streamARN.
	StreamName *string
}

type GetMediaOutput struct {

	// The content type of the requested media.
	ContentType *string

	// The payload Kinesis Video Streams returns is a sequence of chunks from the
	// specified stream. For information about the chunks, see . The chunks that
	// Kinesis Video Streams returns in the GetMedia call also include the following
	// additional Matroska (MKV) tags:
	//
	//     * AWS_KINESISVIDEO_CONTINUATION_TOKEN
	// (UTF-8 string) - In the event your GetMedia call terminates, you can use this
	// continuation token in your next request to get the next chunk where the last
	// request terminated.
	//
	//     * AWS_KINESISVIDEO_MILLIS_BEHIND_NOW (UTF-8 string) -
	// Client applications can use this tag value to determine how far behind the chunk
	// returned in the response is from the latest chunk on the stream.
	//
	//     *
	// AWS_KINESISVIDEO_FRAGMENT_NUMBER - Fragment number returned in the chunk.
	//
	//     *
	// AWS_KINESISVIDEO_SERVER_TIMESTAMP - Server timestamp of the fragment.
	//
	//     *
	// AWS_KINESISVIDEO_PRODUCER_TIMESTAMP - Producer timestamp of the fragment.
	//
	// The
	// following tags will be present if an error occurs:
	//
	//     *
	// AWS_KINESISVIDEO_ERROR_CODE - String description of an error that caused
	// GetMedia to stop.
	//
	//     * AWS_KINESISVIDEO_ERROR_ID: Integer code of the
	// error.
	//
	// The error codes are as follows:
	//
	//     * 3002 - Error writing to the
	// stream
	//
	//     * 4000 - Requested fragment is not found
	//
	//     * 4500 - Access denied
	// for the stream's KMS key
	//
	//     * 4501 - Stream's KMS key is disabled
	//
	//     * 4502
	// - Validation error on the stream's KMS key
	//
	//     * 4503 - KMS key specified in
	// the stream is unavailable
	//
	//     * 4504 - Invalid usage of the KMS key specified
	// in the stream
	//
	//     * 4505 - Invalid state of the KMS key specified in the
	// stream
	//
	//     * 4506 - Unable to find the KMS key specified in the stream
	//
	//     *
	// 5000 - Internal error
	Payload io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetMediaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetMedia{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMedia{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMedia(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "GetMedia",
	}
}
