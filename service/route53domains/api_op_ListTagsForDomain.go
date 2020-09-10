// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation returns all of the tags that are associated with the specified
// domain. All tag operations are eventually consistent; subsequent operations
// might not immediately represent all issued operations.
func (c *Client) ListTagsForDomain(ctx context.Context, params *ListTagsForDomainInput, optFns ...func(*Options)) (*ListTagsForDomainOutput, error) {
	stack := middleware.NewStack("ListTagsForDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTagsForDomainMiddlewares(stack)
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
	addOpListTagsForDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForDomain(options.Region), middleware.Before)

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
			OperationName: "ListTagsForDomain",
			Err:           err,
		}
	}
	out := result.(*ListTagsForDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListTagsForDomainRequest includes the following elements.
type ListTagsForDomainInput struct {
	// The domain for which you want to get a list of tags.
	DomainName *string
}

// The ListTagsForDomain response includes the following elements.
type ListTagsForDomainOutput struct {
	// A list of the tags that are associated with the specified domain.
	TagList []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTagsForDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTagsForDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ListTagsForDomain",
	}
}