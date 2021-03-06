// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Puts the specified workflow run properties for the given workflow run. If a
// property already exists for the specified run, then it overrides the value
// otherwise adds the property to existing properties.
func (c *Client) PutWorkflowRunProperties(ctx context.Context, params *PutWorkflowRunPropertiesInput, optFns ...func(*Options)) (*PutWorkflowRunPropertiesOutput, error) {
	stack := middleware.NewStack("PutWorkflowRunProperties", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutWorkflowRunPropertiesMiddlewares(stack)
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
	addOpPutWorkflowRunPropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutWorkflowRunProperties(options.Region), middleware.Before)
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
			OperationName: "PutWorkflowRunProperties",
			Err:           err,
		}
	}
	out := result.(*PutWorkflowRunPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutWorkflowRunPropertiesInput struct {

	// Name of the workflow which was run.
	//
	// This member is required.
	Name *string

	// The ID of the workflow run for which the run properties should be updated.
	//
	// This member is required.
	RunId *string

	// The properties to put for the specified run.
	//
	// This member is required.
	RunProperties map[string]*string
}

type PutWorkflowRunPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutWorkflowRunPropertiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutWorkflowRunProperties{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutWorkflowRunProperties{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutWorkflowRunProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "PutWorkflowRunProperties",
	}
}
