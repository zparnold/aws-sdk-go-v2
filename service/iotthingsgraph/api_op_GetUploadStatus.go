// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets the status of the specified upload.
func (c *Client) GetUploadStatus(ctx context.Context, params *GetUploadStatusInput, optFns ...func(*Options)) (*GetUploadStatusOutput, error) {
	if params == nil {
		params = &GetUploadStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUploadStatus", params, optFns, addOperationGetUploadStatusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUploadStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUploadStatusInput struct {

	// The ID of the upload. This value is returned by the UploadEntityDefinitions
	// action.
	//
	// This member is required.
	UploadId *string
}

type GetUploadStatusOutput struct {

	// The date at which the upload was created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The ID of the upload.
	//
	// This member is required.
	UploadId *string

	// The status of the upload. The initial status is IN_PROGRESS. The response show
	// all validation failures if the upload fails.
	//
	// This member is required.
	UploadStatus types.UploadStatus

	// The reason for an upload failure.
	FailureReason []*string

	// The ARN of the upload.
	NamespaceArn *string

	// The name of the upload's namespace.
	NamespaceName *string

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUploadStatusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUploadStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUploadStatus{}, middleware.After)
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
	addOpGetUploadStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUploadStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetUploadStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetUploadStatus",
	}
}
