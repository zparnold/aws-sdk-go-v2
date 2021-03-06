// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels an update on the specified stack. If the call completes successfully,
// the stack rolls back the update and reverts to the previous stack configuration.
// You can cancel only stacks that are in the UPDATE_IN_PROGRESS state.
func (c *Client) CancelUpdateStack(ctx context.Context, params *CancelUpdateStackInput, optFns ...func(*Options)) (*CancelUpdateStackOutput, error) {
	stack := middleware.NewStack("CancelUpdateStack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCancelUpdateStackMiddlewares(stack)
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
	addOpCancelUpdateStackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelUpdateStack(options.Region), middleware.Before)
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
			OperationName: "CancelUpdateStack",
			Err:           err,
		}
	}
	out := result.(*CancelUpdateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CancelUpdateStack () action.
type CancelUpdateStackInput struct {

	// The name or the unique stack ID that is associated with the stack.
	//
	// This member is required.
	StackName *string

	// A unique identifier for this CancelUpdateStack request. Specify this token if
	// you plan to retry requests so that AWS CloudFormation knows that you're not
	// attempting to cancel an update on a stack with the same name. You might retry
	// CancelUpdateStack requests to ensure that AWS CloudFormation successfully
	// received them.
	ClientRequestToken *string
}

type CancelUpdateStackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCancelUpdateStackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCancelUpdateStack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCancelUpdateStack{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelUpdateStack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "CancelUpdateStack",
	}
}
