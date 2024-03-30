/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_kiegroup_org_v1beta1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppKiegroupOrgKogitoBuildV1Beta1Manifest{}
)

func NewAppKiegroupOrgKogitoBuildV1Beta1Manifest() datasource.DataSource {
	return &AppKiegroupOrgKogitoBuildV1Beta1Manifest{}
}

type AppKiegroupOrgKogitoBuildV1Beta1Manifest struct{}

type AppKiegroupOrgKogitoBuildV1Beta1ManifestData struct {
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

func (r *AppKiegroupOrgKogitoBuildV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_kiegroup_org_kogito_build_v1beta1_manifest"
}

func (r *AppKiegroupOrgKogitoBuildV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
		MarkdownDescription: "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
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
								Optional:            true,
								Computed:            false,
							},

							"group_id": schema.StringAttribute{
								Description:         "Indicates the unique identifier of the organization or group that created the project.",
								MarkdownDescription: "Indicates the unique identifier of the organization or group that created the project.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Indicates the version of the artifact generated by the project.",
								MarkdownDescription: "Indicates the version of the artifact generated by the project.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"build_image": schema.StringAttribute{
						Description:         "Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						MarkdownDescription: "Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_incremental": schema.BoolAttribute{
						Description:         "DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.",
						MarkdownDescription: "DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_maven_download_output": schema.BoolAttribute{
						Description:         "If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.",
						MarkdownDescription: "If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.ListNestedAttribute{
						Description:         "Environment variables used during build time.",
						MarkdownDescription: "Environment variables used during build time.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
									MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
									Required:            false,
									Optional:            true,
									Computed:            false,
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

										"field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "Path of the field to select in the specified API version.",
													MarkdownDescription: "Path of the field to select in the specified API version.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resource_field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											Attributes: map[string]schema.Attribute{
												"container_name": schema.StringAttribute{
													Description:         "Container name: required for volumes, optional for env vars",
													MarkdownDescription: "Container name: required for volumes, optional for env vars",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"divisor": schema.StringAttribute{
													Description:         "Specifies the output format of the exposed resources, defaults to '1'",
													MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource": schema.StringAttribute{
													Description:         "Required: resource to select",
													MarkdownDescription: "Required: resource to select",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a secret in the pod's namespace",
											MarkdownDescription: "Selects a key of a secret in the pod's namespace",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
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
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"git_source": schema.SingleNestedAttribute{
						Description:         "Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds.",
						MarkdownDescription: "Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds.",
						Attributes: map[string]schema.Attribute{
							"context_dir": schema.StringAttribute{
								Description:         "Context/subdirectory where the code is located, relative to the repo root.",
								MarkdownDescription: "Context/subdirectory where the code is located, relative to the repo root.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reference": schema.StringAttribute{
								Description:         "Branch to use in the Git repository.",
								MarkdownDescription: "Branch to use in the Git repository.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uri": schema.StringAttribute{
								Description:         "Git URI for the s2i source.",
								MarkdownDescription: "Git URI for the s2i source.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"maven_mirror_url": schema.StringAttribute{
						Description:         "Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.",
						MarkdownDescription: "Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"native": schema.BoolAttribute{
						Description:         "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.",
						MarkdownDescription: "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"runtime": schema.StringAttribute{
						Description:         "Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.",
						MarkdownDescription: "Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("quarkus", "springboot"),
						},
					},

					"runtime_image": schema.StringAttribute{
						Description:         "Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						MarkdownDescription: "Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_kogito_runtime": schema.StringAttribute{
						Description:         "Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.",
						MarkdownDescription: "Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.",
						MarkdownDescription: "Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Binary", "RemoteSource", "LocalSource"),
						},
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
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "WebHook type, either GitHub or Generic.",
									MarkdownDescription: "WebHook type, either GitHub or Generic.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("GitHub", "Generic"),
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

func (r *AppKiegroupOrgKogitoBuildV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_kiegroup_org_kogito_build_v1beta1_manifest")

	var model AppKiegroupOrgKogitoBuildV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("app.kiegroup.org/v1beta1")
	model.Kind = pointer.String("KogitoBuild")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
