// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	cryptorand "crypto/rand"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	smithyrand "github.com/awslabs/smithy-go/rand"
	"net/http"
	"time"
)

const ServiceID = "Chime"
const ServiceAPIVersion = "2018-05-01"

// The Amazon Chime API (application programming interface) is designed for
// developers to perform key tasks, such as creating and managing Amazon Chime
// accounts, users, and Voice Connectors. This guide provides detailed information
// about the Amazon Chime API, including operations, types, inputs and outputs, and
// error codes. It also includes some server-side API actions to use with the
// Amazon Chime SDK. For more information about the Amazon Chime SDK, see Using the
// Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide. You can use an AWS SDK, the AWS Command
// Line Interface (AWS CLI), or the REST API to make API calls. We recommend using
// an AWS SDK or the AWS CLI. Each API operation includes links to information
// about using it with a language-specific AWS SDK or the AWS CLI. Using an AWS SDK
// You don't need to write code to calculate a signature for request
// authentication. The SDK clients authenticate your requests by using access keys
// that you provide. For more information about AWS SDKs, see the AWS Developer
// Center (http://aws.amazon.com/developer/). Using the AWS CLI Use your access
// keys with the AWS CLI to make API calls. For information about setting up the
// AWS CLI, see Installing the AWS Command Line Interface
// (https://docs.aws.amazon.com/cli/latest/userguide/installing.html) in the AWS
// Command Line Interface User Guide. For a list of available Amazon Chime
// commands, see the Amazon Chime commands
// (https://docs.aws.amazon.com/cli/latest/reference/chime/index.html) in the AWS
// CLI Command Reference. Using REST API If you use REST to make API calls, you
// must authenticate your request by providing a signature. Amazon Chime supports
// signature version 4. For more information, see Signature Version 4 Signing
// Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
// in the Amazon Web Services General Reference. When making REST API calls, use
// the service name chime and REST endpoint https://service.chime.aws.amazon.com.
// <p>Administrative permissions are controlled using AWS Identity and Access
// Management (IAM). For more information, see <a
// href="https://docs.aws.amazon.com/chime/latest/ag/security-iam.html">Identity
// and Access Management for Amazon Chime</a> in the <i>Amazon Chime Administration
// Guide</i>.</p>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	resolveIdempotencyTokenProvider(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// Provides idempotency tokens values that will be automatically populated into
	// idempotent API operations.
	IdempotencyTokenProvider IdempotencyTokenProvider

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
		APIOptions:  cfg.APIOptions,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("chime")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}

func resolveIdempotencyTokenProvider(o *Options) {
	if o.IdempotencyTokenProvider != nil {
		return
	}
	o.IdempotencyTokenProvider = smithyrand.NewUUIDIdempotencyToken(cryptorand.Reader)
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer: o.Retryer,
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

// IdempotencyTokenProvider interface for providing idempotency token
type IdempotencyTokenProvider interface {
	GetIdempotencyToken() (string, error)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}
