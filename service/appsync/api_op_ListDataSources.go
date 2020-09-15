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

// Lists the data sources for a given API.
func (c *Client) ListDataSources(ctx context.Context, params *ListDataSourcesInput, optFns ...func(*Options)) (*ListDataSourcesOutput, error) {
	stack := middleware.NewStack("ListDataSources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDataSourcesMiddlewares(stack)
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
	addOpListDataSourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSources(options.Region), middleware.Before)

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
			OperationName: "ListDataSources",
			Err:           err,
		}
	}
	out := result.(*ListDataSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSourcesInput struct {
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
	// The API ID.
	ApiId *string
	// The maximum number of results you want the request to return.
	MaxResults *int32
}

type ListDataSourcesOutput struct {
	// An identifier to be passed in the next request to this operation to return the
	// next set of items in the list.
	NextToken *string
	// The DataSource objects.
	DataSources []*types.DataSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDataSourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDataSources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDataSources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "ListDataSources",
	}
}
