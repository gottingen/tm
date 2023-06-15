// Copyright 2020 PingCAP, Inc.
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

import "github.com/pingcap/errors"

const (
	// NotLeaderErr indicates the non-leader member received the requests which should be received by leader.
	// Note: keep the same as the ones defined on the client side, because the client side checks if an error message
	// contains this string to judge whether the leader is changed.
	NotLeaderErr = "is not leader"
	// MismatchLeaderErr indicates the non-leader member received the requests which should be received by leader.
	// Note: keep the same as the ones defined on the client side, because the client side checks if an error message
	// contains this string to judge whether the leader is changed.
	MismatchLeaderErr = "mismatch leader id"
	// NotServedErr indicates an tso node/pod received the requests for the keyspace groups which are not served by it.
	// Note: keep the same as the ones defined on the client side, because the client side checks if an error message
	// contains this string to judge whether the leader is changed.
	NotServedErr = "is not served"
)

// common error in multiple packages
var (
	ErrGetSourceStore      = errors.Normalize("failed to get the source store", errors.RFCCodeText("TM:common:ErrGetSourceStore"))
	ErrIncorrectSystemTime = errors.Normalize("incorrect system time", errors.RFCCodeText("TM:common:ErrIncorrectSystemTime"))
)

// tso errors
var (
	ErrSetLocalTSOConfig                = errors.Normalize("set local tso config failed, %s", errors.RFCCodeText("TM:tso:ErrSetLocalTSOConfig"))
	ErrGetAllocator                     = errors.Normalize("get allocator failed, %s", errors.RFCCodeText("TM:tso:ErrGetAllocator"))
	ErrGetLocalAllocator                = errors.Normalize("get local allocator failed, %s", errors.RFCCodeText("TM:tso:ErrGetLocalAllocator"))
	ErrSyncMaxTS                        = errors.Normalize("sync max ts failed, %s", errors.RFCCodeText("TM:tso:ErrSyncMaxTS"))
	ErrResetUserTimestamp               = errors.Normalize("reset user timestamp failed, %s", errors.RFCCodeText("TM:tso:ErrResetUserTimestamp"))
	ErrGenerateTimestamp                = errors.Normalize("generate timestamp failed, %s", errors.RFCCodeText("TM:tso:ErrGenerateTimestamp"))
	ErrLogicOverflow                    = errors.Normalize("logic part overflow", errors.RFCCodeText("TM:tso:ErrLogicOverflow"))
	ErrProxyTSOTimeout                  = errors.Normalize("proxy tso timeout", errors.RFCCodeText("TM:tso:ErrProxyTSOTimeout"))
	ErrKeyspaceGroupIDInvalid           = errors.Normalize("the keyspace group id is invalid, %s", errors.RFCCodeText("TM:tso:ErrKeyspaceGroupIDInvalid"))
	ErrGetAllocatorManager              = errors.Normalize("get allocator manager failed, %s", errors.RFCCodeText("TM:tso:ErrGetAllocatorManager"))
	ErrLoadKeyspaceGroupsTimeout        = errors.Normalize("load keyspace groups timeout", errors.RFCCodeText("TM:tso:ErrLoadKeyspaceGroupsTimeout"))
	ErrLoadKeyspaceGroupsTerminated     = errors.Normalize("load keyspace groups terminated", errors.RFCCodeText("TM:tso:ErrLoadKeyspaceGroupsTerminated"))
	ErrLoadKeyspaceGroupsRetryExhausted = errors.Normalize("load keyspace groups retry exhausted, %s", errors.RFCCodeText("TM:tso:ErrLoadKeyspaceGroupsRetryExhausted"))
	ErrKeyspaceGroupNotInitialized      = errors.Normalize("the keyspace group %d isn't initialized", errors.RFCCodeText("TM:tso:ErrKeyspaceGroupNotInitialized"))
	ErrKeyspaceNotAssigned              = errors.Normalize("the keyspace %d isn't assigned to any keyspace group", errors.RFCCodeText("TM:tso:ErrKeyspaceNotAssigned"))
)

