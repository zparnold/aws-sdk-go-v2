// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of ARNs of the resources that are members of a specified resource
// group.
func (c *Client) ListGroupResources(ctx context.Context, params *ListGroupResourcesInput, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) {
	stack := middleware.NewStack("ListGroupResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListGroupResourcesMiddlewares(stack)
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
	addOpListGroupResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupResources(options.Region), middleware.Before)
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
			OperationName: "ListGroupResources",
			Err:           err,
		}
	}
	out := result.(*ListGroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupResourcesInput struct {

	// Filters, formatted as ResourceFilter () objects, that you want to apply to a
	// ListGroupResources operation. Filters the results to include only those of the
	// specified resource types.
	//
	//     * resource-type - Filter resources by their type.
	// Specify up to five resource types in the format AWS::ServiceCode::ResourceType.
	// For example, AWS::EC2::Instance, or AWS::S3::Bucket.
	//
	// When you specify a
	// resource-type filter for ListGroupResources, AWS Resource Groups validates your
	// filter resource types against the types that are defined in the query associated
	// with the group. For example, if a group contains only S3 buckets because its
	// query specifies only that resource type, but your resource-type filter includes
	// EC2 instances, AWS Resource Groups does not filter for EC2 instances. In this
	// case, a ListGroupResources request returns a BadRequestException error with a
	// message similar to the following: The resource types specified as filters in the
	// request are not valid. The error includes a list of resource types that failed
	// the validation because they are not part of the query associated with the group.
	// This validation doesn't occur when the group query specifies AWS::AllSupported,
	// because a group based on such a query can contain any of the allowed resource
	// types for the query type (tag-based or AWS CloudFormation stack-based queries).
	Filters []*types.ResourceFilter

	// The name or the ARN of the resource group
	Group *string

	// Don't use this parameter. Use Group instead.
	GroupName *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that the service might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value provided by a previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string
}

type ListGroupResourcesOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// A list of QueryError objects. Each error is an object that contains ErrorCode
	// and Message structures. Possible values for ErrorCode are
	// CLOUDFORMATION_STACK_INACTIVE and CLOUDFORMATION_STACK_NOT_EXISTING.
	QueryErrors []*types.QueryError

	// The ARNs and resource types of resources that are members of the group that you
	// specified.
	ResourceIdentifiers []*types.ResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListGroupResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListGroupResources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroupResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListGroupResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "ListGroupResources",
	}
}
