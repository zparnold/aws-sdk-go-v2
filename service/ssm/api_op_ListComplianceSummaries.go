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

// Returns a summary count of compliant and non-compliant resources for a
// compliance type. For example, this call can return State Manager associations,
// patches, or custom compliance types according to the filter criteria that you
// specify.
func (c *Client) ListComplianceSummaries(ctx context.Context, params *ListComplianceSummariesInput, optFns ...func(*Options)) (*ListComplianceSummariesOutput, error) {
	if params == nil {
		params = &ListComplianceSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComplianceSummaries", params, optFns, addOperationListComplianceSummariesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComplianceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceSummariesInput struct {

	// One or more compliance or inventory filters. Use a filter to return a more
	// specific list of results.
	Filters []*types.ComplianceStringFilter

	// The maximum number of items to return for this call. Currently, you can specify
	// null or 50. The call also returns a token that you can specify in a subsequent
	// call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type ListComplianceSummariesOutput struct {

	// A list of compliant and non-compliant summary counts based on compliance types.
	// For example, this call returns State Manager associations, patches, or custom
	// compliance types according to the filter criteria that you specified.
	ComplianceSummaryItems []*types.ComplianceSummaryItem

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListComplianceSummariesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceSummaries{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceSummaries(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListComplianceSummaries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListComplianceSummaries",
	}
}
