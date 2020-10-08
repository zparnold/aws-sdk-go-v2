// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the details for a particular increase request.
func (c *Client) GetRequestedServiceQuotaChange(ctx context.Context, params *GetRequestedServiceQuotaChangeInput, optFns ...func(*Options)) (*GetRequestedServiceQuotaChangeOutput, error) {
	if params == nil {
		params = &GetRequestedServiceQuotaChangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRequestedServiceQuotaChange", params, optFns, addOperationGetRequestedServiceQuotaChangeMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRequestedServiceQuotaChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRequestedServiceQuotaChangeInput struct {

	// Identifies the quota increase request.
	//
	// This member is required.
	RequestId *string
}

type GetRequestedServiceQuotaChangeOutput struct {

	// Returns the RequestedServiceQuotaChange object for the specific increase
	// request.
	RequestedQuota *types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRequestedServiceQuotaChangeMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRequestedServiceQuotaChange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRequestedServiceQuotaChange{}, middleware.After)
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
	addOpGetRequestedServiceQuotaChangeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRequestedServiceQuotaChange(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRequestedServiceQuotaChange(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetRequestedServiceQuotaChange",
	}
}
