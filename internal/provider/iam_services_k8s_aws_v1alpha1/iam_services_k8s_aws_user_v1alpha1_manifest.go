/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package iam_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &IamServicesK8SAwsUserV1Alpha1Manifest{}
)

func NewIamServicesK8SAwsUserV1Alpha1Manifest() datasource.DataSource {
	return &IamServicesK8SAwsUserV1Alpha1Manifest{}
}

type IamServicesK8SAwsUserV1Alpha1Manifest struct{}

type IamServicesK8SAwsUserV1Alpha1ManifestData struct {
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
		InlinePolicies         *map[string]string `tfsdk:"inline_policies" json:"inlinePolicies,omitempty"`
		Name                   *string            `tfsdk:"name" json:"name,omitempty"`
		Path                   *string            `tfsdk:"path" json:"path,omitempty"`
		PermissionsBoundary    *string            `tfsdk:"permissions_boundary" json:"permissionsBoundary,omitempty"`
		PermissionsBoundaryRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"permissions_boundary_ref" json:"permissionsBoundaryRef,omitempty"`
		Policies   *[]string `tfsdk:"policies" json:"policies,omitempty"`
		PolicyRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"policy_refs" json:"policyRefs,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *IamServicesK8SAwsUserV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_iam_services_k8s_aws_user_v1alpha1_manifest"
}

func (r *IamServicesK8SAwsUserV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "User is the Schema for the Users API",
		MarkdownDescription: "User is the Schema for the Users API",
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
				Description:         "UserSpec defines the desired state of User.Contains information about an IAM user entity.This data type is used as a response element in the following operations:   * CreateUser   * GetUser   * ListUsers",
				MarkdownDescription: "UserSpec defines the desired state of User.Contains information about an IAM user entity.This data type is used as a response element in the following operations:   * CreateUser   * GetUser   * ListUsers",
				Attributes: map[string]schema.Attribute{
					"inline_policies": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the user to create.IAM user, group, role, and policy names must be unique within the account.Names are not distinguished by case. For example, you cannot create resourcesnamed both 'MyResource' and 'myresource'.",
						MarkdownDescription: "The name of the user to create.IAM user, group, role, and policy names must be unique within the account.Names are not distinguished by case. For example, you cannot create resourcesnamed both 'MyResource' and 'myresource'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "The path for the user name. For more information about paths, see IAM identifiers(https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)in the IAM User Guide.This parameter is optional. If it is not included, it defaults to a slash(/).This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))a string of characters consisting of either a forward slash (/) by itselfor a string that must begin and end with forward slashes. In addition, itcan contain any ASCII character from the ! (u0021) through the DEL character(u007F), including most punctuation characters, digits, and upper and lowercasedletters.",
						MarkdownDescription: "The path for the user name. For more information about paths, see IAM identifiers(https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)in the IAM User Guide.This parameter is optional. If it is not included, it defaults to a slash(/).This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))a string of characters consisting of either a forward slash (/) by itselfor a string that must begin and end with forward slashes. In addition, itcan contain any ASCII character from the ! (u0021) through the DEL character(u007F), including most punctuation characters, digits, and upper and lowercasedletters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"permissions_boundary": schema.StringAttribute{
						Description:         "The ARN of the managed policy that is used to set the permissions boundaryfor the user.A permissions boundary policy defines the maximum permissions that identity-basedpolicies can grant to an entity, but does not grant permissions. Permissionsboundaries do not define the maximum permissions that a resource-based policycan grant to an entity. To learn more, see Permissions boundaries for IAMentities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)in the IAM User Guide.For more information about policy types, see Policy types (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types)in the IAM User Guide.",
						MarkdownDescription: "The ARN of the managed policy that is used to set the permissions boundaryfor the user.A permissions boundary policy defines the maximum permissions that identity-basedpolicies can grant to an entity, but does not grant permissions. Permissionsboundaries do not define the maximum permissions that a resource-based policycan grant to an entity. To learn more, see Permissions boundaries for IAMentities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)in the IAM User Guide.For more information about policy types, see Policy types (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types)in the IAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"permissions_boundary_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"policies": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of tags that you want to attach to the new user. Each tag consistsof a key name and an associated value. For more information about tagging,see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)in the IAM User Guide.If any one of the tags is invalid or if you exceed the allowed maximum numberof tags, then the entire request fails and the resource is not created.",
						MarkdownDescription: "A list of tags that you want to attach to the new user. Each tag consistsof a key name and an associated value. For more information about tagging,see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)in the IAM User Guide.If any one of the tags is invalid or if you exceed the allowed maximum numberof tags, then the entire request fails and the resource is not created.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *IamServicesK8SAwsUserV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_iam_services_k8s_aws_user_v1alpha1_manifest")

	var model IamServicesK8SAwsUserV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("iam.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("User")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
