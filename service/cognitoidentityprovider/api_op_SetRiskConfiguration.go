// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures actions on detected risks. To delete the risk configuration for
// UserPoolId or ClientId, pass null values for all four configuration types. To
// enable Amazon Cognito advanced security features, update the user pool to
// include the UserPoolAddOns keyAdvancedSecurityMode. See .
func (c *Client) SetRiskConfiguration(ctx context.Context, params *SetRiskConfigurationInput, optFns ...func(*Options)) (*SetRiskConfigurationOutput, error) {
	stack := middleware.NewStack("SetRiskConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetRiskConfigurationMiddlewares(stack)
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
	addOpSetRiskConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetRiskConfiguration(options.Region), middleware.Before)
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
			OperationName: "SetRiskConfiguration",
			Err:           err,
		}
	}
	out := result.(*SetRiskConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetRiskConfigurationInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The account takeover risk configuration.
	AccountTakeoverRiskConfiguration *types.AccountTakeoverRiskConfigurationType

	// The app client ID. If ClientId is null, then the risk configuration is mapped to
	// userPoolId. When the client ID is null, the same risk configuration is applied
	// to all the clients in the userPool. Otherwise, ClientId is mapped to the client.
	// When the client ID is not null, the user pool configuration is overridden and
	// the risk configuration for the client is used instead.
	ClientId *string

	// The compromised credentials risk configuration.
	CompromisedCredentialsRiskConfiguration *types.CompromisedCredentialsRiskConfigurationType

	// The configuration to override the risk decision.
	RiskExceptionConfiguration *types.RiskExceptionConfigurationType
}

type SetRiskConfigurationOutput struct {

	// The risk configuration.
	//
	// This member is required.
	RiskConfiguration *types.RiskConfigurationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetRiskConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetRiskConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetRiskConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetRiskConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetRiskConfiguration",
	}
}
