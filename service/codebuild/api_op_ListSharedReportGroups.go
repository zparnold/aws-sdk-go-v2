// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of report groups that are shared with other AWS accounts or users.
func (c *Client) ListSharedReportGroups(ctx context.Context, params *ListSharedReportGroupsInput, optFns ...func(*Options)) (*ListSharedReportGroupsOutput, error) {
	stack := middleware.NewStack("ListSharedReportGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListSharedReportGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSharedReportGroups(options.Region), middleware.Before)

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
			OperationName: "ListSharedReportGroups",
			Err:           err,
		}
	}
	out := result.(*ListSharedReportGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSharedReportGroupsInput struct {
	// The maximum number of paginated shared report groups per response. Use nextToken
	// to iterate pages in the list of returned ReportGroup objects. The default value
	// is 100.
	MaxResults *int32
	// The criterion to be used to list report groups shared with the current AWS
	// account or user. Valid values include:
	//
	//     * ARN: List based on the ARN.
	//
	//     *
	// MODIFIED_TIME: List based on when information about the shared report group was
	// last changed.
	SortBy types.SharedResourceSortByType
	// The order in which to list shared report groups. Valid values include:
	//
	//     *
	// ASCENDING: List in ascending order.
	//
	//     * DESCENDING: List in descending order.
	SortOrder types.SortOrderType
	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string
}

type ListSharedReportGroupsOutput struct {
	// The list of ARNs for the report groups shared with the current AWS account or
	// user.
	ReportGroups []*string
	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListSharedReportGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListSharedReportGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSharedReportGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSharedReportGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListSharedReportGroups",
	}
}
