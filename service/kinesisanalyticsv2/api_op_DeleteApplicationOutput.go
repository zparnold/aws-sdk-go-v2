// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the output destination configuration from your SQL-based Amazon Kinesis
// Data Analytics application's configuration. Kinesis Data Analytics will no
// longer write data from the corresponding in-application stream to the external
// output destination.
func (c *Client) DeleteApplicationOutput(ctx context.Context, params *DeleteApplicationOutputInput, optFns ...func(*Options)) (*DeleteApplicationOutputOutput, error) {
	stack := middleware.NewStack("DeleteApplicationOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationOutputMiddlewares(stack)
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
	addOpDeleteApplicationOutputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationOutput(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationOutput",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationOutputInput struct {

	// The application name.
	//
	// This member is required.
	ApplicationName *string

	// The application version. You can use the DescribeApplication () operation to get
	// the current application version. If the version specified is not the current
	// version, the ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the configuration to delete. Each output configuration that is added
	// to the application (either when the application is created or later) using the
	// AddApplicationOutput () operation has a unique ID. You need to provide the ID to
	// uniquely identify the output configuration that you want to delete from the
	// application configuration. You can use the DescribeApplication () operation to
	// get the specific OutputId.
	//
	// This member is required.
	OutputId *string
}

type DeleteApplicationOutputOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string

	// The current application version ID.
	ApplicationVersionId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationOutput{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationOutput",
	}
}