// member errors
var (
	ErrEtcdLeaderNotFound = errors.Normalize("etcd leader not found", errors.RFCCodeText("TM:member:ErrEtcdLeaderNotFound"))
	ErrMarshalLeader      = errors.Normalize("marshal leader failed", errors.RFCCodeText("TM:member:ErrMarshalLeader"))
	ErrPreCheckCampaign   = errors.Normalize("pre-check campaign failed", errors.RFCCodeText("TM:member:ErrPreCheckCampaign"))
)

// core errors
var (
	ErrWrongRangeKeys         = errors.Normalize("wrong range keys", errors.RFCCodeText("TM:core:ErrWrongRangeKeys"))
	ErrStoreNotFound          = errors.Normalize("store %v not found", errors.RFCCodeText("TM:core:ErrStoreNotFound"))
	ErrPauseLeaderTransfer    = errors.Normalize("store %v is paused for leader transfer", errors.RFCCodeText("TM:core:ErrPauseLeaderTransfer"))
	ErrStoreRemoved           = errors.Normalize("store %v has been removed", errors.RFCCodeText("TM:core:ErrStoreRemoved"))
	ErrStoreDestroyed         = errors.Normalize("store %v has been physically destroyed", errors.RFCCodeText("TM:core:ErrStoreDestroyed"))
	ErrStoreUnhealthy         = errors.Normalize("store %v is unhealthy", errors.RFCCodeText("TM:core:ErrStoreUnhealthy"))
	ErrStoreServing           = errors.Normalize("store %v has been serving", errors.RFCCodeText("TM:core:ErrStoreServing"))
	ErrSlowStoreEvicted       = errors.Normalize("store %v is evicted as a slow store", errors.RFCCodeText("TM:core:ErrSlowStoreEvicted"))
	ErrSlowTrendEvicted       = errors.Normalize("store %v is evicted as a slow store by trend", errors.RFCCodeText("TM:core:ErrSlowTrendEvicted"))
	ErrStoresNotEnough        = errors.Normalize("can not remove store %v since the number of up stores would be %v while need %v", errors.RFCCodeText("TM:core:ErrStoresNotEnough"))
	ErrNoStoreForRegionLeader = errors.Normalize("can not remove store %d since there are no extra up store to store the leader", errors.RFCCodeText("TM:core:ErrNoStoreForRegionLeader"))
)

// client errors
var (
	ErrClientCreateTSOStream = errors.Normalize("create TSO stream failed, %s", errors.RFCCodeText("TM:client:ErrClientCreateTSOStream"))
	ErrClientGetTSOTimeout   = errors.Normalize("get TSO timeout", errors.RFCCodeText("TM:client:ErrClientGetTSOTimeout"))
	ErrClientGetTSO          = errors.Normalize("get TSO failed, %v", errors.RFCCodeText("TM:client:ErrClientGetTSO"))
	ErrClientGetLeader       = errors.Normalize("get leader from %v error", errors.RFCCodeText("TM:client:ErrClientGetLeader"))
	ErrClientGetMember       = errors.Normalize("get member failed", errors.RFCCodeText("TM:client:ErrClientGetMember"))
)

// schedule errors
var (
	ErrUnexpectedOperatorStatus = errors.Normalize("operator with unexpected status", errors.RFCCodeText("TM:schedule:ErrUnexpectedOperatorStatus"))
	ErrUnknownOperatorStep      = errors.Normalize("unknown operator step found", errors.RFCCodeText("TM:schedule:ErrUnknownOperatorStep"))
	ErrMergeOperator            = errors.Normalize("merge operator error, %s", errors.RFCCodeText("TM:schedule:ErrMergeOperator"))
	ErrCreateOperator           = errors.Normalize("unable to create operator, %s", errors.RFCCodeText("TM:schedule:ErrCreateOperator"))
)

