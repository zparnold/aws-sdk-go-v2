// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a URL that you can use to connect to the Jupyter server from a notebook
// instance. In the Amazon SageMaker console, when you choose Open next to a
// notebook instance, Amazon SageMaker opens a new tab showing the Jupyter server
// home page from the notebook instance. The console uses this API to get the URL
// and show the page. The IAM role or user used to call this API defines the
// permissions to access the notebook instance. Once the presigned URL is created,
// no additional permission is required to access this URL. IAM authorization
// policies for this API are also enforced for every HTTP request and WebSocket
// frame that attempts to connect to the notebook instance. You can restrict access
// to this API and to the URL that it returns to a list of IP addresses that you
// specify. Use the NotIpAddress condition operator and the aws:SourceIP condition
// context key to specify the list of IP addresses that you want to have access to
// the notebook instance. For more information, see Limit Access to a Notebook
// Instance by IP Address
// (https://docs.aws.amazon.com/sagemaker/latest/dg/security_iam_id-based-policy-examples.html#nbi-ip-filter).
// The URL that you get from a call to CreatePresignedNotebookInstanceUrl () is
// valid only for 5 minutes. If you try to use the URL after the 5-minute limit
// expires, you are directed to the AWS console sign-in page.
func (c *Client) CreatePresignedNotebookInstanceUrl(ctx context.Context, params *CreatePresignedNotebookInstanceUrlInput, optFns ...func(*Options)) (*CreatePresignedNotebookInstanceUrlOutput, error) {
	stack := middleware.NewStack("CreatePresignedNotebookInstanceUrl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePresignedNotebookInstanceUrlMiddlewares(stack)
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
	addOpCreatePresignedNotebookInstanceUrlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePresignedNotebookInstanceUrl(options.Region), middleware.Before)
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
			OperationName: "CreatePresignedNotebookInstanceUrl",
			Err:           err,
		}
	}
	out := result.(*CreatePresignedNotebookInstanceUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePresignedNotebookInstanceUrlInput struct {

	// The name of the notebook instance.
	//
	// This member is required.
	NotebookInstanceName *string

	// The duration of the session, in seconds. The default is 12 hours.
	SessionExpirationDurationInSeconds *int32
}

type CreatePresignedNotebookInstanceUrlOutput struct {

	// A JSON object that contains the URL string.
	AuthorizedUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePresignedNotebookInstanceUrlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePresignedNotebookInstanceUrl{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePresignedNotebookInstanceUrl{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePresignedNotebookInstanceUrl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreatePresignedNotebookInstanceUrl",
	}
}
