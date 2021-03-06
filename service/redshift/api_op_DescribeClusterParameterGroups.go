// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Amazon Redshift parameter groups, including parameter groups
// you created and the default parameter group. For each parameter group, the
// response includes the parameter group name, description, and parameter group
// family name. You can optionally specify a name to retrieve the description of a
// specific parameter group. For more information about parameters and parameter
// groups, go to Amazon Redshift Parameter Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide. If you specify both tag keys
// and tag values in the same request, Amazon Redshift returns all parameter groups
// that match any combination of the specified keys and values. For example, if you
// have owner and environment for tag keys, and admin and test for tag values, all
// parameter groups that have any combination of those values are returned. If both
// tag keys and values are omitted from the request, parameter groups are returned
// regardless of whether they have tag keys or values associated with them.
func (c *Client) DescribeClusterParameterGroups(ctx context.Context, params *DescribeClusterParameterGroupsInput, optFns ...func(*Options)) (*DescribeClusterParameterGroupsOutput, error) {
	stack := middleware.NewStack("DescribeClusterParameterGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeClusterParameterGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterParameterGroups(options.Region), middleware.Before)
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
			OperationName: "DescribeClusterParameterGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeClusterParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeClusterParameterGroupsInput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterParameterGroups ()
	// request exceed the value specified in MaxRecords, AWS returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The name of a specific parameter group for which to return details. By default,
	// details about all parameter groups and the default parameter group are returned.
	ParameterGroupName *string

	// A tag key or keys for which you want to return all matching cluster parameter
	// groups that are associated with the specified key or keys. For example, suppose
	// that you have parameter groups that are tagged with keys called owner and
	// environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the parameter groups that have either or both
	// of these tag keys associated with them.
	TagKeys []*string

	// A tag value or values for which you want to return all matching cluster
	// parameter groups that are associated with the specified tag value or values. For
	// example, suppose that you have parameter groups that are tagged with values
	// called admin and test. If you specify both of these tag values in the request,
	// Amazon Redshift returns a response with the parameter groups that have either or
	// both of these tag values associated with them.
	TagValues []*string
}

// Contains the output from the DescribeClusterParameterGroups () action.
type DescribeClusterParameterGroupsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of ClusterParameterGroup () instances. Each instance describes one
	// cluster parameter group.
	ParameterGroups []*types.ClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeClusterParameterGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterParameterGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterParameterGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusterParameterGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterParameterGroups",
	}
}