// scheduler errors
var (
	ErrSchedulerExisted                 = errors.Normalize("scheduler existed", errors.RFCCodeText("TM:scheduler:ErrSchedulerExisted"))
	ErrSchedulerDuplicated              = errors.Normalize("scheduler duplicated", errors.RFCCodeText("TM:scheduler:ErrSchedulerDuplicated"))
	ErrSchedulerNotFound                = errors.Normalize("scheduler not found", errors.RFCCodeText("TM:scheduler:ErrSchedulerNotFound"))
	ErrScheduleConfigNotExist           = errors.Normalize("the config does not exist", errors.RFCCodeText("TM:scheduler:ErrScheduleConfigNotExist"))
	ErrSchedulerConfig                  = errors.Normalize("wrong scheduler config %s", errors.RFCCodeText("TM:scheduler:ErrSchedulerConfig"))
	ErrCacheOverflow                    = errors.Normalize("cache overflow", errors.RFCCodeText("TM:scheduler:ErrCacheOverflow"))
	ErrInternalGrowth                   = errors.Normalize("unknown interval growth type error", errors.RFCCodeText("TM:scheduler:ErrInternalGrowth"))
	ErrSchedulerCreateFuncNotRegistered = errors.Normalize("create func of %v is not registered", errors.RFCCodeText("TM:scheduler:ErrSchedulerCreateFuncNotRegistered"))
	ErrSchedulerTiKVSplitDisabled       = errors.Normalize("tikv split region disabled", errors.RFCCodeText("TM:scheduler:ErrSchedulerTiKVSplitDisabled"))
)

// checker errors
var (
	ErrCheckerNotFound   = errors.Normalize("checker not found", errors.RFCCodeText("TM:checker:ErrCheckerNotFound"))
	ErrCheckerMergeAgain = errors.Normalize("region will be merged again, %s", errors.RFCCodeText("TM:checker:ErrCheckerMergeAgain"))
)

// diagnostic errors
var (
	ErrDiagnosticDisabled     = errors.Normalize("diagnostic is disabled", errors.RFCCodeText("TM:diagnostic:ErrDiagnosticDisabled"))
	ErrSchedulerUndiagnosable = errors.Normalize("%v hasn't supported diagnostic", errors.RFCCodeText("TM:diagnostic:ErrSchedulerUndiagnosable"))
	ErrNoDiagnosticResult     = errors.Normalize("%v has no diagnostic result", errors.RFCCodeText("TM:diagnostic:ErrNoDiagnosticResult"))
	ErrDiagnosticLoadPlan     = errors.Normalize("load plan failed", errors.RFCCodeText("TM:diagnostic:ErrDiagnosticLoadPlan"))
)

// placement errors
var (
	ErrRuleContent   = errors.Normalize("invalid rule content, %s", errors.RFCCodeText("TM:placement:ErrRuleContent"))
	ErrLoadRule      = errors.Normalize("load rule failed", errors.RFCCodeText("TM:placement:ErrLoadRule"))
	ErrLoadRuleGroup = errors.Normalize("load rule group failed", errors.RFCCodeText("TM:placement:ErrLoadRuleGroup"))
	ErrBuildRuleList = errors.Normalize("build rule list failed, %s", errors.RFCCodeText("TM:placement:ErrBuildRuleList"))
)

// region label errors
var (
	ErrRegionRuleContent  = errors.Normalize("invalid region rule content, %s", errors.RFCCodeText("TM:region:ErrRegionRuleContent"))
	ErrRegionRuleNotFound = errors.Normalize("region label rule not found for id %s", errors.RFCCodeText("TM:region:ErrRegionRuleNotFound"))
)

// cluster errors
var (
	ErrNotBootstrapped = errors.Normalize("TiKV cluster not bootstrapped, please start TiKV first", errors.RFCCodeText("TM:cluster:ErrNotBootstrapped"))
	ErrStoreIsUp       = errors.Normalize("store is still up, please remove store gracefully", errors.RFCCodeText("TM:cluster:ErrStoreIsUp"))
	ErrInvalidStoreID  = errors.Normalize("invalid store id %d, not found", errors.RFCCodeText("TM:cluster:ErrInvalidStoreID"))
)

// versioninfo errors
var (
	ErrFeatureNotExisted = errors.Normalize("feature not existed", errors.RFCCodeText("TM:versioninfo:ErrFeatureNotExisted"))
)

