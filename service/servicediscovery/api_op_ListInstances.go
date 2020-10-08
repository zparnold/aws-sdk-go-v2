// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists summary information about the instances that you registered by using a
// specified service.
func (c *Client) ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) {
	if params == nil {
		params = &ListInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstances", params, optFns, addOperationListInstancesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstancesInput struct {

	// The ID of the service that you want to list instances for.
	//
	// This member is required.
	ServiceId *string

	// The maximum number of instances that you want AWS Cloud Map to return in the
	// response to a ListInstances request. If you don't specify a value for
	// MaxResults, AWS Cloud Map returns up to 100 instances.
	MaxResults *int32

	// For the first ListInstances request, omit this value. If more than MaxResults
	// instances match the specified criteria, you can submit another ListInstances
	// request to get the next group of results. Specify the value of NextToken from
	// the previous response in the next request.
	NextToken *string
}

type ListInstancesOutput struct {

	// Summary information about the instances that are associated with the specified
	// service.
	Instances []*types.InstanceSummary

	// If more than MaxResults instances match the specified criteria, you can submit
	// another ListInstances request to get the next group of results. Specify the
	// value of NextToken from the previous response in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstancesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstances{}, middleware.After)
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
	addOpListInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "ListInstances",
	}
}
