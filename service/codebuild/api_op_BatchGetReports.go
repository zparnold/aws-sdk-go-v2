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

// Returns an array of reports.
func (c *Client) BatchGetReports(ctx context.Context, params *BatchGetReportsInput, optFns ...func(*Options)) (*BatchGetReportsOutput, error) {
	if params == nil {
		params = &BatchGetReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetReports", params, optFns, addOperationBatchGetReportsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetReportsInput struct {

	// An array of ARNs that identify the Report objects to return.
	//
	// This member is required.
	ReportArns []*string
}

type BatchGetReportsOutput struct {

	// The array of Report objects returned by BatchGetReports.
	Reports []*types.Report

	// An array of ARNs passed to BatchGetReportGroups that are not associated with a
	// Report.
	ReportsNotFound []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetReportsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetReports{}, middleware.After)
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
	addOpBatchGetReportsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetReports(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchGetReports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "BatchGetReports",
	}
}
