/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 *
 */

package software.amazon.smithy.aws.go.codegen;

import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoStackStepMiddlewareGenerator;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.SymbolUtils;
import software.amazon.smithy.go.codegen.integration.ProtocolUtils;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.utils.SmithyBuilder;

public class AwsHttpPresignURLClientGenerator {
    private static final String CONVERT_TO_PRESIGN_MIDDLEWARE_NAME = "convertToPresignMiddleware";

    private final Model model;
    private final SymbolProvider symbolProvider;

    private final Symbol clientSymbol;
    private final Symbol newClientSymbol;
    private final Symbol clientAPISymbol;

    private final OperationShape operation;
    private final Symbol operationSymbol;
    private final Shape operationInput;
    private final Symbol operationInputSymbol;
    private final Shape operationOutput;
    private final Symbol operationOutputSymbol;

    private final boolean exported;

    private AwsHttpPresignURLClientGenerator(Builder builder) {
        this.exported = builder.exported;

        this.model = SmithyBuilder.requiredState("model", builder.model);
        this.symbolProvider = SmithyBuilder.requiredState("symbolProvider", builder.symbolProvider);

        this.operation = SmithyBuilder.requiredState("operation", builder.operation);
        this.operationSymbol = symbolProvider.toSymbol(operation);

        this.operationInput = ProtocolUtils.expectInput(model, operation);
        this.operationInputSymbol = symbolProvider.toSymbol(operationInput);

        this.operationOutput = ProtocolUtils.expectOutput(model, operation);
        this.operationOutputSymbol = symbolProvider.toSymbol(operationOutput);

        this.clientSymbol = buildClientSymbol(operationSymbol, exported);
        this.newClientSymbol = buildNewClientSymbol(operationSymbol, exported);
        this.clientAPISymbol = buildAPIClientSymbol(operationSymbol, exported);
    }

    /**
     * Writes the Presign client's type and methods.
     *
     * @param writer writer to write to
     */
    public void writePresignClientType(GoWriter writer) {
        writer.addUseImports(SmithyGoDependency.CONTEXT);

        writer.openBlock("type $T interface {", "}", clientAPISymbol, () -> {
            writer.write("$T(context.Context, $P, ...func(*Options)) ($P, error)", operationSymbol,
                    operationInputSymbol, operationOutputSymbol);
        });

        writer.openBlock("type $T struct {", "}", clientSymbol, () -> {
            // TODO if presigner client can depend directly on client
            writer.write("client $T", clientAPISymbol);
            writer.openBlock("signer interface{", "}", () -> {
                writer.addUseImports(AwsGoDependency.AWS_CORE);
                writer.addUseImports(SmithyGoDependency.NET_HTTP);
                writer.addUseImports(SmithyGoDependency.TIME);
                writer.write(
                        "PresignHTTP(context.Context, aws.Credentials, *http.Request, string, string, string, time.Duration, time.Time) (string, http.Header, error)");
            });
        });

        // TODO consider having the new presigner take options, not a SDK client, this allows the presigner to create its own client with its own configuration, e.g. stub http client.
        writer.openBlock("func $L(client $T) *$T {", "}", newClientSymbol.getName(), clientAPISymbol, clientSymbol,
                () -> {
                    writer.openBlock("return &$T{", "}", clientSymbol, () -> {
                        writer.write("client: client,");
                        writer.write("signer: v4.NewSigner(),");
                    });
                });

        writer.openBlock(
                // TODO consider creating type for presign parameters for future compatibility.
                "func (c *$T) Presign$T(ctx context.Context, input $P, optFns ...func(*Options)) "
                        + "(string, http.Header, error) {",
                "}",
                clientSymbol, operationSymbol, operationInputSymbol,
                () -> {
                    Symbol nopClient = SymbolUtils.createPointableSymbolBuilder("NopClient",
                            SmithyGoDependency.SMITHY_HTTP_TRANSPORT)
                            .build();

                    // TODO could be replaced with a `WithAPIOptions` client option helper.
                    // TODO could be replaced with a `WithHTTPClient` client option helper.
                    writer.openBlock("optFns = append(optFns, func(o *Options) {", "})", () -> {
                        writer.write("o.APIOptions = append(o.APIOptions, $L)", CONVERT_TO_PRESIGN_MIDDLEWARE_NAME);
                        writer.write("o.HTTPClient = &$T{}", nopClient);
                    });
                    writer.write("_, err := c.client.$T(ctx, input, optFns...)", operationSymbol);
                    writer.write("if err != nil { return ``, nil, err }");

                    // TODO need way to extract presigned URL from presign middleware. Probably pre-create the middleware
                    // TODO then query it when operation completes. alternative is to inject presigned as response metadata
                    // TODO on result shape.
                    writer.write("return ``, http.Header{}, fmt.Errorf(\"not implemented\")");
                });
    }

