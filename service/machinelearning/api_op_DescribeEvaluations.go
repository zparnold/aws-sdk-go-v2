// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of DescribeEvaluations that match the search criteria in the
// request.
func (c *Client) DescribeEvaluations(ctx context.Context, params *DescribeEvaluationsInput, optFns ...func(*Options)) (*DescribeEvaluationsOutput, error) {
	stack := middleware.NewStack("DescribeEvaluations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEvaluationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvaluations(options.Region), middleware.Before)
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
			OperationName: "DescribeEvaluations",
			Err:           err,
		}
	}
	out := result.(*DescribeEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEvaluationsInput struct {

	// The equal to operator. The Evaluation results will have FilterVariable values
	// that exactly match the value specified with EQ.
	EQ *string

	// Use one of the following variable to filter a list of Evaluation objects:
	//
	//     *
	// CreatedAt - Sets the search criteria to the Evaluation creation date.
	//
	//     *
	// Status - Sets the search criteria to the Evaluation status.
	//
	//     * Name - Sets
	// the search criteria to the contents of EvaluationName.
	//
	//     * IAMUser - Sets the
	// search criteria to the user account that invoked an Evaluation.
	//
	//     * MLModelId
	// - Sets the search criteria to the MLModel that was evaluated.
	//
	//     *
	// DataSourceId - Sets the search criteria to the DataSource used in Evaluation.
	//
	//
	// * DataUri - Sets the search criteria to the data file(s) used in Evaluation. The
	// URL can identify either a file or an Amazon Simple Storage Solution (Amazon S3)
	// bucket or directory.
	FilterVariable types.EvaluationFilterVariable

	// The greater than or equal to operator. The Evaluation results will have
	// FilterVariable values that are greater than or equal to the value specified with
	// GE.
	GE *string

	// The greater than operator. The Evaluation results will have FilterVariable
	// values that are greater than the value specified with GT.
	GT *string

	// The less than or equal to operator. The Evaluation results will have
	// FilterVariable values that are less than or equal to the value specified with
	// LE.
	LE *string

	// The less than operator. The Evaluation results will have FilterVariable values
	// that are less than the value specified with LT.
	LT *string

	// The maximum number of Evaluation to include in the result.
	Limit *int32

	// The not equal to operator. The Evaluation results will have FilterVariable
	// values not equal to the value specified with NE.
	NE *string

	// The ID of the page in the paginated results.
	NextToken *string

	// A string that is found at the beginning of a variable, such as Name or Id. For
	// example, an Evaluation could have the Name2014-09-09-HolidayGiftMailer. To
	// search for this Evaluation, select Name for the FilterVariable and any of the
	// following strings for the Prefix:  <ul> <li> <p>2014-09</p> </li> <li>
	// <p>2014-09-09</p> </li> <li> <p>2014-09-09-Holiday</p> </li> </ul>
	Prefix *string

	// A two-value parameter that determines the sequence of the resulting list of
	// Evaluation.
	//
	//     * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//     *
	// dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by
	// FilterVariable.
	SortOrder types.SortOrder
}

// Represents the query results from a DescribeEvaluations operation. The content
// is essentially a list of Evaluation.
type DescribeEvaluationsOutput struct {

	// The ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string

	// A list of Evaluation that meet the search criteria.
	Results []*types.Evaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEvaluationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEvaluations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEvaluations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEvaluations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DescribeEvaluations",
	}
}
