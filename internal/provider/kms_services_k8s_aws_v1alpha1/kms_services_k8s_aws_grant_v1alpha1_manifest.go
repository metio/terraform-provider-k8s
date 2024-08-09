/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kms_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &KmsServicesK8SAwsGrantV1Alpha1Manifest{}
)

func NewKmsServicesK8SAwsGrantV1Alpha1Manifest() datasource.DataSource {
	return &KmsServicesK8SAwsGrantV1Alpha1Manifest{}
}

type KmsServicesK8SAwsGrantV1Alpha1Manifest struct{}

type KmsServicesK8SAwsGrantV1Alpha1ManifestData struct {
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
		Constraints *struct {
			EncryptionContextEquals *map[string]string `tfsdk:"encryption_context_equals" json:"encryptionContextEquals,omitempty"`
			EncryptionContextSubset *map[string]string `tfsdk:"encryption_context_subset" json:"encryptionContextSubset,omitempty"`
		} `tfsdk:"constraints" json:"constraints,omitempty"`
		GrantTokens      *[]string `tfsdk:"grant_tokens" json:"grantTokens,omitempty"`
		GranteePrincipal *string   `tfsdk:"grantee_principal" json:"granteePrincipal,omitempty"`
		KeyID            *string   `tfsdk:"key_id" json:"keyID,omitempty"`
		KeyRef           *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"key_ref" json:"keyRef,omitempty"`
		Name              *string   `tfsdk:"name" json:"name,omitempty"`
		Operations        *[]string `tfsdk:"operations" json:"operations,omitempty"`
		RetiringPrincipal *string   `tfsdk:"retiring_principal" json:"retiringPrincipal,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kms_services_k8s_aws_grant_v1alpha1_manifest"
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Grant is the Schema for the Grants API",
		MarkdownDescription: "Grant is the Schema for the Grants API",
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
				Description:         "GrantSpec defines the desired state of Grant.",
				MarkdownDescription: "GrantSpec defines the desired state of Grant.",
				Attributes: map[string]schema.Attribute{
					"constraints": schema.SingleNestedAttribute{
						Description:         "Specifies a grant constraint.KMS supports the EncryptionContextEquals and EncryptionContextSubset grantconstraints. Each constraint value can include up to 8 encryption contextpairs. The encryption context value in each constraint cannot exceed 384characters. For information about grant constraints, see Using grant constraints(https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-constraints)in the Key Management Service Developer Guide. For more information aboutencryption context, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)in the Key Management Service Developer Guide .The encryption context grant constraints allow the permissions in the grantonly when the encryption context in the request matches (EncryptionContextEquals)or includes (EncryptionContextSubset) the encryption context specified inthis structure.The encryption context grant constraints are supported only on grant operations(https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations)that include an EncryptionContext parameter, such as cryptographic operationson symmetric encryption KMS keys. Grants with grant constraints can includethe DescribeKey and RetireGrant operations, but the constraint doesn't applyto these operations. If a grant with a grant constraint includes the CreateGrantoperation, the constraint requires that any grants created with the CreateGrantpermission have an equally strict or stricter encryption context constraint.You cannot use an encryption context grant constraint for cryptographic operationswith asymmetric KMS keys or HMAC KMS keys. These keys don't support an encryptioncontext.",
						MarkdownDescription: "Specifies a grant constraint.KMS supports the EncryptionContextEquals and EncryptionContextSubset grantconstraints. Each constraint value can include up to 8 encryption contextpairs. The encryption context value in each constraint cannot exceed 384characters. For information about grant constraints, see Using grant constraints(https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-constraints)in the Key Management Service Developer Guide. For more information aboutencryption context, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)in the Key Management Service Developer Guide .The encryption context grant constraints allow the permissions in the grantonly when the encryption context in the request matches (EncryptionContextEquals)or includes (EncryptionContextSubset) the encryption context specified inthis structure.The encryption context grant constraints are supported only on grant operations(https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations)that include an EncryptionContext parameter, such as cryptographic operationson symmetric encryption KMS keys. Grants with grant constraints can includethe DescribeKey and RetireGrant operations, but the constraint doesn't applyto these operations. If a grant with a grant constraint includes the CreateGrantoperation, the constraint requires that any grants created with the CreateGrantpermission have an equally strict or stricter encryption context constraint.You cannot use an encryption context grant constraint for cryptographic operationswith asymmetric KMS keys or HMAC KMS keys. These keys don't support an encryptioncontext.",
						Attributes: map[string]schema.Attribute{
							"encryption_context_equals": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"encryption_context_subset": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"grant_tokens": schema.ListAttribute{
						Description:         "A list of grant tokens.Use a grant token when your permission to call this operation comes froma new grant that has not yet achieved eventual consistency. For more information,see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)in the Key Management Service Developer Guide.",
						MarkdownDescription: "A list of grant tokens.Use a grant token when your permission to call this operation comes froma new grant that has not yet achieved eventual consistency. For more information,see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)in the Key Management Service Developer Guide.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grantee_principal": schema.StringAttribute{
						Description:         "The identity that gets the permissions specified in the grant.To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)of an Amazon Web Services principal. Valid Amazon Web Services principalsinclude Amazon Web Services accounts (root), IAM users, IAM roles, federatedusers, and assumed role users. For examples of the ARN syntax to use forspecifying a principal, see Amazon Web Services Identity and Access Management(IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)in the Example ARNs section of the Amazon Web Services General Reference.",
						MarkdownDescription: "The identity that gets the permissions specified in the grant.To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)of an Amazon Web Services principal. Valid Amazon Web Services principalsinclude Amazon Web Services accounts (root), IAM users, IAM roles, federatedusers, and assumed role users. For examples of the ARN syntax to use forspecifying a principal, see Amazon Web Services Identity and Access Management(IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)in the Example ARNs section of the Amazon Web Services General Reference.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"key_id": schema.StringAttribute{
						Description:         "Identifies the KMS key for the grant. The grant gives principals permissionto use this KMS key.Specify the key ID or key ARN of the KMS key. To specify a KMS key in a differentAmazon Web Services account, you must use the key ARN.For example:   * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab   * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890abTo get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.",
						MarkdownDescription: "Identifies the KMS key for the grant. The grant gives principals permissionto use this KMS key.Specify the key ID or key ARN of the KMS key. To specify a KMS key in a differentAmazon Web Services account, you must use the key ARN.For example:   * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab   * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890abTo get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_ref": schema.SingleNestedAttribute{
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

									"namespace": schema.StringAttribute{
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

					"name": schema.StringAttribute{
						Description:         "A friendly name for the grant. Use this value to prevent the unintended creationof duplicate grants when retrying this request.When this value is absent, all CreateGrant requests result in a new grantwith a unique GrantId even if all the supplied parameters are identical.This can result in unintended duplicates when you retry the CreateGrant request.When this value is present, you can retry a CreateGrant request with identicalparameters; if the grant already exists, the original GrantId is returnedwithout creating a new grant. Note that the returned grant token is uniquewith every CreateGrant request, even when a duplicate GrantId is returned.All grant tokens for the same grant ID can be used interchangeably.",
						MarkdownDescription: "A friendly name for the grant. Use this value to prevent the unintended creationof duplicate grants when retrying this request.When this value is absent, all CreateGrant requests result in a new grantwith a unique GrantId even if all the supplied parameters are identical.This can result in unintended duplicates when you retry the CreateGrant request.When this value is present, you can retry a CreateGrant request with identicalparameters; if the grant already exists, the original GrantId is returnedwithout creating a new grant. Note that the returned grant token is uniquewith every CreateGrant request, even when a duplicate GrantId is returned.All grant tokens for the same grant ID can be used interchangeably.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operations": schema.ListAttribute{
						Description:         "A list of operations that the grant permits.This list must include only operations that are permitted in a grant. Also,the operation must be supported on the KMS key. For example, you cannot createa grant for a symmetric encryption KMS key that allows the Sign operation,or a grant for an asymmetric KMS key that allows the GenerateDataKey operation.If you try, KMS returns a ValidationError exception. For details, see Grantoperations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations)in the Key Management Service Developer Guide.",
						MarkdownDescription: "A list of operations that the grant permits.This list must include only operations that are permitted in a grant. Also,the operation must be supported on the KMS key. For example, you cannot createa grant for a symmetric encryption KMS key that allows the Sign operation,or a grant for an asymmetric KMS key that allows the GenerateDataKey operation.If you try, KMS returns a ValidationError exception. For details, see Grantoperations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations)in the Key Management Service Developer Guide.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"retiring_principal": schema.StringAttribute{
						Description:         "The principal that has permission to use the RetireGrant operation to retirethe grant.To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)of an Amazon Web Services principal. Valid Amazon Web Services principalsinclude Amazon Web Services accounts (root), IAM users, federated users,and assumed role users. For examples of the ARN syntax to use for specifyinga principal, see Amazon Web Services Identity and Access Management (IAM)(https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)in the Example ARNs section of the Amazon Web Services General Reference.The grant determines the retiring principal. Other principals might havepermission to retire the grant or revoke the grant. For details, see RevokeGrantand Retiring and revoking grants (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete)in the Key Management Service Developer Guide.",
						MarkdownDescription: "The principal that has permission to use the RetireGrant operation to retirethe grant.To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)of an Amazon Web Services principal. Valid Amazon Web Services principalsinclude Amazon Web Services accounts (root), IAM users, federated users,and assumed role users. For examples of the ARN syntax to use for specifyinga principal, see Amazon Web Services Identity and Access Management (IAM)(https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)in the Example ARNs section of the Amazon Web Services General Reference.The grant determines the retiring principal. Other principals might havepermission to retire the grant or revoke the grant. For details, see RevokeGrantand Retiring and revoking grants (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete)in the Key Management Service Developer Guide.",
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

func (r *KmsServicesK8SAwsGrantV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kms_services_k8s_aws_grant_v1alpha1_manifest")

	var model KmsServicesK8SAwsGrantV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kms.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Grant")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
