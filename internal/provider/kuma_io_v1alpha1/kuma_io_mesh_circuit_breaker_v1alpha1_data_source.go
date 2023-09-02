/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &KumaIoMeshCircuitBreakerV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshCircuitBreakerV1Alpha1DataSource{}
)

func NewKumaIoMeshCircuitBreakerV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshCircuitBreakerV1Alpha1DataSource{}
}

type KumaIoMeshCircuitBreakerV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshCircuitBreakerV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		From *[]struct {
			Default *struct {
				ConnectionLimits *struct {
					MaxConnectionPools *int64 `tfsdk:"max_connection_pools" json:"maxConnectionPools,omitempty"`
					MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
					MaxRequests        *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
					MaxRetries         *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
				} `tfsdk:"connection_limits" json:"connectionLimits,omitempty"`
				OutlierDetection *struct {
					BaseEjectionTime *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
					Detectors        *struct {
						FailurePercentage *struct {
							MinimumHosts  *int64 `tfsdk:"minimum_hosts" json:"minimumHosts,omitempty"`
							RequestVolume *int64 `tfsdk:"request_volume" json:"requestVolume,omitempty"`
							Threshold     *int64 `tfsdk:"threshold" json:"threshold,omitempty"`
						} `tfsdk:"failure_percentage" json:"failurePercentage,omitempty"`
						GatewayFailures *struct {
							Consecutive *int64 `tfsdk:"consecutive" json:"consecutive,omitempty"`
						} `tfsdk:"gateway_failures" json:"gatewayFailures,omitempty"`
						LocalOriginFailures *struct {
							Consecutive *int64 `tfsdk:"consecutive" json:"consecutive,omitempty"`
						} `tfsdk:"local_origin_failures" json:"localOriginFailures,omitempty"`
						SuccessRate *struct {
							MinimumHosts            *int64  `tfsdk:"minimum_hosts" json:"minimumHosts,omitempty"`
							RequestVolume           *int64  `tfsdk:"request_volume" json:"requestVolume,omitempty"`
							StandardDeviationFactor *string `tfsdk:"standard_deviation_factor" json:"standardDeviationFactor,omitempty"`
						} `tfsdk:"success_rate" json:"successRate,omitempty"`
						TotalFailures *struct {
							Consecutive *int64 `tfsdk:"consecutive" json:"consecutive,omitempty"`
						} `tfsdk:"total_failures" json:"totalFailures,omitempty"`
					} `tfsdk:"detectors" json:"detectors,omitempty"`
					Disabled                    *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					Interval                    *string `tfsdk:"interval" json:"interval,omitempty"`
					MaxEjectionPercent          *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
					SplitExternalAndLocalErrors *bool   `tfsdk:"split_external_and_local_errors" json:"splitExternalAndLocalErrors,omitempty"`
				} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		TargetRef *struct {
			Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		To *[]struct {
			Default *struct {
				ConnectionLimits *struct {
					MaxConnectionPools *int64 `tfsdk:"max_connection_pools" json:"maxConnectionPools,omitempty"`
					MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
					MaxRequests        *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
					MaxRetries         *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
				} `tfsdk:"connection_limits" json:"connectionLimits,omitempty"`
				OutlierDetection *struct {
					BaseEjectionTime *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
					Detectors        *struct {
						FailurePercentage *struct {
							MinimumHosts  *int64 `tfsdk:"minimum_hosts" json:"minimumHosts,omitempty"`
							RequestVolume *int64 `tfsdk:"request_volume" json:"requestVolume,omitempty"`
							Threshold     *int64 `tfsdk:"threshold" json:"threshold,omitempty"`
						} `tfsdk:"failure_percentage" json:"failurePercentage,omitempty"`
						GatewayFailures *struct {
							Consecutive *int64 `tfsdk:"consecutive" json:"consecutive,omitempty"`
						} `tfsdk:"gateway_failures" json:"gatewayFailures,omitempty"`
						LocalOriginFailures *struct {
							Consecutive *int64 `tfsdk:"consecutive" json:"consecutive,omitempty"`
						} `tfsdk:"local_origin_failures" json:"localOriginFailures,omitempty"`
						SuccessRate *struct {
							MinimumHosts            *int64  `tfsdk:"minimum_hosts" json:"minimumHosts,omitempty"`
							RequestVolume           *int64  `tfsdk:"request_volume" json:"requestVolume,omitempty"`
							StandardDeviationFactor *string `tfsdk:"standard_deviation_factor" json:"standardDeviationFactor,omitempty"`
						} `tfsdk:"success_rate" json:"successRate,omitempty"`
						TotalFailures *struct {
							Consecutive *int64 `tfsdk:"consecutive" json:"consecutive,omitempty"`
						} `tfsdk:"total_failures" json:"totalFailures,omitempty"`
					} `tfsdk:"detectors" json:"detectors,omitempty"`
					Disabled                    *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					Interval                    *string `tfsdk:"interval" json:"interval,omitempty"`
					MaxEjectionPercent          *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
					SplitExternalAndLocalErrors *bool   `tfsdk:"split_external_and_local_errors" json:"splitExternalAndLocalErrors,omitempty"`
				} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_circuit_breaker_v1alpha1"
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec is the specification of the Kuma MeshCircuitBreaker resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshCircuitBreaker resource.",
				Attributes: map[string]schema.Attribute{
					"from": schema.ListNestedAttribute{
						Description:         "From list makes a match between clients and corresponding configurations",
						MarkdownDescription: "From list makes a match between clients and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"connection_limits": schema.SingleNestedAttribute{
											Description:         "ConnectionLimits contains configuration of each circuit breaking limit, which when exceeded makes the circuit breaker to become open (no traffic is allowed like no current is allowed in the circuits when physical circuit breaker ir open)",
											MarkdownDescription: "ConnectionLimits contains configuration of each circuit breaking limit, which when exceeded makes the circuit breaker to become open (no traffic is allowed like no current is allowed in the circuits when physical circuit breaker ir open)",
											Attributes: map[string]schema.Attribute{
												"max_connection_pools": schema.Int64Attribute{
													Description:         "The maximum number of connection pools per cluster that are concurrently supported at once. Set this for clusters which create a large number of connection pools.",
													MarkdownDescription: "The maximum number of connection pools per cluster that are concurrently supported at once. Set this for clusters which create a large number of connection pools.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_connections": schema.Int64Attribute{
													Description:         "The maximum number of connections allowed to be made to the upstream cluster.",
													MarkdownDescription: "The maximum number of connections allowed to be made to the upstream cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_pending_requests": schema.Int64Attribute{
													Description:         "The maximum number of pending requests that are allowed to the upstream cluster. This limit is applied as a connection limit for non-HTTP traffic.",
													MarkdownDescription: "The maximum number of pending requests that are allowed to the upstream cluster. This limit is applied as a connection limit for non-HTTP traffic.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_requests": schema.Int64Attribute{
													Description:         "The maximum number of parallel requests that are allowed to be made to the upstream cluster. This limit does not apply to non-HTTP traffic.",
													MarkdownDescription: "The maximum number of parallel requests that are allowed to be made to the upstream cluster. This limit does not apply to non-HTTP traffic.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of parallel retries that will be allowed to the upstream cluster.",
													MarkdownDescription: "The maximum number of parallel retries that will be allowed to the upstream cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"outlier_detection": schema.SingleNestedAttribute{
											Description:         "OutlierDetection contains the configuration of the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Performance might be along different axes such as consecutive failures, temporal success rate, temporal latency, etc. Outlier detection is a form of passive health checking.",
											MarkdownDescription: "OutlierDetection contains the configuration of the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Performance might be along different axes such as consecutive failures, temporal success rate, temporal latency, etc. Outlier detection is a form of passive health checking.",
											Attributes: map[string]schema.Attribute{
												"base_ejection_time": schema.StringAttribute{
													Description:         "The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected.",
													MarkdownDescription: "The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"detectors": schema.SingleNestedAttribute{
													Description:         "Contains configuration for supported outlier detectors",
													MarkdownDescription: "Contains configuration for supported outlier detectors",
													Attributes: map[string]schema.Attribute{
														"failure_percentage": schema.SingleNestedAttribute{
															Description:         "Failure Percentage based outlier detection functions similarly to success rate detection, in that it relies on success rate data from each host in a cluster. However, rather than compare those values to the mean success rate of the cluster as a whole, they are compared to a flat user-configured threshold. This threshold is configured via the outlierDetection.failurePercentageThreshold field. The other configuration fields for failure percentage based detection are similar to the fields for success rate detection. As with success rate detection, detection will not be performed for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.failurePercentage.requestVolume value. Detection also will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.failurePercentage.minimumHosts value.",
															MarkdownDescription: "Failure Percentage based outlier detection functions similarly to success rate detection, in that it relies on success rate data from each host in a cluster. However, rather than compare those values to the mean success rate of the cluster as a whole, they are compared to a flat user-configured threshold. This threshold is configured via the outlierDetection.failurePercentageThreshold field. The other configuration fields for failure percentage based detection are similar to the fields for success rate detection. As with success rate detection, detection will not be performed for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.failurePercentage.requestVolume value. Detection also will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.failurePercentage.minimumHosts value.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The minimum number of hosts in a cluster in order to perform failure percentage-based ejection. If the total number of hosts in the cluster is less than this value, failure percentage-based ejection will not be performed.",
																	MarkdownDescription: "The minimum number of hosts in a cluster in order to perform failure percentage-based ejection. If the total number of hosts in the cluster is less than this value, failure percentage-based ejection will not be performed.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in one interval (as defined by the interval duration above) to perform failure percentage-based ejection for this host. If the volume is lower than this setting, failure percentage-based ejection will not be performed for this host.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in one interval (as defined by the interval duration above) to perform failure percentage-based ejection for this host. If the volume is lower than this setting, failure percentage-based ejection will not be performed for this host.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"threshold": schema.Int64Attribute{
																	Description:         "The failure percentage to use when determining failure percentage-based outlier detection. If the failure percentage of a given host is greater than or equal to this value, it will be ejected.",
																	MarkdownDescription: "The failure percentage to use when determining failure percentage-based outlier detection. If the failure percentage of a given host is greater than or equal to this value, it will be ejected.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"gateway_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and local origin failures, such as timeout, TCP reset etc. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and is supported only by the http router.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and local origin failures, such as timeout, TCP reset etc. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and is supported only by the http router.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive gateway failures (502, 503, 504 status codes) before a consecutive gateway failure ejection occurs.",
																	MarkdownDescription: "The number of consecutive gateway failures (502, 503, 504 status codes) before a consecutive gateway failure ejection occurs.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"local_origin_failures": schema.SingleNestedAttribute{
															Description:         "This detection type is enabled only when outlierDetection.splitExternalLocalOriginErrors is true and takes into account only locally originated errors (timeout, reset, etc). If Envoy repeatedly cannot connect to an upstream host or communication with the upstream host is repeatedly interrupted, it will be ejected. Various locally originated problems are detected: timeout, TCP reset, ICMP errors, etc. This detection type is supported by http router and tcp proxy.",
															MarkdownDescription: "This detection type is enabled only when outlierDetection.splitExternalLocalOriginErrors is true and takes into account only locally originated errors (timeout, reset, etc). If Envoy repeatedly cannot connect to an upstream host or communication with the upstream host is repeatedly interrupted, it will be ejected. Various locally originated problems are detected: timeout, TCP reset, ICMP errors, etc. This detection type is supported by http router and tcp proxy.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive locally originated failures before ejection occurs. Parameter takes effect only when splitExternalAndLocalErrors is set to true.",
																	MarkdownDescription: "The number of consecutive locally originated failures before ejection occurs. Parameter takes effect only when splitExternalAndLocalErrors is set to true.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"success_rate": schema.SingleNestedAttribute{
															Description:         "Success Rate based outlier detection aggregates success rate data from every host in a cluster. Then at given intervals ejects hosts based on statistical outlier detection. Success Rate outlier detection will not be calculated for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.successRate.requestVolume value. Moreover, detection will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.successRate.minimumHosts value. In the default configuration mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account all types of errors: locally and externally originated. In split mode (outlierDetection.splitExternalLocalOriginErrors is true), locally originated errors and externally originated (transaction) errors are counted and treated separately.",
															MarkdownDescription: "Success Rate based outlier detection aggregates success rate data from every host in a cluster. Then at given intervals ejects hosts based on statistical outlier detection. Success Rate outlier detection will not be calculated for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.successRate.requestVolume value. Moreover, detection will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.successRate.minimumHosts value. In the default configuration mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account all types of errors: locally and externally originated. In split mode (outlierDetection.splitExternalLocalOriginErrors is true), locally originated errors and externally originated (transaction) errors are counted and treated separately.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The number of hosts in a cluster that must have enough request volume to detect success rate outliers. If the number of hosts is less than this setting, outlier detection via success rate statistics is not performed for any host in the cluster.",
																	MarkdownDescription: "The number of hosts in a cluster that must have enough request volume to detect success rate outliers. If the number of hosts is less than this setting, outlier detection via success rate statistics is not performed for any host in the cluster.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in one interval (as defined by the interval duration configured in outlierDetection section) to include this host in success rate based outlier detection. If the volume is lower than this setting, outlier detection via success rate statistics is not performed for that host.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in one interval (as defined by the interval duration configured in outlierDetection section) to include this host in success rate based outlier detection. If the volume is lower than this setting, outlier detection via success rate statistics is not performed for that host.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"standard_deviation_factor": schema.StringAttribute{
																	Description:         "This factor is used to determine the ejection threshold for success rate outlier ejection. The ejection threshold is the difference between the mean success rate, and the product of this factor and the standard deviation of the mean success rate: mean - (standard_deviation * success_rate_standard_deviation_factor). Either int or decimal represented as string.",
																	MarkdownDescription: "This factor is used to determine the ejection threshold for success rate outlier ejection. The ejection threshold is the difference between the mean success rate, and the product of this factor and the standard deviation of the mean success rate: mean - (standard_deviation * success_rate_standard_deviation_factor). Either int or decimal represented as string.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"total_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalAndLocalErrors is false) this detection type takes into account all generated errors: locally originated and externally originated (transaction) errors. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account only externally originated (transaction) errors, ignoring locally originated errors. If an upstream host is an HTTP-server, only 5xx types of error are taken into account (see Consecutive Gateway Failure for exceptions). Properly formatted responses, even when they carry an operational error (like index not found, access denied) are not taken into account.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalAndLocalErrors is false) this detection type takes into account all generated errors: locally originated and externally originated (transaction) errors. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account only externally originated (transaction) errors, ignoring locally originated errors. If an upstream host is an HTTP-server, only 5xx types of error are taken into account (see Consecutive Gateway Failure for exceptions). Properly formatted responses, even when they carry an operational error (like index not found, access denied) are not taken into account.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive server-side error responses (for HTTP traffic, 5xx responses; for TCP traffic, connection failures; for Redis, failure to respond PONG; etc.) before a consecutive total failure ejection occurs.",
																	MarkdownDescription: "The number of consecutive server-side error responses (for HTTP traffic, 5xx responses; for TCP traffic, connection failures; for Redis, failure to respond PONG; etc.) before a consecutive total failure ejection occurs.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"disabled": schema.BoolAttribute{
													Description:         "When set to true, outlierDetection configuration won't take any effect",
													MarkdownDescription: "When set to true, outlierDetection configuration won't take any effect",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"interval": schema.StringAttribute{
													Description:         "The time interval between ejection analysis sweeps. This can result in both new ejections and hosts being returned to service.",
													MarkdownDescription: "The time interval between ejection analysis sweeps. This can result in both new ejections and hosts being returned to service.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_ejection_percent": schema.Int64Attribute{
													Description:         "The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10% but will eject at least one host regardless of the value.",
													MarkdownDescription: "The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10% but will eject at least one host regardless of the value.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"split_external_and_local_errors": schema.BoolAttribute{
													Description:         "Determines whether to distinguish local origin failures from external errors. If set to true the following configuration parameters are taken into account: detectors.localOriginFailures.consecutive",
													MarkdownDescription: "Determines whether to distinguish local origin failures from external errors. If set to true the following configuration parameters are taken into account: detectors.localOriginFailures.consecutive",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of destinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined in place.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined in place.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between the consumed services and corresponding configurations",
						MarkdownDescription: "To list makes a match between the consumed services and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"connection_limits": schema.SingleNestedAttribute{
											Description:         "ConnectionLimits contains configuration of each circuit breaking limit, which when exceeded makes the circuit breaker to become open (no traffic is allowed like no current is allowed in the circuits when physical circuit breaker ir open)",
											MarkdownDescription: "ConnectionLimits contains configuration of each circuit breaking limit, which when exceeded makes the circuit breaker to become open (no traffic is allowed like no current is allowed in the circuits when physical circuit breaker ir open)",
											Attributes: map[string]schema.Attribute{
												"max_connection_pools": schema.Int64Attribute{
													Description:         "The maximum number of connection pools per cluster that are concurrently supported at once. Set this for clusters which create a large number of connection pools.",
													MarkdownDescription: "The maximum number of connection pools per cluster that are concurrently supported at once. Set this for clusters which create a large number of connection pools.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_connections": schema.Int64Attribute{
													Description:         "The maximum number of connections allowed to be made to the upstream cluster.",
													MarkdownDescription: "The maximum number of connections allowed to be made to the upstream cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_pending_requests": schema.Int64Attribute{
													Description:         "The maximum number of pending requests that are allowed to the upstream cluster. This limit is applied as a connection limit for non-HTTP traffic.",
													MarkdownDescription: "The maximum number of pending requests that are allowed to the upstream cluster. This limit is applied as a connection limit for non-HTTP traffic.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_requests": schema.Int64Attribute{
													Description:         "The maximum number of parallel requests that are allowed to be made to the upstream cluster. This limit does not apply to non-HTTP traffic.",
													MarkdownDescription: "The maximum number of parallel requests that are allowed to be made to the upstream cluster. This limit does not apply to non-HTTP traffic.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of parallel retries that will be allowed to the upstream cluster.",
													MarkdownDescription: "The maximum number of parallel retries that will be allowed to the upstream cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"outlier_detection": schema.SingleNestedAttribute{
											Description:         "OutlierDetection contains the configuration of the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Performance might be along different axes such as consecutive failures, temporal success rate, temporal latency, etc. Outlier detection is a form of passive health checking.",
											MarkdownDescription: "OutlierDetection contains the configuration of the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Performance might be along different axes such as consecutive failures, temporal success rate, temporal latency, etc. Outlier detection is a form of passive health checking.",
											Attributes: map[string]schema.Attribute{
												"base_ejection_time": schema.StringAttribute{
													Description:         "The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected.",
													MarkdownDescription: "The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"detectors": schema.SingleNestedAttribute{
													Description:         "Contains configuration for supported outlier detectors",
													MarkdownDescription: "Contains configuration for supported outlier detectors",
													Attributes: map[string]schema.Attribute{
														"failure_percentage": schema.SingleNestedAttribute{
															Description:         "Failure Percentage based outlier detection functions similarly to success rate detection, in that it relies on success rate data from each host in a cluster. However, rather than compare those values to the mean success rate of the cluster as a whole, they are compared to a flat user-configured threshold. This threshold is configured via the outlierDetection.failurePercentageThreshold field. The other configuration fields for failure percentage based detection are similar to the fields for success rate detection. As with success rate detection, detection will not be performed for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.failurePercentage.requestVolume value. Detection also will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.failurePercentage.minimumHosts value.",
															MarkdownDescription: "Failure Percentage based outlier detection functions similarly to success rate detection, in that it relies on success rate data from each host in a cluster. However, rather than compare those values to the mean success rate of the cluster as a whole, they are compared to a flat user-configured threshold. This threshold is configured via the outlierDetection.failurePercentageThreshold field. The other configuration fields for failure percentage based detection are similar to the fields for success rate detection. As with success rate detection, detection will not be performed for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.failurePercentage.requestVolume value. Detection also will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.failurePercentage.minimumHosts value.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The minimum number of hosts in a cluster in order to perform failure percentage-based ejection. If the total number of hosts in the cluster is less than this value, failure percentage-based ejection will not be performed.",
																	MarkdownDescription: "The minimum number of hosts in a cluster in order to perform failure percentage-based ejection. If the total number of hosts in the cluster is less than this value, failure percentage-based ejection will not be performed.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in one interval (as defined by the interval duration above) to perform failure percentage-based ejection for this host. If the volume is lower than this setting, failure percentage-based ejection will not be performed for this host.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in one interval (as defined by the interval duration above) to perform failure percentage-based ejection for this host. If the volume is lower than this setting, failure percentage-based ejection will not be performed for this host.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"threshold": schema.Int64Attribute{
																	Description:         "The failure percentage to use when determining failure percentage-based outlier detection. If the failure percentage of a given host is greater than or equal to this value, it will be ejected.",
																	MarkdownDescription: "The failure percentage to use when determining failure percentage-based outlier detection. If the failure percentage of a given host is greater than or equal to this value, it will be ejected.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"gateway_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and local origin failures, such as timeout, TCP reset etc. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and is supported only by the http router.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and local origin failures, such as timeout, TCP reset etc. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account a subset of 5xx errors, called 'gateway errors' (502, 503 or 504 status code) and is supported only by the http router.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive gateway failures (502, 503, 504 status codes) before a consecutive gateway failure ejection occurs.",
																	MarkdownDescription: "The number of consecutive gateway failures (502, 503, 504 status codes) before a consecutive gateway failure ejection occurs.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"local_origin_failures": schema.SingleNestedAttribute{
															Description:         "This detection type is enabled only when outlierDetection.splitExternalLocalOriginErrors is true and takes into account only locally originated errors (timeout, reset, etc). If Envoy repeatedly cannot connect to an upstream host or communication with the upstream host is repeatedly interrupted, it will be ejected. Various locally originated problems are detected: timeout, TCP reset, ICMP errors, etc. This detection type is supported by http router and tcp proxy.",
															MarkdownDescription: "This detection type is enabled only when outlierDetection.splitExternalLocalOriginErrors is true and takes into account only locally originated errors (timeout, reset, etc). If Envoy repeatedly cannot connect to an upstream host or communication with the upstream host is repeatedly interrupted, it will be ejected. Various locally originated problems are detected: timeout, TCP reset, ICMP errors, etc. This detection type is supported by http router and tcp proxy.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive locally originated failures before ejection occurs. Parameter takes effect only when splitExternalAndLocalErrors is set to true.",
																	MarkdownDescription: "The number of consecutive locally originated failures before ejection occurs. Parameter takes effect only when splitExternalAndLocalErrors is set to true.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"success_rate": schema.SingleNestedAttribute{
															Description:         "Success Rate based outlier detection aggregates success rate data from every host in a cluster. Then at given intervals ejects hosts based on statistical outlier detection. Success Rate outlier detection will not be calculated for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.successRate.requestVolume value. Moreover, detection will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.successRate.minimumHosts value. In the default configuration mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account all types of errors: locally and externally originated. In split mode (outlierDetection.splitExternalLocalOriginErrors is true), locally originated errors and externally originated (transaction) errors are counted and treated separately.",
															MarkdownDescription: "Success Rate based outlier detection aggregates success rate data from every host in a cluster. Then at given intervals ejects hosts based on statistical outlier detection. Success Rate outlier detection will not be calculated for a host if its request volume over the aggregation interval is less than the outlierDetection.detectors.successRate.requestVolume value. Moreover, detection will not be performed for a cluster if the number of hosts with the minimum required request volume in an interval is less than the outlierDetection.detectors.successRate.minimumHosts value. In the default configuration mode (outlierDetection.splitExternalLocalOriginErrors is false) this detection type takes into account all types of errors: locally and externally originated. In split mode (outlierDetection.splitExternalLocalOriginErrors is true), locally originated errors and externally originated (transaction) errors are counted and treated separately.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The number of hosts in a cluster that must have enough request volume to detect success rate outliers. If the number of hosts is less than this setting, outlier detection via success rate statistics is not performed for any host in the cluster.",
																	MarkdownDescription: "The number of hosts in a cluster that must have enough request volume to detect success rate outliers. If the number of hosts is less than this setting, outlier detection via success rate statistics is not performed for any host in the cluster.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in one interval (as defined by the interval duration configured in outlierDetection section) to include this host in success rate based outlier detection. If the volume is lower than this setting, outlier detection via success rate statistics is not performed for that host.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in one interval (as defined by the interval duration configured in outlierDetection section) to include this host in success rate based outlier detection. If the volume is lower than this setting, outlier detection via success rate statistics is not performed for that host.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"standard_deviation_factor": schema.StringAttribute{
																	Description:         "This factor is used to determine the ejection threshold for success rate outlier ejection. The ejection threshold is the difference between the mean success rate, and the product of this factor and the standard deviation of the mean success rate: mean - (standard_deviation * success_rate_standard_deviation_factor). Either int or decimal represented as string.",
																	MarkdownDescription: "This factor is used to determine the ejection threshold for success rate outlier ejection. The ejection threshold is the difference between the mean success rate, and the product of this factor and the standard deviation of the mean success rate: mean - (standard_deviation * success_rate_standard_deviation_factor). Either int or decimal represented as string.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"total_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalAndLocalErrors is false) this detection type takes into account all generated errors: locally originated and externally originated (transaction) errors. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account only externally originated (transaction) errors, ignoring locally originated errors. If an upstream host is an HTTP-server, only 5xx types of error are taken into account (see Consecutive Gateway Failure for exceptions). Properly formatted responses, even when they carry an operational error (like index not found, access denied) are not taken into account.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalAndLocalErrors is false) this detection type takes into account all generated errors: locally originated and externally originated (transaction) errors. In split mode (outlierDetection.splitExternalLocalOriginErrors is true) this detection type takes into account only externally originated (transaction) errors, ignoring locally originated errors. If an upstream host is an HTTP-server, only 5xx types of error are taken into account (see Consecutive Gateway Failure for exceptions). Properly formatted responses, even when they carry an operational error (like index not found, access denied) are not taken into account.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive server-side error responses (for HTTP traffic, 5xx responses; for TCP traffic, connection failures; for Redis, failure to respond PONG; etc.) before a consecutive total failure ejection occurs.",
																	MarkdownDescription: "The number of consecutive server-side error responses (for HTTP traffic, 5xx responses; for TCP traffic, connection failures; for Redis, failure to respond PONG; etc.) before a consecutive total failure ejection occurs.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"disabled": schema.BoolAttribute{
													Description:         "When set to true, outlierDetection configuration won't take any effect",
													MarkdownDescription: "When set to true, outlierDetection configuration won't take any effect",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"interval": schema.StringAttribute{
													Description:         "The time interval between ejection analysis sweeps. This can result in both new ejections and hosts being returned to service.",
													MarkdownDescription: "The time interval between ejection analysis sweeps. This can result in both new ejections and hosts being returned to service.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_ejection_percent": schema.Int64Attribute{
													Description:         "The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10% but will eject at least one host regardless of the value.",
													MarkdownDescription: "The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10% but will eject at least one host regardless of the value.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"split_external_and_local_errors": schema.BoolAttribute{
													Description:         "Determines whether to distinguish local origin failures from external errors. If set to true the following configuration parameters are taken into account: detectors.localOriginFailures.consecutive",
													MarkdownDescription: "Determines whether to distinguish local origin failures from external errors. If set to true the following configuration parameters are taken into account: detectors.localOriginFailures.consecutive",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of destinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_circuit_breaker_v1alpha1")

	var data KumaIoMeshCircuitBreakerV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshCircuitBreaker"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KumaIoMeshCircuitBreakerV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("MeshCircuitBreaker")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
