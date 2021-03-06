// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the resources or principals for the resource shares that you own.
func (c *Client) GetResourceShareAssociations(ctx context.Context, params *GetResourceShareAssociationsInput, optFns ...func(*Options)) (*GetResourceShareAssociationsOutput, error) {
	stack := middleware.NewStack("GetResourceShareAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetResourceShareAssociationsMiddlewares(stack)
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
	addOpGetResourceShareAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceShareAssociations(options.Region), middleware.Before)
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
			OperationName: "GetResourceShareAssociations",
			Err:           err,
		}
	}
	out := result.(*GetResourceShareAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceShareAssociationsInput struct {

	// The association type. Specify PRINCIPAL to list the principals that are
	// associated with the specified resource share. Specify RESOURCE to list the
	// resources that are associated with the specified resource share.
	//
	// This member is required.
	AssociationType types.ResourceShareAssociationType

	// The association status.
	AssociationStatus types.ResourceShareAssociationStatus

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The principal. You cannot specify this parameter if the association type is
	// RESOURCE.
	Principal *string

	// The Amazon Resource Name (ARN) of the resource. You cannot specify this
	// parameter if the association type is PRINCIPAL.
	ResourceArn *string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []*string
}

type GetResourceShareAssociationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the associations.
	ResourceShareAssociations []*types.ResourceShareAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetResourceShareAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceShareAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceShareAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetResourceShareAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourceShareAssociations",
	}
}
