// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new cache security group. Use a cache security group to control access
// to one or more clusters. Cache security groups are only used when you are
// creating a cluster outside of an Amazon Virtual Private Cloud (Amazon VPC). If
// you are creating a cluster inside of a VPC, use a cache subnet group instead.
// For more information, see CreateCacheSubnetGroup
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSubnetGroup.html).
func (c *Client) CreateCacheSecurityGroup(ctx context.Context, params *CreateCacheSecurityGroupInput, optFns ...func(*Options)) (*CreateCacheSecurityGroupOutput, error) {
	if params == nil {
		params = &CreateCacheSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCacheSecurityGroup", params, optFns, addOperationCreateCacheSecurityGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCacheSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheSecurityGroup operation.
type CreateCacheSecurityGroupInput struct {

	// A name for the cache security group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 alphanumeric characters. Cannot be
	// the word "Default". Example: mysecuritygroup
	//
	// This member is required.
	CacheSecurityGroupName *string

	// A description for the cache security group.
	//
	// This member is required.
	Description *string
}

type CreateCacheSecurityGroupOutput struct {

	// Represents the output of one of the following operations:
	//
	//     *
	// AuthorizeCacheSecurityGroupIngress
	//
	//     * CreateCacheSecurityGroup
	//
	//     *
	// RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *types.CacheSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCacheSecurityGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheSecurityGroup{}, middleware.After)
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
	addOpCreateCacheSecurityGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheSecurityGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCacheSecurityGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateCacheSecurityGroup",
	}
}
