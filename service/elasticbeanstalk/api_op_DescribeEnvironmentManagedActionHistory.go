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

// Lists an environment's completed and failed managed actions.
func (c *Client) DescribeEnvironmentManagedActionHistory(ctx context.Context, params *DescribeEnvironmentManagedActionHistoryInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionHistoryOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentManagedActionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentManagedActionHistory", params, optFns, addOperationDescribeEnvironmentManagedActionHistoryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentManagedActionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list completed and failed managed actions.
type DescribeEnvironmentManagedActionHistoryInput struct {

	// The environment ID of the target environment.
	EnvironmentId *string

	// The name of the target environment.
	EnvironmentName *string

	// The maximum number of items to return for a single request.
	MaxItems *int32

	// The pagination token returned by a previous request.
	NextToken *string
}

// A result message containing a list of completed and failed managed actions.
type DescribeEnvironmentManagedActionHistoryOutput struct {

	// A list of completed and failed managed actions.
	ManagedActionHistoryItems []*types.ManagedActionHistoryItem

	// A pagination token that you pass to DescribeEnvironmentManagedActionHistory ()
	// to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentManagedActionHistoryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentManagedActionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentManagedActionHistory{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentManagedActionHistory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEnvironmentManagedActionHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentManagedActionHistory",
	}
}
