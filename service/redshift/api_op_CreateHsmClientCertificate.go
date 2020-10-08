// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an HSM client certificate that an Amazon Redshift cluster will use to
// connect to the client's HSM in order to store and retrieve the keys used to
// encrypt the cluster databases. The command returns a public key, which you must
// store in the HSM. In addition to creating the HSM certificate, you must create
// an Amazon Redshift HSM configuration that provides a cluster the information
// needed to store and use encryption keys in the HSM. For more information, go to
// Hardware Security Modules
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-HSM.html) in the
// Amazon Redshift Cluster Management Guide.
func (c *Client) CreateHsmClientCertificate(ctx context.Context, params *CreateHsmClientCertificateInput, optFns ...func(*Options)) (*CreateHsmClientCertificateOutput, error) {
	if params == nil {
		params = &CreateHsmClientCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHsmClientCertificate", params, optFns, addOperationCreateHsmClientCertificateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHsmClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateHsmClientCertificateInput struct {

	// The identifier to be assigned to the new HSM client certificate that the cluster
	// will use to connect to the HSM to use the database encryption keys.
	//
	// This member is required.
	HsmClientCertificateIdentifier *string

	// A list of tag instances.
	Tags []*types.Tag
}

type CreateHsmClientCertificateOutput struct {

	// Returns information about an HSM client certificate. The certificate is stored
	// in a secure Hardware Storage Module (HSM), and used by the Amazon Redshift
	// cluster to encrypt data files.
	HsmClientCertificate *types.HsmClientCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHsmClientCertificateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateHsmClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateHsmClientCertificate{}, middleware.After)
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
	addOpCreateHsmClientCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsmClientCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateHsmClientCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateHsmClientCertificate",
	}
}
