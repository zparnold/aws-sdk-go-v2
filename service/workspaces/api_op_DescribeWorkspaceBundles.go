// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes the available WorkSpace bundles. You can filter
// the results using either bundle ID or owner, but not both.
func (c *Client) DescribeWorkspaceBundles(ctx context.Context, params *DescribeWorkspaceBundlesInput, optFns ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error) {
	stack := middleware.NewStack("DescribeWorkspaceBundles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeWorkspaceBundlesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceBundles(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkspaceBundles",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkspaceBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceBundlesInput struct {

	// The identifiers of the bundles. You cannot combine this parameter with any other
	// filter.
	BundleIds []*string

	// The token for the next set of results. (You received this token from a previous
	// call.)
	NextToken *string

	// The owner of the bundles. You cannot combine this parameter with any other
	// filter. Specify AMAZON to describe the bundles provided by AWS or null to
	// describe the bundles that belong to your account.
	Owner *string
}

type DescribeWorkspaceBundlesOutput struct {

	// Information about the bundles.
	Bundles []*types.WorkspaceBundle

	// The token to use to retrieve the next set of results, or null if there are no
	// more results available. This token is valid for one day and must be used within
	// that time frame.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeWorkspaceBundlesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceBundles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceBundles{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkspaceBundles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceBundles",
	}
}