    /**
     * Writes the Presign client's conversion middleware to convert a regular API request to a presigned request.
     *
     * @param writer writer to write to
     */
    public static void writeConvertToPresignMiddleware(GoWriter writer) {
        Symbol smithyStack = SymbolUtils.createPointableSymbolBuilder("Stack", SmithyGoDependency.SMITHY_MIDDLEWARE)
                .build();
        Symbol smithyAfter = SymbolUtils.createValueSymbolBuilder("After", SmithyGoDependency.SMITHY_MIDDLEWARE)
                .build();

        // Middleware to remove
        Symbol requestInvocationIDMiddleware = SymbolUtils.createValueSymbolBuilder("RequestInvocationIDMiddleware",
                AwsGoDependency.AWS_MIDDLEWARE)
                .build();

        Symbol presignMiddleware = SymbolUtils.createValueSymbolBuilder("NewPresignHTTPRequestMiddleware",
                AwsGoDependency.AWS_SIGNER_V4)
                .build();

        // Middleware to add
        // TODO only for query protocols.
        Symbol queryAsGetMiddleware = SymbolUtils.createValueSymbolBuilder("AddAsGetRequestMiddleware",
                AwsGoDependency.AWS_QUERY_PROTOCOL)
                .build();

        writer.openBlock("func $L(stack $P) error {", "}",
                CONVERT_TO_PRESIGN_MIDDLEWARE_NAME,
                smithyStack,
                () -> {
                    writer.write("stack.Finalize.Clear()");
                    writer.write("stack.Deserialize.Clear()");
                    writer.write("stack.Build.Remove($T{}.ID())", requestInvocationIDMiddleware);

                    writer.write("stack.Finalize.Add($T(), $T)", presignMiddleware, smithyAfter);

                    // TODO only for query protocols.
                    writer.write("$T(stack)", queryAsGetMiddleware);

                    writer.write("return nil");
                });
    }

    // TODO probably want to rename these to genPresignClientSymbol etc.
    public Symbol getClientSymbol() {
        return clientSymbol;
    }

    public Symbol getNewClientSymbol() {
        return newClientSymbol;
    }

    private static Symbol buildNewClientSymbol(Symbol operation, boolean exported) {
        String name = String.format("New%sHTTPPresignURLClient", operation.getName());
        return buildSymbol(name, exported);
    }

    private static Symbol buildClientSymbol(Symbol operation, boolean exported) {
        String name = String.format("%sHTTPPresignURLClient", operation.getName());
        return buildSymbol(name, exported);
    }

    private static Symbol buildAPIClientSymbol(Symbol operation, boolean exported) {
        String name = String.format("%sAPIClient", operation.getName());
        return buildSymbol(name, exported);
    }

    private static Symbol buildSymbol(String name, boolean exported) {
        if (!exported) {
            name = Character.toLowerCase(name.charAt(0)) + name.substring(1);
        }
        return SymbolUtils.createValueSymbolBuilder(name).
                build();
    }

    public static class Builder implements SmithyBuilder<AwsHttpPresignURLClientGenerator> {
        private Model model;
        private SymbolProvider symbolProvider;
        private OperationShape operation;
        private boolean exported;

        public Builder model(Model model) {
            this.model = model;
            return this;
        }

        public Builder symbolProvider(SymbolProvider symbolProvider) {
            this.symbolProvider = symbolProvider;
            return this;
        }

        public Builder operation(OperationShape operation) {
            this.operation = operation;
            return this;
        }

        public Builder exported() {
            return this.exported(true);
        }

        public Builder exported(boolean exported) {
            this.exported = exported;
            return this;
        }

        public AwsHttpPresignURLClientGenerator build() {
            return new AwsHttpPresignURLClientGenerator(this);
        }
    }
}
