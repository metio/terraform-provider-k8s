/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_banzaicloud_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &KafkaBanzaicloudIoKafkaUserV1Alpha1Manifest{}
)

func NewKafkaBanzaicloudIoKafkaUserV1Alpha1Manifest() datasource.DataSource {
	return &KafkaBanzaicloudIoKafkaUserV1Alpha1Manifest{}
}

type KafkaBanzaicloudIoKafkaUserV1Alpha1Manifest struct{}

type KafkaBanzaicloudIoKafkaUserV1Alpha1ManifestData struct {
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
		Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
		ClusterRef  *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		CreateCert        *bool     `tfsdk:"create_cert" json:"createCert,omitempty"`
		DnsNames          *[]string `tfsdk:"dns_names" json:"dnsNames,omitempty"`
		ExpirationSeconds *int64    `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
		IncludeJKS        *bool     `tfsdk:"include_jks" json:"includeJKS,omitempty"`
		PkiBackendSpec    *struct {
			IssuerRef *struct {
				Group *string `tfsdk:"group" json:"group,omitempty"`
				Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"issuer_ref" json:"issuerRef,omitempty"`
			PkiBackend *string `tfsdk:"pki_backend" json:"pkiBackend,omitempty"`
			SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
		} `tfsdk:"pki_backend_spec" json:"pkiBackendSpec,omitempty"`
		SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		TopicGrants *[]struct {
			AccessType  *string `tfsdk:"access_type" json:"accessType,omitempty"`
			PatternType *string `tfsdk:"pattern_type" json:"patternType,omitempty"`
			TopicName   *string `tfsdk:"topic_name" json:"topicName,omitempty"`
		} `tfsdk:"topic_grants" json:"topicGrants,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaBanzaicloudIoKafkaUserV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest"
}

func (r *KafkaBanzaicloudIoKafkaUserV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KafkaUser is the Schema for the kafka users API",
		MarkdownDescription: "KafkaUser is the Schema for the kafka users API",
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
				Description:         "KafkaUserSpec defines the desired state of KafkaUser",
				MarkdownDescription: "KafkaUserSpec defines the desired state of KafkaUser",
				Attributes: map[string]schema.Attribute{
					"annotations": schema.MapAttribute{
						Description:         "Annotations defines the annotations placed on the certificate or certificate signing request object",
						MarkdownDescription: "Annotations defines the annotations placed on the certificate or certificate signing request object",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_ref": schema.SingleNestedAttribute{
						Description:         "ClusterReference states a reference to a cluster for topic/user provisioning",
						MarkdownDescription: "ClusterReference states a reference to a cluster for topic/user provisioning",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"create_cert": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dns_names": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"expiration_seconds": schema.Int64Attribute{
						Description:         "expirationSeconds is the requested duration of validity of the issued certificate. The minimum valid value for expirationSeconds is 3600 i.e. 1h. When it is not specified the default validation duration is 90 days",
						MarkdownDescription: "expirationSeconds is the requested duration of validity of the issued certificate. The minimum valid value for expirationSeconds is 3600 i.e. 1h. When it is not specified the default validation duration is 90 days",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(3600),
						},
					},

					"include_jks": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pki_backend_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"issuer_ref": schema.SingleNestedAttribute{
								Description:         "ObjectReference is a reference to an object with a given name, kind and group.",
								MarkdownDescription: "ObjectReference is a reference to an object with a given name, kind and group.",
								Attributes: map[string]schema.Attribute{
									"group": schema.StringAttribute{
										Description:         "Group of the resource being referred to.",
										MarkdownDescription: "Group of the resource being referred to.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind of the resource being referred to.",
										MarkdownDescription: "Kind of the resource being referred to.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the resource being referred to.",
										MarkdownDescription: "Name of the resource being referred to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pki_backend": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cert-manager", "k8s-csr"),
								},
							},

							"signer_name": schema.StringAttribute{
								Description:         "SignerName indicates requested signer, and is a qualified name.",
								MarkdownDescription: "SignerName indicates requested signer, and is a qualified name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "secretName is used as the name of the K8S secret that contains the certificate of the KafkaUser. SecretName should be unique inside the namespace where KafkaUser is located.",
						MarkdownDescription: "secretName is used as the name of the K8S secret that contains the certificate of the KafkaUser. SecretName should be unique inside the namespace where KafkaUser is located.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"topic_grants": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_type": schema.StringAttribute{
									Description:         "KafkaAccessType hold info about Kafka ACL",
									MarkdownDescription: "KafkaAccessType hold info about Kafka ACL",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("read", "write"),
									},
								},

								"pattern_type": schema.StringAttribute{
									Description:         "KafkaPatternType hold the Resource Pattern Type of kafka ACL",
									MarkdownDescription: "KafkaPatternType hold the Resource Pattern Type of kafka ACL",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("literal", "match", "prefixed", "any"),
									},
								},

								"topic_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
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

func (r *KafkaBanzaicloudIoKafkaUserV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest")

	var model KafkaBanzaicloudIoKafkaUserV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kafka.banzaicloud.io/v1alpha1")
	model.Kind = pointer.String("KafkaUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
