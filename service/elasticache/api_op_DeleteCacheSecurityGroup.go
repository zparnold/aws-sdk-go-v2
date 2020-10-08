// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cache security group. You cannot delete a cache security group if it
// is associated with any clusters.
func (c *Client) DeleteCacheSecurityGroup(ctx context.Context, params *DeleteCacheSecurityGroupInput, optFns ...func(*Options)) (*DeleteCacheSecurityGroupOutput, error) {
	if params == nil {
		params = &DeleteCacheSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCacheSecurityGroup", params, optFns, addOperationDeleteCacheSecurityGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCacheSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteCacheSecurityGroup operation.
type DeleteCacheSecurityGroupInput struct {

	// The name of the cache security group to delete. You cannot delete the default
	// security group.
	//
	// This member is required.
	CacheSecurityGroupName *string
}

type DeleteCacheSecurityGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCacheSecurityGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCacheSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCacheSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
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
	addOpDeleteCacheSecurityGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCacheSecurityGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCacheSecurityGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteCacheSecurityGroup",
	}
}
