/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_kiegroup_org_v1beta1

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
	_ datasource.DataSource              = &AppKiegroupOrgKogitoInfraV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &AppKiegroupOrgKogitoInfraV1Beta1DataSource{}
)

func NewAppKiegroupOrgKogitoInfraV1Beta1DataSource() datasource.DataSource {
	return &AppKiegroupOrgKogitoInfraV1Beta1DataSource{}
}

type AppKiegroupOrgKogitoInfraV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppKiegroupOrgKogitoInfraV1Beta1DataSourceData struct {
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
		ConfigMapEnvFromReferences *[]string `tfsdk:"config_map_env_from_references" json:"configMapEnvFromReferences,omitempty"`
		ConfigMapVolumeReferences  *[]struct {
			FileMode  *int64  `tfsdk:"file_mode" json:"fileMode,omitempty"`
			MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Optional  *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"config_map_volume_references" json:"configMapVolumeReferences,omitempty"`
		Envs *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
				ResourceFieldRef *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
					Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
				} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
				SecretKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"envs" json:"envs,omitempty"`
		InfraProperties *map[string]string `tfsdk:"infra_properties" json:"infraProperties,omitempty"`
		Resource        *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"resource" json:"resource,omitempty"`
		SecretEnvFromReferences *[]string `tfsdk:"secret_env_from_references" json:"secretEnvFromReferences,omitempty"`
		SecretVolumeReferences  *[]struct {
			FileMode  *int64  `tfsdk:"file_mode" json:"fileMode,omitempty"`
			MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Optional  *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"secret_volume_references" json:"secretVolumeReferences,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_kiegroup_org_kogito_infra_v1beta1"
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service.  It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index).  Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.",
		MarkdownDescription: "KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service.  It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index).  Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "KogitoInfraSpec defines the desired state of KogitoInfra.",
				MarkdownDescription: "KogitoInfraSpec defines the desired state of KogitoInfra.",
				Attributes: map[string]schema.Attribute{
					"config_map_env_from_references": schema.ListAttribute{
						Description:         "List of secret that should be mounted to the services as envs",
						MarkdownDescription: "List of secret that should be mounted to the services as envs",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"config_map_volume_references": schema.ListNestedAttribute{
						Description:         "List of configmap that should be added to the services bound to this infra instance",
						MarkdownDescription: "List of configmap that should be added to the services bound to this infra instance",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"file_mode": schema.Int64Attribute{
									Description:         "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",
									MarkdownDescription: "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"mount_path": schema.StringAttribute{
									Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",
									MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "This must match the Name of a ConfigMap.",
									MarkdownDescription: "This must match the Name of a ConfigMap.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"optional": schema.BoolAttribute{
									Description:         "Specify whether the Secret or its keys must be defined",
									MarkdownDescription: "Specify whether the Secret or its keys must be defined",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"envs": schema.ListNestedAttribute{
						Description:         "Environment variables to be added to the runtime container. Keys must be a C_IDENTIFIER.",
						MarkdownDescription: "Environment variables to be added to the runtime container. Keys must be a C_IDENTIFIER.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
									MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
									MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a ConfigMap.",
											MarkdownDescription: "Selects a key of a ConfigMap.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap or its key must be defined",
													MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"field_path": schema.StringAttribute{
													Description:         "Path of the field to select in the specified API version.",
													MarkdownDescription: "Path of the field to select in the specified API version.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"resource_field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											Attributes: map[string]schema.Attribute{
												"container_name": schema.StringAttribute{
													Description:         "Container name: required for volumes, optional for env vars",
													MarkdownDescription: "Container name: required for volumes, optional for env vars",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"divisor": schema.StringAttribute{
													Description:         "Specifies the output format of the exposed resources, defaults to '1'",
													MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"resource": schema.StringAttribute{
													Description:         "Required: resource to select",
													MarkdownDescription: "Required: resource to select",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a secret in the pod's namespace",
											MarkdownDescription: "Selects a key of a secret in the pod's namespace",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"infra_properties": schema.MapAttribute{
						Description:         "Optional properties which would be needed to setup correct runtime/service configuration, based on the resource type.  For example, MongoDB will require 'username' and 'database' as properties for a correct setup, else it will fail",
						MarkdownDescription: "Optional properties which would be needed to setup correct runtime/service configuration, based on the resource type.  For example, MongoDB will require 'username' and 'database' as properties for a correct setup, else it will fail",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource": schema.SingleNestedAttribute{
						Description:         "Resource for the service. Example: Infinispan/Kafka/Keycloak.",
						MarkdownDescription: "Resource for the service. Example: Infinispan/Kafka/Keycloak.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion describes the API Version of referred Kubernetes resource for example, infinispan.org/v1",
								MarkdownDescription: "APIVersion describes the API Version of referred Kubernetes resource for example, infinispan.org/v1",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind describes the kind of referred Kubernetes resource for example, Infinispan",
								MarkdownDescription: "Kind describes the kind of referred Kubernetes resource for example, Infinispan",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of referred resource.",
								MarkdownDescription: "Name of referred resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace where referred resource exists.",
								MarkdownDescription: "Namespace where referred resource exists.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"secret_env_from_references": schema.ListAttribute{
						Description:         "List of secret that should be mounted to the services as envs",
						MarkdownDescription: "List of secret that should be mounted to the services as envs",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"secret_volume_references": schema.ListNestedAttribute{
						Description:         "List of secret that should be munted to the services bound to this infra instance",
						MarkdownDescription: "List of secret that should be munted to the services bound to this infra instance",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"file_mode": schema.Int64Attribute{
									Description:         "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",
									MarkdownDescription: "Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"mount_path": schema.StringAttribute{
									Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",
									MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'. Default mount path is /home/kogito/config",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "This must match the Name of a ConfigMap.",
									MarkdownDescription: "This must match the Name of a ConfigMap.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"optional": schema.BoolAttribute{
									Description:         "Specify whether the Secret or its keys must be defined",
									MarkdownDescription: "Specify whether the Secret or its keys must be defined",
									Required:            false,
									Optional:            false,
									Computed:            true,
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

func (r *AppKiegroupOrgKogitoInfraV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *AppKiegroupOrgKogitoInfraV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_app_kiegroup_org_kogito_infra_v1beta1")

	var data AppKiegroupOrgKogitoInfraV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "app.kiegroup.org", Version: "v1beta1", Resource: "kogitoinfras"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AppKiegroupOrgKogitoInfraV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("app.kiegroup.org/v1beta1")
	data.Kind = pointer.String("KogitoInfra")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
