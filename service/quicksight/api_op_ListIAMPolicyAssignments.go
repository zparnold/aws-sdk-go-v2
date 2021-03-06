// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists IAM policy assignments in the current Amazon QuickSight account.
func (c *Client) ListIAMPolicyAssignments(ctx context.Context, params *ListIAMPolicyAssignmentsInput, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsOutput, error) {
	stack := middleware.NewStack("ListIAMPolicyAssignments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListIAMPolicyAssignmentsMiddlewares(stack)
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
	addOpListIAMPolicyAssignmentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIAMPolicyAssignments(options.Region), middleware.Before)
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
			OperationName: "ListIAMPolicyAssignments",
			Err:           err,
		}
	}
	out := result.(*ListIAMPolicyAssignmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIAMPolicyAssignmentsInput struct {

	// The ID of the AWS account that contains these IAM policy assignments.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace for the assignments.
	//
	// This member is required.
	Namespace *string

	// The status of the assignments.
	AssignmentStatus types.AssignmentStatus

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type ListIAMPolicyAssignmentsOutput struct {

	// Information describing the IAM policy assignments.
	IAMPolicyAssignments []*types.IAMPolicyAssignmentSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListIAMPolicyAssignmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListIAMPolicyAssignments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListIAMPolicyAssignments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIAMPolicyAssignments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListIAMPolicyAssignments",
	}
}
