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

// Returns a collection of MediaPackage VOD PackagingConfiguration resources.
func (c *Client) ListPackagingConfigurations(ctx context.Context, params *ListPackagingConfigurationsInput, optFns ...func(*Options)) (*ListPackagingConfigurationsOutput, error) {
	stack := middleware.NewStack("ListPackagingConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPackagingConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackagingConfigurations(options.Region), middleware.Before)
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
			OperationName: "ListPackagingConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListPackagingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagingConfigurationsInput struct {

	// Upper bound on number of records to return.
	MaxResults *int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	// Returns MediaPackage VOD PackagingConfigurations associated with the specified
	// PackagingGroup.
	PackagingGroupId *string
}

type ListPackagingConfigurationsOutput struct {

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// A list of MediaPackage VOD PackagingConfiguration resources.
	PackagingConfigurations []*types.PackagingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPackagingConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPackagingConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackagingConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPackagingConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "ListPackagingConfigurations",
	}
}
