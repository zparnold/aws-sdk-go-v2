// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all Systems Manager (SSM) documents in the current AWS account and
// Region. You can limit the results of this request by using a filter.
func (c *Client) ListDocuments(ctx context.Context, params *ListDocumentsInput, optFns ...func(*Options)) (*ListDocumentsOutput, error) {
	stack := middleware.NewStack("ListDocuments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDocumentsMiddlewares(stack)
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
	addOpListDocumentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDocuments(options.Region), middleware.Before)

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
			OperationName: "ListDocuments",
			Err:           err,
		}
	}
	out := result.(*ListDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentsInput struct {
	// This data type is deprecated. Instead, use Filters.
	DocumentFilterList []*types.DocumentFilter
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// One or more DocumentKeyValuesFilter objects. Use a filter to return a more
	// specific list of results. For keys, you can specify one or more key-value pair
	// tags that have been applied to a document. Other valid keys include Owner, Name,
	// PlatformTypes, DocumentType, and TargetType. For example, to return documents
	// you own use Key=Owner,Values=Self. To specify a custom key-value pair, use the
	// format Key=tag:tagName,Values=valueName.
	Filters []*types.DocumentKeyValuesFilter
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
}

type ListDocumentsOutput struct {
	// The names of the Systems Manager documents.
	DocumentIdentifiers []*types.DocumentIdentifier
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDocumentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDocuments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocuments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDocuments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListDocuments",
	}
}
