// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a report group. A report group contains a collection of reports.
func (c *Client) CreateReportGroup(ctx context.Context, params *CreateReportGroupInput, optFns ...func(*Options)) (*CreateReportGroupOutput, error) {
	if params == nil {
		params = &CreateReportGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReportGroup", params, optFns, addOperationCreateReportGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReportGroupInput struct {

	// A ReportExportConfig object that contains information about where the report
	// group test results are exported.
	//
	// This member is required.
	ExportConfig *types.ReportExportConfig

	// The name of the report group.
	//
	// This member is required.
	Name *string

	// The type of report group.
	//
	// This member is required.
	Type types.ReportType

	// A list of tag key and value pairs associated with this report group. These tags
	// are available for use by AWS services that support AWS CodeBuild report group
	// tags.
	Tags []*types.Tag
}

type CreateReportGroupOutput struct {

	// Information about the report group that was created.
	ReportGroup *types.ReportGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateReportGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReportGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReportGroup{}, middleware.After)
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
	addOpCreateReportGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReportGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateReportGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "CreateReportGroup",
	}
}
