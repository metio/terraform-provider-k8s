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
	_ datasource.DataSource = &IamServicesK8SAwsPolicyV1Alpha1Manifest{}
)

func NewIamServicesK8SAwsPolicyV1Alpha1Manifest() datasource.DataSource {
	return &IamServicesK8SAwsPolicyV1Alpha1Manifest{}
}

type IamServicesK8SAwsPolicyV1Alpha1Manifest struct{}

type IamServicesK8SAwsPolicyV1Alpha1ManifestData struct {
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
		Description    *string `tfsdk:"description" json:"description,omitempty"`
		Name           *string `tfsdk:"name" json:"name,omitempty"`
		Path           *string `tfsdk:"path" json:"path,omitempty"`
		PolicyDocument *string `tfsdk:"policy_document" json:"policyDocument,omitempty"`
		Tags           *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *IamServicesK8SAwsPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_iam_services_k8s_aws_policy_v1alpha1_manifest"
}

func (r *IamServicesK8SAwsPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Policy is the Schema for the Policies API",
		MarkdownDescription: "Policy is the Schema for the Policies API",
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
				Description:         "PolicySpec defines the desired state of Policy.Contains information about a managed policy.This data type is used as a response element in the CreatePolicy, GetPolicy,and ListPolicies operations.For more information about managed policies, refer to Managed policies andinline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)in the IAM User Guide.",
				MarkdownDescription: "PolicySpec defines the desired state of Policy.Contains information about a managed policy.This data type is used as a response element in the CreatePolicy, GetPolicy,and ListPolicies operations.For more information about managed policies, refer to Managed policies andinline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)in the IAM User Guide.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "A friendly description of the policy.Typically used to store information about the permissions defined in thepolicy. For example, 'Grants access to production DynamoDB tables.'The policy description is immutable. After a value is assigned, it cannotbe changed.",
						MarkdownDescription: "A friendly description of the policy.Typically used to store information about the permissions defined in thepolicy. For example, 'Grants access to production DynamoDB tables.'The policy description is immutable. After a value is assigned, it cannotbe changed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The friendly name of the policy.IAM user, group, role, and policy names must be unique within the account.Names are not distinguished by case. For example, you cannot create resourcesnamed both 'MyResource' and 'myresource'.",
						MarkdownDescription: "The friendly name of the policy.IAM user, group, role, and policy names must be unique within the account.Names are not distinguished by case. For example, you cannot create resourcesnamed both 'MyResource' and 'myresource'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "The path for the policy.For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)in the IAM User Guide.This parameter is optional. If it is not included, it defaults to a slash(/).This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))a string of characters consisting of either a forward slash (/) by itselfor a string that must begin and end with forward slashes. In addition, itcan contain any ASCII character from the ! (u0021) through the DEL character(u007F), including most punctuation characters, digits, and upper and lowercasedletters.You cannot use an asterisk (*) in the path name.",
						MarkdownDescription: "The path for the policy.For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)in the IAM User Guide.This parameter is optional. If it is not included, it defaults to a slash(/).This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))a string of characters consisting of either a forward slash (/) by itselfor a string that must begin and end with forward slashes. In addition, itcan contain any ASCII character from the ! (u0021) through the DEL character(u007F), including most punctuation characters, digits, and upper and lowercasedletters.You cannot use an asterisk (*) in the path name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy_document": schema.StringAttribute{
						Description:         "The JSON policy document that you want to use as the content for the newpolicy.You must provide policies in JSON format in IAM. However, for CloudFormationtemplates formatted in YAML, you can provide the policy in JSON or YAML format.CloudFormation always converts a YAML policy to JSON format before submittingit to IAM.The maximum length of the policy document that you can pass in this operation,including whitespace, is listed below. To view the maximum character countsof a managed policy with no whitespaces, see IAM and STS character quotas(https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length).To learn more about JSON policy grammar, see Grammar of the IAM JSON policylanguage (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html)in the IAM User Guide.The regex pattern (http://wikipedia.org/wiki/regex) used to validate thisparameter is a string of characters consisting of the following:   * Any printable ASCII character ranging from the space character (u0020)   through the end of the ASCII character range   * The printable characters in the Basic Latin and Latin-1 Supplement character   set (through u00FF)   * The special characters tab (u0009), line feed (u000A), and carriage   return (u000D)",
						MarkdownDescription: "The JSON policy document that you want to use as the content for the newpolicy.You must provide policies in JSON format in IAM. However, for CloudFormationtemplates formatted in YAML, you can provide the policy in JSON or YAML format.CloudFormation always converts a YAML policy to JSON format before submittingit to IAM.The maximum length of the policy document that you can pass in this operation,including whitespace, is listed below. To view the maximum character countsof a managed policy with no whitespaces, see IAM and STS character quotas(https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length).To learn more about JSON policy grammar, see Grammar of the IAM JSON policylanguage (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html)in the IAM User Guide.The regex pattern (http://wikipedia.org/wiki/regex) used to validate thisparameter is a string of characters consisting of the following:   * Any printable ASCII character ranging from the space character (u0020)   through the end of the ASCII character range   * The printable characters in the Basic Latin and Latin-1 Supplement character   set (through u00FF)   * The special characters tab (u0009), line feed (u000A), and carriage   return (u000D)",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of tags that you want to attach to the new IAM customer managed policy.Each tag consists of a key name and an associated value. For more informationabout tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)in the IAM User Guide.If any one of the tags is invalid or if you exceed the allowed maximum numberof tags, then the entire request fails and the resource is not created.",
						MarkdownDescription: "A list of tags that you want to attach to the new IAM customer managed policy.Each tag consists of a key name and an associated value. For more informationabout tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)in the IAM User Guide.If any one of the tags is invalid or if you exceed the allowed maximum numberof tags, then the entire request fails and the resource is not created.",
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

func (r *IamServicesK8SAwsPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_iam_services_k8s_aws_policy_v1alpha1_manifest")

	var model IamServicesK8SAwsPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("iam.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Policy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
