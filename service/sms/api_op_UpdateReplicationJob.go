// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates the specified settings for the specified replication job.
func (c *Client) UpdateReplicationJob(ctx context.Context, params *UpdateReplicationJobInput, optFns ...func(*Options)) (*UpdateReplicationJobOutput, error) {
	stack := middleware.NewStack("UpdateReplicationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateReplicationJobMiddlewares(stack)
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
	addOpUpdateReplicationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationJob(options.Region), middleware.Before)
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
			OperationName: "UpdateReplicationJob",
			Err:           err,
		}
	}
	out := result.(*UpdateReplicationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationJobInput struct {

	// The identifier of the replication job.
	//
	// This member is required.
	ReplicationJobId *string

	// The description of the replication job.
	Description *string

	// When true, the replication job produces encrypted AMIs . See also KmsKeyId
	// below.
	Encrypted *bool

	// The time between consecutive replication runs, in hours.
	Frequency *int32

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of the
	// following:
	//
	//     * KMS key ID
	//
	//     * KMS key alias
	//
	//     * ARN referring to KMS
	// key ID
	//
	//     * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key
	// id is not specified, the customer's default KMS key for EBS is used.
	KmsKeyId *string

	// The license type to be used for the AMI created by a successful replication run.
	LicenseType types.LicenseType

	// The start time of the next replication run.
	NextReplicationRunStartTime *time.Time

	// The maximum number of SMS-created AMIs to retain. The oldest will be deleted
	// once the maximum number is reached and a new AMI is created.
	NumberOfRecentAmisToKeep *int32

	// The name of the IAM role to be used by AWS SMS.
	RoleName *string
}

type UpdateReplicationJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateReplicationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateReplicationJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateReplicationJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateReplicationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "UpdateReplicationJob",
	}
}
