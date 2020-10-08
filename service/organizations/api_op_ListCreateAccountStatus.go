// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the account creation requests that match the specified status that is
// currently being tracked for the organization. Always check the NextToken
// response parameter for a null value when calling a List* operation. These
// operations can occasionally return an empty set of results even when there are
// more results available. The NextToken response parameter value is null only when
// there are no more results to display. This operation can be called only from the
// organization's master account or by a member account that is a delegated
// administrator for an AWS service.
func (c *Client) ListCreateAccountStatus(ctx context.Context, params *ListCreateAccountStatusInput, optFns ...func(*Options)) (*ListCreateAccountStatusOutput, error) {
	if params == nil {
		params = &ListCreateAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCreateAccountStatus", params, optFns, addOperationListCreateAccountStatusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCreateAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCreateAccountStatusInput struct {

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	// A list of one or more states that you want included in the response. If this
	// parameter isn't present, all requests are included in the response.
	States []types.CreateAccountState
}

type ListCreateAccountStatusOutput struct {

	// A list of objects with details about the requests. Certain elements, such as the
	// accountId number, are present in the output only after the account has been
	// successfully created.
	CreateAccountStatuses []*types.CreateAccountStatus

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCreateAccountStatusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCreateAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCreateAccountStatus{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCreateAccountStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCreateAccountStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListCreateAccountStatus",
	}
}
