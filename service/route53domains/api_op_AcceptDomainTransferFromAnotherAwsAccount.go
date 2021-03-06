// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts the transfer of a domain from another AWS account to the current AWS
// account. You initiate a transfer between AWS accounts using
// TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
// <p>Use either <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html">ListOperations</a>
// or <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html">GetOperationDetail</a>
// to determine whether the operation succeeded. <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html">GetOperationDetail</a>
// provides additional information, for example, <code>Domain Transfer from Aws
// Account 111122223333 has been cancelled</code>. </p>
func (c *Client) AcceptDomainTransferFromAnotherAwsAccount(ctx context.Context, params *AcceptDomainTransferFromAnotherAwsAccountInput, optFns ...func(*Options)) (*AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
	stack := middleware.NewStack("AcceptDomainTransferFromAnotherAwsAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAcceptDomainTransferFromAnotherAwsAccountMiddlewares(stack)
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
	addOpAcceptDomainTransferFromAnotherAwsAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptDomainTransferFromAnotherAwsAccount(options.Region), middleware.Before)
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
			OperationName: "AcceptDomainTransferFromAnotherAwsAccount",
			Err:           err,
		}
	}
	out := result.(*AcceptDomainTransferFromAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The AcceptDomainTransferFromAnotherAwsAccount request includes the following
// elements.
type AcceptDomainTransferFromAnotherAwsAccountInput struct {

	// The name of the domain that was specified when another AWS account submitted a
	// TransferDomainToAnotherAwsAccount
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	//
	// This member is required.
	DomainName *string

	// The password that was returned by the TransferDomainToAnotherAwsAccount
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	//
	// This member is required.
	Password *string
}

// The AcceptDomainTransferFromAnotherAwsAccount response includes the following
// element.
type AcceptDomainTransferFromAnotherAwsAccountOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAcceptDomainTransferFromAnotherAwsAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptDomainTransferFromAnotherAwsAccount{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptDomainTransferFromAnotherAwsAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptDomainTransferFromAnotherAwsAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "AcceptDomainTransferFromAnotherAwsAccount",
	}
}
