// Copyright 2022 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errs

import (
	"fmt"

	"github.com/pingcap/errors"
)

const (
	// NotLeaderErr indicates the non-leader member received the requests which should be received by leader.
	// Note: keep the same as the ones defined on the server side, because the client side checks if an error message
	// contains this string to judge whether the leader is changed.
	NotLeaderErr = "is not leader"
	// MismatchLeaderErr indicates the non-leader member received the requests which should be received by leader.
	// Note: keep the same as the ones defined on the server side, because the client side checks if an error message
	// contains this string to judge whether the leader is changed.
	MismatchLeaderErr = "mismatch leader id"
	// NotServedErr indicates an tso node/pod received the requests for the keyspace groups which are not served by it.
	// Note: keep the same as the ones defined on the server side, because the client side checks if an error message
	// contains this string to judge whether the leader is changed.
	NotServedErr    = "is not served"
	RetryTimeoutErr = "retry timeout"
)

// client errors
var (
	ErrClientGetProtoClient     = errors.Normalize("failed to get proto client", errors.RFCCodeText("TM:client:ErrClientGetProtoClient"))
	ErrClientCreateTSOStream    = errors.Normalize("create TSO stream failed, %s", errors.RFCCodeText("TM:client:ErrClientCreateTSOStream"))
	ErrClientTSOStreamClosed    = errors.Normalize("encountered TSO stream being closed unexpectedly", errors.RFCCodeText("TM:client:ErrClientTSOStreamClosed"))
	ErrClientGetTSOTimeout      = errors.Normalize("get TSO timeout", errors.RFCCodeText("TM:client:ErrClientGetTSOTimeout"))
	ErrClientGetTSO             = errors.Normalize("get TSO failed, %v", errors.RFCCodeText("TM:client:ErrClientGetTSO"))
	ErrClientGetLeader          = errors.Normalize("get leader from %v error", errors.RFCCodeText("TM:client:ErrClientGetLeader"))
	ErrClientGetMember          = errors.Normalize("get member failed", errors.RFCCodeText("TM:client:ErrClientGetMember"))
	ErrClientGetClusterInfo     = errors.Normalize("get cluster info failed", errors.RFCCodeText("TM:client:ErrClientGetClusterInfo"))
	ErrClientUpdateMember       = errors.Normalize("update member failed, %v", errors.RFCCodeText("TM:client:ErrUpdateMember"))
	ErrClientProtoUnmarshal     = errors.Normalize("failed to unmarshal proto", errors.RFCCodeText("TM:proto:ErrClientProtoUnmarshal"))
	ErrClientGetMultiResponse   = errors.Normalize("get invalid value response %v, must only one", errors.RFCCodeText("TM:client:ErrClientGetMultiResponse"))
	ErrClientGetServingEndpoint = errors.Normalize("get serving endpoint failed", errors.RFCCodeText("TM:client:ErrClientGetServingEndpoint"))
)

// grpcutil errors
var (
	ErrSecurityConfig = errors.Normalize("security config error: %s", errors.RFCCodeText("TM:grpcutil:ErrSecurityConfig"))
)

// The third-party project error.
// url errors
var (
	ErrURLParse = errors.Normalize("parse url error", errors.RFCCodeText("TM:url:ErrURLParse"))
)

// grpc errors
var (
	ErrGRPCDial      = errors.Normalize("dial error", errors.RFCCodeText("TM:grpc:ErrGRPCDial"))
	ErrCloseGRPCConn = errors.Normalize("close gRPC connection failed", errors.RFCCodeText("TM:grpc:ErrCloseGRPCConn"))
)

// etcd errors
var (
	ErrEtcdTLSConfig = errors.Normalize("etcd TLS config error", errors.RFCCodeText("TM:etcd:ErrEtcdTLSConfig"))
)

// crypto
var (
	ErrCryptoX509KeyPair        = errors.Normalize("x509 keypair error", errors.RFCCodeText("TM:crypto:ErrCryptoX509KeyPair"))
	ErrCryptoAppendCertsFromPEM = errors.Normalize("cert pool append certs error", errors.RFCCodeText("TM:crypto:ErrCryptoAppendCertsFromPEM"))
)

// resource group errors
var (
	ErrClientListResourceGroup              = errors.Normalize("get all resource group failed, %v", errors.RFCCodeText("TM:client:ErrClientListResourceGroup"))
	ErrClientResourceGroupConfigUnavailable = errors.Normalize("resource group config is unavailable, %v", errors.RFCCodeText("TM:client:ErrClientResourceGroupConfigUnavailable"))
	ErrClientResourceGroupThrottled         = errors.Normalize("exceeded resource group quota limitation", errors.RFCCodeText("TM:client:ErrClientResourceGroupThrottled"))
)

type ErrClientGetResourceGroup struct {
	ResourceGroupName string
	Cause             string
}

func (e *ErrClientGetResourceGroup) Error() string {
	return fmt.Sprintf("get resource group %v failed, %v", e.ResourceGroupName, e.Cause)
}
