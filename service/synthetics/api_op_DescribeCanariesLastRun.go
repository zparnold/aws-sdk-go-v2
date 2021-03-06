// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to see information from the most recent run of each canary
// that you have created.
func (c *Client) DescribeCanariesLastRun(ctx context.Context, params *DescribeCanariesLastRunInput, optFns ...func(*Options)) (*DescribeCanariesLastRunOutput, error) {
	stack := middleware.NewStack("DescribeCanariesLastRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeCanariesLastRunMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCanariesLastRun(options.Region), middleware.Before)
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
			OperationName: "DescribeCanariesLastRun",
			Err:           err,
		}
	}
	out := result.(*DescribeCanariesLastRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCanariesLastRunInput struct {

	// Specify this parameter to limit how many runs are returned each time you use the
	// DescribeLastRun operation. If you omit this parameter, the default of 100 is
	// used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent DescribeCanaries operation to retrieve the next set of results.
	NextToken *string
}

type DescribeCanariesLastRunOutput struct {

	// An array that contains the information from the most recent run of each canary.
	CanariesLastRun []*types.CanaryLastRun

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent DescribeCanariesLastRun operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeCanariesLastRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCanariesLastRun{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCanariesLastRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCanariesLastRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "DescribeCanariesLastRun",
	}
}