// autoscaling errors
var (
	ErrUnsupportedMetricsType   = errors.Normalize("unsupported metrics type %v", errors.RFCCodeText("TM:autoscaling:ErrUnsupportedMetricsType"))
	ErrUnsupportedComponentType = errors.Normalize("unsupported component type %v", errors.RFCCodeText("TM:autoscaling:ErrUnsupportedComponentType"))
	ErrUnexpectedType           = errors.Normalize("unexpected type %v", errors.RFCCodeText("TM:autoscaling:ErrUnexpectedType"))
	ErrTypeConversion           = errors.Normalize("type conversion error", errors.RFCCodeText("TM:autoscaling:ErrTypeConversion"))
	ErrEmptyMetricsResponse     = errors.Normalize("metrics response from Prometheus is empty", errors.RFCCodeText("TM:autoscaling:ErrEmptyMetricsResponse"))
	ErrEmptyMetricsResult       = errors.Normalize("result from Prometheus is empty, %s", errors.RFCCodeText("TM:autoscaling:ErrEmptyMetricsResult"))
)

// apiutil errors
var (
	ErrRedirect       = errors.Normalize("redirect failed", errors.RFCCodeText("TM:apiutil:ErrRedirect"))
	ErrOptionNotExist = errors.Normalize("the option %s does not exist", errors.RFCCodeText("TM:apiutil:ErrOptionNotExist"))
)

// grpcutil errors
var (
	ErrSecurityConfig = errors.Normalize("security config error: %s", errors.RFCCodeText("TM:grpcutil:ErrSecurityConfig"))
)

// server errors
var (
	ErrServiceRegistered     = errors.Normalize("service with path [%s] already registered", errors.RFCCodeText("TM:server:ErrServiceRegistered"))
	ErrAPIInformationInvalid = errors.Normalize("invalid api information, group %s version %s", errors.RFCCodeText("TM:server:ErrAPIInformationInvalid"))
	ErrClientURLEmpty        = errors.Normalize("client url empty", errors.RFCCodeText("TM:server:ErrClientEmpty"))
	ErrLeaderNil             = errors.Normalize("leader is nil", errors.RFCCodeText("TM:server:ErrLeaderNil"))
	ErrCancelStartEtcd       = errors.Normalize("etcd start canceled", errors.RFCCodeText("TM:server:ErrCancelStartEtcd"))
	ErrConfigItem            = errors.Normalize("cannot set invalid configuration", errors.RFCCodeText("TM:server:ErrConfiguration"))
	ErrServerNotStarted      = errors.Normalize("server not started", errors.RFCCodeText("TM:server:ErrServerNotStarted"))
)

// logutil errors
var (
	ErrInitFileLog = errors.Normalize("init file log error, %s", errors.RFCCodeText("TM:logutil:ErrInitFileLog"))
)

// typeutil errors
var (
	ErrBytesToUint64 = errors.Normalize("invalid data, must 8 bytes, but %d", errors.RFCCodeText("TM:typeutil:ErrBytesToUint64"))
)

// The third-party project error.
// url errors
var (
	ErrURLParse      = errors.Normalize("parse url error", errors.RFCCodeText("TM:url:ErrURLParse"))
	ErrQueryUnescape = errors.Normalize("inverse transformation of QueryEscape error", errors.RFCCodeText("TM:url:ErrQueryUnescape"))
)

// grpc errors
var (
	ErrGRPCDial         = errors.Normalize("dial error", errors.RFCCodeText("TM:grpc:ErrGRPCDial"))
	ErrCloseGRPCConn    = errors.Normalize("close gRPC connection failed", errors.RFCCodeText("TM:grpc:ErrCloseGRPCConn"))
	ErrGRPCSend         = errors.Normalize("send request error", errors.RFCCodeText("TM:grpc:ErrGRPCSend"))
	ErrGRPCRecv         = errors.Normalize("receive response error", errors.RFCCodeText("TM:grpc:ErrGRPCRecv"))
	ErrGRPCCloseSend    = errors.Normalize("close send error", errors.RFCCodeText("TM:grpc:ErrGRPCCloseSend"))
	ErrGRPCCreateStream = errors.Normalize("create stream error", errors.RFCCodeText("TM:grpc:ErrGRPCCreateStream"))
)

// proto errors
var (
	ErrProtoUnmarshal = errors.Normalize("failed to unmarshal proto", errors.RFCCodeText("TM:proto:ErrProtoUnmarshal"))
	ErrProtoMarshal   = errors.Normalize("failed to marshal proto", errors.RFCCodeText("TM:proto:ErrProtoMarshal"))
)

