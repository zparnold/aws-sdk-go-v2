// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an input.
func (c *Client) CreateInput(ctx context.Context, params *CreateInputInput, optFns ...func(*Options)) (*CreateInputOutput, error) {
	stack := middleware.NewStack("CreateInput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateInputMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateInputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInput(options.Region), middleware.Before)

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
			OperationName: "CreateInput",
			Err:           err,
		}
	}
	out := result.(*CreateInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInputInput struct {
	// A brief description of the input.
	InputDescription *string
	// The name you want to give to the input.
	InputName *string
	// Metadata that can be used to manage the input.
	Tags []*types.Tag
	// The definition of the input.
	InputDefinition *types.InputDefinition
}

type CreateInputOutput struct {
	// Information about the configuration of the input.
	InputConfiguration *types.InputConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateInputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateInput{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateInput{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateInput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "CreateInput",
	}
}
