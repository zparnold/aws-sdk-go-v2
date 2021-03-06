// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// ResolveCustomer is called by a SaaS application during the registration process.
// When a buyer visits your website during the registration process, the buyer
// submits a registration token through their browser. The registration token is
// resolved through this API to obtain a CustomerIdentifier and product code.
func (c *Client) ResolveCustomer(ctx context.Context, params *ResolveCustomerInput, optFns ...func(*Options)) (*ResolveCustomerOutput, error) {
	stack := middleware.NewStack("ResolveCustomer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResolveCustomerMiddlewares(stack)
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
	addOpResolveCustomerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResolveCustomer(options.Region), middleware.Before)
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
			OperationName: "ResolveCustomer",
			Err:           err,
		}
	}
	out := result.(*ResolveCustomerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains input to the ResolveCustomer operation.
type ResolveCustomerInput struct {

	// When a buyer visits your website during the registration process, the buyer
	// submits a registration token through the browser. The registration token is
	// resolved to obtain a CustomerIdentifier and product code.
	//
	// This member is required.
	RegistrationToken *string
}

// The result of the ResolveCustomer operation. Contains the CustomerIdentifier and
// product code.
type ResolveCustomerOutput struct {

	// The CustomerIdentifier is used to identify an individual customer in your
	// application. Calls to BatchMeterUsage require CustomerIdentifiers for each
	// UsageRecord.
	CustomerIdentifier *string

	// The product code is returned to confirm that the buyer is registering for your
	// product. Subsequent BatchMeterUsage calls should be made using this product
	// code.
	ProductCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResolveCustomerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResolveCustomer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResolveCustomer{}, middleware.After)
}

func newServiceMetadataMiddleware_opResolveCustomer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "ResolveCustomer",
	}
}
