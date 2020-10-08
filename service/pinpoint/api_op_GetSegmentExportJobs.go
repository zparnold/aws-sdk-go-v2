// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the status and settings of the export jobs for a
// segment.
func (c *Client) GetSegmentExportJobs(ctx context.Context, params *GetSegmentExportJobsInput, optFns ...func(*Options)) (*GetSegmentExportJobsOutput, error) {
	if params == nil {
		params = &GetSegmentExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSegmentExportJobs", params, optFns, addOperationGetSegmentExportJobsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSegmentExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSegmentExportJobsInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the segment.
	//
	// This member is required.
	SegmentId *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string

	// The NextToken string that specifies which page of results to return in a
	// paginated response.
	Token *string
}

type GetSegmentExportJobsOutput struct {

	// Provides information about all the export jobs that are associated with an
	// application or segment. An export job is a job that exports endpoint definitions
	// to a file.
	//
	// This member is required.
	ExportJobsResponse *types.ExportJobsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSegmentExportJobsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSegmentExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSegmentExportJobs{}, middleware.After)
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
	addOpGetSegmentExportJobsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSegmentExportJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSegmentExportJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetSegmentExportJobs",
	}
}
