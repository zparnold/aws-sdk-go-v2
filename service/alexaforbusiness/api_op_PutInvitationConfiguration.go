// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures the email template for the user enrollment invitation with the
// specified attributes.
func (c *Client) PutInvitationConfiguration(ctx context.Context, params *PutInvitationConfigurationInput, optFns ...func(*Options)) (*PutInvitationConfigurationOutput, error) {
	stack := middleware.NewStack("PutInvitationConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutInvitationConfigurationMiddlewares(stack)
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
	addOpPutInvitationConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutInvitationConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutInvitationConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutInvitationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInvitationConfigurationInput struct {

	// The name of the organization sending the enrollment invite to a user.
	//
	// This member is required.
	OrganizationName *string

	// The email ID of the organization or individual contact that the enrolled user
	// can use.
	ContactEmail *string

	// The list of private skill IDs that you want to recommend to the user to enable
	// in the invitation.
	PrivateSkillIds []*string
}

type PutInvitationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutInvitationConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutInvitationConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInvitationConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutInvitationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "PutInvitationConfiguration",
	}
}
