/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type AppKiegroupOrgKogitoBuildV1Beta1Resource struct{}

var (
	_ resource.Resource = (*AppKiegroupOrgKogitoBuildV1Beta1Resource)(nil)
)

type AppKiegroupOrgKogitoBuildV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppKiegroupOrgKogitoBuildV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Artifact *struct {
			ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

			GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"artifact" yaml:"artifact,omitempty"`

		BuildImage *string `tfsdk:"build_image" yaml:"buildImage,omitempty"`

		DisableIncremental *bool `tfsdk:"disable_incremental" yaml:"disableIncremental,omitempty"`

		EnableMavenDownloadOutput *bool `tfsdk:"enable_maven_download_output" yaml:"enableMavenDownloadOutput,omitempty"`

		Env *[]struct {
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
					ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

					Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

					Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
				} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

				SecretKeyRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
		} `tfsdk:"env" yaml:"env,omitempty"`

		GitSource *struct {
			ContextDir *string `tfsdk:"context_dir" yaml:"contextDir,omitempty"`

			Reference *string `tfsdk:"reference" yaml:"reference,omitempty"`

			Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
		} `tfsdk:"git_source" yaml:"gitSource,omitempty"`

		MavenMirrorURL *string `tfsdk:"maven_mirror_url" yaml:"mavenMirrorURL,omitempty"`

		Native *bool `tfsdk:"native" yaml:"native,omitempty"`

		Resources *struct {
			Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

			Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		Runtime *string `tfsdk:"runtime" yaml:"runtime,omitempty"`

		RuntimeImage *string `tfsdk:"runtime_image" yaml:"runtimeImage,omitempty"`

		TargetKogitoRuntime *string `tfsdk:"target_kogito_runtime" yaml:"targetKogitoRuntime,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`

		WebHooks *[]struct {
			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"web_hooks" yaml:"webHooks,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppKiegroupOrgKogitoBuildV1Beta1Resource() resource.Resource {
	return &AppKiegroupOrgKogitoBuildV1Beta1Resource{}
}

func (r *AppKiegroupOrgKogitoBuildV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_kiegroup_org_kogito_build_v1beta1"
}

func (r *AppKiegroupOrgKogitoBuildV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
		MarkdownDescription: "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
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
				Description:         "KogitoBuildSpec defines the desired state of KogitoBuild.",
				MarkdownDescription: "KogitoBuildSpec defines the desired state of KogitoBuild.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"artifact": {
						Description:         "Artifact contains override information for building the Maven artifact (used for Local Source builds).  You might want to override this information when building from decisions, rules or process files. In this scenario the Kogito Images will generate a new Java project for you underneath. This information will be used to generate this project.",
						MarkdownDescription: "Artifact contains override information for building the Maven artifact (used for Local Source builds).  You might want to override this information when building from decisions, rules or process files. In this scenario the Kogito Images will generate a new Java project for you underneath. This information will be used to generate this project.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"artifact_id": {
								Description:         "Indicates the unique base name of the primary artifact being generated.",
								MarkdownDescription: "Indicates the unique base name of the primary artifact being generated.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_id": {
								Description:         "Indicates the unique identifier of the organization or group that created the project.",
								MarkdownDescription: "Indicates the unique identifier of the organization or group that created the project.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Indicates the version of the artifact generated by the project.",
								MarkdownDescription: "Indicates the version of the artifact generated by the project.",

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

					"build_image": {
						Description:         "Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						MarkdownDescription: "Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_incremental": {
						Description:         "DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.",
						MarkdownDescription: "DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_maven_download_output": {
						Description:         "If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.",
						MarkdownDescription: "If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"env": {
						Description:         "Environment variables used during build time.",
						MarkdownDescription: "Environment variables used during build time.",

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

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource": {
												Description:         "Required: resource to select",
												MarkdownDescription: "Required: resource to select",

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

									"secret_key_ref": {
										Description:         "Selects a key of a secret in the pod's namespace",
										MarkdownDescription: "Selects a key of a secret in the pod's namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

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

					"git_source": {
						Description:         "Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds.",
						MarkdownDescription: "Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"context_dir": {
								Description:         "Context/subdirectory where the code is located, relative to the repo root.",
								MarkdownDescription: "Context/subdirectory where the code is located, relative to the repo root.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"reference": {
								Description:         "Branch to use in the Git repository.",
								MarkdownDescription: "Branch to use in the Git repository.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"uri": {
								Description:         "Git URI for the s2i source.",
								MarkdownDescription: "Git URI for the s2i source.",

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

					"maven_mirror_url": {
						Description:         "Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.",
						MarkdownDescription: "Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"native": {
						Description:         "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.",
						MarkdownDescription: "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Resources Requirements for builder pods.",
						MarkdownDescription: "Resources Requirements for builder pods.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"limits": {
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"requests": {
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"runtime": {
						Description:         "Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.",
						MarkdownDescription: "Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("quarkus", "springboot"),
						},
					},

					"runtime_image": {
						Description:         "Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
						MarkdownDescription: "Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_kogito_runtime": {
						Description:         "Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.",
						MarkdownDescription: "Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": {
						Description:         "Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.",
						MarkdownDescription: "Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Binary", "RemoteSource", "LocalSource"),
						},
					},

					"web_hooks": {
						Description:         "WebHooks secrets for source to image builds based on Git repositories (Remote Sources).",
						MarkdownDescription: "WebHooks secrets for source to image builds based on Git repositories (Remote Sources).",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"secret": {
								Description:         "Secret value for webHook",
								MarkdownDescription: "Secret value for webHook",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "WebHook type, either GitHub or Generic.",
								MarkdownDescription: "WebHook type, either GitHub or Generic.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("GitHub", "Generic"),
								},
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

func (r *AppKiegroupOrgKogitoBuildV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_app_kiegroup_org_kogito_build_v1beta1")

	var state AppKiegroupOrgKogitoBuildV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppKiegroupOrgKogitoBuildV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.kiegroup.org/v1beta1")
	goModel.Kind = utilities.Ptr("KogitoBuild")

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

func (r *AppKiegroupOrgKogitoBuildV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_kiegroup_org_kogito_build_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppKiegroupOrgKogitoBuildV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_app_kiegroup_org_kogito_build_v1beta1")

	var state AppKiegroupOrgKogitoBuildV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppKiegroupOrgKogitoBuildV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.kiegroup.org/v1beta1")
	goModel.Kind = utilities.Ptr("KogitoBuild")

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

func (r *AppKiegroupOrgKogitoBuildV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_app_kiegroup_org_kogito_build_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
