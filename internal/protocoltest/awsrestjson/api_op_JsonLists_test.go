// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestClient_JsonLists_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonListsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes JSON lists
		"RestJsonLists": {
			Params: &JsonListsInput{
				StringList: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				StringSet: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				IntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
				},
				BooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
				},
				TimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
				},
				EnumList: []types.FooEnum{
					"Foo",
					"0",
				},
				NestedStringList: [][]*string{
					{
						ptr.String("foo"),
						ptr.String("bar"),
					},
					{
						ptr.String("baz"),
						ptr.String("qux"),
					},
				},
				StructureList: []*types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
					},
				},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonLists",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "stringList": [
			        "foo",
			        "bar"
			    ],
			    "stringSet": [
			        "foo",
			        "bar"
			    ],
			    "integerList": [
			        1,
			        2
			    ],
			    "booleanList": [
			        true,
			        false
			    ],
			    "timestampList": [
			        1398796238,
			        1398796238
			    ],
			    "enumList": [
			        "Foo",
			        "0"
			    ],
			    "nestedStringList": [
			        [
			            "foo",
			            "bar"
			        ],
			        [
			            "baz",
			            "qux"
			        ]
			    ],
			    "myStructureList": [
			        {
			            "value": "1",
			            "other": "2"
			        },
			        {
			            "value": "3",
			            "other": "4"
			        }
			    ]
			}`))
			},
		},
		// Serializes empty JSON lists
		"RestJsonListsEmpty": {
			Params: &JsonListsInput{
				StringList: []*string{},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonLists",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "stringList": []
			}`))
			},
		},
		// Serializes null values in lists
		"RestJsonListsSerializeNull": {
			Params: &JsonListsInput{
				StringList: []*string{
					nil,
				},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonLists",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "stringList": [
			        null
			    ]
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					actualReq = r
					return &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			result, err := client.JsonLists(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.Path; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_JsonLists_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonListsOutput
	}{
		// Serializes JSON lists
		"RestJsonLists": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "stringList": [
			        "foo",
			        "bar"
			    ],
			    "stringSet": [
			        "foo",
			        "bar"
			    ],
			    "integerList": [
			        1,
			        2
			    ],
			    "booleanList": [
			        true,
			        false
			    ],
			    "timestampList": [
			        1398796238,
			        1398796238
			    ],
			    "enumList": [
			        "Foo",
			        "0"
			    ],
			    "nestedStringList": [
			        [
			            "foo",
			            "bar"
			        ],
			        [
			            "baz",
			            "qux"
			        ]
			    ],
			    "myStructureList": [
			        {
			            "value": "1",
			            "other": "2"
			        },
			        {
			            "value": "3",
			            "other": "4"
			        }
			    ]
			}`),
			ExpectResult: &JsonListsOutput{
				StringList: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				StringSet: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				IntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
				},
				BooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
				},
				TimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
				},
				EnumList: []types.FooEnum{
					"Foo",
					"0",
				},
				NestedStringList: [][]*string{
					{
						ptr.String("foo"),
						ptr.String("bar"),
					},
					{
						ptr.String("baz"),
						ptr.String("qux"),
					},
				},
				StructureList: []*types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
					},
				},
			},
		},
		// Serializes empty JSON lists
		"RestJsonListsEmpty": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "stringList": []
			}`),
			ExpectResult: &JsonListsOutput{
				StringList: []*string{},
			},
		},
		// Serializes null values in lists
		"RestJsonListsSerializeNull": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "stringList": [
			        null
			    ]
			}`),
			ExpectResult: &JsonListsOutput{
				StringList: []*string{
					nil,
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: c.StatusCode,
						Header:     c.Header.Clone(),
						Body:       ioutil.NopCloser(bytes.NewReader(c.Body)),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			var params JsonListsInput
			result, err := client.JsonLists(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if diff := cmp.Diff(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); len(diff) != 0 {
				t.Errorf("expect c.ExpectResult value match:\n%s", diff)
			}
		})
	}
}