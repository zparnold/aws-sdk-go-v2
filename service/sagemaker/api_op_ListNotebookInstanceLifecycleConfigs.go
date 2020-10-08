// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists notebook instance lifestyle configurations created with the
// CreateNotebookInstanceLifecycleConfig () API.
func (c *Client) ListNotebookInstanceLifecycleConfigs(ctx context.Context, params *ListNotebookInstanceLifecycleConfigsInput, optFns ...func(*Options)) (*ListNotebookInstanceLifecycleConfigsOutput, error) {
	if params == nil {
		params = &ListNotebookInstanceLifecycleConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotebookInstanceLifecycleConfigs", params, optFns, addOperationListNotebookInstanceLifecycleConfigsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotebookInstanceLifecycleConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotebookInstanceLifecycleConfigsInput struct {

	// A filter that returns only lifecycle configurations that were created after the
	// specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only lifecycle configurations that were created before the
	// specified time (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only lifecycle configurations that were modified after the
	// specified time (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only lifecycle configurations that were modified before
	// the specified time (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of lifecycle configurations to return in the response.
	MaxResults *int32

	// A string in the lifecycle configuration name. This filter returns only lifecycle
	// configurations whose name contains the specified string.
	NameContains *string

	// If the result of a ListNotebookInstanceLifecycleConfigs request was truncated,
	// the response includes a NextToken. To get the next set of lifecycle
	// configurations, use the token in the next request.
	NextToken *string

	// Sorts the list of results. The default is CreationTime.
	SortBy types.NotebookInstanceLifecycleConfigSortKey

	// The sort order for results.
	SortOrder types.NotebookInstanceLifecycleConfigSortOrder
}

type ListNotebookInstanceLifecycleConfigsOutput struct {

	// If the response is truncated, Amazon SageMaker returns this token. To get the
	// next set of lifecycle configurations, use it in the next request.
	NextToken *string

	// An array of NotebookInstanceLifecycleConfiguration objects, each listing a
	// lifecycle configuration.
	NotebookInstanceLifecycleConfigs []*types.NotebookInstanceLifecycleConfigSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListNotebookInstanceLifecycleConfigsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNotebookInstanceLifecycleConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNotebookInstanceLifecycleConfigs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNotebookInstanceLifecycleConfigs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListNotebookInstanceLifecycleConfigs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListNotebookInstanceLifecycleConfigs",
	}
}
