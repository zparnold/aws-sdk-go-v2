// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a transcription job generated by Amazon Transcribe Medical and any
// related information.
func (c *Client) DeleteMedicalTranscriptionJob(ctx context.Context, params *DeleteMedicalTranscriptionJobInput, optFns ...func(*Options)) (*DeleteMedicalTranscriptionJobOutput, error) {
	if params == nil {
		params = &DeleteMedicalTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMedicalTranscriptionJob", params, optFns, addOperationDeleteMedicalTranscriptionJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMedicalTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMedicalTranscriptionJobInput struct {

	// The name you provide to the DeleteMedicalTranscriptionJob object to delete a
	// transcription job.
	//
	// This member is required.
	MedicalTranscriptionJobName *string
}

type DeleteMedicalTranscriptionJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteMedicalTranscriptionJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMedicalTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMedicalTranscriptionJob{}, middleware.After)
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
	addOpDeleteMedicalTranscriptionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMedicalTranscriptionJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteMedicalTranscriptionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "DeleteMedicalTranscriptionJob",
	}
}
