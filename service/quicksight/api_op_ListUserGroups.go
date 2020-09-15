// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the Amazon QuickSight groups that an Amazon QuickSight user is a member
// of.
func (c *Client) ListUserGroups(ctx context.Context, params *ListUserGroupsInput, optFns ...func(*Options)) (*ListUserGroupsOutput, error) {
	stack := middleware.NewStack("ListUserGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListUserGroupsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListUserGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUserGroups(options.Region), middleware.Before)

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
			OperationName: "ListUserGroups",
			Err:           err,
		}
	}
	out := result.(*ListUserGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserGroupsInput struct {
	// The maximum number of results to return from this request.
	MaxResults *int32
	// The namespace. Currently, you should set this to default.
	Namespace *string
	// A pagination token that can be used in a subsequent request.
	NextToken *string
	// The AWS account ID that the user is in. Currently, you use the ID for the AWS
	// account that contains your Amazon QuickSight account.
	AwsAccountId *string
	// The Amazon QuickSight user name that you want to list group memberships for.
	UserName *string
}

type ListUserGroupsOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The list of groups the user is a member of.
	GroupList []*types.Group
	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListUserGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListUserGroups{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListUserGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListUserGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListUserGroups",
	}
}
