// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of a MediaPackage VOD Asset resource.
func (c *Client) DescribeAsset(ctx context.Context, params *DescribeAssetInput, optFns ...func(*Options)) (*DescribeAssetOutput, error) {
	stack := middleware.NewStack("DescribeAsset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAssetMiddlewares(stack)
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
	addOpDescribeAssetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAsset(options.Region), middleware.Before)
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
			OperationName: "DescribeAsset",
			Err:           err,
		}
	}
	out := result.(*DescribeAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetInput struct {

	// The ID of an MediaPackage VOD Asset resource.
	//
	// This member is required.
	Id *string
}

type DescribeAssetOutput struct {

	// The ARN of the Asset.
	Arn *string

	// The time the Asset was initially submitted for Ingest.
	CreatedAt *string

	// The list of egress endpoints available for the Asset.
	EgressEndpoints []*types.EgressEndpoint

	// The unique identifier for the Asset.
	Id *string

	// The ID of the PackagingGroup for the Asset.
	PackagingGroupId *string

	// The resource ID to include in SPEKE key requests.
	ResourceId *string

	// ARN of the source object in S3.
	SourceArn *string

	// The IAM role_arn used to access the source S3 bucket.
	SourceRoleArn *string

	// A collection of tags associated with a resource
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAssetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAsset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAsset{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAsset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "DescribeAsset",
	}
}
