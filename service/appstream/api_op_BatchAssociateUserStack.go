// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the specified users with the specified stacks. Users in a user pool
// cannot be assigned to stacks with fleets that are joined to an Active Directory
// domain.
func (c *Client) BatchAssociateUserStack(ctx context.Context, params *BatchAssociateUserStackInput, optFns ...func(*Options)) (*BatchAssociateUserStackOutput, error) {
	stack := middleware.NewStack("BatchAssociateUserStack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchAssociateUserStackMiddlewares(stack)
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
	addOpBatchAssociateUserStackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateUserStack(options.Region), middleware.Before)
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
			OperationName: "BatchAssociateUserStack",
			Err:           err,
		}
	}
	out := result.(*BatchAssociateUserStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateUserStackInput struct {

	// The list of UserStackAssociation objects.
	//
	// This member is required.
	UserStackAssociations []*types.UserStackAssociation
}

type BatchAssociateUserStackOutput struct {

	// The list of UserStackAssociationError objects.
	Errors []*types.UserStackAssociationError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchAssociateUserStackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchAssociateUserStack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchAssociateUserStack{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchAssociateUserStack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "BatchAssociateUserStack",
	}
}
