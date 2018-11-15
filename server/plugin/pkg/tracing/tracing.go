/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package tracing

import (
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"
	"golang.org/x/net/context"
)

const CTX_TRACE_SPAN = "x-trace-span"

type Request interface{}
type Span interface{}

type Tracing interface {
	ServerBegin(operationName string, r Request) Span
	ServerEnd(span Span, code int, message string)
	ClientBegin(operationName string, r Request) Span
	ClientEnd(span Span, code int, message string)
}

type RegistryRequest struct {
	Ctx      context.Context
	Endpoint string
	Options  registry.PluginOp
}