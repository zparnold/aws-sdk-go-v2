// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For a specified resource ID, this API action returns a list of compliance
// statuses for different resource types. Currently, you can only specify one
// resource ID per call. List results depend on the criteria specified in the
// filter.
func (c *Client) ListComplianceItems(ctx context.Context, params *ListComplianceItemsInput, optFns ...func(*Options)) (*ListComplianceItemsOutput, error) {
	if params == nil {
		params = &ListComplianceItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComplianceItems", params, optFns, addOperationListComplianceItemsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComplianceItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceItemsInput struct {

	// One or more compliance filters. Use a filter to return a more specific list of
	// results.
	Filters []*types.ComplianceStringFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// The ID for the resources from which to get compliance information. Currently,
	// you can only specify one resource ID.
	ResourceIds []*string

	// The type of resource from which to get compliance information. Currently, the
	// only supported resource type is ManagedInstance.
	ResourceTypes []*string
}

type ListComplianceItemsOutput struct {

	// A list of compliance information for the specified resource ID.
	ComplianceItems []*types.ComplianceItem

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListComplianceItemsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceItems{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceItems(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListComplianceItems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListComplianceItems",
	}
}
