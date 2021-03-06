// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the data source permanently. This action breaks all the datasets that
// reference the deleted data source.
func (c *Client) DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) {
	stack := middleware.NewStack("DeleteDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteDataSourceMiddlewares(stack)
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
	addOpDeleteDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataSource(options.Region), middleware.Before)
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
			OperationName: "DeleteDataSource",
			Err:           err,
		}
	}
	out := result.(*DeleteDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataSourceInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	//
	// This member is required.
	DataSourceId *string
}

type DeleteDataSourceOutput struct {

	// The Amazon Resource Name (ARN) of the data source that you deleted.
	Arn *string

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteDataSource",
	}
}
