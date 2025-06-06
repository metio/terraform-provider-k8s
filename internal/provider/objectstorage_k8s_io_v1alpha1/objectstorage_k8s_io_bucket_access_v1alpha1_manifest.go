/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package objectstorage_k8s_io_v1alpha1

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
	_ datasource.DataSource = &ObjectstorageK8SIoBucketAccessV1Alpha1Manifest{}
)

func NewObjectstorageK8SIoBucketAccessV1Alpha1Manifest() datasource.DataSource {
	return &ObjectstorageK8SIoBucketAccessV1Alpha1Manifest{}
}

type ObjectstorageK8SIoBucketAccessV1Alpha1Manifest struct{}

type ObjectstorageK8SIoBucketAccessV1Alpha1ManifestData struct {
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
		BucketAccessClassName *string `tfsdk:"bucket_access_class_name" json:"bucketAccessClassName,omitempty"`
		BucketClaimName       *string `tfsdk:"bucket_claim_name" json:"bucketClaimName,omitempty"`
		CredentialsSecretName *string `tfsdk:"credentials_secret_name" json:"credentialsSecretName,omitempty"`
		Protocol              *string `tfsdk:"protocol" json:"protocol,omitempty"`
		ServiceAccountName    *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ObjectstorageK8SIoBucketAccessV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_objectstorage_k8s_io_bucket_access_v1alpha1_manifest"
}

func (r *ObjectstorageK8SIoBucketAccessV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"bucket_access_class_name": schema.StringAttribute{
						Description:         "BucketAccessClassName is the name of the BucketAccessClass",
						MarkdownDescription: "BucketAccessClassName is the name of the BucketAccessClass",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"bucket_claim_name": schema.StringAttribute{
						Description:         "BucketClaimName is the name of the BucketClaim.",
						MarkdownDescription: "BucketClaimName is the name of the BucketClaim.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"credentials_secret_name": schema.StringAttribute{
						Description:         "CredentialsSecretName is the name of the secret that COSI should populate with the credentials. If a secret by this name already exists, then it is assumed that credentials have already been generated. It is not overridden. This secret is deleted when the BucketAccess is delted.",
						MarkdownDescription: "CredentialsSecretName is the name of the secret that COSI should populate with the credentials. If a secret by this name already exists, then it is assumed that credentials have already been generated. It is not overridden. This secret is deleted when the BucketAccess is delted.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"protocol": schema.StringAttribute{
						Description:         "Protocol is the name of the Protocol that this access credential is supposed to support If left empty, it will choose the protocol supported by the bucket. If the bucket supports multiple protocols, the end protocol is determined by the driver.",
						MarkdownDescription: "Protocol is the name of the Protocol that this access credential is supposed to support If left empty, it will choose the protocol supported by the bucket. If the bucket supports multiple protocols, the end protocol is determined by the driver.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the serviceAccount that COSI will map to the OSP service account when IAM styled authentication is specified",
						MarkdownDescription: "ServiceAccountName is the name of the serviceAccount that COSI will map to the OSP service account when IAM styled authentication is specified",
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
	}
}

func (r *ObjectstorageK8SIoBucketAccessV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_objectstorage_k8s_io_bucket_access_v1alpha1_manifest")

	var model ObjectstorageK8SIoBucketAccessV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("objectstorage.k8s.io/v1alpha1")
	model.Kind = pointer.String("BucketAccess")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
