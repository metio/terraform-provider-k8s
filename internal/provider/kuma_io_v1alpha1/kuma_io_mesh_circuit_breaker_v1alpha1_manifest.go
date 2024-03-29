/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KumaIoMeshCircuitBreakerV1Alpha1Manifest{}
)

func NewKumaIoMeshCircuitBreakerV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshCircuitBreakerV1Alpha1Manifest{}
}

type KumaIoMeshCircuitBreakerV1Alpha1Manifest struct{}

type KumaIoMeshCircuitBreakerV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
				Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name       *string            `tfsdk:"name" json:"name,omitempty"`
				ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
				Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		TargetRef *struct {
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
			Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
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
				Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name       *string            `tfsdk:"name" json:"name,omitempty"`
				ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
				Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_circuit_breaker_v1alpha1_manifest"
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec is the specification of the Kuma MeshCircuitBreaker resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshCircuitBreaker resource.",
				Attributes: map[string]schema.Attribute{
					"from": schema.ListNestedAttribute{
						Description:         "From list makes a match between clients and corresponding configurations",
						MarkdownDescription: "From list makes a match between clients and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinationsreferenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinationsreferenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"connection_limits": schema.SingleNestedAttribute{
											Description:         "ConnectionLimits contains configuration of each circuit breaking limit,which when exceeded makes the circuit breaker to become open (no trafficis allowed like no current is allowed in the circuits when physicalcircuit breaker ir open)",
											MarkdownDescription: "ConnectionLimits contains configuration of each circuit breaking limit,which when exceeded makes the circuit breaker to become open (no trafficis allowed like no current is allowed in the circuits when physicalcircuit breaker ir open)",
											Attributes: map[string]schema.Attribute{
												"max_connection_pools": schema.Int64Attribute{
													Description:         "The maximum number of connection pools per cluster that are concurrentlysupported at once. Set this for clusters which create a large number ofconnection pools.",
													MarkdownDescription: "The maximum number of connection pools per cluster that are concurrentlysupported at once. Set this for clusters which create a large number ofconnection pools.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_connections": schema.Int64Attribute{
													Description:         "The maximum number of connections allowed to be made to the upstreamcluster.",
													MarkdownDescription: "The maximum number of connections allowed to be made to the upstreamcluster.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_pending_requests": schema.Int64Attribute{
													Description:         "The maximum number of pending requests that are allowed to the upstreamcluster. This limit is applied as a connection limit for non-HTTPtraffic.",
													MarkdownDescription: "The maximum number of pending requests that are allowed to the upstreamcluster. This limit is applied as a connection limit for non-HTTPtraffic.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_requests": schema.Int64Attribute{
													Description:         "The maximum number of parallel requests that are allowed to be madeto the upstream cluster. This limit does not apply to non-HTTP traffic.",
													MarkdownDescription: "The maximum number of parallel requests that are allowed to be madeto the upstream cluster. This limit does not apply to non-HTTP traffic.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of parallel retries that will be allowed tothe upstream cluster.",
													MarkdownDescription: "The maximum number of parallel retries that will be allowed tothe upstream cluster.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"outlier_detection": schema.SingleNestedAttribute{
											Description:         "OutlierDetection contains the configuration of the process of dynamicallydetermining whether some number of hosts in an upstream cluster areperforming unlike the others and removing them from the healthy loadbalancing set. Performance might be along different axes such asconsecutive failures, temporal success rate, temporal latency, etc.Outlier detection is a form of passive health checking.",
											MarkdownDescription: "OutlierDetection contains the configuration of the process of dynamicallydetermining whether some number of hosts in an upstream cluster areperforming unlike the others and removing them from the healthy loadbalancing set. Performance might be along different axes such asconsecutive failures, temporal success rate, temporal latency, etc.Outlier detection is a form of passive health checking.",
											Attributes: map[string]schema.Attribute{
												"base_ejection_time": schema.StringAttribute{
													Description:         "The base time that a host is ejected for. The real time is equal tothe base time multiplied by the number of times the host has beenejected.",
													MarkdownDescription: "The base time that a host is ejected for. The real time is equal tothe base time multiplied by the number of times the host has beenejected.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"detectors": schema.SingleNestedAttribute{
													Description:         "Contains configuration for supported outlier detectors",
													MarkdownDescription: "Contains configuration for supported outlier detectors",
													Attributes: map[string]schema.Attribute{
														"failure_percentage": schema.SingleNestedAttribute{
															Description:         "Failure Percentage based outlier detection functions similarly to successrate detection, in that it relies on success rate data from each host ina cluster. However, rather than compare those values to the mean successrate of the cluster as a whole, they are compared to a flatuser-configured threshold. This threshold is configured via theoutlierDetection.failurePercentageThreshold field.The other configuration fields for failure percentage based detection aresimilar to the fields for success rate detection. As with success ratedetection, detection will not be performed for a host if its requestvolume over the aggregation interval is less than theoutlierDetection.detectors.failurePercentage.requestVolume value.Detection also will not be performed for a cluster if the number of hostswith the minimum required request volume in an interval is less than theoutlierDetection.detectors.failurePercentage.minimumHosts value.",
															MarkdownDescription: "Failure Percentage based outlier detection functions similarly to successrate detection, in that it relies on success rate data from each host ina cluster. However, rather than compare those values to the mean successrate of the cluster as a whole, they are compared to a flatuser-configured threshold. This threshold is configured via theoutlierDetection.failurePercentageThreshold field.The other configuration fields for failure percentage based detection aresimilar to the fields for success rate detection. As with success ratedetection, detection will not be performed for a host if its requestvolume over the aggregation interval is less than theoutlierDetection.detectors.failurePercentage.requestVolume value.Detection also will not be performed for a cluster if the number of hostswith the minimum required request volume in an interval is less than theoutlierDetection.detectors.failurePercentage.minimumHosts value.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The minimum number of hosts in a cluster in order to perform failurepercentage-based ejection. If the total number of hosts in the cluster isless than this value, failure percentage-based ejection will not beperformed.",
																	MarkdownDescription: "The minimum number of hosts in a cluster in order to perform failurepercentage-based ejection. If the total number of hosts in the cluster isless than this value, failure percentage-based ejection will not beperformed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration above) to perform failurepercentage-based ejection for this host. If the volume is lower than thissetting, failure percentage-based ejection will not be performed for thishost.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration above) to perform failurepercentage-based ejection for this host. If the volume is lower than thissetting, failure percentage-based ejection will not be performed for thishost.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"threshold": schema.Int64Attribute{
																	Description:         "The failure percentage to use when determining failure percentage-basedoutlier detection. If the failure percentage of a given host is greaterthan or equal to this value, it will be ejected.",
																	MarkdownDescription: "The failure percentage to use when determining failure percentage-basedoutlier detection. If the failure percentage of a given host is greaterthan or equal to this value, it will be ejected.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"gateway_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalLocalOriginErrors isfalse) this detection type takes into account a subset of 5xx errors,called 'gateway errors' (502, 503 or 504 status code) and local originfailures, such as timeout, TCP reset etc.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account a subset of 5xx errors, called'gateway errors' (502, 503 or 504 status code) and is supported only bythe http router.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalLocalOriginErrors isfalse) this detection type takes into account a subset of 5xx errors,called 'gateway errors' (502, 503 or 504 status code) and local originfailures, such as timeout, TCP reset etc.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account a subset of 5xx errors, called'gateway errors' (502, 503 or 504 status code) and is supported only bythe http router.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive gateway failures (502, 503, 504 status codes)before a consecutive gateway failure ejection occurs.",
																	MarkdownDescription: "The number of consecutive gateway failures (502, 503, 504 status codes)before a consecutive gateway failure ejection occurs.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"local_origin_failures": schema.SingleNestedAttribute{
															Description:         "This detection type is enabled only whenoutlierDetection.splitExternalLocalOriginErrors is true and takes intoaccount only locally originated errors (timeout, reset, etc).If Envoy repeatedly cannot connect to an upstream host or communicationwith the upstream host is repeatedly interrupted, it will be ejected.Various locally originated problems are detected: timeout, TCP reset,ICMP errors, etc. This detection type is supported by http router andtcp proxy.",
															MarkdownDescription: "This detection type is enabled only whenoutlierDetection.splitExternalLocalOriginErrors is true and takes intoaccount only locally originated errors (timeout, reset, etc).If Envoy repeatedly cannot connect to an upstream host or communicationwith the upstream host is repeatedly interrupted, it will be ejected.Various locally originated problems are detected: timeout, TCP reset,ICMP errors, etc. This detection type is supported by http router andtcp proxy.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive locally originated failures before ejectionoccurs. Parameter takes effect only when splitExternalAndLocalErrorsis set to true.",
																	MarkdownDescription: "The number of consecutive locally originated failures before ejectionoccurs. Parameter takes effect only when splitExternalAndLocalErrorsis set to true.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"success_rate": schema.SingleNestedAttribute{
															Description:         "Success Rate based outlier detection aggregates success rate data fromevery host in a cluster. Then at given intervals ejects hosts based onstatistical outlier detection. Success Rate outlier detection will not becalculated for a host if its request volume over the aggregation intervalis less than the outlierDetection.detectors.successRate.requestVolumevalue.Moreover, detection will not be performed for a cluster if the number ofhosts with the minimum required request volume in an interval is lessthan the outlierDetection.detectors.successRate.minimumHosts value.In the default configuration mode(outlierDetection.splitExternalLocalOriginErrors is false) this detectiontype takes into account all types of errors: locally and externallyoriginated.In split mode (outlierDetection.splitExternalLocalOriginErrors is true),locally originated errors and externally originated (transaction) errorsare counted and treated separately.",
															MarkdownDescription: "Success Rate based outlier detection aggregates success rate data fromevery host in a cluster. Then at given intervals ejects hosts based onstatistical outlier detection. Success Rate outlier detection will not becalculated for a host if its request volume over the aggregation intervalis less than the outlierDetection.detectors.successRate.requestVolumevalue.Moreover, detection will not be performed for a cluster if the number ofhosts with the minimum required request volume in an interval is lessthan the outlierDetection.detectors.successRate.minimumHosts value.In the default configuration mode(outlierDetection.splitExternalLocalOriginErrors is false) this detectiontype takes into account all types of errors: locally and externallyoriginated.In split mode (outlierDetection.splitExternalLocalOriginErrors is true),locally originated errors and externally originated (transaction) errorsare counted and treated separately.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The number of hosts in a cluster that must have enough request volume todetect success rate outliers. If the number of hosts is less than thissetting, outlier detection via success rate statistics is not performedfor any host in the cluster.",
																	MarkdownDescription: "The number of hosts in a cluster that must have enough request volume todetect success rate outliers. If the number of hosts is less than thissetting, outlier detection via success rate statistics is not performedfor any host in the cluster.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration configured inoutlierDetection section) to include this host in success rate basedoutlier detection. If the volume is lower than this setting, outlierdetection via success rate statistics is not performed for that host.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration configured inoutlierDetection section) to include this host in success rate basedoutlier detection. If the volume is lower than this setting, outlierdetection via success rate statistics is not performed for that host.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"standard_deviation_factor": schema.StringAttribute{
																	Description:         "This factor is used to determine the ejection threshold for success rateoutlier ejection. The ejection threshold is the difference betweenthe mean success rate, and the product of this factor and the standarddeviation of the mean success rate: mean - (standard_deviation *success_rate_standard_deviation_factor).Either int or decimal represented as string.",
																	MarkdownDescription: "This factor is used to determine the ejection threshold for success rateoutlier ejection. The ejection threshold is the difference betweenthe mean success rate, and the product of this factor and the standarddeviation of the mean success rate: mean - (standard_deviation *success_rate_standard_deviation_factor).Either int or decimal represented as string.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"total_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalAndLocalErrors isfalse) this detection type takes into account all generated errors:locally originated and externally originated (transaction) errors.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account only externally originated(transaction) errors, ignoring locally originated errors.If an upstream host is an HTTP-server, only 5xx types of error are takeninto account (see Consecutive Gateway Failure for exceptions).Properly formatted responses, even when they carry an operational error(like index not found, access denied) are not taken into account.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalAndLocalErrors isfalse) this detection type takes into account all generated errors:locally originated and externally originated (transaction) errors.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account only externally originated(transaction) errors, ignoring locally originated errors.If an upstream host is an HTTP-server, only 5xx types of error are takeninto account (see Consecutive Gateway Failure for exceptions).Properly formatted responses, even when they carry an operational error(like index not found, access denied) are not taken into account.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive server-side error responses (for HTTP traffic,5xx responses; for TCP traffic, connection failures; for Redis, failureto respond PONG; etc.) before a consecutive total failure ejectionoccurs.",
																	MarkdownDescription: "The number of consecutive server-side error responses (for HTTP traffic,5xx responses; for TCP traffic, connection failures; for Redis, failureto respond PONG; etc.) before a consecutive total failure ejectionoccurs.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"disabled": schema.BoolAttribute{
													Description:         "When set to true, outlierDetection configuration won't take any effect",
													MarkdownDescription: "When set to true, outlierDetection configuration won't take any effect",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "The time interval between ejection analysis sweeps. This can result inboth new ejections and hosts being returned to service.",
													MarkdownDescription: "The time interval between ejection analysis sweeps. This can result inboth new ejections and hosts being returned to service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_ejection_percent": schema.Int64Attribute{
													Description:         "The maximum % of an upstream cluster that can be ejected due to outlierdetection. Defaults to 10% but will eject at least one host regardless ofthe value.",
													MarkdownDescription: "The maximum % of an upstream cluster that can be ejected due to outlierdetection. Defaults to 10% but will eject at least one host regardless ofthe value.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"split_external_and_local_errors": schema.BoolAttribute{
													Description:         "Determines whether to distinguish local origin failures from externalerrors. If set to true the following configuration parameters are takeninto account: detectors.localOriginFailures.consecutive",
													MarkdownDescription: "Determines whether to distinguish local origin failures from externalerrors. If set to true the following configuration parameters are takeninto account: detectors.localOriginFailures.consecutive",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group ofdestinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofdestinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
											},
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_types": schema.ListAttribute{
											Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined in place.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined in place.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_types": schema.ListAttribute{
								Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between the consumed services and correspondingconfigurations",
						MarkdownDescription: "To list makes a match between the consumed services and correspondingconfigurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinationsreferenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinationsreferenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"connection_limits": schema.SingleNestedAttribute{
											Description:         "ConnectionLimits contains configuration of each circuit breaking limit,which when exceeded makes the circuit breaker to become open (no trafficis allowed like no current is allowed in the circuits when physicalcircuit breaker ir open)",
											MarkdownDescription: "ConnectionLimits contains configuration of each circuit breaking limit,which when exceeded makes the circuit breaker to become open (no trafficis allowed like no current is allowed in the circuits when physicalcircuit breaker ir open)",
											Attributes: map[string]schema.Attribute{
												"max_connection_pools": schema.Int64Attribute{
													Description:         "The maximum number of connection pools per cluster that are concurrentlysupported at once. Set this for clusters which create a large number ofconnection pools.",
													MarkdownDescription: "The maximum number of connection pools per cluster that are concurrentlysupported at once. Set this for clusters which create a large number ofconnection pools.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_connections": schema.Int64Attribute{
													Description:         "The maximum number of connections allowed to be made to the upstreamcluster.",
													MarkdownDescription: "The maximum number of connections allowed to be made to the upstreamcluster.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_pending_requests": schema.Int64Attribute{
													Description:         "The maximum number of pending requests that are allowed to the upstreamcluster. This limit is applied as a connection limit for non-HTTPtraffic.",
													MarkdownDescription: "The maximum number of pending requests that are allowed to the upstreamcluster. This limit is applied as a connection limit for non-HTTPtraffic.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_requests": schema.Int64Attribute{
													Description:         "The maximum number of parallel requests that are allowed to be madeto the upstream cluster. This limit does not apply to non-HTTP traffic.",
													MarkdownDescription: "The maximum number of parallel requests that are allowed to be madeto the upstream cluster. This limit does not apply to non-HTTP traffic.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of parallel retries that will be allowed tothe upstream cluster.",
													MarkdownDescription: "The maximum number of parallel retries that will be allowed tothe upstream cluster.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"outlier_detection": schema.SingleNestedAttribute{
											Description:         "OutlierDetection contains the configuration of the process of dynamicallydetermining whether some number of hosts in an upstream cluster areperforming unlike the others and removing them from the healthy loadbalancing set. Performance might be along different axes such asconsecutive failures, temporal success rate, temporal latency, etc.Outlier detection is a form of passive health checking.",
											MarkdownDescription: "OutlierDetection contains the configuration of the process of dynamicallydetermining whether some number of hosts in an upstream cluster areperforming unlike the others and removing them from the healthy loadbalancing set. Performance might be along different axes such asconsecutive failures, temporal success rate, temporal latency, etc.Outlier detection is a form of passive health checking.",
											Attributes: map[string]schema.Attribute{
												"base_ejection_time": schema.StringAttribute{
													Description:         "The base time that a host is ejected for. The real time is equal tothe base time multiplied by the number of times the host has beenejected.",
													MarkdownDescription: "The base time that a host is ejected for. The real time is equal tothe base time multiplied by the number of times the host has beenejected.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"detectors": schema.SingleNestedAttribute{
													Description:         "Contains configuration for supported outlier detectors",
													MarkdownDescription: "Contains configuration for supported outlier detectors",
													Attributes: map[string]schema.Attribute{
														"failure_percentage": schema.SingleNestedAttribute{
															Description:         "Failure Percentage based outlier detection functions similarly to successrate detection, in that it relies on success rate data from each host ina cluster. However, rather than compare those values to the mean successrate of the cluster as a whole, they are compared to a flatuser-configured threshold. This threshold is configured via theoutlierDetection.failurePercentageThreshold field.The other configuration fields for failure percentage based detection aresimilar to the fields for success rate detection. As with success ratedetection, detection will not be performed for a host if its requestvolume over the aggregation interval is less than theoutlierDetection.detectors.failurePercentage.requestVolume value.Detection also will not be performed for a cluster if the number of hostswith the minimum required request volume in an interval is less than theoutlierDetection.detectors.failurePercentage.minimumHosts value.",
															MarkdownDescription: "Failure Percentage based outlier detection functions similarly to successrate detection, in that it relies on success rate data from each host ina cluster. However, rather than compare those values to the mean successrate of the cluster as a whole, they are compared to a flatuser-configured threshold. This threshold is configured via theoutlierDetection.failurePercentageThreshold field.The other configuration fields for failure percentage based detection aresimilar to the fields for success rate detection. As with success ratedetection, detection will not be performed for a host if its requestvolume over the aggregation interval is less than theoutlierDetection.detectors.failurePercentage.requestVolume value.Detection also will not be performed for a cluster if the number of hostswith the minimum required request volume in an interval is less than theoutlierDetection.detectors.failurePercentage.minimumHosts value.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The minimum number of hosts in a cluster in order to perform failurepercentage-based ejection. If the total number of hosts in the cluster isless than this value, failure percentage-based ejection will not beperformed.",
																	MarkdownDescription: "The minimum number of hosts in a cluster in order to perform failurepercentage-based ejection. If the total number of hosts in the cluster isless than this value, failure percentage-based ejection will not beperformed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration above) to perform failurepercentage-based ejection for this host. If the volume is lower than thissetting, failure percentage-based ejection will not be performed for thishost.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration above) to perform failurepercentage-based ejection for this host. If the volume is lower than thissetting, failure percentage-based ejection will not be performed for thishost.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"threshold": schema.Int64Attribute{
																	Description:         "The failure percentage to use when determining failure percentage-basedoutlier detection. If the failure percentage of a given host is greaterthan or equal to this value, it will be ejected.",
																	MarkdownDescription: "The failure percentage to use when determining failure percentage-basedoutlier detection. If the failure percentage of a given host is greaterthan or equal to this value, it will be ejected.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"gateway_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalLocalOriginErrors isfalse) this detection type takes into account a subset of 5xx errors,called 'gateway errors' (502, 503 or 504 status code) and local originfailures, such as timeout, TCP reset etc.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account a subset of 5xx errors, called'gateway errors' (502, 503 or 504 status code) and is supported only bythe http router.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalLocalOriginErrors isfalse) this detection type takes into account a subset of 5xx errors,called 'gateway errors' (502, 503 or 504 status code) and local originfailures, such as timeout, TCP reset etc.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account a subset of 5xx errors, called'gateway errors' (502, 503 or 504 status code) and is supported only bythe http router.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive gateway failures (502, 503, 504 status codes)before a consecutive gateway failure ejection occurs.",
																	MarkdownDescription: "The number of consecutive gateway failures (502, 503, 504 status codes)before a consecutive gateway failure ejection occurs.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"local_origin_failures": schema.SingleNestedAttribute{
															Description:         "This detection type is enabled only whenoutlierDetection.splitExternalLocalOriginErrors is true and takes intoaccount only locally originated errors (timeout, reset, etc).If Envoy repeatedly cannot connect to an upstream host or communicationwith the upstream host is repeatedly interrupted, it will be ejected.Various locally originated problems are detected: timeout, TCP reset,ICMP errors, etc. This detection type is supported by http router andtcp proxy.",
															MarkdownDescription: "This detection type is enabled only whenoutlierDetection.splitExternalLocalOriginErrors is true and takes intoaccount only locally originated errors (timeout, reset, etc).If Envoy repeatedly cannot connect to an upstream host or communicationwith the upstream host is repeatedly interrupted, it will be ejected.Various locally originated problems are detected: timeout, TCP reset,ICMP errors, etc. This detection type is supported by http router andtcp proxy.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive locally originated failures before ejectionoccurs. Parameter takes effect only when splitExternalAndLocalErrorsis set to true.",
																	MarkdownDescription: "The number of consecutive locally originated failures before ejectionoccurs. Parameter takes effect only when splitExternalAndLocalErrorsis set to true.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"success_rate": schema.SingleNestedAttribute{
															Description:         "Success Rate based outlier detection aggregates success rate data fromevery host in a cluster. Then at given intervals ejects hosts based onstatistical outlier detection. Success Rate outlier detection will not becalculated for a host if its request volume over the aggregation intervalis less than the outlierDetection.detectors.successRate.requestVolumevalue.Moreover, detection will not be performed for a cluster if the number ofhosts with the minimum required request volume in an interval is lessthan the outlierDetection.detectors.successRate.minimumHosts value.In the default configuration mode(outlierDetection.splitExternalLocalOriginErrors is false) this detectiontype takes into account all types of errors: locally and externallyoriginated.In split mode (outlierDetection.splitExternalLocalOriginErrors is true),locally originated errors and externally originated (transaction) errorsare counted and treated separately.",
															MarkdownDescription: "Success Rate based outlier detection aggregates success rate data fromevery host in a cluster. Then at given intervals ejects hosts based onstatistical outlier detection. Success Rate outlier detection will not becalculated for a host if its request volume over the aggregation intervalis less than the outlierDetection.detectors.successRate.requestVolumevalue.Moreover, detection will not be performed for a cluster if the number ofhosts with the minimum required request volume in an interval is lessthan the outlierDetection.detectors.successRate.minimumHosts value.In the default configuration mode(outlierDetection.splitExternalLocalOriginErrors is false) this detectiontype takes into account all types of errors: locally and externallyoriginated.In split mode (outlierDetection.splitExternalLocalOriginErrors is true),locally originated errors and externally originated (transaction) errorsare counted and treated separately.",
															Attributes: map[string]schema.Attribute{
																"minimum_hosts": schema.Int64Attribute{
																	Description:         "The number of hosts in a cluster that must have enough request volume todetect success rate outliers. If the number of hosts is less than thissetting, outlier detection via success rate statistics is not performedfor any host in the cluster.",
																	MarkdownDescription: "The number of hosts in a cluster that must have enough request volume todetect success rate outliers. If the number of hosts is less than thissetting, outlier detection via success rate statistics is not performedfor any host in the cluster.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"request_volume": schema.Int64Attribute{
																	Description:         "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration configured inoutlierDetection section) to include this host in success rate basedoutlier detection. If the volume is lower than this setting, outlierdetection via success rate statistics is not performed for that host.",
																	MarkdownDescription: "The minimum number of total requests that must be collected in oneinterval (as defined by the interval duration configured inoutlierDetection section) to include this host in success rate basedoutlier detection. If the volume is lower than this setting, outlierdetection via success rate statistics is not performed for that host.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"standard_deviation_factor": schema.StringAttribute{
																	Description:         "This factor is used to determine the ejection threshold for success rateoutlier ejection. The ejection threshold is the difference betweenthe mean success rate, and the product of this factor and the standarddeviation of the mean success rate: mean - (standard_deviation *success_rate_standard_deviation_factor).Either int or decimal represented as string.",
																	MarkdownDescription: "This factor is used to determine the ejection threshold for success rateoutlier ejection. The ejection threshold is the difference betweenthe mean success rate, and the product of this factor and the standarddeviation of the mean success rate: mean - (standard_deviation *success_rate_standard_deviation_factor).Either int or decimal represented as string.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"total_failures": schema.SingleNestedAttribute{
															Description:         "In the default mode (outlierDetection.splitExternalAndLocalErrors isfalse) this detection type takes into account all generated errors:locally originated and externally originated (transaction) errors.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account only externally originated(transaction) errors, ignoring locally originated errors.If an upstream host is an HTTP-server, only 5xx types of error are takeninto account (see Consecutive Gateway Failure for exceptions).Properly formatted responses, even when they carry an operational error(like index not found, access denied) are not taken into account.",
															MarkdownDescription: "In the default mode (outlierDetection.splitExternalAndLocalErrors isfalse) this detection type takes into account all generated errors:locally originated and externally originated (transaction) errors.In split mode (outlierDetection.splitExternalLocalOriginErrors is true)this detection type takes into account only externally originated(transaction) errors, ignoring locally originated errors.If an upstream host is an HTTP-server, only 5xx types of error are takeninto account (see Consecutive Gateway Failure for exceptions).Properly formatted responses, even when they carry an operational error(like index not found, access denied) are not taken into account.",
															Attributes: map[string]schema.Attribute{
																"consecutive": schema.Int64Attribute{
																	Description:         "The number of consecutive server-side error responses (for HTTP traffic,5xx responses; for TCP traffic, connection failures; for Redis, failureto respond PONG; etc.) before a consecutive total failure ejectionoccurs.",
																	MarkdownDescription: "The number of consecutive server-side error responses (for HTTP traffic,5xx responses; for TCP traffic, connection failures; for Redis, failureto respond PONG; etc.) before a consecutive total failure ejectionoccurs.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"disabled": schema.BoolAttribute{
													Description:         "When set to true, outlierDetection configuration won't take any effect",
													MarkdownDescription: "When set to true, outlierDetection configuration won't take any effect",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "The time interval between ejection analysis sweeps. This can result inboth new ejections and hosts being returned to service.",
													MarkdownDescription: "The time interval between ejection analysis sweeps. This can result inboth new ejections and hosts being returned to service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_ejection_percent": schema.Int64Attribute{
													Description:         "The maximum % of an upstream cluster that can be ejected due to outlierdetection. Defaults to 10% but will eject at least one host regardless ofthe value.",
													MarkdownDescription: "The maximum % of an upstream cluster that can be ejected due to outlierdetection. Defaults to 10% but will eject at least one host regardless ofthe value.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"split_external_and_local_errors": schema.BoolAttribute{
													Description:         "Determines whether to distinguish local origin failures from externalerrors. If set to true the following configuration parameters are takeninto account: detectors.localOriginFailures.consecutive",
													MarkdownDescription: "Determines whether to distinguish local origin failures from externalerrors. If set to true the following configuration parameters are takeninto account: detectors.localOriginFailures.consecutive",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group ofdestinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofdestinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
											},
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_types": schema.ListAttribute{
											Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
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
		},
	}
}

func (r *KumaIoMeshCircuitBreakerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_circuit_breaker_v1alpha1_manifest")

	var model KumaIoMeshCircuitBreakerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshCircuitBreaker")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
