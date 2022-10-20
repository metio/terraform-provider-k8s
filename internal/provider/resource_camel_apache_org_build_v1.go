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

type CamelApacheOrgBuildV1Resource struct{}

var (
	_ resource.Resource = (*CamelApacheOrgBuildV1Resource)(nil)
)

type CamelApacheOrgBuildV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CamelApacheOrgBuildV1GoModel struct {
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
		Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

		Tasks *[]struct {
			Buildah *struct {
				BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

				ContextDir *string `tfsdk:"context_dir" yaml:"contextDir,omitempty"`

				ExecutorImage *string `tfsdk:"executor_image" yaml:"executorImage,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Platform *string `tfsdk:"platform" yaml:"platform,omitempty"`

				Registry *struct {
					Address *string `tfsdk:"address" yaml:"address,omitempty"`

					Ca *string `tfsdk:"ca" yaml:"ca,omitempty"`

					Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

					Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

					Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"registry" yaml:"registry,omitempty"`

				Verbose *bool `tfsdk:"verbose" yaml:"verbose,omitempty"`
			} `tfsdk:"buildah" yaml:"buildah,omitempty"`

			Builder *struct {
				BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

				BuildDir *string `tfsdk:"build_dir" yaml:"buildDir,omitempty"`

				Dependencies *[]string `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

				Maven *struct {
					CaSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"ca_secret" yaml:"caSecret,omitempty"`

					CaSecrets *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"ca_secrets" yaml:"caSecrets,omitempty"`

					CliOptions *[]string `tfsdk:"cli_options" yaml:"cliOptions,omitempty"`

					Extension *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

						GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"extension" yaml:"extension,omitempty"`

					LocalRepository *string `tfsdk:"local_repository" yaml:"localRepository,omitempty"`

					Properties *map[string]string `tfsdk:"properties" yaml:"properties,omitempty"`

					Repositories *[]struct {
						Id *string `tfsdk:"id" yaml:"id,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Releases *struct {
							ChecksumPolicy *string `tfsdk:"checksum_policy" yaml:"checksumPolicy,omitempty"`

							Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

							UpdatePolicy *string `tfsdk:"update_policy" yaml:"updatePolicy,omitempty"`
						} `tfsdk:"releases" yaml:"releases,omitempty"`

						Snapshots *struct {
							ChecksumPolicy *string `tfsdk:"checksum_policy" yaml:"checksumPolicy,omitempty"`

							Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

							UpdatePolicy *string `tfsdk:"update_policy" yaml:"updatePolicy,omitempty"`
						} `tfsdk:"snapshots" yaml:"snapshots,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"repositories" yaml:"repositories,omitempty"`

					Servers *[]struct {
						Configuration *map[string]string `tfsdk:"configuration" yaml:"configuration,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"servers" yaml:"servers,omitempty"`

					Settings *struct {
						ConfigMapKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
					} `tfsdk:"settings" yaml:"settings,omitempty"`

					SettingsSecurity *struct {
						ConfigMapKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
					} `tfsdk:"settings_security" yaml:"settingsSecurity,omitempty"`
				} `tfsdk:"maven" yaml:"maven,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Resources *[]struct {
					Compression *bool `tfsdk:"compression" yaml:"compression,omitempty"`

					Content *string `tfsdk:"content" yaml:"content,omitempty"`

					ContentKey *string `tfsdk:"content_key" yaml:"contentKey,omitempty"`

					ContentRef *string `tfsdk:"content_ref" yaml:"contentRef,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					RawContent *string `tfsdk:"raw_content" yaml:"rawContent,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				Runtime *struct {
					ApplicationClass *string `tfsdk:"application_class" yaml:"applicationClass,omitempty"`

					Capabilities *struct {
						Dependencies *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

							GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

							Version *string `tfsdk:"version" yaml:"version,omitempty"`
						} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

						Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

						GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

					Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"runtime" yaml:"runtime,omitempty"`

				Sources *[]struct {
					Compression *bool `tfsdk:"compression" yaml:"compression,omitempty"`

					Content *string `tfsdk:"content" yaml:"content,omitempty"`

					ContentKey *string `tfsdk:"content_key" yaml:"contentKey,omitempty"`

					ContentRef *string `tfsdk:"content_ref" yaml:"contentRef,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					Interceptors *[]string `tfsdk:"interceptors" yaml:"interceptors,omitempty"`

					Language *string `tfsdk:"language" yaml:"language,omitempty"`

					Loader *string `tfsdk:"loader" yaml:"loader,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Property_names *[]string `tfsdk:"property__names" yaml:"property-names,omitempty"`

					RawContent *string `tfsdk:"raw_content" yaml:"rawContent,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"sources" yaml:"sources,omitempty"`

				Steps *[]string `tfsdk:"steps" yaml:"steps,omitempty"`
			} `tfsdk:"builder" yaml:"builder,omitempty"`

			Kaniko *struct {
				BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

				Cache *struct {
					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					PersistentVolumeClaim *string `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`
				} `tfsdk:"cache" yaml:"cache,omitempty"`

				ContextDir *string `tfsdk:"context_dir" yaml:"contextDir,omitempty"`

				ExecutorImage *string `tfsdk:"executor_image" yaml:"executorImage,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Registry *struct {
					Address *string `tfsdk:"address" yaml:"address,omitempty"`

					Ca *string `tfsdk:"ca" yaml:"ca,omitempty"`

					Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

					Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

					Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"registry" yaml:"registry,omitempty"`

				Verbose *bool `tfsdk:"verbose" yaml:"verbose,omitempty"`
			} `tfsdk:"kaniko" yaml:"kaniko,omitempty"`

			S2i *struct {
				ContextDir *string `tfsdk:"context_dir" yaml:"contextDir,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`
			} `tfsdk:"s2i" yaml:"s2i,omitempty"`

			Spectrum *struct {
				BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

				ContextDir *string `tfsdk:"context_dir" yaml:"contextDir,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Registry *struct {
					Address *string `tfsdk:"address" yaml:"address,omitempty"`

					Ca *string `tfsdk:"ca" yaml:"ca,omitempty"`

					Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

					Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

					Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"registry" yaml:"registry,omitempty"`
			} `tfsdk:"spectrum" yaml:"spectrum,omitempty"`
		} `tfsdk:"tasks" yaml:"tasks,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCamelApacheOrgBuildV1Resource() resource.Resource {
	return &CamelApacheOrgBuildV1Resource{}
}

func (r *CamelApacheOrgBuildV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camel_apache_org_build_v1"
}

func (r *CamelApacheOrgBuildV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Build is the Schema for the builds API",
		MarkdownDescription: "Build is the Schema for the builds API",
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
				Description:         "BuildSpec defines the Build operation to be executed",
				MarkdownDescription: "BuildSpec defines the Build operation to be executed",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"strategy": {
						Description:         "The strategy that should be used to perform the Build.",
						MarkdownDescription: "The strategy that should be used to perform the Build.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("routine", "pod"),
						},
					},

					"tasks": {
						Description:         "The sequence of Build tasks to be performed as part of the Build execution.",
						MarkdownDescription: "The sequence of Build tasks to be performed as part of the Build execution.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"buildah": {
								Description:         "a BuildahTask, for Buildah strategy",
								MarkdownDescription: "a BuildahTask, for Buildah strategy",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_image": {
										Description:         "base image layer",
										MarkdownDescription: "base image layer",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"context_dir": {
										Description:         "can be useful to share info with other tasks",
										MarkdownDescription: "can be useful to share info with other tasks",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"executor_image": {
										Description:         "docker image to use",
										MarkdownDescription: "docker image to use",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "final image name",
										MarkdownDescription: "final image name",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "name of the task",
										MarkdownDescription: "name of the task",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"platform": {
										Description:         "The platform of build image",
										MarkdownDescription: "The platform of build image",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"registry": {
										Description:         "where to publish the final image",
										MarkdownDescription: "where to publish the final image",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"address": {
												Description:         "the URI to access",
												MarkdownDescription: "the URI to access",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ca": {
												Description:         "the configmap which stores the Certificate Authority",
												MarkdownDescription: "the configmap which stores the Certificate Authority",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure": {
												Description:         "if the container registry is insecure (ie, http only)",
												MarkdownDescription: "if the container registry is insecure (ie, http only)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "the registry organization",
												MarkdownDescription: "the registry organization",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "the secret where credentials are stored",
												MarkdownDescription: "the secret where credentials are stored",

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

									"verbose": {
										Description:         "log more information",
										MarkdownDescription: "log more information",

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

							"builder": {
								Description:         "a BuilderTask (base task)",
								MarkdownDescription: "a BuilderTask (base task)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_image": {
										Description:         "the base image layer",
										MarkdownDescription: "the base image layer",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"build_dir": {
										Description:         "workspace directory to use",
										MarkdownDescription: "workspace directory to use",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dependencies": {
										Description:         "the list of dependencies to use for this build",
										MarkdownDescription: "the list of dependencies to use for this build",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"maven": {
										Description:         "the configuration required by Maven for the application build phase",
										MarkdownDescription: "the configuration required by Maven for the application build phase",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_secret": {
												Description:         "Deprecated: use CASecrets The Secret name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
												MarkdownDescription: "Deprecated: use CASecrets The Secret name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",

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

											"ca_secrets": {
												Description:         "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
												MarkdownDescription: "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

											"cli_options": {
												Description:         "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
												MarkdownDescription: "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"extension": {
												Description:         "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
												MarkdownDescription: "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"artifact_id": {
														Description:         "Maven Artifact",
														MarkdownDescription: "Maven Artifact",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"group_id": {
														Description:         "Maven Group",
														MarkdownDescription: "Maven Group",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"version": {
														Description:         "Maven Version",
														MarkdownDescription: "Maven Version",

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

											"local_repository": {
												Description:         "The path of the local Maven repository.",
												MarkdownDescription: "The path of the local Maven repository.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"properties": {
												Description:         "The Maven properties.",
												MarkdownDescription: "The Maven properties.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"repositories": {
												Description:         "additional repositories",
												MarkdownDescription: "additional repositories",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"id": {
														Description:         "identifies the repository",
														MarkdownDescription: "identifies the repository",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "name of the repository",
														MarkdownDescription: "name of the repository",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"releases": {
														Description:         "can use stable releases",
														MarkdownDescription: "can use stable releases",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"checksum_policy": {
																Description:         "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																MarkdownDescription: "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"enabled": {
																Description:         "is the policy activated or not",
																MarkdownDescription: "is the policy activated or not",

																Type: types.BoolType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"update_policy": {
																Description:         "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																MarkdownDescription: "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",

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

													"snapshots": {
														Description:         "can use snapshot",
														MarkdownDescription: "can use snapshot",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"checksum_policy": {
																Description:         "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																MarkdownDescription: "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"enabled": {
																Description:         "is the policy activated or not",
																MarkdownDescription: "is the policy activated or not",

																Type: types.BoolType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"update_policy": {
																Description:         "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																MarkdownDescription: "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",

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

													"url": {
														Description:         "location of the repository",
														MarkdownDescription: "location of the repository",

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

											"servers": {
												Description:         "Servers (auth)",
												MarkdownDescription: "Servers (auth)",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"configuration": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"username": {
														Description:         "",
														MarkdownDescription: "",

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

											"settings": {
												Description:         "A reference to the ConfigMap or Secret key that contains the Maven settings.",
												MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the Maven settings.",

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

													"secret_key_ref": {
														Description:         "Selects a key of a secret.",
														MarkdownDescription: "Selects a key of a secret.",

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

											"settings_security": {
												Description:         "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",
												MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",

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

													"secret_key_ref": {
														Description:         "Selects a key of a secret.",
														MarkdownDescription: "Selects a key of a secret.",

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

									"name": {
										Description:         "name of the task",
										MarkdownDescription: "name of the task",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Deprecated: no longer in use",
										MarkdownDescription: "Deprecated: no longer in use",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"compression": {
												Description:         "if the content is compressed (base64 encrypted)",
												MarkdownDescription: "if the content is compressed (base64 encrypted)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content": {
												Description:         "the source code (plain text)",
												MarkdownDescription: "the source code (plain text)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_key": {
												Description:         "the confimap key holding the source content",
												MarkdownDescription: "the confimap key holding the source content",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_ref": {
												Description:         "the confimap reference holding the source content",
												MarkdownDescription: "the confimap reference holding the source content",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "the content type (tipically text or binary)",
												MarkdownDescription: "the content type (tipically text or binary)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mount_path": {
												Description:         "the mount path on destination 'Pod'",
												MarkdownDescription: "the mount path on destination 'Pod'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "the name of the specification",
												MarkdownDescription: "the name of the specification",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "the path where the file is stored",
												MarkdownDescription: "the path where the file is stored",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"raw_content": {
												Description:         "the source code (binary)",
												MarkdownDescription: "the source code (binary)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"type": {
												Description:         "the kind of data to expect",
												MarkdownDescription: "the kind of data to expect",

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

									"runtime": {
										Description:         "the configuration required for the runtime application",
										MarkdownDescription: "the configuration required for the runtime application",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"application_class": {
												Description:         "application entry point (main) to be executed",
												MarkdownDescription: "application entry point (main) to be executed",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"capabilities": {
												Description:         "features offered by this runtime",
												MarkdownDescription: "features offered by this runtime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dependencies": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"artifact_id": {
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"group_id": {
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"version": {
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"metadata": {
														Description:         "Deprecated: not in use",
														MarkdownDescription: "Deprecated: not in use",

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

											"dependencies": {
												Description:         "list of dependencies needed to run the application",
												MarkdownDescription: "list of dependencies needed to run the application",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"artifact_id": {
														Description:         "Maven Artifact",
														MarkdownDescription: "Maven Artifact",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"group_id": {
														Description:         "Maven Group",
														MarkdownDescription: "Maven Group",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"version": {
														Description:         "Maven Version",
														MarkdownDescription: "Maven Version",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"metadata": {
												Description:         "set of metadata",
												MarkdownDescription: "set of metadata",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"provider": {
												Description:         "Camel main application provider, ie, Camel Quarkus",
												MarkdownDescription: "Camel main application provider, ie, Camel Quarkus",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"version": {
												Description:         "Camel K Runtime version",
												MarkdownDescription: "Camel K Runtime version",

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

									"sources": {
										Description:         "Deprecated: no longer in use the source code for the Route(s)",
										MarkdownDescription: "Deprecated: no longer in use the source code for the Route(s)",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"compression": {
												Description:         "if the content is compressed (base64 encrypted)",
												MarkdownDescription: "if the content is compressed (base64 encrypted)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content": {
												Description:         "the source code (plain text)",
												MarkdownDescription: "the source code (plain text)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_key": {
												Description:         "the confimap key holding the source content",
												MarkdownDescription: "the confimap key holding the source content",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_ref": {
												Description:         "the confimap reference holding the source content",
												MarkdownDescription: "the confimap reference holding the source content",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "the content type (tipically text or binary)",
												MarkdownDescription: "the content type (tipically text or binary)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"interceptors": {
												Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
												MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"language": {
												Description:         "specify which is the language (Camel DSL) used to interpret this source code",
												MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"loader": {
												Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
												MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "the name of the specification",
												MarkdownDescription: "the name of the specification",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "the path where the file is stored",
												MarkdownDescription: "the path where the file is stored",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"property__names": {
												Description:         "List of property names defined in the source (e.g. if type is 'template')",
												MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"raw_content": {
												Description:         "the source code (binary)",
												MarkdownDescription: "the source code (binary)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"type": {
												Description:         "Type defines the kind of source described by this object",
												MarkdownDescription: "Type defines the kind of source described by this object",

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

									"steps": {
										Description:         "the list of steps to execute (see pkg/builder/)",
										MarkdownDescription: "the list of steps to execute (see pkg/builder/)",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kaniko": {
								Description:         "a KanikoTask, for Kaniko strategy",
								MarkdownDescription: "a KanikoTask, for Kaniko strategy",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_image": {
										Description:         "base image layer",
										MarkdownDescription: "base image layer",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache": {
										Description:         "use a cache",
										MarkdownDescription: "use a cache",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enabled": {
												Description:         "true if a cache is enabled",
												MarkdownDescription: "true if a cache is enabled",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"persistent_volume_claim": {
												Description:         "the PVC used to store the cache",
												MarkdownDescription: "the PVC used to store the cache",

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

									"context_dir": {
										Description:         "can be useful to share info with other tasks",
										MarkdownDescription: "can be useful to share info with other tasks",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"executor_image": {
										Description:         "docker image to use",
										MarkdownDescription: "docker image to use",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "final image name",
										MarkdownDescription: "final image name",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "name of the task",
										MarkdownDescription: "name of the task",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"registry": {
										Description:         "where to publish the final image",
										MarkdownDescription: "where to publish the final image",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"address": {
												Description:         "the URI to access",
												MarkdownDescription: "the URI to access",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ca": {
												Description:         "the configmap which stores the Certificate Authority",
												MarkdownDescription: "the configmap which stores the Certificate Authority",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure": {
												Description:         "if the container registry is insecure (ie, http only)",
												MarkdownDescription: "if the container registry is insecure (ie, http only)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "the registry organization",
												MarkdownDescription: "the registry organization",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "the secret where credentials are stored",
												MarkdownDescription: "the secret where credentials are stored",

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

									"verbose": {
										Description:         "log more information",
										MarkdownDescription: "log more information",

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

							"s2i": {
								Description:         "a S2iTask, for S2I strategy",
								MarkdownDescription: "a S2iTask, for S2I strategy",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"context_dir": {
										Description:         "can be useful to share info with other tasks",
										MarkdownDescription: "can be useful to share info with other tasks",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "name of the task",
										MarkdownDescription: "name of the task",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tag": {
										Description:         "used by the ImageStream",
										MarkdownDescription: "used by the ImageStream",

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

							"spectrum": {
								Description:         "a SpectrumTask, for Spectrum strategy",
								MarkdownDescription: "a SpectrumTask, for Spectrum strategy",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_image": {
										Description:         "base image layer",
										MarkdownDescription: "base image layer",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"context_dir": {
										Description:         "can be useful to share info with other tasks",
										MarkdownDescription: "can be useful to share info with other tasks",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "final image name",
										MarkdownDescription: "final image name",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "name of the task",
										MarkdownDescription: "name of the task",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"registry": {
										Description:         "where to publish the final image",
										MarkdownDescription: "where to publish the final image",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"address": {
												Description:         "the URI to access",
												MarkdownDescription: "the URI to access",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ca": {
												Description:         "the configmap which stores the Certificate Authority",
												MarkdownDescription: "the configmap which stores the Certificate Authority",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure": {
												Description:         "if the container registry is insecure (ie, http only)",
												MarkdownDescription: "if the container registry is insecure (ie, http only)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "the registry organization",
												MarkdownDescription: "the registry organization",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "the secret where credentials are stored",
												MarkdownDescription: "the secret where credentials are stored",

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

					"timeout": {
						Description:         "Timeout defines the Build maximum execution duration. The Build deadline is set to the Build start time plus the Timeout duration. If the Build deadline is exceeded, the Build context is canceled, and its phase set to BuildPhaseFailed.",
						MarkdownDescription: "Timeout defines the Build maximum execution duration. The Build deadline is set to the Build start time plus the Timeout duration. If the Build deadline is exceeded, the Build context is canceled, and its phase set to BuildPhaseFailed.",

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
		},
	}, nil
}

func (r *CamelApacheOrgBuildV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_build_v1")

	var state CamelApacheOrgBuildV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgBuildV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("Build")

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

func (r *CamelApacheOrgBuildV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_build_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CamelApacheOrgBuildV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_build_v1")

	var state CamelApacheOrgBuildV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgBuildV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("Build")

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

func (r *CamelApacheOrgBuildV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_build_v1")
	// NO-OP: Terraform removes the state automatically for us
}
