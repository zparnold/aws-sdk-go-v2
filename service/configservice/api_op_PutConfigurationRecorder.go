// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new configuration recorder to record the selected resource
// configurations. You can use this action to change the role roleARN or the
// recordingGroup of an existing recorder. To change the role, call the action on
// the existing configuration recorder and specify a role. Currently, you can
// specify only one configuration recorder per region in your account. If
// ConfigurationRecorder does not have the recordingGroup parameter specified, the
// default is to record all supported resource types.
func (c *Client) PutConfigurationRecorder(ctx context.Context, params *PutConfigurationRecorderInput, optFns ...func(*Options)) (*PutConfigurationRecorderOutput, error) {
	if params == nil {
		params = &PutConfigurationRecorderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationRecorder", params, optFns, addOperationPutConfigurationRecorderMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationRecorderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the PutConfigurationRecorder () action.
type PutConfigurationRecorderInput struct {

	// The configuration recorder object that records each configuration change made to
	// the resources.
	//
	// This member is required.
	ConfigurationRecorder *types.ConfigurationRecorder
}

type PutConfigurationRecorderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutConfigurationRecorderMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutConfigurationRecorder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConfigurationRecorder{}, middleware.After)
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
	addOpPutConfigurationRecorderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationRecorder(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutConfigurationRecorder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutConfigurationRecorder",
	}
}
