// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the permission policy for an application. For the list of actions supported
// for this operation, see Application Permissions
// (https://docs.aws.amazon.com/serverlessrepo/latest/devguide/access-control-resource-based.html#application-permissions)
// .
func (c *Client) PutApplicationPolicy(ctx context.Context, params *PutApplicationPolicyInput, optFns ...func(*Options)) (*PutApplicationPolicyOutput, error) {
	if params == nil {
		params = &PutApplicationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutApplicationPolicy", params, optFns, addOperationPutApplicationPolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutApplicationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutApplicationPolicyInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// An array of policy statements applied to the application.
	//
	// This member is required.
	Statements []*types.ApplicationPolicyStatement
}

type PutApplicationPolicyOutput struct {

	// An array of policy statements applied to the application.
	Statements []*types.ApplicationPolicyStatement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutApplicationPolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutApplicationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutApplicationPolicy{}, middleware.After)
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
	addOpPutApplicationPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutApplicationPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutApplicationPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "PutApplicationPolicy",
	}
}
