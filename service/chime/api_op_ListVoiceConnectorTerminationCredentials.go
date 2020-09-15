// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the SIP credentials for the specified Amazon Chime Voice Connector.
func (c *Client) ListVoiceConnectorTerminationCredentials(ctx context.Context, params *ListVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*ListVoiceConnectorTerminationCredentialsOutput, error) {
	stack := middleware.NewStack("ListVoiceConnectorTerminationCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListVoiceConnectorTerminationCredentialsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListVoiceConnectorTerminationCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVoiceConnectorTerminationCredentials(options.Region), middleware.Before)

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
			OperationName: "ListVoiceConnectorTerminationCredentials",
			Err:           err,
		}
	}
	out := result.(*ListVoiceConnectorTerminationCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVoiceConnectorTerminationCredentialsInput struct {
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type ListVoiceConnectorTerminationCredentialsOutput struct {
	// A list of user names.
	Usernames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListVoiceConnectorTerminationCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListVoiceConnectorTerminationCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListVoiceConnectorTerminationCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opListVoiceConnectorTerminationCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListVoiceConnectorTerminationCredentials",
	}
}
