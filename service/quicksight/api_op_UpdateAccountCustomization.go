// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates customizations associated with the QuickSight subscription on your AWS
// account.
func (c *Client) UpdateAccountCustomization(ctx context.Context, params *UpdateAccountCustomizationInput, optFns ...func(*Options)) (*UpdateAccountCustomizationOutput, error) {
	if params == nil {
		params = &UpdateAccountCustomizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountCustomization", params, optFns, addOperationUpdateAccountCustomizationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountCustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountCustomizationInput struct {

	// The customizations you want to update in QuickSight.
	//
	// This member is required.
	AccountCustomization *types.AccountCustomization

	// The ID for the AWS account that you want to update QuickSight customizations
	// for.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace associated with the customization that you're updating.
	Namespace *string
}

type UpdateAccountCustomizationOutput struct {

	// The customizations associated with your QuickSight subscription.
	AccountCustomization *types.AccountCustomization

	// The ID for the AWS account that you want to update QuickSight customizations
	// for.
	AwsAccountId *string

	// The namespace associated with the customization that you're updating.
	Namespace *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAccountCustomizationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccountCustomization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccountCustomization{}, middleware.After)
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
	addOpUpdateAccountCustomizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountCustomization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccountCustomization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateAccountCustomization",
	}
}
