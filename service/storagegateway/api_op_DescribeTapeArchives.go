// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of specified virtual tapes in the virtual tape shelf
// (VTS). This operation is only supported in the tape gateway type.  <p>If a
// specific <code>TapeARN</code> is not specified, AWS Storage Gateway returns a
// description of all virtual tapes found in the VTS associated with your
// account.</p>
func (c *Client) DescribeTapeArchives(ctx context.Context, params *DescribeTapeArchivesInput, optFns ...func(*Options)) (*DescribeTapeArchivesOutput, error) {
	stack := middleware.NewStack("DescribeTapeArchives", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTapeArchivesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapeArchives(options.Region), middleware.Before)

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
			OperationName: "DescribeTapeArchives",
			Err:           err,
		}
	}
	out := result.(*DescribeTapeArchivesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapeArchivesInput
type DescribeTapeArchivesInput struct {
	// Specifies that the number of virtual tapes described be limited to the specified
	// number.
	Limit *int32
	// An opaque string that indicates the position at which to begin describing
	// virtual tapes.
	Marker *string
	// Specifies one or more unique Amazon Resource Names (ARNs) that represent the
	// virtual tapes you want to describe.
	TapeARNs []*string
}

// DescribeTapeArchivesOutput
type DescribeTapeArchivesOutput struct {
	// An opaque string that indicates the position at which the virtual tapes that
	// were fetched for description ended. Use this marker in your next request to
	// fetch the next set of virtual tapes in the virtual tape shelf (VTS). If there
	// are no more virtual tapes to describe, this field does not appear in the
	// response.
	Marker *string
	// An array of virtual tape objects in the virtual tape shelf (VTS). The
	// description includes of the Amazon Resource Name (ARN) of the virtual tapes. The
	// information returned includes the Amazon Resource Names (ARNs) of the tapes,
	// size of the tapes, status of the tapes, progress of the description, and tape
	// barcode.
	TapeArchives []*types.TapeArchive

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTapeArchivesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapeArchives{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapeArchives{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTapeArchives(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeTapeArchives",
	}
}
