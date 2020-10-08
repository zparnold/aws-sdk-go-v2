// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies the specified tags to the specified Amazon Chime SDK meeting.
func (c *Client) TagMeeting(ctx context.Context, params *TagMeetingInput, optFns ...func(*Options)) (*TagMeetingOutput, error) {
	if params == nil {
		params = &TagMeetingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagMeeting", params, optFns, addOperationTagMeetingMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*TagMeetingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagMeetingInput struct {

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string

	// The tag key-value pairs.
	//
	// This member is required.
	Tags []*types.Tag
}

type TagMeetingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagMeetingMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTagMeeting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTagMeeting{}, middleware.After)
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
	addOpTagMeetingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagMeeting(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTagMeeting(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "TagMeeting",
	}
}
