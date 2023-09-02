/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TestsTestkubeIoTestV2Manifest{}
)

func NewTestsTestkubeIoTestV2Manifest() datasource.DataSource {
	return &TestsTestkubeIoTestV2Manifest{}
}

type TestsTestkubeIoTestV2Manifest struct{}

type TestsTestkubeIoTestV2ManifestData struct {
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
		Content *struct {
			Data       *string `tfsdk:"data" json:"data,omitempty"`
			Repository *struct {
				Branch   *string `tfsdk:"branch" json:"branch,omitempty"`
				Commit   *string `tfsdk:"commit" json:"commit,omitempty"`
				Path     *string `tfsdk:"path" json:"path,omitempty"`
				Token    *string `tfsdk:"token" json:"token,omitempty"`
				Type     *string `tfsdk:"type" json:"type,omitempty"`
				Uri      *string `tfsdk:"uri" json:"uri,omitempty"`
				Username *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"repository" json:"repository,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
			Uri  *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"content" json:"content,omitempty"`
		ExecutorArgs *[]string          `tfsdk:"executor_args" json:"executorArgs,omitempty"`
		Name         *string            `tfsdk:"name" json:"name,omitempty"`
		Params       *map[string]string `tfsdk:"params" json:"params,omitempty"`
		Schedule     *string            `tfsdk:"schedule" json:"schedule,omitempty"`
		Type         *string            `tfsdk:"type" json:"type,omitempty"`
		Variables    *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Type      *string `tfsdk:"type" json:"type,omitempty"`
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
		} `tfsdk:"variables" json:"variables,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_v2_manifest"
}

func (r *TestsTestkubeIoTestV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Test is the Schema for the tests API",
		MarkdownDescription: "Test is the Schema for the tests API",
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
				Description:         "TestSpec defines the desired state of Test",
				MarkdownDescription: "TestSpec defines the desired state of Test",
				Attributes: map[string]schema.Attribute{
					"content": schema.SingleNestedAttribute{
						Description:         "test content object",
						MarkdownDescription: "test content object",
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Description:         "test content body",
								MarkdownDescription: "test content body",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.SingleNestedAttribute{
								Description:         "repository of test content",
								MarkdownDescription: "repository of test content",
								Attributes: map[string]schema.Attribute{
									"branch": schema.StringAttribute{
										Description:         "branch/tag name for checkout",
										MarkdownDescription: "branch/tag name for checkout",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"commit": schema.StringAttribute{
										Description:         "commit id (sha) for checkout",
										MarkdownDescription: "commit id (sha) for checkout",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
										MarkdownDescription: "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token": schema.StringAttribute{
										Description:         "git auth token for private repositories",
										MarkdownDescription: "git auth token for private repositories",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "VCS repository type",
										MarkdownDescription: "VCS repository type",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"uri": schema.StringAttribute{
										Description:         "uri of content file or git directory",
										MarkdownDescription: "uri of content file or git directory",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"username": schema.StringAttribute{
										Description:         "git auth username for private repositories",
										MarkdownDescription: "git auth username for private repositories",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "test type",
								MarkdownDescription: "test type",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uri": schema.StringAttribute{
								Description:         "uri of test content",
								MarkdownDescription: "uri of test content",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"executor_args": schema.ListAttribute{
						Description:         "additional executor binary arguments",
						MarkdownDescription: "additional executor binary arguments",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "test execution custom name",
						MarkdownDescription: "test execution custom name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"params": schema.MapAttribute{
						Description:         "DEPRECATED execution params passed to executor",
						MarkdownDescription: "DEPRECATED execution params passed to executor",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schedule": schema.StringAttribute{
						Description:         "schedule in cron job format for scheduled test execution",
						MarkdownDescription: "schedule in cron job format for scheduled test execution",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "test type",
						MarkdownDescription: "test type",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"variables": schema.SingleNestedAttribute{
						Description:         "Variables are new params with secrets attached",
						MarkdownDescription: "Variables are new params with secrets attached",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "variable name",
								MarkdownDescription: "variable name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "variable type",
								MarkdownDescription: "variable type",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"value": schema.StringAttribute{
								Description:         "variable string value",
								MarkdownDescription: "variable string value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"value_from": schema.SingleNestedAttribute{
								Description:         "or load it from var source",
								MarkdownDescription: "or load it from var source",
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

func (r *TestsTestkubeIoTestV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tests_testkube_io_test_v2_manifest")

	var model TestsTestkubeIoTestV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("tests.testkube.io/v2")
	model.Kind = pointer.String("Test")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
