// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves valid VPC peering authorizations that are pending for the AWS account.
// This operation returns all VPC peering authorizations and requests for peering.
// This includes those initiated and received by this account.
//
//     *
// CreateVpcPeeringAuthorization ()
//
//     * DescribeVpcPeeringAuthorizations ()
//
//
// * DeleteVpcPeeringAuthorization ()
//
//     * CreateVpcPeeringConnection ()
//
//     *
// DescribeVpcPeeringConnections ()
//
//     * DeleteVpcPeeringConnection ()
func (c *Client) DescribeVpcPeeringAuthorizations(ctx context.Context, params *DescribeVpcPeeringAuthorizationsInput, optFns ...func(*Options)) (*DescribeVpcPeeringAuthorizationsOutput, error) {
	if params == nil {
		params = &DescribeVpcPeeringAuthorizationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcPeeringAuthorizations", params, optFns, addOperationDescribeVpcPeeringAuthorizationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcPeeringAuthorizationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcPeeringAuthorizationsInput struct {
}

type DescribeVpcPeeringAuthorizationsOutput struct {

	// A collection of objects that describe all valid VPC peering operations for the
	// current AWS account.
	VpcPeeringAuthorizations []*types.VpcPeeringAuthorization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcPeeringAuthorizationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVpcPeeringAuthorizations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVpcPeeringAuthorizations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringAuthorizations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeVpcPeeringAuthorizations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeVpcPeeringAuthorizations",
	}
}
