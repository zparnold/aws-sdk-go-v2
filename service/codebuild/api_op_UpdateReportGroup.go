// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a report group.
func (c *Client) UpdateReportGroup(ctx context.Context, params *UpdateReportGroupInput, optFns ...func(*Options)) (*UpdateReportGroupOutput, error) {
	stack := middleware.NewStack("UpdateReportGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateReportGroupMiddlewares(stack)
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
	addOpUpdateReportGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReportGroup(options.Region), middleware.Before)
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
			OperationName: "UpdateReportGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReportGroupInput struct {

	// The ARN of the report group to update.
	//
	// This member is required.
	Arn *string

	// Used to specify an updated export type. Valid values are:
	//
	//     * S3: The report
	// results are exported to an S3 bucket.
	//
	//     * NO_EXPORT: The report results are
	// not exported.
	ExportConfig *types.ReportExportConfig

	// An updated list of tag key and value pairs associated with this report group.
	// These tags are available for use by AWS services that support AWS CodeBuild
	// report group tags.
	Tags []*types.Tag
}

type UpdateReportGroupOutput struct {

	// Information about the updated report group.
	ReportGroup *types.ReportGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateReportGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateReportGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateReportGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateReportGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "UpdateReportGroup",
	}
}
