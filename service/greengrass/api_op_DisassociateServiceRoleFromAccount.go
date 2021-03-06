// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the service role from your account. Without a service role,
// deployments will not work.
func (c *Client) DisassociateServiceRoleFromAccount(ctx context.Context, params *DisassociateServiceRoleFromAccountInput, optFns ...func(*Options)) (*DisassociateServiceRoleFromAccountOutput, error) {
	stack := middleware.NewStack("DisassociateServiceRoleFromAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateServiceRoleFromAccountMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateServiceRoleFromAccount(options.Region), middleware.Before)
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
			OperationName: "DisassociateServiceRoleFromAccount",
			Err:           err,
		}
	}
	out := result.(*DisassociateServiceRoleFromAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateServiceRoleFromAccountInput struct {
}

type DisassociateServiceRoleFromAccountOutput struct {

	// The time when the service role was disassociated from the account.
	DisassociatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateServiceRoleFromAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateServiceRoleFromAccount{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateServiceRoleFromAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateServiceRoleFromAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "DisassociateServiceRoleFromAccount",
	}
}
