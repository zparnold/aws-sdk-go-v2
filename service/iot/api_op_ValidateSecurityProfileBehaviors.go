// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Validates a Device Defender security profile behaviors specification.
func (c *Client) ValidateSecurityProfileBehaviors(ctx context.Context, params *ValidateSecurityProfileBehaviorsInput, optFns ...func(*Options)) (*ValidateSecurityProfileBehaviorsOutput, error) {
	if params == nil {
		params = &ValidateSecurityProfileBehaviorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateSecurityProfileBehaviors", params, optFns, addOperationValidateSecurityProfileBehaviorsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateSecurityProfileBehaviorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateSecurityProfileBehaviorsInput struct {

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	//
	// This member is required.
	Behaviors []*types.Behavior
}

type ValidateSecurityProfileBehaviorsOutput struct {

	// True if the behaviors were valid.
	Valid *bool

	// The list of any errors found in the behaviors.
	ValidationErrors []*types.ValidationError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationValidateSecurityProfileBehaviorsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidateSecurityProfileBehaviors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidateSecurityProfileBehaviors{}, middleware.After)
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
	addOpValidateSecurityProfileBehaviorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opValidateSecurityProfileBehaviors(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opValidateSecurityProfileBehaviors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ValidateSecurityProfileBehaviors",
	}
}
