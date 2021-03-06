// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use the GetUtterancesView operation to get information about the utterances that
// your users have made to your bot. You can use this list to tune the utterances
// that your bot responds to. For example, say that you have created a bot to order
// flowers. After your users have used your bot for a while, use the
// GetUtterancesView operation to see the requests that they have made and whether
// they have been successful. You might find that the utterance "I want flowers" is
// not being recognized. You could add this utterance to the OrderFlowers intent so
// that your bot recognizes that utterance. After you publish a new version of a
// bot, you can get information about the old version and the new so that you can
// compare the performance across the two versions. Utterance statistics are
// generated once a day. Data is available for the last 15 days. You can request
// information for up to 5 versions of your bot in each request. Amazon Lex returns
// the most frequent utterances received by the bot in the last 15 days. The
// response contains information about a maximum of 100 utterances for each
// version. If you set childDirected field to true when you created your bot, or if
// you opted out of participating in improving Amazon Lex, utterances are not
// available. This operation requires permissions for the lex:GetUtterancesView
// action.
func (c *Client) GetUtterancesView(ctx context.Context, params *GetUtterancesViewInput, optFns ...func(*Options)) (*GetUtterancesViewOutput, error) {
	stack := middleware.NewStack("GetUtterancesView", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUtterancesViewMiddlewares(stack)
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
	addOpGetUtterancesViewValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUtterancesView(options.Region), middleware.Before)
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
			OperationName: "GetUtterancesView",
			Err:           err,
		}
	}
	out := result.(*GetUtterancesViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUtterancesViewInput struct {

	// The name of the bot for which utterance information should be returned.
	//
	// This member is required.
	BotName *string

	// An array of bot versions for which utterance information should be returned. The
	// limit is 5 versions per request.
	//
	// This member is required.
	BotVersions []*string

	// To return utterances that were recognized and handled, use Detected. To return
	// utterances that were not recognized, use Missed.
	//
	// This member is required.
	StatusType types.StatusType
}

type GetUtterancesViewOutput struct {

	// The name of the bot for which utterance information was returned.
	BotName *string

	// An array of UtteranceList () objects, each containing a list of UtteranceData ()
	// objects describing the utterances that were processed by your bot. The response
	// contains a maximum of 100 UtteranceData objects for each version. Amazon Lex
	// returns the most frequent utterances received by the bot in the last 15 days.
	Utterances []*types.UtteranceList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUtterancesViewMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUtterancesView{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUtterancesView{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUtterancesView(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetUtterancesView",
	}
}
