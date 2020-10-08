// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the status and settings of a specific import job for
// an application.
func (c *Client) GetImportJob(ctx context.Context, params *GetImportJobInput, optFns ...func(*Options)) (*GetImportJobOutput, error) {
	if params == nil {
		params = &GetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImportJob", params, optFns, addOperationGetImportJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImportJobInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the job.
	//
	// This member is required.
	JobId *string
}

type GetImportJobOutput struct {

	// Provides information about the status and settings of a job that imports
	// endpoint definitions from one or more files. The files can be stored in an
	// Amazon Simple Storage Service (Amazon S3) bucket or uploaded directly from a
	// computer by using the Amazon Pinpoint console.
	//
	// This member is required.
	ImportJobResponse *types.ImportJobResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetImportJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImportJob{}, middleware.After)
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
	addOpGetImportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetImportJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetImportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetImportJob",
	}
}
