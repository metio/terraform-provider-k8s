/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type AppKiegroupOrgKogitoInfraV1Beta1Resource struct{}

var (
	_ resource.Resource = (*AppKiegroupOrgKogitoInfraV1Beta1Resource)(nil)
)

type AppKiegroupOrgKogitoInfraV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppKiegroupOrgKogitoInfraV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		SecretEnvFromReferences *[]string `tfsdk:"secret_env_from_references" yaml:"secretEnvFromReferences,omitempty"`

		SecretVolumeReferences *[]struct {
			FileMode *int64 `tfsdk:"file_mode" yaml:"fileMode,omitempty"`

			MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"secret_volume_references" yaml:"secretVolumeReferences,omitempty"`

		ConfigMapEnvFromReferences *[]string `tfsdk:"config_map_env_from_references" yaml:"configMapEnvFromReferences,omitempty"`

		ConfigMapVolumeReferences *[]struct {
			FileMode *int64 `tfsdk:"file_mode" yaml:"fileMode,omitempty"`

			MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"config_map_volume_references" yaml:"configMapVolumeReferences,omitempty"`

		Envs *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`

			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

					FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

				ResourceFieldRef *struct {
					Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`

					ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

					Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`
				} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

				SecretKeyRef *struct {
					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
		} `tfsdk:"envs" yaml:"envs,omitempty"`

		InfraProperties *map[string]string `tfsdk:"infra_properties" yaml:"infraProperties,omitempty"`

		Resource *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"resource" yaml:"resource,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppKiegroupOrgKogitoInfraV1Beta1Resource() resource.Resource {
	return &AppKiegroupOrgKogitoInfraV1Beta1Resource{}
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_kiegroup_org_kogito_infra_v1beta1"
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service.  It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index).  Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.",
		MarkdownDescription: "KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service.  It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index).  Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "KogitoInfraSpec defines the desired state of KogitoInfra.",
				MarkdownDescription: "KogitoInfraSpec defines the desired state of KogitoInfra.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"secret_env_from_references": {
						Description:         "List of secret that should be mounted to the services as envs",
						MarkdownDescription: "List of secret that should be mounted to the services as envs",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_volume_references": {
						Description:         "List of secret that should be munted to the services bound to this infra instance",
						MarkdownDescription: "List of secret that should be munted to the services bound to this infra instance",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"file_mode": {
								Description:         "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",
								MarkdownDescription: "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount_path": {
								Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",
								MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "This must match the Name of a ConfigMap.",
								MarkdownDescription: "This must match the Name of a ConfigMap.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"optional": {
								Description:         "Specify whether the Secret or its keys must be defined",
								MarkdownDescription: "Specify whether the Secret or its keys must be defined",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_map_env_from_references": {
						Description:         "List of secret that should be mounted to the services as envs",
						MarkdownDescription: "List of secret that should be mounted to the services as envs",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_map_volume_references": {
						Description:         "List of configmap that should be added to the services bound to this infra instance",
						MarkdownDescription: "List of configmap that should be added to the services bound to this infra instance",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"file_mode": {
								Description:         "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",
								MarkdownDescription: "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount_path": {
								Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",
								MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "This must match the Name of a ConfigMap.",
								MarkdownDescription: "This must match the Name of a ConfigMap.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"optional": {
								Description:         "Specify whether the Secret or its keys must be defined",
								MarkdownDescription: "Specify whether the Secret or its keys must be defined",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"envs": {
						Description:         "Environment variables to be added to the runtime container. Keys must be a C_IDENTIFIER.",
						MarkdownDescription: "Environment variables to be added to the runtime container. Keys must be a C_IDENTIFIER.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
								MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
								MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value_from": {
								Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
								MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map_key_ref": {
										Description:         "Selects a key of a ConfigMap.",
										MarkdownDescription: "Selects a key of a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"field_ref": {
										Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
										MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api_version": {
												Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
												MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_path": {
												Description:         "Path of the field to select in the specified API version.",
												MarkdownDescription: "Path of the field to select in the specified API version.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_field_ref": {
										Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
										MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"resource": {
												Description:         "Required: resource to select",
												MarkdownDescription: "Required: resource to select",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"container_name": {
												Description:         "Container name: required for volumes, optional for env vars",
												MarkdownDescription: "Container name: required for volumes, optional for env vars",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"divisor": {
												Description:         "Specifies the output format of the exposed resources, defaults to '1'",
												MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_key_ref": {
										Description:         "Selects a key of a secret in the pod's namespace",
										MarkdownDescription: "Selects a key of a secret in the pod's namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"infra_properties": {
						Description:         "Optional properties which would be needed to setup correct runtime/service configuration, based on the resource type.  For example, MongoDB will require 'username' and 'database' as properties for a correct setup, else it will fail",
						MarkdownDescription: "Optional properties which would be needed to setup correct runtime/service configuration, based on the resource type.  For example, MongoDB will require 'username' and 'database' as properties for a correct setup, else it will fail",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource": {
						Description:         "Resource for the service. Example: Infinispan/Kafka/Keycloak.",
						MarkdownDescription: "Resource for the service. Example: Infinispan/Kafka/Keycloak.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion describes the API Version of referred Kubernetes resource for example, infinispan.org/v1",
								MarkdownDescription: "APIVersion describes the API Version of referred Kubernetes resource for example, infinispan.org/v1",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "Kind describes the kind of referred Kubernetes resource for example, Infinispan",
								MarkdownDescription: "Kind describes the kind of referred Kubernetes resource for example, Infinispan",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of referred resource.",
								MarkdownDescription: "Name of referred resource.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace where referred resource exists.",
								MarkdownDescription: "Namespace where referred resource exists.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_app_kiegroup_org_kogito_infra_v1beta1")

	var state AppKiegroupOrgKogitoInfraV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppKiegroupOrgKogitoInfraV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.kiegroup.org/v1beta1")
	goModel.Kind = utilities.Ptr("KogitoInfra")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_kiegroup_org_kogito_infra_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_app_kiegroup_org_kogito_infra_v1beta1")

	var state AppKiegroupOrgKogitoInfraV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppKiegroupOrgKogitoInfraV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.kiegroup.org/v1beta1")
	goModel.Kind = utilities.Ptr("KogitoInfra")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_app_kiegroup_org_kogito_infra_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
