// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detects custom labels in a supplied image by using an Amazon Rekognition Custom
// Labels model. You specify which version of a model version to use by using the
// ProjectVersionArn input parameter. You pass the input image as base64-encoded
// image bytes or as a reference to an image in an Amazon S3 bucket. If you use the
// AWS CLI to call Amazon Rekognition operations, passing image bytes is not
// supported. The image must be either a PNG or JPEG formatted file. For each
// object that the model version detects on an image, the API returns a
// (CustomLabel) object in an array (CustomLabels). Each CustomLabel object
// provides the label name (Name), the level of confidence that the image contains
// the object (Confidence), and object location information, if it exists, for the
// label on the image (Geometry). During training model calculates a threshold
// value that determines if a prediction for a label is true. By default,
// DetectCustomLabels doesn't return labels whose confidence value is below the
// model's calculated threshold value. To filter labels that are returned, specify
// a value for MinConfidence that is higher than the model's calculated threshold.
// You can get the model's calculated threshold from the model's training results
// shown in the Amazon Rekognition Custom Labels console. To get all labels,
// regardless of confidence, specify a MinConfidence value of 0. You can also add
// the MaxResults parameter to limit the number of labels returned.  <p>This is a
// stateless API operation. That is, the operation does not persist any data.</p>
// <p>This operation requires permissions to perform the
// <code>rekognition:DetectCustomLabels</code> action. </p>
func (c *Client) DetectCustomLabels(ctx context.Context, params *DetectCustomLabelsInput, optFns ...func(*Options)) (*DetectCustomLabelsOutput, error) {
	if params == nil {
		params = &DetectCustomLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectCustomLabels", params, optFns, addOperationDetectCustomLabelsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectCustomLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectCustomLabelsInput struct {

	// Provides the input image either as bytes or an S3 object. You pass image bytes
	// to an Amazon Rekognition API operation by using the Bytes property. For example,
	// you would use the Bytes property to pass an image loaded from a local file
	// system. Image bytes passed by using the Bytes property must be base64-encoded.
	// Your code may not need to encode image bytes if you are using an AWS SDK to call
	// Amazon Rekognition API operations.  <p>For more information, see Analyzing an
	// Image Loaded from a Local File System in the Amazon Rekognition Developer
	// Guide.</p> <p> You pass images stored in an S3 bucket to an Amazon Rekognition
	// API operation by using the <code>S3Object</code> property. Images stored in an
	// S3 bucket do not need to be base64-encoded.</p> <p>The region for the S3 bucket
	// containing the S3 object must match the region you use for Amazon Rekognition
	// operations.</p> <p>If you use the AWS CLI to call Amazon Rekognition operations,
	// passing image bytes using the Bytes property is not supported. You must first
	// upload the image to an Amazon S3 bucket and then call the operation using the
	// S3Object property.</p> <p>For Amazon Rekognition to process an S3 object, the
	// user must have permission to access the S3 object. For more information, see
	// Resource Based Policies in the Amazon Rekognition Developer Guide. </p>
	//
	// This member is required.
	Image *types.Image

	// The ARN of the model version that you want to use.
	//
	// This member is required.
	ProjectVersionArn *string

	// Maximum number of results you want the service to return in the response. The
	// service returns the specified number of highest confidence labels ranked from
	// highest confidence to lowest.
	MaxResults *int32

	// Specifies the minimum confidence level for the labels to return. Amazon
	// Rekognition doesn't return any labels with a confidence lower than this
	// specified value. If you specify a value of 0, all labels are return, regardless
	// of the default thresholds that the model version applies.
	MinConfidence *float32
}

type DetectCustomLabelsOutput struct {

	// An array of custom labels detected in the input image.
	CustomLabels []*types.CustomLabel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetectCustomLabelsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectCustomLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectCustomLabels{}, middleware.After)
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
	addOpDetectCustomLabelsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectCustomLabels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetectCustomLabels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectCustomLabels",
	}
}
