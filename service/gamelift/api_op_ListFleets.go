// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a collection of fleet resources for this AWS account. You can filter
// the result set to find only those fleets that are deployed with a specific build
// or script. Use the pagination parameters to retrieve results in sequential
// pages. Fleet resources are not listed in a particular order. Learn more Setting
// up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * DescribeFleetAttributes ()
//
//     * UpdateFleetAttributes ()
//
//     *
// StartFleetActions () or StopFleetActions ()
func (c *Client) ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if params == nil {
		params = &ListFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleets", params, optFns, addOperationListFleetsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type ListFleetsInput struct {

	// A unique identifier for a build to return fleets for. Use this parameter to
	// return only fleets using a specified build. Use either the build ID or ARN
	// value. To retrieve all fleets, do not include either a BuildId and ScriptID
	// parameter.
	BuildId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// A unique identifier for a Realtime script to return fleets for. Use this
	// parameter to return only fleets using a specified script. Use either the script
	// ID or ARN value. To retrieve all fleets, leave this parameter empty.
	ScriptId *string
}

// Represents the returned data in response to a request action.
type ListFleetsOutput struct {

	// Set of fleet IDs matching the list request. You can retrieve additional
	// information about all returned fleets by passing this result set to a call to
	// DescribeFleetAttributes (), DescribeFleetCapacity (), or
	// DescribeFleetUtilization ().
	FleetIds []*string

	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFleetsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFleets{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFleets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListFleets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListFleets",
	}
}
