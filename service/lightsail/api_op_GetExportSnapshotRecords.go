// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the export snapshot record created as a result of the export snapshot
// operation. An export snapshot record can be used to create a new Amazon EC2
// instance and its related resources with the create cloud formation stack
// operation.
func (c *Client) GetExportSnapshotRecords(ctx context.Context, params *GetExportSnapshotRecordsInput, optFns ...func(*Options)) (*GetExportSnapshotRecordsOutput, error) {
	if params == nil {
		params = &GetExportSnapshotRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExportSnapshotRecords", params, optFns, addOperationGetExportSnapshotRecordsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExportSnapshotRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExportSnapshotRecordsInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetExportSnapshotRecords request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string
}

type GetExportSnapshotRecordsOutput struct {

	// A list of objects describing the export snapshot records.
	ExportSnapshotRecords []*types.ExportSnapshotRecord

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetExportSnapshotRecords request and specify
	// the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetExportSnapshotRecordsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetExportSnapshotRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExportSnapshotRecords{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetExportSnapshotRecords(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetExportSnapshotRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetExportSnapshotRecords",
	}
}
