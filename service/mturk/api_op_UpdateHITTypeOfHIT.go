// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The UpdateHITTypeOfHIT operation allows you to change the HITType properties of
// a HIT. This operation disassociates the HIT from its old HITType properties and
// associates it with the new HITType properties. The HIT takes on the properties
// of the new HITType in place of the old ones.
func (c *Client) UpdateHITTypeOfHIT(ctx context.Context, params *UpdateHITTypeOfHITInput, optFns ...func(*Options)) (*UpdateHITTypeOfHITOutput, error) {
	stack := middleware.NewStack("UpdateHITTypeOfHIT", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateHITTypeOfHITMiddlewares(stack)
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
	addOpUpdateHITTypeOfHITValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHITTypeOfHIT(options.Region), middleware.Before)
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
			OperationName: "UpdateHITTypeOfHIT",
			Err:           err,
		}
	}
	out := result.(*UpdateHITTypeOfHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHITTypeOfHITInput struct {

	// The HIT to update.
	//
	// This member is required.
	HITId *string

	// The ID of the new HIT type.
	//
	// This member is required.
	HITTypeId *string
}

type UpdateHITTypeOfHITOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateHITTypeOfHITMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateHITTypeOfHIT{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateHITTypeOfHIT{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateHITTypeOfHIT(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "UpdateHITTypeOfHIT",
	}
}
