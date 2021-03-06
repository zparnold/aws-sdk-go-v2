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

// Retrieves a list of runs for a specified canary.
func (c *Client) GetCanaryRuns(ctx context.Context, params *GetCanaryRunsInput, optFns ...func(*Options)) (*GetCanaryRunsOutput, error) {
	stack := middleware.NewStack("GetCanaryRuns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCanaryRunsMiddlewares(stack)
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
	addOpGetCanaryRunsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCanaryRuns(options.Region), middleware.Before)
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
			OperationName: "GetCanaryRuns",
			Err:           err,
		}
	}
	out := result.(*GetCanaryRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCanaryRunsInput struct {

	// The name of the canary that you want to see runs for.
	//
	// This member is required.
	Name *string

	// Specify this parameter to limit how many runs are returned each time you use the
	// GetCanaryRuns operation. If you omit this parameter, the default of 100 is used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent GetCanaryRuns operation to retrieve the next set of results.
	NextToken *string
}

type GetCanaryRunsOutput struct {

	// An array of structures. Each structure contains the details of one of the
	// retrieved canary runs.
	CanaryRuns []*types.CanaryRun

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent GetCanaryRuns operation to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCanaryRunsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCanaryRuns{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCanaryRuns{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCanaryRuns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "GetCanaryRuns",
	}
}
