/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package remediation_medik8s_io_v1alpha1

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
	_ datasource.DataSource = &RemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest{}
)

func NewRemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest() datasource.DataSource {
	return &RemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest{}
}

type RemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest struct{}

type RemediationMedik8SIoNodeHealthCheckV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		EscalatingRemediations *[]struct {
			Order               *int64 `tfsdk:"order" json:"order,omitempty"`
			RemediationTemplate *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"remediation_template" json:"remediationTemplate,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"escalating_remediations" json:"escalatingRemediations,omitempty"`
		MinHealthy          *string   `tfsdk:"min_healthy" json:"minHealthy,omitempty"`
		PauseRequests       *[]string `tfsdk:"pause_requests" json:"pauseRequests,omitempty"`
		RemediationTemplate *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"remediation_template" json:"remediationTemplate,omitempty"`
		Selector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		UnhealthyConditions *[]struct {
			Duration *string `tfsdk:"duration" json:"duration,omitempty"`
			Status   *string `tfsdk:"status" json:"status,omitempty"`
			Type     *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"unhealthy_conditions" json:"unhealthyConditions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_remediation_medik8s_io_node_health_check_v1alpha1_manifest"
}

func (r *RemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeHealthCheck is the Schema for the nodehealthchecks API",
		MarkdownDescription: "NodeHealthCheck is the Schema for the nodehealthchecks API",
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
				Description:         "NodeHealthCheckSpec defines the desired state of NodeHealthCheck",
				MarkdownDescription: "NodeHealthCheckSpec defines the desired state of NodeHealthCheck",
				Attributes: map[string]schema.Attribute{
					"escalating_remediations": schema.ListNestedAttribute{
						Description:         "EscalatingRemediations contain a list of ordered remediation templates with a timeout.The remediation templates will be used one after another, until the unhealthy nodegets healthy within the timeout of the currently processed remediation. The order ofremediation is defined by the 'order' field of each 'escalatingRemediation'.Mutually exclusive with RemediationTemplate",
						MarkdownDescription: "EscalatingRemediations contain a list of ordered remediation templates with a timeout.The remediation templates will be used one after another, until the unhealthy nodegets healthy within the timeout of the currently processed remediation. The order ofremediation is defined by the 'order' field of each 'escalatingRemediation'.Mutually exclusive with RemediationTemplate",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"order": schema.Int64Attribute{
									Description:         "Order defines the order for this remediation.Remediations with lower order will be used before remediations with higher order.Remediations must not have the same order.",
									MarkdownDescription: "Order defines the order for this remediation.Remediations with lower order will be used before remediations with higher order.Remediations must not have the same order.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"remediation_template": schema.SingleNestedAttribute{
									Description:         "RemediationTemplate is a reference to a remediation templateprovided by a remediation provider.If a node needs remediation the controller will create an object from this templateand then it should be picked up by a remediation provider.",
									MarkdownDescription: "RemediationTemplate is a reference to a remediation templateprovided by a remediation provider.If a node needs remediation the controller will create an object from this templateand then it should be picked up by a remediation provider.",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"field_path": schema.StringAttribute{
											Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
											MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
											MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_version": schema.StringAttribute{
											Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
											MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uid": schema.StringAttribute{
											Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
											MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"timeout": schema.StringAttribute{
									Description:         "Timeout defines how long NHC will wait for the node getting healthybefore the next remediation (if any) will be used. When the last remediation times out,the overall remediation is considered as failed.As a safeguard for preventing parallel remediations, a minimum of 60s is enforced.Expects a string of decimal numbers each with optionalfraction and a unit suffix, eg '300ms', '1.5h' or '2h45m'.Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
									MarkdownDescription: "Timeout defines how long NHC will wait for the node getting healthybefore the next remediation (if any) will be used. When the last remediation times out,the overall remediation is considered as failed.As a safeguard for preventing parallel remediations, a minimum of 60s is enforced.Expects a string of decimal numbers each with optionalfraction and a unit suffix, eg '300ms', '1.5h' or '2h45m'.Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_healthy": schema.StringAttribute{
						Description:         "Remediation is allowed if at least 'MinHealthy' nodes selected by 'selector' are healthy.Expects either a positive integer value or a percentage value.Percentage values must be positive whole numbers and are capped at 100%.100% is valid and will block all remediation.",
						MarkdownDescription: "Remediation is allowed if at least 'MinHealthy' nodes selected by 'selector' are healthy.Expects either a positive integer value or a percentage value.Percentage values must be positive whole numbers and are capped at 100%.100% is valid and will block all remediation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pause_requests": schema.ListAttribute{
						Description:         "PauseRequests will prevent any new remediation to start, while in-flight remediationskeep running. Each entry is free form, and ideally represents the requested party reasonfor this pausing - i.e:    'imaginary-cluster-upgrade-manager-operator'",
						MarkdownDescription: "PauseRequests will prevent any new remediation to start, while in-flight remediationskeep running. Each entry is free form, and ideally represents the requested party reasonfor this pausing - i.e:    'imaginary-cluster-upgrade-manager-operator'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remediation_template": schema.SingleNestedAttribute{
						Description:         "RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.If a node needs remediation the controller will create an object from this templateand then it should be picked up by a remediation provider.Mutually exclusive with EscalatingRemediations",
						MarkdownDescription: "RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.If a node needs remediation the controller will create an object from this templateand then it should be picked up by a remediation provider.Mutually exclusive with EscalatingRemediations",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Label selector to match nodes whose health will be exercised.Selecting both control-plane and worker nodes in one NHC CR ishighly discouraged and can result in undesired behaviour.Note: mandatory now for above reason, but for backwards compatibility existingCRs will continue to work with an empty selector, which matches all nodes.",
						MarkdownDescription: "Label selector to match nodes whose health will be exercised.Selecting both control-plane and worker nodes in one NHC CR ishighly discouraged and can result in undesired behaviour.Note: mandatory now for above reason, but for backwards compatibility existingCRs will continue to work with an empty selector, which matches all nodes.",
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
											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"unhealthy_conditions": schema.ListNestedAttribute{
						Description:         "UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy.  The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy.",
						MarkdownDescription: "UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy.  The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"duration": schema.StringAttribute{
									Description:         "Duration of the condition specified when a node is considered unhealthy.Expects a string of decimal numbers each with optionalfraction and a unit suffix, eg '300ms', '1.5h' or '2h45m'.Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
									MarkdownDescription: "Duration of the condition specified when a node is considered unhealthy.Expects a string of decimal numbers each with optionalfraction and a unit suffix, eg '300ms', '1.5h' or '2h45m'.Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
									},
								},

								"status": schema.StringAttribute{
									Description:         "The condition status in the node's status to watch for.Typically False, True or Unknown.",
									MarkdownDescription: "The condition status in the node's status to watch for.Typically False, True or Unknown.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"type": schema.StringAttribute{
									Description:         "The condition type in the node's status to watch for.",
									MarkdownDescription: "The condition type in the node's status to watch for.",
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
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *RemediationMedik8SIoNodeHealthCheckV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_remediation_medik8s_io_node_health_check_v1alpha1_manifest")

	var model RemediationMedik8SIoNodeHealthCheckV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("remediation.medik8s.io/v1alpha1")
	model.Kind = pointer.String("NodeHealthCheck")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