// etcd errors
var (
	ErrNewEtcdClient     = errors.Normalize("new etcd client failed", errors.RFCCodeText("TM:etcd:ErrNewEtcdClient"))
	ErrStartEtcd         = errors.Normalize("start etcd failed", errors.RFCCodeText("TM:etcd:ErrStartEtcd"))
	ErrEtcdURLMap        = errors.Normalize("etcd url map error", errors.RFCCodeText("TM:etcd:ErrEtcdURLMap"))
	ErrEtcdGrantLease    = errors.Normalize("etcd lease failed", errors.RFCCodeText("TM:etcd:ErrEtcdGrantLease"))
	ErrEtcdTxnInternal   = errors.Normalize("internal etcd transaction error occurred", errors.RFCCodeText("TM:etcd:ErrEtcdTxnInternal"))
	ErrEtcdTxnConflict   = errors.Normalize("etcd transaction failed, conflicted and rolled back", errors.RFCCodeText("TM:etcd:ErrEtcdTxnConflict"))
	ErrEtcdKVPut         = errors.Normalize("etcd KV put failed", errors.RFCCodeText("TM:etcd:ErrEtcdKVPut"))
	ErrEtcdKVDelete      = errors.Normalize("etcd KV delete failed", errors.RFCCodeText("TM:etcd:ErrEtcdKVDelete"))
	ErrEtcdKVGet         = errors.Normalize("etcd KV get failed", errors.RFCCodeText("TM:etcd:ErrEtcdKVGet"))
	ErrEtcdKVGetResponse = errors.Normalize("etcd invalid get value response %v, must only one", errors.RFCCodeText("TM:etcd:ErrEtcdKVGetResponse"))
	ErrEtcdGetCluster    = errors.Normalize("etcd get cluster from remote peer failed", errors.RFCCodeText("TM:etcd:ErrEtcdGetCluster"))
	ErrEtcdMoveLeader    = errors.Normalize("etcd move leader error", errors.RFCCodeText("TM:etcd:ErrEtcdMoveLeader"))
	ErrEtcdTLSConfig     = errors.Normalize("etcd TLS config error", errors.RFCCodeText("TM:etcd:ErrEtcdTLSConfig"))
	ErrEtcdWatcherCancel = errors.Normalize("watcher canceled", errors.RFCCodeText("TM:etcd:ErrEtcdWatcherCancel"))
	ErrCloseEtcdClient   = errors.Normalize("close etcd client failed", errors.RFCCodeText("TM:etcd:ErrCloseEtcdClient"))
	ErrEtcdMemberList    = errors.Normalize("etcd member list failed", errors.RFCCodeText("TM:etcd:ErrEtcdMemberList"))
	ErrEtcdMemberRemove  = errors.Normalize("etcd remove member failed", errors.RFCCodeText("TM:etcd:ErrEtcdMemberRemove"))
)

// dashboard errors
var (
	ErrDashboardStart = errors.Normalize("start dashboard failed", errors.RFCCodeText("TM:dashboard:ErrDashboardStart"))
	ErrDashboardStop  = errors.Normalize("stop dashboard failed", errors.RFCCodeText("TM:dashboard:ErrDashboardStop"))
)

// strconv errors
var (
	ErrStrconvParseBool  = errors.Normalize("parse bool error", errors.RFCCodeText("TM:strconv:ErrStrconvParseBool"))
	ErrStrconvParseInt   = errors.Normalize("parse int error", errors.RFCCodeText("TM:strconv:ErrStrconvParseInt"))
	ErrStrconvParseUint  = errors.Normalize("parse uint error", errors.RFCCodeText("TM:strconv:ErrStrconvParseUint"))
	ErrStrconvParseFloat = errors.Normalize("parse float error", errors.RFCCodeText("TM:strconv:ErrStrconvParseFloat"))
)

// prometheus errors
var (
	ErrPrometheusPushMetrics  = errors.Normalize("push metrics to gateway failed", errors.RFCCodeText("TM:prometheus:ErrPrometheusPushMetrics"))
	ErrPrometheusCreateClient = errors.Normalize("create client error", errors.RFCCodeText("TM:prometheus:ErrPrometheusCreateClient"))
	ErrPrometheusQuery        = errors.Normalize("query error", errors.RFCCodeText("TM:prometheus:ErrPrometheusQuery"))
)

