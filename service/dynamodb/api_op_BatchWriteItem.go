// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The BatchWriteItem operation puts or deletes multiple items in one or more
// tables. A single call to BatchWriteItem can write up to 16 MB of data, which can
// comprise as many as 25 put or delete requests. Individual items to be written
// can be as large as 400 KB. BatchWriteItem cannot update items. To update items,
// use the UpdateItem action. The individual PutItem and DeleteItem operations
// specified in BatchWriteItem are atomic; however BatchWriteItem as a whole is
// not. If any requested operations fail because the table's provisioned throughput
// is exceeded or an internal processing failure occurs, the failed operations are
// returned in the UnprocessedItems response parameter. You can investigate and
// optionally resend the requests. Typically, you would call BatchWriteItem in a
// loop. Each iteration would check for unprocessed items and submit a new
// BatchWriteItem request with those unprocessed items until all items have been
// processed. If none of the items can be processed due to insufficient provisioned
// throughput on all of the tables in the request, then BatchWriteItem returns a
// ProvisionedThroughputExceededException. If DynamoDB returns any unprocessed
// items, you should retry the batch operation on those items. However, we strongly
// recommend that you use an exponential backoff algorithm. If you retry the batch
// operation immediately, the underlying read or write requests can still fail due
// to throttling on the individual tables. If you delay the batch operation using
// exponential backoff, the individual requests in the batch are much more likely
// to succeed. For more information, see Batch Operations and Error Handling
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#Programming.Errors.BatchOperations)
// in the Amazon DynamoDB Developer Guide.  <p>With <code>BatchWriteItem</code>,
// you can efficiently write or delete large amounts of data, such as from Amazon
// EMR, or copy data from another database into DynamoDB. In order to improve
// performance with these large-scale operations, <code>BatchWriteItem</code> does
// not behave in the same way as individual <code>PutItem</code> and
// <code>DeleteItem</code> calls would. For example, you cannot specify conditions
// on individual put and delete requests, and <code>BatchWriteItem</code> does not
// return deleted items in the response.</p> <p>If you use a programming language
// that supports concurrency, you can use threads to write items in parallel. Your
// application must include the necessary logic to manage the threads. With
// languages that don't support threading, you must update or delete the specified
// items one at a time. In both situations, <code>BatchWriteItem</code> performs
// the specified put and delete operations in parallel, giving you the power of the
// thread pool approach without having to introduce complexity into your
// application.</p> <p>Parallel processing reduces latency, but each specified put
// and delete request consumes the same number of write capacity units whether it
// is processed in parallel or not. Delete operations on nonexistent items consume
// one write capacity unit.</p> <p>If one or more of the following is true,
// DynamoDB rejects the entire batch write operation:</p> <ul> <li> <p>One or more
// tables specified in the <code>BatchWriteItem</code> request does not exist.</p>
// </li> <li> <p>Primary key attributes specified on an item in the request do not
// match those in the corresponding table's primary key schema.</p> </li> <li>
// <p>You try to perform multiple operations on the same item in the same
// <code>BatchWriteItem</code> request. For example, you cannot put and delete the
// same item in the same <code>BatchWriteItem</code> request. </p> </li> <li> <p>
// Your request contains at least two items with identical hash and range keys
// (which essentially is two put operations). </p> </li> <li> <p>There are more
// than 25 requests in the batch.</p> </li> <li> <p>Any individual item in a batch
// exceeds 400 KB.</p> </li> <li> <p>The total request size exceeds 16 MB.</p>
// </li> </ul>
func (c *Client) BatchWriteItem(ctx context.Context, params *BatchWriteItemInput, optFns ...func(*Options)) (*BatchWriteItemOutput, error) {
	if params == nil {
		params = &BatchWriteItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchWriteItem", params, optFns, addOperationBatchWriteItemMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchWriteItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchWriteItem operation.
type BatchWriteItemInput struct {

	// A map of one or more table names and, for each table, a list of operations to be
	// performed (DeleteRequest or PutRequest). Each element in the map consists of the
	// following:
	//
	//     * DeleteRequest - Perform a DeleteItem operation on the
	// specified item. The item to be deleted is identified by a Key subelement:
	//
	//
	// * Key - A map of primary key attribute values that uniquely identify the item.
	// Each entry in this map consists of an attribute name and an attribute value. For
	// each primary key, you must provide all of the key attributes. For example, with
	// a simple primary key, you only need to provide a value for the partition key.
	// For a composite primary key, you must provide values for both the partition key
	// and the sort key.
	//
	//     * PutRequest - Perform a PutItem operation on the
	// specified item. The item to be put is identified by an Item subelement:
	//
	//
	// * Item - A map of attributes and their values. Each entry in this map consists
	// of an attribute name and an attribute value. Attribute values must not be null;
	// string and binary type attributes must have lengths greater than zero; and set
	// type attributes must not be empty. Requests that contain empty values are
	// rejected with a ValidationException exception. If you specify any attributes
	// that are part of an index key, then the data types for those attributes must
	// match those of the schema in the table's attribute definition.
	//
	// This member is required.
	RequestItems map[string][]*types.WriteRequest

	// Determines the level of detail about provisioned throughput consumption that is
	// returned in the response:
	//
	//     * INDEXES - The response includes the aggregate
	// ConsumedCapacity for the operation, together with ConsumedCapacity for each
	// table and secondary index that was accessed. Note that some operations, such as
	// GetItem and BatchGetItem, do not access any indexes at all. In these cases,
	// specifying INDEXES will only return ConsumedCapacity information for table(s).
	//
	//
	// * TOTAL - The response includes only the aggregate ConsumedCapacity for the
	// operation.
	//
	//     * NONE - No ConsumedCapacity details are included in the
	// response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity

	// Determines whether item collection metrics are returned. If set to SIZE, the
	// response includes statistics about item collections, if any, that were modified
	// during the operation are returned in the response. If set to NONE (the default),
	// no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics
}

// Represents the output of a BatchWriteItem operation.
type BatchWriteItemOutput struct {

	// The capacity units consumed by the entire BatchWriteItem operation. Each element
	// consists of:
	//
	//     * TableName - The table that consumed the provisioned
	// throughput.
	//
	//     * CapacityUnits - The total number of capacity units consumed.
	ConsumedCapacity []*types.ConsumedCapacity

	// A list of tables that were processed by BatchWriteItem and, for each table,
	// information about any item collections that were affected by individual
	// DeleteItem or PutItem operations. Each entry consists of the following
	// subelements:
	//
	//     * ItemCollectionKey - The partition key value of the item
	// collection. This is the same as the partition key value of the item.
	//
	//     *
	// SizeEstimateRangeGB - An estimate of item collection size, expressed in GB. This
	// is a two-element array containing a lower bound and an upper bound for the
	// estimate. The estimate includes the size of all the items in the table, plus the
	// size of all attributes projected into all of the local secondary indexes on the
	// table. Use this estimate to measure whether a local secondary index is
	// approaching its size limit. The estimate is subject to change over time;
	// therefore, do not rely on the precision or accuracy of the estimate.
	ItemCollectionMetrics map[string][]*types.ItemCollectionMetrics

	// A map of tables and requests against those tables that were not processed. The
	// UnprocessedItems value is in the same form as RequestItems, so you can provide
	// this value directly to a subsequent BatchGetItem operation. For more
	// information, see RequestItems in the Request Parameters section. Each
	// UnprocessedItems entry consists of a table name and, for that table, a list of
	// operations to perform (DeleteRequest or PutRequest).
	//
	//     * DeleteRequest -
	// Perform a DeleteItem operation on the specified item. The item to be deleted is
	// identified by a Key subelement:
	//
	//         * Key - A map of primary key attribute
	// values that uniquely identify the item. Each entry in this map consists of an
	// attribute name and an attribute value.
	//
	//     * PutRequest - Perform a PutItem
	// operation on the specified item. The item to be put is identified by an Item
	// subelement:
	//
	//         * Item - A map of attributes and their values. Each entry
	// in this map consists of an attribute name and an attribute value. Attribute
	// values must not be null; string and binary type attributes must have lengths
	// greater than zero; and set type attributes must not be empty. Requests that
	// contain empty values will be rejected with a ValidationException exception. If
	// you specify any attributes that are part of an index key, then the data types
	// for those attributes must match those of the schema in the table's attribute
	// definition.
	//
	// If there are no unprocessed items remaining, the response contains
	// an empty UnprocessedItems map.
	UnprocessedItems map[string][]*types.WriteRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchWriteItemMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpBatchWriteItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchWriteItem{}, middleware.After)
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
	addOpBatchWriteItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchWriteItem(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opBatchWriteItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "BatchWriteItem",
	}
}
