// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about all block storage disks in your AWS account and
// region.
func (c *Client) GetDisks(ctx context.Context, params *GetDisksInput, optFns ...func(*Options)) (*GetDisksOutput, error) {
	stack := middleware.NewStack("GetDisks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDisksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDisks(options.Region), middleware.Before)
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
			OperationName: "GetDisks",
			Err:           err,
		}
	}
	out := result.(*GetDisksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDisksInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetDisks request. If your results are paginated,
	// the response will return a next page token that you can specify as the page
	// token in a subsequent request.
	PageToken *string
}

type GetDisksOutput struct {

	// An array of objects containing information about all block storage disks.
	Disks []*types.Disk

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetDisks request and specify the next page
	// token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDisksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDisks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDisks{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDisks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetDisks",
	}
}
