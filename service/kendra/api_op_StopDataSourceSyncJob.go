// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a running synchronization job. You can't stop a scheduled synchronization
// job.
func (c *Client) StopDataSourceSyncJob(ctx context.Context, params *StopDataSourceSyncJobInput, optFns ...func(*Options)) (*StopDataSourceSyncJobOutput, error) {
	if params == nil {
		params = &StopDataSourceSyncJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDataSourceSyncJob", params, optFns, addOperationStopDataSourceSyncJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDataSourceSyncJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDataSourceSyncJobInput struct {

	// The identifier of the data source for which to stop the synchronization jobs.
	//
	// This member is required.
	Id *string

	// The identifier of the index that contains the data source.
	//
	// This member is required.
	IndexId *string
}

type StopDataSourceSyncJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDataSourceSyncJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDataSourceSyncJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDataSourceSyncJob{}, middleware.After)
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
	addOpStopDataSourceSyncJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDataSourceSyncJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopDataSourceSyncJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "StopDataSourceSyncJob",
	}
}
