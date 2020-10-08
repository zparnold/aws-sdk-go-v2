// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns summary information about the versions of a type.
func (c *Client) ListTypeVersions(ctx context.Context, params *ListTypeVersionsInput, optFns ...func(*Options)) (*ListTypeVersionsOutput, error) {
	if params == nil {
		params = &ListTypeVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypeVersions", params, optFns, addOperationListTypeVersionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypeVersionsInput struct {

	// The Amazon Resource Name (ARN) of the type for which you want version summary
	// information. Conditional: You must specify either TypeName and Type, or Arn.
	Arn *string

	// The deprecation status of the type versions that you want to get summary
	// information about. Valid values include:
	//
	//     * LIVE: The type version is
	// registered and can be used in CloudFormation operations, dependent on its
	// provisioning behavior and visibility scope.
	//
	//     * DEPRECATED: The type version
	// has been deregistered and can no longer be used in CloudFormation
	// operations.
	//
	// The default is LIVE.
	DeprecatedStatus types.DeprecatedStatus

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	// The kind of the type. Currently the only valid value is RESOURCE. Conditional:
	// You must specify either TypeName and Type, or Arn.
	Type types.RegistryType

	// The name of the type for which you want version summary information.
	// Conditional: You must specify either TypeName and Type, or Arn.
	TypeName *string
}

type ListTypeVersionsOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call this action again and assign
	// that token to the request object's NextToken parameter. If the request returns
	// all results, NextToken is set to null.
	NextToken *string

	// A list of TypeVersionSummary structures that contain information about the
	// specified type's versions.
	TypeVersionSummaries []*types.TypeVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTypeVersionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListTypeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListTypeVersions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTypeVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTypeVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListTypeVersions",
	}
}
