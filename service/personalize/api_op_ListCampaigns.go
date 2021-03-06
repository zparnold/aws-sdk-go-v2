// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of campaigns that use the given solution. When a solution is not
// specified, all the campaigns associated with the account are listed. The
// response provides the properties for each campaign, including the Amazon
// Resource Name (ARN). For more information on campaigns, see CreateCampaign ().
func (c *Client) ListCampaigns(ctx context.Context, params *ListCampaignsInput, optFns ...func(*Options)) (*ListCampaignsOutput, error) {
	stack := middleware.NewStack("ListCampaigns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListCampaignsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCampaigns(options.Region), middleware.Before)
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
			OperationName: "ListCampaigns",
			Err:           err,
		}
	}
	out := result.(*ListCampaignsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCampaignsInput struct {

	// The maximum number of campaigns to return.
	MaxResults *int32

	// A token returned from the previous call to ListCampaigns for getting the next
	// set of campaigns (if they exist).
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution to list the campaigns for. When a
	// solution is not specified, all the campaigns associated with the account are
	// listed.
	SolutionArn *string
}

type ListCampaignsOutput struct {

	// A list of the campaigns.
	Campaigns []*types.CampaignSummary

	// A token for getting the next set of campaigns (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListCampaignsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListCampaigns{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCampaigns{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCampaigns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListCampaigns",
	}
}
