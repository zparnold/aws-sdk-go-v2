// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates an existing vocabulary with new values in a different text file. The
// UpdateMedicalVocabulary operation overwrites all of the existing information
// with the values that you provide in the request.
func (c *Client) UpdateMedicalVocabulary(ctx context.Context, params *UpdateMedicalVocabularyInput, optFns ...func(*Options)) (*UpdateMedicalVocabularyOutput, error) {
	if params == nil {
		params = &UpdateMedicalVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMedicalVocabulary", params, optFns, addOperationUpdateMedicalVocabularyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMedicalVocabularyInput struct {

	// The language code of the entries in the updated vocabulary. US English (en-US)
	// is the only valid language code in Amazon Transcribe Medical.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The name of the vocabulary to update. The name is case-sensitive. If you try to
	// update a vocabulary with the same name as a previous vocabulary you will receive
	// a ConflictException error.
	//
	// This member is required.
	VocabularyName *string

	// The Amazon S3 location of the text file containing the definition of the custom
	// vocabulary. The URI must be in the same AWS region as the API endpoint you are
	// calling. You can see the fields you need to enter for you Amazon S3 location in
	// the example URI here:  https://s3..amazonaws.com///  For example:
	// https://s3.us-east-1.amazonaws.com/AWSDOC-EXAMPLE-BUCKET/vocab.txt For more
	// information about S3 object names, see Object Keys
	// (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide. For more information about custom vocabularies
	// in Amazon Transcribe Medical, see Medical Custom Vocabularies
	// (http://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary).
	VocabularyFileUri *string
}

type UpdateMedicalVocabularyOutput struct {

	// The language code for the text file used to update the custom vocabulary. US
	// English (en-US) is the only language supported in Amazon Transcribe Medical.
	LanguageCode types.LanguageCode

	// The date and time the vocabulary was updated.
	LastModifiedTime *time.Time

	// The name of the updated vocabulary.
	VocabularyName *string

	// The processing state of the update to the vocabulary. When the VocabularyState
	// field is READY the vocabulary is ready to be used in a
	// StartMedicalTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMedicalVocabularyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMedicalVocabulary{}, middleware.After)
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
	addOpUpdateMedicalVocabularyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMedicalVocabulary(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateMedicalVocabulary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "UpdateMedicalVocabulary",
	}
}
