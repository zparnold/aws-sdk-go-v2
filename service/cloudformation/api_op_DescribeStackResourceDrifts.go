// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns drift information for the resources that have been checked for drift in
// the specified stack. This includes actual and expected configuration values for
// resources where AWS CloudFormation detects configuration drift. For a given
// stack, there will be one StackResourceDrift for each stack resource that has
// been checked for drift. Resources that have not yet been checked for drift are
// not included. Resources that do not currently support drift detection are not
// checked, and so not included. For a list of resources that support drift
// detection, see Resources that Support Drift Detection
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html).
// Use DetectStackResourceDrift () to detect drift on individual resources, or
// DetectStackDrift () to detect drift on all supported resources for a given
// stack.
func (c *Client) DescribeStackResourceDrifts(ctx context.Context, params *DescribeStackResourceDriftsInput, optFns ...func(*Options)) (*DescribeStackResourceDriftsOutput, error) {
	stack := middleware.NewStack("DescribeStackResourceDrifts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeStackResourceDriftsMiddlewares(stack)
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
	addOpDescribeStackResourceDriftsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackResourceDrifts(options.Region), middleware.Before)
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
			OperationName: "DescribeStackResourceDrifts",
			Err:           err,
		}
	}
	out := result.(*DescribeStackResourceDriftsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackResourceDriftsInput struct {

	// The name of the stack for which you want drift information.
	//
	// This member is required.
	StackName *string

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// A string that identifies the next page of stack resource drift results.
	NextToken *string

	// The resource drift status values to use as filters for the resource drift
	// results returned.
	//
	//     * DELETED: The resource differs from its expected
	// template configuration in that the resource has been deleted.
	//
	//     * MODIFIED:
	// One or more resource properties differ from their expected template values.
	//
	//
	// * IN_SYNC: The resources's actual configuration matches its expected template
	// configuration.
	//
	//     * NOT_CHECKED: AWS CloudFormation does not currently return
	// this value.
	StackResourceDriftStatusFilters []types.StackResourceDriftStatus
}

type DescribeStackResourceDriftsOutput struct {

	// Drift information for the resources that have been checked for drift in the
	// specified stack. This includes actual and expected configuration values for
	// resources where AWS CloudFormation detects drift. For a given stack, there will
	// be one StackResourceDrift for each stack resource that has been checked for
	// drift. Resources that have not yet been checked for drift are not included.
	// Resources that do not currently support drift detection are not checked, and so
	// not included. For a list of resources that support drift detection, see
	// Resources that Support Drift Detection
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html).
	//
	// This member is required.
	StackResourceDrifts []*types.StackResourceDrift

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call DescribeStackResourceDrifts
	// again and assign that token to the request object's NextToken parameter. If the
	// request returns all results, NextToken is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeStackResourceDriftsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackResourceDrifts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackResourceDrifts{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStackResourceDrifts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackResourceDrifts",
	}
}
