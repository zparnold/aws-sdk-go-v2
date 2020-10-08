// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified replication job. After you delete a replication job, there
// are no further replication runs. AWS deletes the contents of the Amazon S3
// bucket used to store AWS SMS artifacts. The AMIs created by the replication runs
// are not deleted.
func (c *Client) DeleteReplicationJob(ctx context.Context, params *DeleteReplicationJobInput, optFns ...func(*Options)) (*DeleteReplicationJobOutput, error) {
	if params == nil {
		params = &DeleteReplicationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReplicationJob", params, optFns, addOperationDeleteReplicationJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReplicationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteReplicationJobInput struct {

	// The identifier of the replication job.
	//
	// This member is required.
	ReplicationJobId *string
}

type DeleteReplicationJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteReplicationJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteReplicationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteReplicationJob{}, middleware.After)
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
	addOpDeleteReplicationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReplicationJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteReplicationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "DeleteReplicationJob",
	}
}
