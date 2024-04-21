/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentbit_fluent_io_v1alpha2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &FluentbitFluentIoFilterV1Alpha2Manifest{}
)

func NewFluentbitFluentIoFilterV1Alpha2Manifest() datasource.DataSource {
	return &FluentbitFluentIoFilterV1Alpha2Manifest{}
}

type FluentbitFluentIoFilterV1Alpha2Manifest struct{}

type FluentbitFluentIoFilterV1Alpha2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Filters *[]struct {
			Aws *struct {
				AccountID       *bool   `tfsdk:"account_id" json:"accountID,omitempty"`
				Alias           *string `tfsdk:"alias" json:"alias,omitempty"`
				AmiID           *bool   `tfsdk:"ami_id" json:"amiID,omitempty"`
				Az              *bool   `tfsdk:"az" json:"az,omitempty"`
				Ec2InstanceID   *bool   `tfsdk:"ec2_instance_id" json:"ec2InstanceID,omitempty"`
				Ec2InstanceType *bool   `tfsdk:"ec2_instance_type" json:"ec2InstanceType,omitempty"`
				HostName        *bool   `tfsdk:"host_name" json:"hostName,omitempty"`
				ImdsVersion     *string `tfsdk:"imds_version" json:"imdsVersion,omitempty"`
				PrivateIP       *bool   `tfsdk:"private_ip" json:"privateIP,omitempty"`
				RetryLimit      *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				VpcID           *bool   `tfsdk:"vpc_id" json:"vpcID,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			CustomPlugin *struct {
				Config *string `tfsdk:"config" json:"config,omitempty"`
			} `tfsdk:"custom_plugin" json:"customPlugin,omitempty"`
			Grep *struct {
				Alias      *string `tfsdk:"alias" json:"alias,omitempty"`
				Exclude    *string `tfsdk:"exclude" json:"exclude,omitempty"`
				Regex      *string `tfsdk:"regex" json:"regex,omitempty"`
				RetryLimit *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
			} `tfsdk:"grep" json:"grep,omitempty"`
			Kubernetes *struct {
				Alias                   *string `tfsdk:"alias" json:"alias,omitempty"`
				Annotations             *bool   `tfsdk:"annotations" json:"annotations,omitempty"`
				BufferSize              *string `tfsdk:"buffer_size" json:"bufferSize,omitempty"`
				CacheUseDockerId        *bool   `tfsdk:"cache_use_docker_id" json:"cacheUseDockerId,omitempty"`
				DnsRetries              *int64  `tfsdk:"dns_retries" json:"dnsRetries,omitempty"`
				DnsWaitTime             *int64  `tfsdk:"dns_wait_time" json:"dnsWaitTime,omitempty"`
				DummyMeta               *bool   `tfsdk:"dummy_meta" json:"dummyMeta,omitempty"`
				K8sLoggingExclude       *bool   `tfsdk:"k8s_logging_exclude" json:"k8sLoggingExclude,omitempty"`
				K8sLoggingParser        *bool   `tfsdk:"k8s_logging_parser" json:"k8sLoggingParser,omitempty"`
				KeepLog                 *bool   `tfsdk:"keep_log" json:"keepLog,omitempty"`
				KubeCAFile              *string `tfsdk:"kube_ca_file" json:"kubeCAFile,omitempty"`
				KubeCAPath              *string `tfsdk:"kube_ca_path" json:"kubeCAPath,omitempty"`
				KubeMetaCacheTTL        *string `tfsdk:"kube_meta_cache_ttl" json:"kubeMetaCacheTTL,omitempty"`
				KubeMetaPreloadCacheDir *string `tfsdk:"kube_meta_preload_cache_dir" json:"kubeMetaPreloadCacheDir,omitempty"`
				KubeTagPrefix           *string `tfsdk:"kube_tag_prefix" json:"kubeTagPrefix,omitempty"`
				KubeTokenFile           *string `tfsdk:"kube_token_file" json:"kubeTokenFile,omitempty"`
				KubeTokenTTL            *string `tfsdk:"kube_token_ttl" json:"kubeTokenTTL,omitempty"`
				KubeURL                 *string `tfsdk:"kube_url" json:"kubeURL,omitempty"`
				KubeletHost             *string `tfsdk:"kubelet_host" json:"kubeletHost,omitempty"`
				KubeletPort             *int64  `tfsdk:"kubelet_port" json:"kubeletPort,omitempty"`
				Labels                  *bool   `tfsdk:"labels" json:"labels,omitempty"`
				MergeLog                *bool   `tfsdk:"merge_log" json:"mergeLog,omitempty"`
				MergeLogKey             *string `tfsdk:"merge_log_key" json:"mergeLogKey,omitempty"`
				MergeLogTrim            *bool   `tfsdk:"merge_log_trim" json:"mergeLogTrim,omitempty"`
				MergeParser             *string `tfsdk:"merge_parser" json:"mergeParser,omitempty"`
				RegexParser             *string `tfsdk:"regex_parser" json:"regexParser,omitempty"`
				RetryLimit              *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				TlsDebug                *int64  `tfsdk:"tls_debug" json:"tlsDebug,omitempty"`
				TlsVerify               *bool   `tfsdk:"tls_verify" json:"tlsVerify,omitempty"`
				UseJournal              *bool   `tfsdk:"use_journal" json:"useJournal,omitempty"`
				UseKubelet              *bool   `tfsdk:"use_kubelet" json:"useKubelet,omitempty"`
			} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			Lua *struct {
				Alias         *string `tfsdk:"alias" json:"alias,omitempty"`
				Call          *string `tfsdk:"call" json:"call,omitempty"`
				Code          *string `tfsdk:"code" json:"code,omitempty"`
				ProtectedMode *bool   `tfsdk:"protected_mode" json:"protectedMode,omitempty"`
				RetryLimit    *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				Script        *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"script" json:"script,omitempty"`
				TimeAsTable *bool     `tfsdk:"time_as_table" json:"timeAsTable,omitempty"`
				TypeIntKey  *[]string `tfsdk:"type_int_key" json:"typeIntKey,omitempty"`
			} `tfsdk:"lua" json:"lua,omitempty"`
			Modify *struct {
				Alias      *string `tfsdk:"alias" json:"alias,omitempty"`
				Conditions *[]struct {
					AKeyMatches                         *string            `tfsdk:"a_key_matches" json:"aKeyMatches,omitempty"`
					KeyDoesNotExist                     *map[string]string `tfsdk:"key_does_not_exist" json:"keyDoesNotExist,omitempty"`
					KeyExists                           *string            `tfsdk:"key_exists" json:"keyExists,omitempty"`
					KeyValueDoesNotEqual                *map[string]string `tfsdk:"key_value_does_not_equal" json:"keyValueDoesNotEqual,omitempty"`
					KeyValueDoesNotMatch                *map[string]string `tfsdk:"key_value_does_not_match" json:"keyValueDoesNotMatch,omitempty"`
					KeyValueEquals                      *map[string]string `tfsdk:"key_value_equals" json:"keyValueEquals,omitempty"`
					KeyValueMatches                     *map[string]string `tfsdk:"key_value_matches" json:"keyValueMatches,omitempty"`
					MatchingKeysDoNotHaveMatchingValues *map[string]string `tfsdk:"matching_keys_do_not_have_matching_values" json:"matchingKeysDoNotHaveMatchingValues,omitempty"`
					MatchingKeysHaveMatchingValues      *map[string]string `tfsdk:"matching_keys_have_matching_values" json:"matchingKeysHaveMatchingValues,omitempty"`
					NoKeyMatches                        *string            `tfsdk:"no_key_matches" json:"noKeyMatches,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				RetryLimit *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				Rules      *[]struct {
					Add            *map[string]string `tfsdk:"add" json:"add,omitempty"`
					Copy           *map[string]string `tfsdk:"copy" json:"copy,omitempty"`
					HardCopy       *map[string]string `tfsdk:"hard_copy" json:"hardCopy,omitempty"`
					HardRename     *map[string]string `tfsdk:"hard_rename" json:"hardRename,omitempty"`
					Remove         *string            `tfsdk:"remove" json:"remove,omitempty"`
					RemoveRegex    *string            `tfsdk:"remove_regex" json:"removeRegex,omitempty"`
					RemoveWildcard *string            `tfsdk:"remove_wildcard" json:"removeWildcard,omitempty"`
					Rename         *map[string]string `tfsdk:"rename" json:"rename,omitempty"`
					Set            *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
			} `tfsdk:"modify" json:"modify,omitempty"`
			Multiline *struct {
				Alias              *string `tfsdk:"alias" json:"alias,omitempty"`
				Buffer             *bool   `tfsdk:"buffer" json:"buffer,omitempty"`
				EmitterMemBufLimit *int64  `tfsdk:"emitter_mem_buf_limit" json:"emitterMemBufLimit,omitempty"`
				EmitterName        *string `tfsdk:"emitter_name" json:"emitterName,omitempty"`
				EmitterType        *string `tfsdk:"emitter_type" json:"emitterType,omitempty"`
				FlushMs            *int64  `tfsdk:"flush_ms" json:"flushMs,omitempty"`
				KeyContent         *string `tfsdk:"key_content" json:"keyContent,omitempty"`
				Mode               *string `tfsdk:"mode" json:"mode,omitempty"`
				Parser             *string `tfsdk:"parser" json:"parser,omitempty"`
				RetryLimit         *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
			} `tfsdk:"multiline" json:"multiline,omitempty"`
			Nest *struct {
				AddPrefix    *string   `tfsdk:"add_prefix" json:"addPrefix,omitempty"`
				Alias        *string   `tfsdk:"alias" json:"alias,omitempty"`
				NestUnder    *string   `tfsdk:"nest_under" json:"nestUnder,omitempty"`
				NestedUnder  *string   `tfsdk:"nested_under" json:"nestedUnder,omitempty"`
				Operation    *string   `tfsdk:"operation" json:"operation,omitempty"`
				RemovePrefix *string   `tfsdk:"remove_prefix" json:"removePrefix,omitempty"`
				RetryLimit   *string   `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				Wildcard     *[]string `tfsdk:"wildcard" json:"wildcard,omitempty"`
			} `tfsdk:"nest" json:"nest,omitempty"`
			Parser *struct {
				Alias       *string `tfsdk:"alias" json:"alias,omitempty"`
				KeyName     *string `tfsdk:"key_name" json:"keyName,omitempty"`
				Parser      *string `tfsdk:"parser" json:"parser,omitempty"`
				PreserveKey *bool   `tfsdk:"preserve_key" json:"preserveKey,omitempty"`
				ReserveData *bool   `tfsdk:"reserve_data" json:"reserveData,omitempty"`
				RetryLimit  *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				UnescapeKey *bool   `tfsdk:"unescape_key" json:"unescapeKey,omitempty"`
			} `tfsdk:"parser" json:"parser,omitempty"`
			RecordModifier *struct {
				Alias         *string   `tfsdk:"alias" json:"alias,omitempty"`
				AllowlistKeys *[]string `tfsdk:"allowlist_keys" json:"allowlistKeys,omitempty"`
				Records       *[]string `tfsdk:"records" json:"records,omitempty"`
				RemoveKeys    *[]string `tfsdk:"remove_keys" json:"removeKeys,omitempty"`
				RetryLimit    *string   `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				UuidKeys      *[]string `tfsdk:"uuid_keys" json:"uuidKeys,omitempty"`
				WhitelistKeys *[]string `tfsdk:"whitelist_keys" json:"whitelistKeys,omitempty"`
			} `tfsdk:"record_modifier" json:"recordModifier,omitempty"`
			RewriteTag *struct {
				Alias              *string   `tfsdk:"alias" json:"alias,omitempty"`
				EmitterMemBufLimit *string   `tfsdk:"emitter_mem_buf_limit" json:"emitterMemBufLimit,omitempty"`
				EmitterName        *string   `tfsdk:"emitter_name" json:"emitterName,omitempty"`
				EmitterStorageType *string   `tfsdk:"emitter_storage_type" json:"emitterStorageType,omitempty"`
				RetryLimit         *string   `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				Rules              *[]string `tfsdk:"rules" json:"rules,omitempty"`
			} `tfsdk:"rewrite_tag" json:"rewriteTag,omitempty"`
			Throttle *struct {
				Alias       *string `tfsdk:"alias" json:"alias,omitempty"`
				Interval    *string `tfsdk:"interval" json:"interval,omitempty"`
				PrintStatus *bool   `tfsdk:"print_status" json:"printStatus,omitempty"`
				Rate        *int64  `tfsdk:"rate" json:"rate,omitempty"`
				RetryLimit  *string `tfsdk:"retry_limit" json:"retryLimit,omitempty"`
				Window      *int64  `tfsdk:"window" json:"window,omitempty"`
			} `tfsdk:"throttle" json:"throttle,omitempty"`
		} `tfsdk:"filters" json:"filters,omitempty"`
		LogLevel   *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		Match      *string `tfsdk:"match" json:"match,omitempty"`
		MatchRegex *string `tfsdk:"match_regex" json:"matchRegex,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentbitFluentIoFilterV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentbit_fluent_io_filter_v1alpha2_manifest"
}

func (r *FluentbitFluentIoFilterV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Filter is the Schema for namespace level filter API",
		MarkdownDescription: "Filter is the Schema for namespace level filter API",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "FilterSpec defines the desired state of ClusterFilter",
				MarkdownDescription: "FilterSpec defines the desired state of ClusterFilter",
				Attributes: map[string]schema.Attribute{
					"filters": schema.ListNestedAttribute{
						Description:         "A set of filter plugins in order.",
						MarkdownDescription: "A set of filter plugins in order.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws": schema.SingleNestedAttribute{
									Description:         "Aws defines a Aws configuration.",
									MarkdownDescription: "Aws defines a Aws configuration.",
									Attributes: map[string]schema.Attribute{
										"account_id": schema.BoolAttribute{
											Description:         "The account ID for current EC2 instance.Default is false.",
											MarkdownDescription: "The account ID for current EC2 instance.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ami_id": schema.BoolAttribute{
											Description:         "The EC2 instance image id.Default is false.",
											MarkdownDescription: "The EC2 instance image id.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"az": schema.BoolAttribute{
											Description:         "The availability zone; for example, 'us-east-1a'. Default is true.",
											MarkdownDescription: "The availability zone; for example, 'us-east-1a'. Default is true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ec2_instance_id": schema.BoolAttribute{
											Description:         "The EC2 instance ID.Default is true.",
											MarkdownDescription: "The EC2 instance ID.Default is true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ec2_instance_type": schema.BoolAttribute{
											Description:         "The EC2 instance type.Default is false.",
											MarkdownDescription: "The EC2 instance type.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"host_name": schema.BoolAttribute{
											Description:         "The hostname for current EC2 instance.Default is false.",
											MarkdownDescription: "The hostname for current EC2 instance.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"imds_version": schema.StringAttribute{
											Description:         "Specify which version of the instance metadata service to use. Valid values are 'v1' or 'v2'.",
											MarkdownDescription: "Specify which version of the instance metadata service to use. Valid values are 'v1' or 'v2'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("v1", "v2"),
											},
										},

										"private_ip": schema.BoolAttribute{
											Description:         "The EC2 instance private ip.Default is false.",
											MarkdownDescription: "The EC2 instance private ip.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"vpc_id": schema.BoolAttribute{
											Description:         "The VPC ID for current EC2 instance.Default is false.",
											MarkdownDescription: "The VPC ID for current EC2 instance.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"custom_plugin": schema.SingleNestedAttribute{
									Description:         "CustomPlugin defines a Custom plugin configuration.",
									MarkdownDescription: "CustomPlugin defines a Custom plugin configuration.",
									Attributes: map[string]schema.Attribute{
										"config": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"grep": schema.SingleNestedAttribute{
									Description:         "Grep defines Grep Filter configuration.",
									MarkdownDescription: "Grep defines Grep Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": schema.StringAttribute{
											Description:         "Exclude records which field matches the regular expression. Value Format: FIELD REGEX",
											MarkdownDescription: "Exclude records which field matches the regular expression. Value Format: FIELD REGEX",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
											Description:         "Keep records which field matches the regular expression. Value Format: FIELD REGEX",
											MarkdownDescription: "Keep records which field matches the regular expression. Value Format: FIELD REGEX",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kubernetes": schema.SingleNestedAttribute{
									Description:         "Kubernetes defines Kubernetes Filter configuration.",
									MarkdownDescription: "Kubernetes defines Kubernetes Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"annotations": schema.BoolAttribute{
											Description:         "Include Kubernetes resource annotations in the extra metadata.",
											MarkdownDescription: "Include Kubernetes resource annotations in the extra metadata.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"buffer_size": schema.StringAttribute{
											Description:         "Set the buffer size for HTTP client when reading responses from Kubernetes API server.",
											MarkdownDescription: "Set the buffer size for HTTP client when reading responses from Kubernetes API server.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
											},
										},

										"cache_use_docker_id": schema.BoolAttribute{
											Description:         "When enabled, metadata will be fetched from K8s when docker_id is changed.",
											MarkdownDescription: "When enabled, metadata will be fetched from K8s when docker_id is changed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dns_retries": schema.Int64Attribute{
											Description:         "DNS lookup retries N times until the network start working",
											MarkdownDescription: "DNS lookup retries N times until the network start working",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dns_wait_time": schema.Int64Attribute{
											Description:         "DNS lookup interval between network status checks",
											MarkdownDescription: "DNS lookup interval between network status checks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dummy_meta": schema.BoolAttribute{
											Description:         "If set, use dummy-meta data (for test/dev purposes)",
											MarkdownDescription: "If set, use dummy-meta data (for test/dev purposes)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"k8s_logging_exclude": schema.BoolAttribute{
											Description:         "Allow Kubernetes Pods to exclude their logs from the log processor (read more about it in Kubernetes Annotations section).",
											MarkdownDescription: "Allow Kubernetes Pods to exclude their logs from the log processor (read more about it in Kubernetes Annotations section).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"k8s_logging_parser": schema.BoolAttribute{
											Description:         "Allow Kubernetes Pods to suggest a pre-defined Parser (read more about it in Kubernetes Annotations section)",
											MarkdownDescription: "Allow Kubernetes Pods to suggest a pre-defined Parser (read more about it in Kubernetes Annotations section)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keep_log": schema.BoolAttribute{
											Description:         "When Keep_Log is disabled, the log field is removed from the incoming message once it has been successfully merged (Merge_Log must be enabled as well).",
											MarkdownDescription: "When Keep_Log is disabled, the log field is removed from the incoming message once it has been successfully merged (Merge_Log must be enabled as well).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_ca_file": schema.StringAttribute{
											Description:         "CA certificate file",
											MarkdownDescription: "CA certificate file",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_ca_path": schema.StringAttribute{
											Description:         "Absolute path to scan for certificate files",
											MarkdownDescription: "Absolute path to scan for certificate files",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_meta_cache_ttl": schema.StringAttribute{
											Description:         "configurable TTL for K8s cached metadata. By default, it is set to 0 which means TTL for cache entries is disabled and cache entries are evicted at random when capacity is reached. In order to enable this option, you should set the number to a time interval. For example, set this value to 60 or 60s and cache entries which have been created more than 60s will be evicted.",
											MarkdownDescription: "configurable TTL for K8s cached metadata. By default, it is set to 0 which means TTL for cache entries is disabled and cache entries are evicted at random when capacity is reached. In order to enable this option, you should set the number to a time interval. For example, set this value to 60 or 60s and cache entries which have been created more than 60s will be evicted.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_meta_preload_cache_dir": schema.StringAttribute{
											Description:         "If set, Kubernetes meta-data can be cached/pre-loaded from files in JSON format in this directory, named as namespace-pod.meta",
											MarkdownDescription: "If set, Kubernetes meta-data can be cached/pre-loaded from files in JSON format in this directory, named as namespace-pod.meta",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_tag_prefix": schema.StringAttribute{
											Description:         "When the source records comes from Tail input plugin, this option allows to specify what's the prefix used in Tail configuration.",
											MarkdownDescription: "When the source records comes from Tail input plugin, this option allows to specify what's the prefix used in Tail configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_token_file": schema.StringAttribute{
											Description:         "Token file",
											MarkdownDescription: "Token file",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_token_ttl": schema.StringAttribute{
											Description:         "configurable 'time to live' for the K8s token. By default, it is set to 600 seconds. After this time, the token is reloaded from Kube_Token_File or the Kube_Token_Command.",
											MarkdownDescription: "configurable 'time to live' for the K8s token. By default, it is set to 600 seconds. After this time, the token is reloaded from Kube_Token_File or the Kube_Token_Command.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kube_url": schema.StringAttribute{
											Description:         "API Server end-point",
											MarkdownDescription: "API Server end-point",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kubelet_host": schema.StringAttribute{
											Description:         "kubelet host using for HTTP request, this only works when Use_Kubelet set to On.",
											MarkdownDescription: "kubelet host using for HTTP request, this only works when Use_Kubelet set to On.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kubelet_port": schema.Int64Attribute{
											Description:         "kubelet port using for HTTP request, this only works when useKubelet is set to On.",
											MarkdownDescription: "kubelet port using for HTTP request, this only works when useKubelet is set to On.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.BoolAttribute{
											Description:         "Include Kubernetes resource labels in the extra metadata.",
											MarkdownDescription: "Include Kubernetes resource labels in the extra metadata.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"merge_log": schema.BoolAttribute{
											Description:         "When enabled, it checks if the log field content is a JSON string map, if so, it append the map fields as part of the log structure.",
											MarkdownDescription: "When enabled, it checks if the log field content is a JSON string map, if so, it append the map fields as part of the log structure.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"merge_log_key": schema.StringAttribute{
											Description:         "When Merge_Log is enabled, the filter tries to assume the log field from the incoming message is a JSON string message and make a structured representation of it at the same level of the log field in the map. Now if Merge_Log_Key is set (a string name), all the new structured fields taken from the original log content are inserted under the new key.",
											MarkdownDescription: "When Merge_Log is enabled, the filter tries to assume the log field from the incoming message is a JSON string message and make a structured representation of it at the same level of the log field in the map. Now if Merge_Log_Key is set (a string name), all the new structured fields taken from the original log content are inserted under the new key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"merge_log_trim": schema.BoolAttribute{
											Description:         "When Merge_Log is enabled, trim (remove possible n or r) field values.",
											MarkdownDescription: "When Merge_Log is enabled, trim (remove possible n or r) field values.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"merge_parser": schema.StringAttribute{
											Description:         "Optional parser name to specify how to parse the data contained in the log key. Recommended use is for developers or testing only.",
											MarkdownDescription: "Optional parser name to specify how to parse the data contained in the log key. Recommended use is for developers or testing only.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex_parser": schema.StringAttribute{
											Description:         "Set an alternative Parser to process record Tag and extract pod_name, namespace_name, container_name and docker_id. The parser must be registered in a parsers file (refer to parser filter-kube-test as an example).",
											MarkdownDescription: "Set an alternative Parser to process record Tag and extract pod_name, namespace_name, container_name and docker_id. The parser must be registered in a parsers file (refer to parser filter-kube-test as an example).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"tls_debug": schema.Int64Attribute{
											Description:         "Debug level between 0 (nothing) and 4 (every detail).",
											MarkdownDescription: "Debug level between 0 (nothing) and 4 (every detail).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_verify": schema.BoolAttribute{
											Description:         "When enabled, turns on certificate validation when connecting to the Kubernetes API server.",
											MarkdownDescription: "When enabled, turns on certificate validation when connecting to the Kubernetes API server.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_journal": schema.BoolAttribute{
											Description:         "When enabled, the filter reads logs coming in Journald format.",
											MarkdownDescription: "When enabled, the filter reads logs coming in Journald format.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_kubelet": schema.BoolAttribute{
											Description:         "This is an optional feature flag to get metadata information from kubelet instead of calling Kube Server API to enhance the log. This could mitigate the Kube API heavy traffic issue for large cluster.",
											MarkdownDescription: "This is an optional feature flag to get metadata information from kubelet instead of calling Kube Server API to enhance the log. This could mitigate the Kube API heavy traffic issue for large cluster.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"lua": schema.SingleNestedAttribute{
									Description:         "Lua defines Lua Filter configuration.",
									MarkdownDescription: "Lua defines Lua Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"call": schema.StringAttribute{
											Description:         "Lua function name that will be triggered to do filtering. It's assumed that the function is declared inside the Script defined above.",
											MarkdownDescription: "Lua function name that will be triggered to do filtering. It's assumed that the function is declared inside the Script defined above.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"code": schema.StringAttribute{
											Description:         "Inline LUA code instead of loading from a path via script.",
											MarkdownDescription: "Inline LUA code instead of loading from a path via script.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protected_mode": schema.BoolAttribute{
											Description:         "If enabled, Lua script will be executed in protected mode. It prevents to crash when invalid Lua script is executed. Default is true.",
											MarkdownDescription: "If enabled, Lua script will be executed in protected mode. It prevents to crash when invalid Lua script is executed. Default is true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"script": schema.SingleNestedAttribute{
											Description:         "Path to the Lua script that will be used.",
											MarkdownDescription: "Path to the Lua script that will be used.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap or its key must be defined",
													MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"time_as_table": schema.BoolAttribute{
											Description:         "By default when the Lua script is invoked, the record timestamp is passed as a Floating number which might lead to loss precision when the data is converted back. If you desire timestamp precision enabling this option will pass the timestamp as a Lua table with keys sec for seconds since epoch and nsec for nanoseconds.",
											MarkdownDescription: "By default when the Lua script is invoked, the record timestamp is passed as a Floating number which might lead to loss precision when the data is converted back. If you desire timestamp precision enabling this option will pass the timestamp as a Lua table with keys sec for seconds since epoch and nsec for nanoseconds.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type_int_key": schema.ListAttribute{
											Description:         "If these keys are matched, the fields are converted to integer. If more than one key, delimit by space. Note that starting from Fluent Bit v1.6 integer data types are preserved and not converted to double as in previous versions.",
											MarkdownDescription: "If these keys are matched, the fields are converted to integer. If more than one key, delimit by space. Note that starting from Fluent Bit v1.6 integer data types are preserved and not converted to double as in previous versions.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"modify": schema.SingleNestedAttribute{
									Description:         "Modify defines Modify Filter configuration.",
									MarkdownDescription: "Modify defines Modify Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"conditions": schema.ListNestedAttribute{
											Description:         "All conditions have to be true for the rules to be applied.",
											MarkdownDescription: "All conditions have to be true for the rules to be applied.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"a_key_matches": schema.StringAttribute{
														Description:         "Is true if a key matches regex KEY",
														MarkdownDescription: "Is true if a key matches regex KEY",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_does_not_exist": schema.MapAttribute{
														Description:         "Is true if KEY does not exist",
														MarkdownDescription: "Is true if KEY does not exist",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_exists": schema.StringAttribute{
														Description:         "Is true if KEY exists",
														MarkdownDescription: "Is true if KEY exists",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_value_does_not_equal": schema.MapAttribute{
														Description:         "Is true if KEY exists and its value is not VALUE",
														MarkdownDescription: "Is true if KEY exists and its value is not VALUE",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_value_does_not_match": schema.MapAttribute{
														Description:         "Is true if key KEY exists and its value does not match VALUE",
														MarkdownDescription: "Is true if key KEY exists and its value does not match VALUE",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_value_equals": schema.MapAttribute{
														Description:         "Is true if KEY exists and its value is VALUE",
														MarkdownDescription: "Is true if KEY exists and its value is VALUE",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_value_matches": schema.MapAttribute{
														Description:         "Is true if key KEY exists and its value matches VALUE",
														MarkdownDescription: "Is true if key KEY exists and its value matches VALUE",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"matching_keys_do_not_have_matching_values": schema.MapAttribute{
														Description:         "Is true if all keys matching KEY have values that do not match VALUE",
														MarkdownDescription: "Is true if all keys matching KEY have values that do not match VALUE",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"matching_keys_have_matching_values": schema.MapAttribute{
														Description:         "Is true if all keys matching KEY have values that match VALUE",
														MarkdownDescription: "Is true if all keys matching KEY have values that match VALUE",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"no_key_matches": schema.StringAttribute{
														Description:         "Is true if no key matches regex KEY",
														MarkdownDescription: "Is true if no key matches regex KEY",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"rules": schema.ListNestedAttribute{
											Description:         "Rules are applied in the order they appear, with each rule operating on the result of the previous rule.",
											MarkdownDescription: "Rules are applied in the order they appear, with each rule operating on the result of the previous rule.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"add": schema.MapAttribute{
														Description:         "Add a key/value pair with key KEY and value VALUE if KEY does not exist",
														MarkdownDescription: "Add a key/value pair with key KEY and value VALUE if KEY does not exist",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"copy": schema.MapAttribute{
														Description:         "Copy a key/value pair with key KEY to COPIED_KEY if KEY exists AND COPIED_KEY does not exist",
														MarkdownDescription: "Copy a key/value pair with key KEY to COPIED_KEY if KEY exists AND COPIED_KEY does not exist",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hard_copy": schema.MapAttribute{
														Description:         "Copy a key/value pair with key KEY to COPIED_KEY if KEY exists. If COPIED_KEY already exists, this field is overwritten",
														MarkdownDescription: "Copy a key/value pair with key KEY to COPIED_KEY if KEY exists. If COPIED_KEY already exists, this field is overwritten",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hard_rename": schema.MapAttribute{
														Description:         "Rename a key/value pair with key KEY to RENAMED_KEY if KEY exists. If RENAMED_KEY already exists, this field is overwritten",
														MarkdownDescription: "Rename a key/value pair with key KEY to RENAMED_KEY if KEY exists. If RENAMED_KEY already exists, this field is overwritten",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"remove": schema.StringAttribute{
														Description:         "Remove a key/value pair with key KEY if it exists",
														MarkdownDescription: "Remove a key/value pair with key KEY if it exists",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"remove_regex": schema.StringAttribute{
														Description:         "Remove all key/value pairs with key matching regexp KEY",
														MarkdownDescription: "Remove all key/value pairs with key matching regexp KEY",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"remove_wildcard": schema.StringAttribute{
														Description:         "Remove all key/value pairs with key matching wildcard KEY",
														MarkdownDescription: "Remove all key/value pairs with key matching wildcard KEY",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"rename": schema.MapAttribute{
														Description:         "Rename a key/value pair with key KEY to RENAMED_KEY if KEY exists AND RENAMED_KEY does not exist",
														MarkdownDescription: "Rename a key/value pair with key KEY to RENAMED_KEY if KEY exists AND RENAMED_KEY does not exist",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set": schema.MapAttribute{
														Description:         "Add a key/value pair with key KEY and value VALUE. If KEY already exists, this field is overwritten",
														MarkdownDescription: "Add a key/value pair with key KEY and value VALUE. If KEY already exists, this field is overwritten",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"multiline": schema.SingleNestedAttribute{
									Description:         "Multiline defines a Multiline configuration.",
									MarkdownDescription: "Multiline defines a Multiline configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"buffer": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emitter_mem_buf_limit": schema.Int64Attribute{
											Description:         "Set a limit on the amount of memory in MB the emitter can consume if the outputs provide backpressure. The default for this limit is 10M. The pipeline will pause once the buffer exceeds the value of this setting. For example, if the value is set to 10MB then the pipeline will pause if the buffer exceeds 10M. The pipeline will remain paused until the output drains the buffer below the 10M limit.",
											MarkdownDescription: "Set a limit on the amount of memory in MB the emitter can consume if the outputs provide backpressure. The default for this limit is 10M. The pipeline will pause once the buffer exceeds the value of this setting. For example, if the value is set to 10MB then the pipeline will pause if the buffer exceeds 10M. The pipeline will remain paused until the output drains the buffer below the 10M limit.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emitter_name": schema.StringAttribute{
											Description:         "Name for the emitter input instance which re-emits the completed records at the beginning of the pipeline.",
											MarkdownDescription: "Name for the emitter input instance which re-emits the completed records at the beginning of the pipeline.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emitter_type": schema.StringAttribute{
											Description:         "The storage type for the emitter input instance. This option supports the values memory (default) and filesystem.",
											MarkdownDescription: "The storage type for the emitter input instance. This option supports the values memory (default) and filesystem.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("memory", "filesystem"),
											},
										},

										"flush_ms": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_content": schema.StringAttribute{
											Description:         "Key name that holds the content to process. Note that a Multiline Parser definition can already specify the key_content to use, but this option allows to overwrite that value for the purpose of the filter.",
											MarkdownDescription: "Key name that holds the content to process. Note that a Multiline Parser definition can already specify the key_content to use, but this option allows to overwrite that value for the purpose of the filter.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("parser", "partial_message"),
											},
										},

										"parser": schema.StringAttribute{
											Description:         "Specify one or multiple Multiline Parsing definitions to apply to the content. You can specify multiple multiline parsers to detect different formats by separating them with a comma.",
											MarkdownDescription: "Specify one or multiple Multiline Parsing definitions to apply to the content. You can specify multiple multiline parsers to detect different formats by separating them with a comma.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"nest": schema.SingleNestedAttribute{
									Description:         "Nest defines Nest Filter configuration.",
									MarkdownDescription: "Nest defines Nest Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"add_prefix": schema.StringAttribute{
											Description:         "Prefix affected keys with this string",
											MarkdownDescription: "Prefix affected keys with this string",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nest_under": schema.StringAttribute{
											Description:         "Nest records matching the Wildcard under this key",
											MarkdownDescription: "Nest records matching the Wildcard under this key",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nested_under": schema.StringAttribute{
											Description:         "Lift records nested under the Nested_under key",
											MarkdownDescription: "Lift records nested under the Nested_under key",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operation": schema.StringAttribute{
											Description:         "Select the operation nest or lift",
											MarkdownDescription: "Select the operation nest or lift",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("nest", "lift"),
											},
										},

										"remove_prefix": schema.StringAttribute{
											Description:         "Remove prefix from affected keys if it matches this string",
											MarkdownDescription: "Remove prefix from affected keys if it matches this string",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"wildcard": schema.ListAttribute{
											Description:         "Nest records which field matches the wildcard",
											MarkdownDescription: "Nest records which field matches the wildcard",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"parser": schema.SingleNestedAttribute{
									Description:         "Parser defines Parser Filter configuration.",
									MarkdownDescription: "Parser defines Parser Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_name": schema.StringAttribute{
											Description:         "Specify field name in record to parse.",
											MarkdownDescription: "Specify field name in record to parse.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parser": schema.StringAttribute{
											Description:         "Specify the parser name to interpret the field. Multiple Parser entries are allowed (split by comma).",
											MarkdownDescription: "Specify the parser name to interpret the field. Multiple Parser entries are allowed (split by comma).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"preserve_key": schema.BoolAttribute{
											Description:         "Keep original Key_Name field in the parsed result. If false, the field will be removed.",
											MarkdownDescription: "Keep original Key_Name field in the parsed result. If false, the field will be removed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reserve_data": schema.BoolAttribute{
											Description:         "Keep all other original fields in the parsed result. If false, all other original fields will be removed.",
											MarkdownDescription: "Keep all other original fields in the parsed result. If false, all other original fields will be removed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"unescape_key": schema.BoolAttribute{
											Description:         "If the key is a escaped string (e.g: stringify JSON), unescape the string before to apply the parser.",
											MarkdownDescription: "If the key is a escaped string (e.g: stringify JSON), unescape the string before to apply the parser.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"record_modifier": schema.SingleNestedAttribute{
									Description:         "RecordModifier defines Record Modifier Filter configuration.",
									MarkdownDescription: "RecordModifier defines Record Modifier Filter configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allowlist_keys": schema.ListAttribute{
											Description:         "If the key is not matched, that field is removed.",
											MarkdownDescription: "If the key is not matched, that field is removed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"records": schema.ListAttribute{
											Description:         "Append fields. This parameter needs key and value pair.",
											MarkdownDescription: "Append fields. This parameter needs key and value pair.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_keys": schema.ListAttribute{
											Description:         "If the key is matched, that field is removed.",
											MarkdownDescription: "If the key is matched, that field is removed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"uuid_keys": schema.ListAttribute{
											Description:         "If set, the plugin appends uuid to each record. The value assigned becomes the key in the map.",
											MarkdownDescription: "If set, the plugin appends uuid to each record. The value assigned becomes the key in the map.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"whitelist_keys": schema.ListAttribute{
											Description:         "An alias of allowlistKeys for backwards compatibility.",
											MarkdownDescription: "An alias of allowlistKeys for backwards compatibility.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"rewrite_tag": schema.SingleNestedAttribute{
									Description:         "RewriteTag defines a RewriteTag configuration.",
									MarkdownDescription: "RewriteTag defines a RewriteTag configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emitter_mem_buf_limit": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emitter_name": schema.StringAttribute{
											Description:         "When the filter emits a record under the new Tag, there is an internal emitter plugin that takes care of the job. Since this emitter expose metrics as any other component of the pipeline, you can use this property to configure an optional name for it.",
											MarkdownDescription: "When the filter emits a record under the new Tag, there is an internal emitter plugin that takes care of the job. Since this emitter expose metrics as any other component of the pipeline, you can use this property to configure an optional name for it.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emitter_storage_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"rules": schema.ListAttribute{
											Description:         "Defines the matching criteria and the format of the Tag for the matching record. The Rule format have four components: KEY REGEX NEW_TAG KEEP.",
											MarkdownDescription: "Defines the matching criteria and the format of the Tag for the matching record. The Rule format have four components: KEY REGEX NEW_TAG KEEP.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"throttle": schema.SingleNestedAttribute{
									Description:         "Throttle defines a Throttle configuration.",
									MarkdownDescription: "Throttle defines a Throttle configuration.",
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias for the plugin",
											MarkdownDescription: "Alias for the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interval": schema.StringAttribute{
											Description:         "Interval is the time interval expressed in 'sleep' format. e.g. 3s, 1.5m, 0.5h, etc.",
											MarkdownDescription: "Interval is the time interval expressed in 'sleep' format. e.g. 3s, 1.5m, 0.5h, etc.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"print_status": schema.BoolAttribute{
											Description:         "PrintStatus represents whether to print status messages with current rate and the limits to information logs.",
											MarkdownDescription: "PrintStatus represents whether to print status messages with current rate and the limits to information logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rate": schema.Int64Attribute{
											Description:         "Rate is the amount of messages for the time.",
											MarkdownDescription: "Rate is the amount of messages for the time.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_limit": schema.StringAttribute{
											Description:         "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											MarkdownDescription: "RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$`), ""),
											},
										},

										"window": schema.Int64Attribute{
											Description:         "Window is the amount of intervals to calculate average over.",
											MarkdownDescription: "Window is the amount of intervals to calculate average over.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("off", "error", "warning", "info", "debug", "trace"),
						},
					},

					"match": schema.StringAttribute{
						Description:         "A pattern to match against the tags of incoming records. It's case-sensitive and support the star (*) character as a wildcard.",
						MarkdownDescription: "A pattern to match against the tags of incoming records. It's case-sensitive and support the star (*) character as a wildcard.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"match_regex": schema.StringAttribute{
						Description:         "A regular expression to match against the tags of incoming records. Use this option if you want to use the full regex syntax.",
						MarkdownDescription: "A regular expression to match against the tags of incoming records. Use this option if you want to use the full regex syntax.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *FluentbitFluentIoFilterV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentbit_fluent_io_filter_v1alpha2_manifest")

	var model FluentbitFluentIoFilterV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentbit.fluent.io/v1alpha2")
	model.Kind = pointer.String("Filter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
