// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates phone numbers with the specified Amazon Chime Voice Connector.
func (c *Client) AssociatePhoneNumbersWithVoiceConnector(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorOutput, error) {
	stack := middleware.NewStack("AssociatePhoneNumbersWithVoiceConnector", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociatePhoneNumbersWithVoiceConnectorMiddlewares(stack)
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
	addOpAssociatePhoneNumbersWithVoiceConnectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePhoneNumbersWithVoiceConnector(options.Region), middleware.Before)
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
			OperationName: "AssociatePhoneNumbersWithVoiceConnector",
			Err:           err,
		}
	}
	out := result.(*AssociatePhoneNumbersWithVoiceConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePhoneNumbersWithVoiceConnectorInput struct {

	// List of phone numbers, in E.164 format.
	//
	// This member is required.
	E164PhoneNumbers []*string

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// If true, associates the provided phone numbers with the provided Amazon Chime
	// Voice Connector and removes any previously existing associations. If false, does
	// not associate any phone numbers that have previously existing associations.
	ForceAssociate *bool
}

type AssociatePhoneNumbersWithVoiceConnectorOutput struct {

	// If the action fails for one or more of the phone numbers in the request, a list
	// of the phone numbers is returned, along with error codes and error messages.
	PhoneNumberErrors []*types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociatePhoneNumbersWithVoiceConnectorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociatePhoneNumbersWithVoiceConnector{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociatePhoneNumbersWithVoiceConnector{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociatePhoneNumbersWithVoiceConnector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "AssociatePhoneNumbersWithVoiceConnector",
	}
}