// http errors
var (
	ErrSendRequest    = errors.Normalize("send HTTP request failed", errors.RFCCodeText("TM:http:ErrSendRequest"))
	ErrWriteHTTPBody  = errors.Normalize("write HTTP body failed", errors.RFCCodeText("TM:http:ErrWriteHTTPBody"))
	ErrNewHTTPRequest = errors.Normalize("new HTTP request failed", errors.RFCCodeText("TM:http:ErrNewHTTPRequest"))
)

// ioutil error
var (
	ErrIORead = errors.Normalize("IO read error", errors.RFCCodeText("TM:ioutil:ErrIORead"))
)

// os error
var (
	ErrOSOpen = errors.Normalize("open error", errors.RFCCodeText("TM:os:ErrOSOpen"))
)

// dir error
var (
	ErrReadDirName = errors.Normalize("read dir name error", errors.RFCCodeText("TM:dir:ErrReadDirName"))
)

// netstat error
var (
	ErrNetstatTCPSocks = errors.Normalize("TCP socks error", errors.RFCCodeText("TM:netstat:ErrNetstatTCPSocks"))
)

// hex error
var (
	ErrHexDecodingString = errors.Normalize("decode string %s error", errors.RFCCodeText("TM:hex:ErrHexDecodingString"))
)

// filepath errors
var (
	ErrFilePathAbs = errors.Normalize("failed to convert a path to absolute path", errors.RFCCodeText("TM:filepath:ErrFilePathAbs"))
)

// plugin errors
var (
	ErrLoadPlugin       = errors.Normalize("failed to load plugin", errors.RFCCodeText("TM:plugin:ErrLoadPlugin"))
	ErrLookupPluginFunc = errors.Normalize("failed to lookup plugin function", errors.RFCCodeText("TM:plugin:ErrLookupPluginFunc"))
)

// json errors
var (
	ErrJSONMarshal   = errors.Normalize("failed to marshal json", errors.RFCCodeText("TM:json:ErrJSONMarshal"))
	ErrJSONUnmarshal = errors.Normalize("failed to unmarshal json", errors.RFCCodeText("TM:json:ErrJSONUnmarshal"))
)

// leveldb errors
var (
	ErrLevelDBClose = errors.Normalize("close leveldb error", errors.RFCCodeText("TM:leveldb:ErrLevelDBClose"))
	ErrLevelDBWrite = errors.Normalize("leveldb write error", errors.RFCCodeText("TM:leveldb:ErrLevelDBWrite"))
	ErrLevelDBOpen  = errors.Normalize("leveldb open file error", errors.RFCCodeText("TM:leveldb:ErrLevelDBOpen"))
)

// semver
var (
	ErrSemverNewVersion = errors.Normalize("new version error", errors.RFCCodeText("TM:semver:ErrSemverNewVersion"))
)

// log
var (
	ErrInitLogger = errors.Normalize("init logger error", errors.RFCCodeText("TM:log:ErrInitLogger"))
)

