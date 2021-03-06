// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard
// (PutDeliverabilityDashboardOption operation) for the domain.
func (c *Client) ListDomainDeliverabilityCampaigns(ctx context.Context, params *ListDomainDeliverabilityCampaignsInput, optFns ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error) {
	stack := middleware.NewStack("ListDomainDeliverabilityCampaigns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDomainDeliverabilityCampaignsMiddlewares(stack)
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
	addOpListDomainDeliverabilityCampaignsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainDeliverabilityCampaigns(options.Region), middleware.Before)
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
			OperationName: "ListDomainDeliverabilityCampaigns",
			Err:           err,
		}
	}
	out := result.(*ListDomainDeliverabilityCampaignsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard
// (PutDeliverabilityDashboardOption operation) for the domain.
type ListDomainDeliverabilityCampaignsInput struct {

	// The last day, in Unix time format, that you want to obtain deliverability data
	// for. This value has to be less than or equal to 30 days after the value of the
	// StartDate parameter.
	//
	// This member is required.
	EndDate *time.Time

	// The first day, in Unix time format, that you want to obtain deliverability data
	// for.
	//
	// This member is required.
	StartDate *time.Time

	// The domain to obtain deliverability data for.
	//
	// This member is required.
	SubscribedDomain *string

	// A token that’s returned from a previous call to the
	// ListDomainDeliverabilityCampaigns operation. This token indicates the position
	// of a campaign in the list of campaigns.
	NextToken *string

	// The maximum number of results to include in response to a single call to the
	// ListDomainDeliverabilityCampaigns operation. If the number of results is larger
	// than the number that you specify in this parameter, the response includes a
	// NextToken element, which you can use to obtain additional results.
	PageSize *int32
}

// An array of objects that provide deliverability data for all the campaigns that
// used a specific domain to send email during a specified time range. This data is
// available for a domain only if you enabled the Deliverability dashboard
// (PutDeliverabilityDashboardOption operation) for the domain.
type ListDomainDeliverabilityCampaignsOutput struct {

	// An array of responses, one for each campaign that used the domain to send email
	// during the specified time range.
	//
	// This member is required.
	DomainDeliverabilityCampaigns []*types.DomainDeliverabilityCampaign

	// A token that’s returned from a previous call to the
	// ListDomainDeliverabilityCampaigns operation. This token indicates the position
	// of the campaign in the list of campaigns.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDomainDeliverabilityCampaignsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDomainDeliverabilityCampaigns{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainDeliverabilityCampaigns{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDomainDeliverabilityCampaigns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDomainDeliverabilityCampaigns",
	}
}
