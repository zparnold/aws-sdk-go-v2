// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current role and list of Amazon S3 log buckets used by the DDoS
// Response Team (DRT) to access your AWS account while assisting with attack
// mitigation.
func (c *Client) DescribeDRTAccess(ctx context.Context, params *DescribeDRTAccessInput, optFns ...func(*Options)) (*DescribeDRTAccessOutput, error) {
	if params == nil {
		params = &DescribeDRTAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDRTAccess", params, optFns, addOperationDescribeDRTAccessMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDRTAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDRTAccessInput struct {
}

type DescribeDRTAccessOutput struct {

	// The list of Amazon S3 buckets accessed by the DRT.
	LogBucketList []*string

	// The Amazon Resource Name (ARN) of the role the DRT used to access your AWS
	// account.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDRTAccessMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDRTAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDRTAccess{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDRTAccess(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDRTAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DescribeDRTAccess",
	}
}
