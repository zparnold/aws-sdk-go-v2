// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the resolvers for a given API and type.
func (c *Client) ListResolvers(ctx context.Context, params *ListResolversInput, optFns ...func(*Options)) (*ListResolversOutput, error) {
	stack := middleware.NewStack("ListResolvers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListResolversMiddlewares(stack)
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
	addOpListResolversValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResolvers(options.Region), middleware.Before)

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
			OperationName: "ListResolvers",
			Err:           err,
		}
	}
	out := result.(*ListResolversOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolversInput struct {
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
	// The type name.
	TypeName *string
	// The API ID.
	ApiId *string
	// The maximum number of results you want the request to return.
	MaxResults *int32
}

type ListResolversOutput struct {
	// The Resolver objects.
	Resolvers []*types.Resolver
	// An identifier to be passed in the next request to this operation to return the
	// next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListResolversMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListResolvers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListResolvers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResolvers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "ListResolvers",
	}
}
