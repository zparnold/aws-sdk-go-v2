// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// DescribeExportConfigurations is deprecated. Use DescribeImportTasks
// (https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html),
// instead.
func (c *Client) DescribeExportConfigurations(ctx context.Context, params *DescribeExportConfigurationsInput, optFns ...func(*Options)) (*DescribeExportConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeExportConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExportConfigurations", params, optFns, addOperationDescribeExportConfigurationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportConfigurationsInput struct {

	// A list of continuous export IDs to search for.
	ExportIds []*string

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults *int32

	// The token from the previous call to describe-export-tasks.
	NextToken *string
}

type DescribeExportConfigurationsOutput struct {

	//
	ExportsInfo []*types.ExportInfo

	// The token from the previous call to describe-export-tasks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeExportConfigurationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExportConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExportConfigurations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportConfigurations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeExportConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeExportConfigurations",
	}
}
