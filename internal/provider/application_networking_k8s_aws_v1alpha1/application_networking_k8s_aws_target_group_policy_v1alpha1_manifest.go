/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package application_networking_k8s_aws_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest{}
)

func NewApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest() datasource.DataSource {
	return &ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest{}
}

type ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest struct{}

type ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1ManifestData struct {
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
		HealthCheck *struct {
			Enabled                 *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			HealthyThresholdCount   *int64  `tfsdk:"healthy_threshold_count" json:"healthyThresholdCount,omitempty"`
			IntervalSeconds         *int64  `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
			Path                    *string `tfsdk:"path" json:"path,omitempty"`
			Port                    *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol                *string `tfsdk:"protocol" json:"protocol,omitempty"`
			ProtocolVersion         *string `tfsdk:"protocol_version" json:"protocolVersion,omitempty"`
			StatusMatch             *string `tfsdk:"status_match" json:"statusMatch,omitempty"`
			TimeoutSeconds          *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			UnhealthyThresholdCount *int64  `tfsdk:"unhealthy_threshold_count" json:"unhealthyThresholdCount,omitempty"`
		} `tfsdk:"health_check" json:"healthCheck,omitempty"`
		Protocol        *string `tfsdk:"protocol" json:"protocol,omitempty"`
		ProtocolVersion *string `tfsdk:"protocol_version" json:"protocolVersion,omitempty"`
		TargetRef       *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_application_networking_k8s_aws_target_group_policy_v1alpha1_manifest"
}

func (r *ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "TargetGroupPolicySpec defines the desired state of TargetGroupPolicy.",
				MarkdownDescription: "TargetGroupPolicySpec defines the desired state of TargetGroupPolicy.",
				Attributes: map[string]schema.Attribute{
					"health_check": schema.SingleNestedAttribute{
						Description:         "The health check configuration.  Changes to this value will update VPC Lattice resource in place.",
						MarkdownDescription: "The health check configuration.  Changes to this value will update VPC Lattice resource in place.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Indicates whether health checking is enabled.",
								MarkdownDescription: "Indicates whether health checking is enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"healthy_threshold_count": schema.Int64Attribute{
								Description:         "The number of consecutive successful health checks required before considering an unhealthy target healthy.",
								MarkdownDescription: "The number of consecutive successful health checks required before considering an unhealthy target healthy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(2),
									int64validator.AtMost(10),
								},
							},

							"interval_seconds": schema.Int64Attribute{
								Description:         "The approximate amount of time, in seconds, between health checks of an individual target.",
								MarkdownDescription: "The approximate amount of time, in seconds, between health checks of an individual target.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(5),
									int64validator.AtMost(300),
								},
							},

							"path": schema.StringAttribute{
								Description:         "The destination for health checks on the targets.",
								MarkdownDescription: "The destination for health checks on the targets.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port used when performing health checks on targets. If not specified, health check defaults to the port that a target receives traffic on.",
								MarkdownDescription: "The port used when performing health checks on targets. If not specified, health check defaults to the port that a target receives traffic on.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"protocol": schema.StringAttribute{
								Description:         "The protocol used when performing health checks on targets.",
								MarkdownDescription: "The protocol used when performing health checks on targets.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("HTTP", "HTTPS"),
								},
							},

							"protocol_version": schema.StringAttribute{
								Description:         "The protocol version used when performing health checks on targets.",
								MarkdownDescription: "The protocol version used when performing health checks on targets.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("HTTP1", "HTTP2"),
								},
							},

							"status_match": schema.StringAttribute{
								Description:         "A regular expression to match HTTP status codes when checking for successful response from a target.",
								MarkdownDescription: "A regular expression to match HTTP status codes when checking for successful response from a target.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
								Description:         "The amount of time, in seconds, to wait before reporting a target as unhealthy.",
								MarkdownDescription: "The amount of time, in seconds, to wait before reporting a target as unhealthy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(120),
								},
							},

							"unhealthy_threshold_count": schema.Int64Attribute{
								Description:         "The number of consecutive failed health checks required before considering a target unhealthy.",
								MarkdownDescription: "The number of consecutive failed health checks required before considering a target unhealthy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(2),
									int64validator.AtMost(10),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol": schema.StringAttribute{
						Description:         "The protocol to use for routing traffic to the targets. Supported values are HTTP (default), HTTPS and TCP.  Changes to this value results in a replacement of VPC Lattice target group.",
						MarkdownDescription: "The protocol to use for routing traffic to the targets. Supported values are HTTP (default), HTTPS and TCP.  Changes to this value results in a replacement of VPC Lattice target group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"protocol_version": schema.StringAttribute{
						Description:         "The protocol version to use. Supported values are HTTP1 (default) and HTTP2. When a policy Protocol is TCP, you should not set this field. Otherwise, the whole TargetGroupPolicy will not take effect. When a policy is behind GRPCRoute, this field value will be ignored as GRPC is only supported through HTTP/2.  Changes to this value results in a replacement of VPC Lattice target group.",
						MarkdownDescription: "The protocol version to use. Supported values are HTTP1 (default) and HTTP2. When a policy Protocol is TCP, you should not set this field. Otherwise, the whole TargetGroupPolicy will not take effect. When a policy is behind GRPCRoute, this field value will be ignored as GRPC is only supported through HTTP/2.  Changes to this value results in a replacement of VPC Lattice target group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef points to the kubernetes Service resource that will have this policy attached.  This field is following the guidelines of Kubernetes Gateway API policy attachment.",
						MarkdownDescription: "TargetRef points to the kubernetes Service resource that will have this policy attached.  This field is following the guidelines of Kubernetes Gateway API policy attachment.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the target resource.",
								MarkdownDescription: "Group is the group of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the target resource.",
								MarkdownDescription: "Kind is kind of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the target resource.",
								MarkdownDescription: "Name is the name of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the referent. When unspecified, the local namespace is inferred. Even when policy targets a resource in a different namespace, it MUST only apply to traffic originating from the same namespace as the policy.",
								MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, the local namespace is inferred. Even when policy targets a resource in a different namespace, it MUST only apply to traffic originating from the same namespace as the policy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_application_networking_k8s_aws_target_group_policy_v1alpha1_manifest")

	var model ApplicationNetworkingK8SAwsTargetGroupPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("application-networking.k8s.aws/v1alpha1")
	model.Kind = pointer.String("TargetGroupPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
