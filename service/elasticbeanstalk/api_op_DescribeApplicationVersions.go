// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a list of application versions.
func (c *Client) DescribeApplicationVersions(ctx context.Context, params *DescribeApplicationVersionsInput, optFns ...func(*Options)) (*DescribeApplicationVersionsOutput, error) {
	if params == nil {
		params = &DescribeApplicationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicationVersions", params, optFns, addOperationDescribeApplicationVersionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe application versions.
type DescribeApplicationVersionsInput struct {

	// Specify an application name to show only application versions for that
	// application.
	ApplicationName *string

	// For a paginated request. Specify a maximum number of application versions to
	// include in each response. If no MaxRecords is specified, all available
	// application versions are retrieved in a single response.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string

	// Specify a version label to show a specific application version.
	VersionLabels []*string
}

// Result message wrapping a list of application version descriptions.
type DescribeApplicationVersionsOutput struct {

	// List of ApplicationVersionDescription objects sorted in order of creation.
	ApplicationVersions []*types.ApplicationVersionDescription

	// In a paginated request, the token that you can pass in a subsequent request to
	// get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeApplicationVersionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeApplicationVersions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeApplicationVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeApplicationVersions",
	}
}
