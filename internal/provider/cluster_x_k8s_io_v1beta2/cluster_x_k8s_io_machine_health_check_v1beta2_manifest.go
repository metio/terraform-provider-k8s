/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ datasource.DataSource = &ClusterXK8SIoMachineHealthCheckV1Beta2Manifest{}
)

func NewClusterXK8SIoMachineHealthCheckV1Beta2Manifest() datasource.DataSource {
	return &ClusterXK8SIoMachineHealthCheckV1Beta2Manifest{}
}

type ClusterXK8SIoMachineHealthCheckV1Beta2Manifest struct{}

type ClusterXK8SIoMachineHealthCheckV1Beta2ManifestData struct {
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
		Checks *struct {
			NodeStartupTimeoutSeconds *int64 `tfsdk:"node_startup_timeout_seconds" json:"nodeStartupTimeoutSeconds,omitempty"`
			UnhealthyNodeConditions   *[]struct {
				Status         *string `tfsdk:"status" json:"status,omitempty"`
				TimeoutSeconds *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				Type           *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"unhealthy_node_conditions" json:"unhealthyNodeConditions,omitempty"`
		} `tfsdk:"checks" json:"checks,omitempty"`
		ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		Remediation *struct {
			TemplateRef *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"template_ref" json:"templateRef,omitempty"`
			TriggerIf *struct {
				UnhealthyInRange           *string `tfsdk:"unhealthy_in_range" json:"unhealthyInRange,omitempty"`
				UnhealthyLessThanOrEqualTo *string `tfsdk:"unhealthy_less_than_or_equal_to" json:"unhealthyLessThanOrEqualTo,omitempty"`
			} `tfsdk:"trigger_if" json:"triggerIf,omitempty"`
		} `tfsdk:"remediation" json:"remediation,omitempty"`
		Selector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterXK8SIoMachineHealthCheckV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_machine_health_check_v1beta2_manifest"
}

func (r *ClusterXK8SIoMachineHealthCheckV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachineHealthCheck is the Schema for the machinehealthchecks API.",
		MarkdownDescription: "MachineHealthCheck is the Schema for the machinehealthchecks API.",
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
				Description:         "spec is the specification of machine health check policy",
				MarkdownDescription: "spec is the specification of machine health check policy",
				Attributes: map[string]schema.Attribute{
					"checks": schema.SingleNestedAttribute{
						Description:         "checks are the checks that are used to evaluate if a Machine is healthy. Independent of this configuration the MachineHealthCheck controller will always flag Machines with 'cluster.x-k8s.io/remediate-machine' annotation and Machines with deleted Nodes as unhealthy. Furthermore, if checks.nodeStartupTimeoutSeconds is not set it is defaulted to 10 minutes and evaluated accordingly.",
						MarkdownDescription: "checks are the checks that are used to evaluate if a Machine is healthy. Independent of this configuration the MachineHealthCheck controller will always flag Machines with 'cluster.x-k8s.io/remediate-machine' annotation and Machines with deleted Nodes as unhealthy. Furthermore, if checks.nodeStartupTimeoutSeconds is not set it is defaulted to 10 minutes and evaluated accordingly.",
						Attributes: map[string]schema.Attribute{
							"node_startup_timeout_seconds": schema.Int64Attribute{
								Description:         "nodeStartupTimeoutSeconds allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
								MarkdownDescription: "nodeStartupTimeoutSeconds allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_node_conditions")),
								},
							},

							"unhealthy_node_conditions": schema.ListNestedAttribute{
								Description:         "unhealthyNodeConditions contains a list of conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
								MarkdownDescription: "unhealthyNodeConditions contains a list of conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"status": schema.StringAttribute{
											Description:         "status of the condition, one of True, False, Unknown.",
											MarkdownDescription: "status of the condition, one of True, False, Unknown.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},

										"timeout_seconds": schema.Int64Attribute{
											Description:         "timeoutSeconds is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
											MarkdownDescription: "timeoutSeconds is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"type": schema.StringAttribute{
											Description:         "type of Node condition",
											MarkdownDescription: "type of Node condition",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.List{
									listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_startup_timeout_seconds")),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "clusterName is the name of the Cluster this object belongs to.",
						MarkdownDescription: "clusterName is the name of the Cluster this object belongs to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(63),
						},
					},

					"remediation": schema.SingleNestedAttribute{
						Description:         "remediation configures if and how remediations are triggered if a Machine is unhealthy. If remediation or remediation.triggerIf is not set, remediation will always be triggered for unhealthy Machines. If remediation or remediation.templateRef is not set, the OwnerRemediated condition will be set on unhealthy Machines to trigger remediation via the owner of the Machines, for example a MachineSet or a KubeadmControlPlane.",
						MarkdownDescription: "remediation configures if and how remediations are triggered if a Machine is unhealthy. If remediation or remediation.triggerIf is not set, remediation will always be triggered for unhealthy Machines. If remediation or remediation.templateRef is not set, the OwnerRemediated condition will be set on unhealthy Machines to trigger remediation via the owner of the Machines, for example a MachineSet or a KubeadmControlPlane.",
						Attributes: map[string]schema.Attribute{
							"template_ref": schema.SingleNestedAttribute{
								Description:         "templateRef is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
								MarkdownDescription: "templateRef is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "apiVersion of the remediation template. apiVersion must be fully qualified domain name followed by / and a version. NOTE: This field must be kept in sync with the APIVersion of the remediation template.",
										MarkdownDescription: "apiVersion of the remediation template. apiVersion must be fully qualified domain name followed by / and a version. NOTE: This field must be kept in sync with the APIVersion of the remediation template.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(317),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*\/[a-z]([-a-z0-9]*[a-z0-9])?$`), ""),
										},
									},

									"kind": schema.StringAttribute{
										Description:         "kind of the remediation template. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
										MarkdownDescription: "kind of the remediation template. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
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
										Description:         "name of the remediation template. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
										MarkdownDescription: "name of the remediation template. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("trigger_if")),
								},
							},

							"trigger_if": schema.SingleNestedAttribute{
								Description:         "triggerIf configures if remediations are triggered. If this field is not set, remediations are always triggered.",
								MarkdownDescription: "triggerIf configures if remediations are triggered. If this field is not set, remediations are always triggered.",
								Attributes: map[string]schema.Attribute{
									"unhealthy_in_range": schema.StringAttribute{
										Description:         "unhealthyInRange specifies that remediations are only triggered if the number of unhealthy Machines is in the configured range. Takes precedence over unhealthyLessThanOrEqualTo. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy Machines (and) (b) there are at most 5 unhealthy Machines",
										MarkdownDescription: "unhealthyInRange specifies that remediations are only triggered if the number of unhealthy Machines is in the configured range. Takes precedence over unhealthyLessThanOrEqualTo. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy Machines (and) (b) there are at most 5 unhealthy Machines",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(32),
											stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
											stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_less_than_or_equal_to")),
										},
									},

									"unhealthy_less_than_or_equal_to": schema.StringAttribute{
										Description:         "unhealthyLessThanOrEqualTo specifies that remediations are only triggered if the number of unhealthy Machines is less than or equal to the configured value. unhealthyInRange takes precedence if set.",
										MarkdownDescription: "unhealthyLessThanOrEqualTo specifies that remediations are only triggered if the number of unhealthy Machines is less than or equal to the configured value. unhealthyInRange takes precedence if set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_in_range")),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("template_ref")),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "selector is a label selector to match machines whose health will be exercised",
						MarkdownDescription: "selector is a label selector to match machines whose health will be exercised",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ClusterXK8SIoMachineHealthCheckV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_x_k8s_io_machine_health_check_v1beta2_manifest")

	var model ClusterXK8SIoMachineHealthCheckV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("MachineHealthCheck")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
