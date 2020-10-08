// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete an existing configuration set. In Amazon Pinpoint, configuration sets are
// groups of rules that you can apply to the emails you send. You apply a
// configuration set to an email by including a reference to the configuration set
// in the headers of the email. When you apply a configuration set to an email, all
// of the rules in that configuration set are applied to the email.
func (c *Client) DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) {
	if params == nil {
		params = &DeleteConfigurationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationSet", params, optFns, addOperationDeleteConfigurationSetMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a configuration set.
type DeleteConfigurationSetInput struct {

	// The name of the configuration set that you want to delete.
	//
	// This member is required.
	ConfigurationSetName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type DeleteConfigurationSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConfigurationSetMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfigurationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfigurationSet{}, middleware.After)
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
	addOpDeleteConfigurationSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteConfigurationSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteConfigurationSet",
	}
}
