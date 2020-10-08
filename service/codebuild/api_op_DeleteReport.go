// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a report.
func (c *Client) DeleteReport(ctx context.Context, params *DeleteReportInput, optFns ...func(*Options)) (*DeleteReportOutput, error) {
	if params == nil {
		params = &DeleteReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReport", params, optFns, addOperationDeleteReportMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteReportInput struct {

	// The ARN of the report to delete.
	//
	// This member is required.
	Arn *string
}

type DeleteReportOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteReportMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteReport{}, middleware.After)
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
	addOpDeleteReportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReport(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteReport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "DeleteReport",
	}
}
