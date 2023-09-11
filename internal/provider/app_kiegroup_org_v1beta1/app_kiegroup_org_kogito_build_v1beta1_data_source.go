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
	_ datasource.DataSource              = &AppKiegroupOrgKogitoBuildV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &AppKiegroupOrgKogitoBuildV1Beta1DataSource{}
)

func NewAppKiegroupOrgKogitoBuildV1Beta1DataSource() datasource.DataSource {
	return &AppKiegroupOrgKogitoBuildV1Beta1DataSource{}
}

type AppKiegroupOrgKogitoBuildV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppKiegroupOrgKogitoBuildV1Beta1DataSourceData struct {
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
		Artifact *struct {
			ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
			GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
			Version    *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"artifact" json:"artifact,omitempty"`
		BuildImage                *string `tfsdk:"build_image" json:"buildImage,omitempty"`
		DisableIncremental        *bool   `tfsdk:"disable_incremental" json:"disableIncremental,omitempty"`
		EnableMavenDownloadOutput *bool   `tfsdk:"enable_maven_download_output" json:"enableMavenDownloadOutput,omitempty"`
		Env                       *[]struct {
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
		} `tfsdk:"env" json:"env,omitempty"`
		GitSource *struct {
			ContextDir *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
			Reference  *string `tfsdk:"reference" json:"reference,omitempty"`
			Uri        *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"git_source" json:"gitSource,omitempty"`
		MavenMirrorURL *string `tfsdk:"maven_mirror_url" json:"mavenMirrorURL,omitempty"`
		Native         *bool   `tfsdk:"native" json:"native,omitempty"`
		Resources      *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Runtime             *string `tfsdk:"runtime" json:"runtime,omitempty"`
		RuntimeImage        *string `tfsdk:"runtime_image" json:"runtimeImage,omitempty"`
		TargetKogitoRuntime *string `tfsdk:"target_kogito_runtime" json:"targetKogitoRuntime,omitempty"`
		Type                *string `tfsdk:"type" json:"type,omitempty"`
		WebHooks            *[]struct {
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"web_hooks" json:"webHooks,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppKiegroupOrgKogitoBuildV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_kiegroup_org_kogito_build_v1beta1"
}

func (r *AppKiegroupOrgKogitoBuildV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
		MarkdownDescription: "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
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
				Description:         "KogitoBuildSpec defines the desired state of KogitoBuild.",
				MarkdownDescription: "KogitoBuildSpec defines the desired state of KogitoBuild.",
				Attributes: map[string]schema.Attribute{
					"artifact": schema.SingleNestedAttribute{
						Description:         "Artifact contains override information for building the Maven artifact (used for Local Source builds).  You might want to override this information when building from decisions, rules or process files. In this scenario the Kogito Images will generate a new Java project for you underneath. This information will be used to generate this project.",
						MarkdownDescription: "Artifact contains override information for building the Maven artifact (used for Local Source builds).  You might want to override this information when building from decisions, rules or process files. In this scenario the Kogito Images will generate a new Java project for you underneath. This information will be used to generate this project.",
						Attributes: map[string]schema.Attribute{
							"artifact_id": schema.StringAttribute{
								Description:         "Indicates the unique base name of the primary artifact being generated.",
								MarkdownDescription: "Indicates the unique base name of the primary artifact being generated.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"group_id": schema.StringAttribute{
								Description:         "Indicates the unique identifier of the organization or group that created the project.",
								MarkdownDescription: "Indicates the unique identifier of the organization or group that created the project.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"version": schema.StringAttribute{
								Description:         "Indicates the version of the artifact generated by the project.",
								MarkdownDescription: "Indicates the version of the artifact generated by the project.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"build_image": schema.StringAttribute{
						Description:         "Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						MarkdownDescription: "Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disable_incremental": schema.BoolAttribute{
						Description:         "DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.",
						MarkdownDescription: "DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_maven_download_output": schema.BoolAttribute{
						Description:         "If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.",
						MarkdownDescription: "If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"env": schema.ListNestedAttribute{
						Description:         "Environment variables used during build time.",
						MarkdownDescription: "Environment variables used during build time.",
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

					"git_source": schema.SingleNestedAttribute{
						Description:         "Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds.",
						MarkdownDescription: "Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds.",
						Attributes: map[string]schema.Attribute{
							"context_dir": schema.StringAttribute{
								Description:         "Context/subdirectory where the code is located, relative to the repo root.",
								MarkdownDescription: "Context/subdirectory where the code is located, relative to the repo root.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"reference": schema.StringAttribute{
								Description:         "Branch to use in the Git repository.",
								MarkdownDescription: "Branch to use in the Git repository.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uri": schema.StringAttribute{
								Description:         "Git URI for the s2i source.",
								MarkdownDescription: "Git URI for the s2i source.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"maven_mirror_url": schema.StringAttribute{
						Description:         "Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.",
						MarkdownDescription: "Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"native": schema.BoolAttribute{
						Description:         "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.",
						MarkdownDescription: "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources Requirements for builder pods.",
						MarkdownDescription: "Resources Requirements for builder pods.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"runtime": schema.StringAttribute{
						Description:         "Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.",
						MarkdownDescription: "Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"runtime_image": schema.StringAttribute{
						Description:         "Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						MarkdownDescription: "Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_kogito_runtime": schema.StringAttribute{
						Description:         "Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.",
						MarkdownDescription: "Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"type": schema.StringAttribute{
						Description:         "Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.",
						MarkdownDescription: "Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"web_hooks": schema.ListNestedAttribute{
						Description:         "WebHooks secrets for source to image builds based on Git repositories (Remote Sources).",
						MarkdownDescription: "WebHooks secrets for source to image builds based on Git repositories (Remote Sources).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"secret": schema.StringAttribute{
									Description:         "Secret value for webHook",
									MarkdownDescription: "Secret value for webHook",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "WebHook type, either GitHub or Generic.",
									MarkdownDescription: "WebHook type, either GitHub or Generic.",
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

func (r *AppKiegroupOrgKogitoBuildV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *AppKiegroupOrgKogitoBuildV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_app_kiegroup_org_kogito_build_v1beta1")

	var data AppKiegroupOrgKogitoBuildV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "app.kiegroup.org", Version: "v1beta1", Resource: "kogitobuilds"}).
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

	var readResponse AppKiegroupOrgKogitoBuildV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("app.kiegroup.org/v1beta1")
	data.Kind = pointer.String("KogitoBuild")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
