// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the DataSourceName of a DataSource. You can use the GetDataSource
// operation to view the contents of the updated data element.
func (c *Client) UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) {
	if params == nil {
		params = &UpdateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSource", params, optFns, addOperationUpdateDataSourceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourceInput struct {

	// The ID assigned to the DataSource during creation.
	//
	// This member is required.
	DataSourceId *string

	// A new user-supplied name or description of the DataSource that will replace the
	// current description.
	//
	// This member is required.
	DataSourceName *string
}

// Represents the output of an UpdateDataSource operation. You can see the updated
// content by using the GetBatchPrediction operation.
type UpdateDataSourceOutput struct {

	// The ID assigned to the DataSource during creation. This value should be
	// identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataSourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataSource{}, middleware.After)
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
	addOpUpdateDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "UpdateDataSource",
	}
}
