/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appmesh_k8s_aws_v1beta2

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
	_ datasource.DataSource = &AppmeshK8SAwsMeshV1Beta2Manifest{}
)

func NewAppmeshK8SAwsMeshV1Beta2Manifest() datasource.DataSource {
	return &AppmeshK8SAwsMeshV1Beta2Manifest{}
}

type AppmeshK8SAwsMeshV1Beta2Manifest struct{}

type AppmeshK8SAwsMeshV1Beta2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AwsName      *string `tfsdk:"aws_name" json:"awsName,omitempty"`
		EgressFilter *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"egress_filter" json:"egressFilter,omitempty"`
		MeshOwner            *string `tfsdk:"mesh_owner" json:"meshOwner,omitempty"`
		MeshServiceDiscovery *struct {
			IpPreference *string `tfsdk:"ip_preference" json:"ipPreference,omitempty"`
		} `tfsdk:"mesh_service_discovery" json:"meshServiceDiscovery,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppmeshK8SAwsMeshV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appmesh_k8s_aws_mesh_v1beta2_manifest"
}

func (r *AppmeshK8SAwsMeshV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Mesh is the Schema for the meshes API",
		MarkdownDescription: "Mesh is the Schema for the meshes API",
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
				Description:         "MeshSpec defines the desired state of Mesh refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_MeshSpec.html",
				MarkdownDescription: "MeshSpec defines the desired state of Mesh refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_MeshSpec.html",
				Attributes: map[string]schema.Attribute{
					"aws_name": schema.StringAttribute{
						Description:         "AWSName is the AppMesh Mesh object's name. If unspecified or empty, it defaults to be '${name}' of k8s Mesh",
						MarkdownDescription: "AWSName is the AppMesh Mesh object's name. If unspecified or empty, it defaults to be '${name}' of k8s Mesh",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"egress_filter": schema.SingleNestedAttribute{
						Description:         "The egress filter rules for the service mesh. If unspecified, default settings from AWS API will be applied. Refer to AWS Docs for default settings.",
						MarkdownDescription: "The egress filter rules for the service mesh. If unspecified, default settings from AWS API will be applied. Refer to AWS Docs for default settings.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "The egress filter type.",
								MarkdownDescription: "The egress filter type.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ALLOW_ALL", "DROP_ALL"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"mesh_owner": schema.StringAttribute{
						Description:         "The AWS IAM account ID of the service mesh owner. Required if the account ID is not your own.",
						MarkdownDescription: "The AWS IAM account ID of the service mesh owner. Required if the account ID is not your own.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mesh_service_discovery": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"ip_preference": schema.StringAttribute{
								Description:         "The ipPreference for the mesh.",
								MarkdownDescription: "The ipPreference for the mesh.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IPv6_ONLY", "IPv4_ONLY"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "NamespaceSelector selects Namespaces using labels to designate mesh membership. This field follows standard label selector semantics: 	if present but empty, it selects all namespaces. 	if absent, it selects no namespace.",
						MarkdownDescription: "NamespaceSelector selects Namespaces using labels to designate mesh membership. This field follows standard label selector semantics: 	if present but empty, it selects all namespaces. 	if absent, it selects no namespace.",
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

func (r *AppmeshK8SAwsMeshV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appmesh_k8s_aws_mesh_v1beta2_manifest")

	var model AppmeshK8SAwsMeshV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("appmesh.k8s.aws/v1beta2")
	model.Kind = pointer.String("Mesh")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
