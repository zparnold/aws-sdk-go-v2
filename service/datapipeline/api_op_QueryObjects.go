// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Queries the specified pipeline for the names of objects that match the specified
// set of conditions.
func (c *Client) QueryObjects(ctx context.Context, params *QueryObjectsInput, optFns ...func(*Options)) (*QueryObjectsOutput, error) {
	if params == nil {
		params = &QueryObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryObjects", params, optFns, addOperationQueryObjectsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for QueryObjects.
type QueryObjectsInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// Indicates whether the query applies to components or instances. The possible
	// values are: COMPONENT, INSTANCE, and ATTEMPT.
	//
	// This member is required.
	Sphere *string

	// The maximum number of object names that QueryObjects will return in a single
	// call. The default value is 100.
	Limit *int32

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// QueryObjects with the marker value from the previous call to retrieve the next
	// set of results.
	Marker *string

	// The query that defines the objects to be returned. The Query object can contain
	// a maximum of ten selectors. The conditions in the query are limited to top-level
	// String fields in the object. These filters can be applied to components,
	// instances, and attempts.
	Query *types.Query
}

// Contains the output of QueryObjects.
type QueryObjectsOutput struct {

	// Indicates whether there are more results that can be obtained by a subsequent
	// call.
	HasMoreResults *bool

	// The identifiers that match the query selectors.
	Ids []*string

	// The starting point for the next page of results. To view the next page of
	// results, call QueryObjects again with this marker value. If the value is null,
	// there are no more results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationQueryObjectsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQueryObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryObjects{}, middleware.After)
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
	addOpQueryObjectsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryObjects(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opQueryObjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "QueryObjects",
	}
}
