// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves an authorization token. An authorization token represents your IAM
// authentication credentials and can be used to access any Amazon ECR registry
// that your IAM principal has access to. The authorization token is valid for 12
// hours. The authorizationToken returned is a base64 encoded string that can be
// decoded and used in a docker login command to authenticate to a registry. The
// AWS CLI offers an get-login-password command that simplifies the login process.
// For more information, see Registry Authentication
// (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth)
// in the Amazon Elastic Container Registry User Guide.
func (c *Client) GetAuthorizationToken(ctx context.Context, params *GetAuthorizationTokenInput, optFns ...func(*Options)) (*GetAuthorizationTokenOutput, error) {
	if params == nil {
		params = &GetAuthorizationTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAuthorizationToken", params, optFns, addOperationGetAuthorizationTokenMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAuthorizationTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAuthorizationTokenInput struct {

	// A list of AWS account IDs that are associated with the registries for which to
	// get AuthorizationData objects. If you do not specify a registry, the default
	// registry is assumed.
	RegistryIds []*string
}

type GetAuthorizationTokenOutput struct {

	// A list of authorization token data objects that correspond to the registryIds
	// values in the request.
	AuthorizationData []*types.AuthorizationData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAuthorizationTokenMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAuthorizationToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAuthorizationToken{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAuthorizationToken(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAuthorizationToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "GetAuthorizationToken",
	}
}
