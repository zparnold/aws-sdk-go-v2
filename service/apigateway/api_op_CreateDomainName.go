// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new domain name.
func (c *Client) CreateDomainName(ctx context.Context, params *CreateDomainNameInput, optFns ...func(*Options)) (*CreateDomainNameOutput, error) {
	stack := middleware.NewStack("CreateDomainName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDomainNameMiddlewares(stack)
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
	addOpCreateDomainNameValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainName(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "CreateDomainName",
			Err:           err,
		}
	}
	out := result.(*CreateDomainNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a new domain name.
type CreateDomainNameInput struct {

	// [Required] The name of the DomainName () resource.
	//
	// This member is required.
	DomainName *string

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string

	// [Deprecated] The body of the server certificate that will be used by
	// edge-optimized endpoint for this domain name provided by your certificate
	// authority.
	CertificateBody *string

	// [Deprecated] The intermediate certificates and optionally the root certificate,
	// one after the other without any blank lines, used by an edge-optimized endpoint
	// for this domain name. If you include the root certificate, your certificate
	// chain must start with intermediate certificates and end with the root
	// certificate. Use the intermediate certificates that were provided by your
	// certificate authority. Do not include any intermediaries that are not in the
	// chain of trust path.
	CertificateChain *string

	// The user-friendly name of the certificate that will be used by edge-optimized
	// endpoint for this domain name.
	CertificateName *string

	// [Deprecated] Your edge-optimized endpoint's domain name certificate's private
	// key.
	CertificatePrivateKey *string

	// The endpoint configuration of this DomainName () showing the endpoint types of
	// the domain name.
	EndpointConfiguration *types.EndpointConfiguration

	Name *string

	// The reference to an AWS-managed certificate that will be used by regional
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	RegionalCertificateArn *string

	// The user-friendly name of the certificate that will be used by regional endpoint
	// for this domain name.
	RegionalCertificateName *string

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName
	// (). The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy types.SecurityPolicy

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]*string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a custom domain name as a user-friendly host name of an API (RestApi
// ()). When you deploy an API, API Gateway creates a default host name for the
// API. This default API host name is of the
// {restapi-id}.execute-api.{region}.amazonaws.com format. With the default host
// name, you can access the API's root resource with the URL of
// https://{restapi-id}.execute-api.{region}.amazonaws.com/{stage}/. When you set
// up a custom domain name of apis.example.com for this API, you can then access
// the same resource using the URL of the https://apis.examples.com/myApi, where
// myApi is the base path mapping (BasePathMapping ()) of your API under the custom
// domain name. Set a Custom Host Name for an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type CreateDomainNameOutput struct {

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string

	// The name of the certificate that will be used by edge-optimized endpoint for
	// this domain name.
	CertificateName *string

	// The timestamp when the certificate that was used by edge-optimized endpoint for
	// this domain name was uploaded.
	CertificateUploadDate *time.Time

	// The domain name of the Amazon CloudFront distribution associated with this
	// custom domain name for an edge-optimized endpoint. You set up this association
	// when adding a DNS record pointing the custom domain name to this distribution
	// name. For more information about CloudFront distributions, see the Amazon
	// CloudFront documentation (https://aws.amazon.com/documentation/cloudfront/).
	DistributionDomainName *string

	// The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized
	// endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more
	// information, see Set up a Regional Custom Domain Name
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html)
	// and AWS Regions and Endpoints for API Gateway
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).
	DistributionHostedZoneId *string

	// The custom domain name as an API host name, for example, my-api.example.com.
	DomainName *string

	// The status of the DomainName () migration. The valid values are AVAILABLE and
	// UPDATING. If the status is UPDATING, the domain cannot be modified further until
	// the existing operation is complete. If it is AVAILABLE, the domain can be
	// updated.
	DomainNameStatus types.DomainNameStatus

	// An optional text message containing detailed information about status of the
	// DomainName () migration.
	DomainNameStatusMessage *string

	// The endpoint configuration of this DomainName () showing the endpoint types of
	// the domain name.
	EndpointConfiguration *types.EndpointConfiguration

	// The reference to an AWS-managed certificate that will be used for validating the
	// regional domain name. AWS Certificate Manager is the only supported source.
	RegionalCertificateArn *string

	// The name of the certificate that will be used for validating the regional domain
	// name.
	RegionalCertificateName *string

	// The domain name associated with the regional endpoint for this custom domain
	// name. You set up this association by adding a DNS record that points the custom
	// domain name to this regional domain name. The regional domain name is returned
	// by API Gateway when you create a regional endpoint.
	RegionalDomainName *string

	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint. For
	// more information, see Set up a Regional Custom Domain Name
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html)
	// and AWS Regions and Endpoints for API Gateway
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).
	RegionalHostedZoneId *string

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName
	// (). The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy types.SecurityPolicy

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDomainNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainName{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainName{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDomainName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateDomainName",
	}
}
