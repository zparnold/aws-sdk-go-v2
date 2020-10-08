// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of solutions that use the given dataset group. When a dataset
// group is not specified, all the solutions associated with the account are
// listed. The response provides the properties for each solution, including the
// Amazon Resource Name (ARN). For more information on solutions, see
// CreateSolution ().
func (c *Client) ListSolutions(ctx context.Context, params *ListSolutionsInput, optFns ...func(*Options)) (*ListSolutionsOutput, error) {
	if params == nil {
		params = &ListSolutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolutions", params, optFns, addOperationListSolutionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolutionsInput struct {

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// The maximum number of solutions to return.
	MaxResults *int32

	// A token returned from the previous call to ListSolutions for getting the next
	// set of solutions (if they exist).
	NextToken *string
}

type ListSolutionsOutput struct {

	// A token for getting the next set of solutions (if they exist).
	NextToken *string

	// A list of the current solutions.
	Solutions []*types.SolutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSolutionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSolutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSolutions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSolutions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSolutions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListSolutions",
	}
}
