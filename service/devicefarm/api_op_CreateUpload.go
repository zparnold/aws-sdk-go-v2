// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Uploads an app or test scripts.
func (c *Client) CreateUpload(ctx context.Context, params *CreateUploadInput, optFns ...func(*Options)) (*CreateUploadOutput, error) {
	stack := middleware.NewStack("CreateUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateUploadMiddlewares(stack)
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
	addOpCreateUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUpload(options.Region), middleware.Before)
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
			OperationName: "CreateUpload",
			Err:           err,
		}
	}
	out := result.(*CreateUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the create upload operation.
type CreateUploadInput struct {

	// The upload's file name. The name should not contain any forward slashes (/). If
	// you are uploading an iOS app, the file name must end with the .ipa extension. If
	// you are uploading an Android app, the file name must end with the .apk
	// extension. For all others, the file name must end with the .zip file extension.
	//
	// This member is required.
	Name *string

	// The ARN of the project for the upload.
	//
	// This member is required.
	ProjectArn *string

	// The upload's upload type. Must be one of the following values:
	//
	//     *
	// ANDROID_APP
	//
	//     * IOS_APP
	//
	//     * WEB_APP
	//
	//     * EXTERNAL_DATA
	//
	//     *
	// APPIUM_JAVA_JUNIT_TEST_PACKAGE
	//
	//     * APPIUM_JAVA_TESTNG_TEST_PACKAGE
	//
	//     *
	// APPIUM_PYTHON_TEST_PACKAGE
	//
	//     * APPIUM_NODE_TEST_PACKAGE
	//
	//     *
	// APPIUM_RUBY_TEST_PACKAGE
	//
	//     * APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE
	//
	//     *
	// APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE
	//
	//     * APPIUM_WEB_PYTHON_TEST_PACKAGE
	//
	//     *
	// APPIUM_WEB_NODE_TEST_PACKAGE
	//
	//     * APPIUM_WEB_RUBY_TEST_PACKAGE
	//
	//     *
	// CALABASH_TEST_PACKAGE
	//
	//     * INSTRUMENTATION_TEST_PACKAGE
	//
	//     *
	// UIAUTOMATION_TEST_PACKAGE
	//
	//     * UIAUTOMATOR_TEST_PACKAGE
	//
	//     *
	// XCTEST_TEST_PACKAGE
	//
	//     * XCTEST_UI_TEST_PACKAGE
	//
	//     *
	// APPIUM_JAVA_JUNIT_TEST_SPEC
	//
	//     * APPIUM_JAVA_TESTNG_TEST_SPEC
	//
	//     *
	// APPIUM_PYTHON_TEST_SPEC
	//
	//     * APPIUM_NODE_TEST_SPEC
	//
	//     *
	// APPIUM_RUBY_TEST_SPEC
	//
	//     * APPIUM_WEB_JAVA_JUNIT_TEST_SPEC
	//
	//     *
	// APPIUM_WEB_JAVA_TESTNG_TEST_SPEC
	//
	//     * APPIUM_WEB_PYTHON_TEST_SPEC
	//
	//     *
	// APPIUM_WEB_NODE_TEST_SPEC
	//
	//     * APPIUM_WEB_RUBY_TEST_SPEC
	//
	//     *
	// INSTRUMENTATION_TEST_SPEC
	//
	//     * XCTEST_UI_TEST_SPEC
	//
	// If you call CreateUpload
	// with WEB_APP specified, AWS Device Farm throws an ArgumentException error.
	//
	// This member is required.
	Type types.UploadType

	// The upload's content type (for example, application/octet-stream).
	ContentType *string
}

// Represents the result of a create upload request.
type CreateUploadOutput struct {

	// The newly created upload.
	Upload *types.Upload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUpload{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateUpload",
	}
}
