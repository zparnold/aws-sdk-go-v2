// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// This operation uploads a part of an archive. You can upload archive parts in any
// order. You can also upload them in parallel. You can upload up to 10,000 parts
// for a multipart upload.  <p>Amazon Glacier rejects your upload part request if
// any of the following conditions is true:</p> <ul> <li> <p> <b>SHA256 tree hash
// does not match</b>To ensure that part data is not corrupted in transmission, you
// compute a SHA256 tree hash of the part and include it in your request. Upon
// receiving the part data, Amazon S3 Glacier also computes a SHA256 tree hash. If
// these hash values don't match, the operation fails. For information about
// computing a SHA256 tree hash, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html">Computing
// Checksums</a>.</p> </li> <li> <p> <b>Part size does not match</b>The size of
// each part except the last must match the size specified in the corresponding
// <a>InitiateMultipartUpload</a> request. The size of the last part must be the
// same size as, or smaller than, the specified size.</p> <note> <p>If you upload a
// part whose size is smaller than the part size you specified in your initiate
// multipart upload request and that part is not the last part, then the upload
// part request will succeed. However, the subsequent Complete Multipart Upload
// request will fail.</p> </note> </li> <li> <p> <b>Range does not align</b>The
// byte range value in the request does not align with the part size specified in
// the corresponding initiate request. For example, if you specify a part size of
// 4194304 bytes (4 MB), then 0 to 4194303 bytes (4 MB - 1) and 4194304 (4 MB) to
// 8388607 (8 MB - 1) are valid part ranges. However, if you set a range value of 2
// MB to 6 MB, the range does not align with the part size and the upload will
// fail. </p> </li> </ul> <p>This operation is idempotent. If you upload the same
// part multiple times, the data included in the most recent request overwrites the
// previously uploaded data.</p> <p>An AWS account has full permission to perform
// all operations (actions). However, AWS Identity and Access Management (IAM)
// users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p> For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html">Uploading
// Large Archives in Parts (Multipart Upload)</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-upload-part.html">Upload
// Part </a> in the <i>Amazon Glacier Developer Guide</i>.</p>
func (c *Client) UploadMultipartPart(ctx context.Context, params *UploadMultipartPartInput, optFns ...func(*Options)) (*UploadMultipartPartOutput, error) {
	if params == nil {
		params = &UploadMultipartPartInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadMultipartPart", params, optFns, addOperationUploadMultipartPartMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadMultipartPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to upload a part of an archive in a multipart upload operation.
type UploadMultipartPartInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload.
	//
	// This member is required.
	UploadId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The data to upload.
	Body io.Reader

	// The SHA256 tree hash of the data being uploaded.
	Checksum *string

	// Identifies the range of bytes in the assembled archive that will be uploaded in
	// this part. Amazon S3 Glacier uses this information to assemble the archive in
	// the proper sequence. The format of this header follows RFC 2616. An example
	// header is Content-Range:bytes 0-4194303/*.
	Range *string
}

// Contains the Amazon S3 Glacier response to your request.
type UploadMultipartPartOutput struct {

	// The SHA256 tree hash that Amazon S3 Glacier computed for the uploaded part.
	Checksum *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadMultipartPartMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUploadMultipartPart{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUploadMultipartPart{}, middleware.After)
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
	addOpUploadMultipartPartValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUploadMultipartPart(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)
	return nil
}

func newServiceMetadataMiddleware_opUploadMultipartPart(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "UploadMultipartPart",
	}
}