// encryption
var (
	ErrEncryptionInvalidMethod      = errors.Normalize("invalid encryption method", errors.RFCCodeText("TM:encryption:ErrEncryptionInvalidMethod"))
	ErrEncryptionInvalidConfig      = errors.Normalize("invalid config", errors.RFCCodeText("TM:encryption:ErrEncryptionInvalidConfig"))
	ErrEncryptionGenerateIV         = errors.Normalize("fail to generate iv", errors.RFCCodeText("TM:encryption:ErrEncryptionGenerateIV"))
	ErrEncryptionGCMEncrypt         = errors.Normalize("GCM encryption fail", errors.RFCCodeText("TM:encryption:ErrEncryptionGCMEncrypt"))
	ErrEncryptionGCMDecrypt         = errors.Normalize("GCM decryption fail", errors.RFCCodeText("TM:encryption:ErrEncryptionGCMDecrypt"))
	ErrEncryptionCTREncrypt         = errors.Normalize("CTR encryption fail", errors.RFCCodeText("TM:encryption:ErrEncryptionCTREncrypt"))
	ErrEncryptionCTRDecrypt         = errors.Normalize("CTR decryption fail", errors.RFCCodeText("TM:encryption:ErrEncryptionCTRDecrypt"))
	ErrEncryptionEncryptRegion      = errors.Normalize("encrypt region fail", errors.RFCCodeText("TM:encryption:ErrEncryptionEncryptRegion"))
	ErrEncryptionDecryptRegion      = errors.Normalize("decrypt region fail", errors.RFCCodeText("TM:encryption:ErrEncryptionDecryptRegion"))
	ErrEncryptionNewDataKey         = errors.Normalize("fail to generate data key", errors.RFCCodeText("TM:encryption:ErrEncryptionNewDataKey"))
	ErrEncryptionNewMasterKey       = errors.Normalize("fail to get master key", errors.RFCCodeText("TM:encryption:ErrEncryptionNewMasterKey"))
	ErrEncryptionCurrentKeyNotFound = errors.Normalize("current data key not found", errors.RFCCodeText("TM:encryption:ErrEncryptionCurrentKeyNotFound"))
	ErrEncryptionKeyNotFound        = errors.Normalize("data key not found", errors.RFCCodeText("TM:encryption:ErrEncryptionKeyNotFound"))
	ErrEncryptionKeysWatcher        = errors.Normalize("data key watcher error", errors.RFCCodeText("TM:encryption:ErrEncryptionKeysWatcher"))
	ErrEncryptionLoadKeys           = errors.Normalize("load data keys error", errors.RFCCodeText("TM:encryption:ErrEncryptionLoadKeys"))
	ErrEncryptionRotateDataKey      = errors.Normalize("failed to rotate data key", errors.RFCCodeText("TM:encryption:ErrEncryptionRotateDataKey"))
	ErrEncryptionSaveDataKeys       = errors.Normalize("failed to save data keys", errors.RFCCodeText("TM:encryption:ErrEncryptionSaveDataKeys"))
	ErrEncryptionKMS                = errors.Normalize("KMS error", errors.RFCCodeText("TM:ErrEncryptionKMS"))
)

// crypto
var (
	ErrCryptoX509KeyPair        = errors.Normalize("x509 keypair error", errors.RFCCodeText("TM:crypto:ErrCryptoX509KeyPair"))
	ErrCryptoAppendCertsFromPEM = errors.Normalize("cert pool append certs error", errors.RFCCodeText("TM:crypto:ErrCryptoAppendCertsFromPEM"))
)

// gin errors
var (
	ErrBindJSON = errors.Normalize("bind JSON error", errors.RFCCodeText("TM:gin:ErrBindJSON"))
)

// unsafe recovery errors
var (
	ErrUnsafeRecoveryIsRunning    = errors.Normalize("unsafe recovery is running", errors.RFCCodeText("TM:unsaferecovery:ErrUnsafeRecoveryIsRunning"))
	ErrUnsafeRecoveryInvalidInput = errors.Normalize("invalid input %s", errors.RFCCodeText("TM:unsaferecovery:ErrUnsafeRecoveryInvalidInput"))
)

// progress errors
var (
	ErrProgressWrongStatus = errors.Normalize("progress status is wrong", errors.RFCCodeText("TM:progress:ErrProgressWrongStatus"))
	ErrProgressNotFound    = errors.Normalize("no progress found for %s", errors.RFCCodeText("TM:progress:ErrProgressNotFound"))
)

// Resource Manager errors
var (
	ErrResourceGroupAlreadyExists = errors.Normalize("the %s resource group already exists", errors.RFCCodeText("TM:resourcemanager:ErrResourceGroupAlreadyExists"))
	ErrResourceGroupNotExists     = errors.Normalize("the %s resource group does not exist", errors.RFCCodeText("TM:resourcemanager:ErrGroupNotExists"))
	ErrDeleteReservedGroup        = errors.Normalize("cannot delete reserved group", errors.RFCCodeText("TM:resourcemanager:ErrDeleteReservedGroup"))
	ErrInvalidGroup               = errors.Normalize("invalid group settings, please check the group name, priority and the number of resources", errors.RFCCodeText("TM:resourcemanager:ErrInvalidGroup"))
)
