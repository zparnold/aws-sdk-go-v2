// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all Amazon Macie Classic member accounts for the current Amazon Macie
// Classic master account.
func (c *Client) ListMemberAccounts(ctx context.Context, params *ListMemberAccountsInput, optFns ...func(*Options)) (*ListMemberAccountsOutput, error) {
	stack := middleware.NewStack("ListMemberAccounts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListMemberAccountsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMemberAccounts(options.Region), middleware.Before)
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
			OperationName: "ListMemberAccounts",
			Err:           err,
		}
	}
	out := result.(*ListMemberAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMemberAccountsInput struct {

	// Use this parameter to indicate the maximum number of items that you want in the
	// response. The default value is 250.
	MaxResults *int32

	// Use this parameter when paginating results. Set the value of this parameter to
	// null on your first call to the ListMemberAccounts action. Subsequent calls to
	// the action fill nextToken in the request with the value of nextToken from the
	// previous response to continue listing data.
	NextToken *string
}

type ListMemberAccountsOutput struct {

	// A list of the Amazon Macie Classic member accounts returned by the action. The
	// current master account is also included in this list.
	MemberAccounts []*types.MemberAccount

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListMemberAccountsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListMemberAccounts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMemberAccounts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMemberAccounts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie",
		OperationName: "ListMemberAccounts",
	}
}
