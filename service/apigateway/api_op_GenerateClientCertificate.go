// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Generates a ClientCertificate () resource.
func (c *Client) GenerateClientCertificate(ctx context.Context, params *GenerateClientCertificateInput, optFns ...func(*Options)) (*GenerateClientCertificateOutput, error) {
	stack := middleware.NewStack("GenerateClientCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGenerateClientCertificateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateClientCertificate(options.Region), middleware.Before)

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
			OperationName: "GenerateClientCertificate",
			Err:           err,
		}
	}
	out := result.(*GenerateClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to generate a ClientCertificate () resource.
type GenerateClientCertificateInput struct {
	Template         *bool
	Title            *string
	TemplateSkipList []*string
	Name             *string
	// The description of the ClientCertificate ().
	Description *string
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]*string
}

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint. Client certificates are used
// to authenticate an API by the backend server. To authenticate an API client (or
// user), use IAM roles and policies, a custom Authorizer () or an Amazon Cognito
// user pool. Use Client-Side Certificate
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type GenerateClientCertificateOutput struct {
	// The timestamp when the client certificate was created.
	CreatedDate *time.Time
	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string
	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time
	// The description of the client certificate.
	Description *string
	// The identifier of the client certificate.
	ClientCertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGenerateClientCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGenerateClientCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateClientCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGenerateClientCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GenerateClientCertificate",
	}
}
