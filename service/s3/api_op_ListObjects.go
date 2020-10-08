// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns some or all (up to 1,000) of the objects in a bucket. You can use the
// request parameters as selection criteria to return a subset of the objects in a
// bucket. A 200 OK response can contain valid or invalid XML. Be sure to design
// your application to parse the contents of the response and handle it
// appropriately. This API has been revised. We recommend that you use the newer
// version, ListObjectsV2 (), when developing applications. For backward
// compatibility, Amazon S3 continues to support ListObjects.  <p>The following
// operations are related to <code>ListObjects</code>:</p> <ul> <li> <p>
// <a>ListObjectsV2</a> </p> </li> <li> <p> <a>GetObject</a> </p> </li> <li> <p>
// <a>PutObject</a> </p> </li> <li> <p> <a>CreateBucket</a> </p> </li> <li> <p>
// <a>ListBuckets</a> </p> </li> </ul>
func (c *Client) ListObjects(ctx context.Context, params *ListObjectsInput, optFns ...func(*Options)) (*ListObjectsOutput, error) {
	if params == nil {
		params = &ListObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjects", params, optFns, addOperationListObjectsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectsInput struct {

	// The name of the bucket containing the objects.
	//
	// This member is required.
	Bucket *string

	// A delimiter is a character you use to group keys.
	Delimiter *string

	// Requests Amazon S3 to encode the object keys in the response and specifies the
	// encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters with an
	// ASCII value from 0 to 10. For characters that are not supported in XML 1.0, you
	// can add this parameter to request that Amazon S3 encode the keys in the
	// response.
	EncodingType types.EncodingType

	// Specifies the key to start with when listing objects in a bucket.
	Marker *string

	// Sets the maximum number of keys returned in the response. By default the API
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	MaxKeys *int32

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string

	// Confirms that the requester knows that she or he will be charged for the list
	// objects request. Bucket owners need not specify this parameter in their
	// requests.
	RequestPayer types.RequestPayer
}

type ListObjectsOutput struct {

	// All of the keys rolled up in a common prefix count as a single return when
	// calculating the number of returns.  <p>A response can contain CommonPrefixes
	// only if you specify a delimiter.</p> <p>CommonPrefixes contains all (if there
	// are any) keys between Prefix and the next occurrence of the string specified by
	// the delimiter.</p> <p> CommonPrefixes lists keys that act like subdirectories in
	// the directory specified by Prefix.</p> <p>For example, if the prefix is notes/
	// and the delimiter is a slash (/) as in notes/summer/july, the common prefix is
	// notes/summer/. All of the keys that roll up into a common prefix count as a
	// single return when calculating the number of returns.</p>
	CommonPrefixes []*types.CommonPrefix

	// Metadata about each object returned.
	Contents []*types.Object

	// Causes keys that contain the same string between the prefix and the first
	// occurrence of the delimiter to be rolled up into a single result element in the
	// CommonPrefixes collection. These rolled-up keys are not returned elsewhere in
	// the response. Each rolled-up result counts as only one return against the
	// MaxKeys value.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType types.EncodingType

	// A flag that indicates whether Amazon S3 returned all of the results that
	// satisfied the search criteria.
	IsTruncated *bool

	// Indicates where in the bucket listing begins. Marker is included in the response
	// if it was sent with the request.
	Marker *string

	// The maximum number of keys returned in the response body.
	MaxKeys *int32

	// Bucket name.
	Name *string

	// When response is truncated (the IsTruncated element value in the response is
	// true), you can use the key name in this field as marker in the subsequent
	// request to get next set of objects. Amazon S3 lists objects in alphabetical
	// order Note: This element is returned only if you have delimiter request
	// parameter specified. If response does not include the NextMaker and it is
	// truncated, you can use the value of the last Key in the response as the marker
	// in the subsequent request to get the next set of object keys.
	NextMarker *string

	// Keys that begin with the indicated prefix.
	Prefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListObjectsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListObjects{}, middleware.After)
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
	addOpListObjectsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjects(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opListObjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListObjects",
	}
}
