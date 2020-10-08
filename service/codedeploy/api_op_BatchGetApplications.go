// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more applications. The maximum number of
// applications that can be returned is 100.
func (c *Client) BatchGetApplications(ctx context.Context, params *BatchGetApplicationsInput, optFns ...func(*Options)) (*BatchGetApplicationsOutput, error) {
	if params == nil {
		params = &BatchGetApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetApplications", params, optFns, addOperationBatchGetApplicationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetApplications operation.
type BatchGetApplicationsInput struct {

	// A list of application names separated by spaces. The maximum number of
	// application names you can specify is 100.
	//
	// This member is required.
	ApplicationNames []*string
}

// Represents the output of a BatchGetApplications operation.
type BatchGetApplicationsOutput struct {

	// Information about the applications.
	ApplicationsInfo []*types.ApplicationInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetApplicationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetApplications{}, middleware.After)
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
	addOpBatchGetApplicationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetApplications(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchGetApplications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "BatchGetApplications",
	}
}
