// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The GetQualificationTypeoperation retrieves information about a Qualification
// type using its ID.
func (c *Client) GetQualificationType(ctx context.Context, params *GetQualificationTypeInput, optFns ...func(*Options)) (*GetQualificationTypeOutput, error) {
	if params == nil {
		params = &GetQualificationTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQualificationType", params, optFns, addOperationGetQualificationTypeMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQualificationTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQualificationTypeInput struct {

	// The ID of the QualificationType.
	//
	// This member is required.
	QualificationTypeId *string
}

type GetQualificationTypeOutput struct {

	// The returned Qualification Type
	QualificationType *types.QualificationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetQualificationTypeMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetQualificationType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQualificationType{}, middleware.After)
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
	addOpGetQualificationTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQualificationType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetQualificationType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "GetQualificationType",
	}
}
