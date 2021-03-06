// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a subnet group. You cannot delete a subnet group if it is associated
// with any DAX clusters.
func (c *Client) DeleteSubnetGroup(ctx context.Context, params *DeleteSubnetGroupInput, optFns ...func(*Options)) (*DeleteSubnetGroupOutput, error) {
	stack := middleware.NewStack("DeleteSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteSubnetGroupMiddlewares(stack)
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
	addOpDeleteSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "DeleteSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSubnetGroupInput struct {

	// The name of the subnet group to delete.
	//
	// This member is required.
	SubnetGroupName *string
}

type DeleteSubnetGroupOutput struct {

	// A user-specified message for this action (i.e., a reason for deleting the subnet
	// group).
	DeletionMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "DeleteSubnetGroup",
	}
}
