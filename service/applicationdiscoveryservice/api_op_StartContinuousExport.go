// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Start the continuous flow of agent's discovered data into Amazon Athena.
func (c *Client) StartContinuousExport(ctx context.Context, params *StartContinuousExportInput, optFns ...func(*Options)) (*StartContinuousExportOutput, error) {
	stack := middleware.NewStack("StartContinuousExport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartContinuousExportMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartContinuousExport(options.Region), middleware.Before)
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
			OperationName: "StartContinuousExport",
			Err:           err,
		}
	}
	out := result.(*StartContinuousExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartContinuousExportInput struct {
}

type StartContinuousExportOutput struct {

	// The type of data collector used to gather this data (currently only offered for
	// AGENT).
	DataSource types.DataSource

	// The unique ID assigned to this export.
	ExportId *string

	// The name of the s3 bucket where the export data parquet files are stored.
	S3Bucket *string

	// A dictionary which describes how the data is stored.
	//
	//     * databaseName - the
	// name of the Glue database used to store the schema.
	SchemaStorageConfig map[string]*string

	// The timestamp representing when the continuous export was started.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartContinuousExportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartContinuousExport{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartContinuousExport{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartContinuousExport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StartContinuousExport",
	}
}
