/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package jetstream_nats_io_v1beta2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &JetstreamNatsIoObjectStoreV1Beta2Manifest{}
)

func NewJetstreamNatsIoObjectStoreV1Beta2Manifest() datasource.DataSource {
	return &JetstreamNatsIoObjectStoreV1Beta2Manifest{}
}

type JetstreamNatsIoObjectStoreV1Beta2Manifest struct{}

type JetstreamNatsIoObjectStoreV1Beta2ManifestData struct {
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
		Account     *string            `tfsdk:"account" json:"account,omitempty"`
		Bucket      *string            `tfsdk:"bucket" json:"bucket,omitempty"`
		Compression *bool              `tfsdk:"compression" json:"compression,omitempty"`
		Creds       *string            `tfsdk:"creds" json:"creds,omitempty"`
		Description *string            `tfsdk:"description" json:"description,omitempty"`
		JsDomain    *string            `tfsdk:"js_domain" json:"jsDomain,omitempty"`
		MaxBytes    *int64             `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
		Metadata    *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
		Nkey        *string            `tfsdk:"nkey" json:"nkey,omitempty"`
		Placement   *struct {
			Cluster *string   `tfsdk:"cluster" json:"cluster,omitempty"`
			Tags    *[]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
		PreventDelete *bool     `tfsdk:"prevent_delete" json:"preventDelete,omitempty"`
		PreventUpdate *bool     `tfsdk:"prevent_update" json:"preventUpdate,omitempty"`
		Replicas      *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
		Servers       *[]string `tfsdk:"servers" json:"servers,omitempty"`
		Storage       *string   `tfsdk:"storage" json:"storage,omitempty"`
		Tls           *struct {
			ClientCert *string   `tfsdk:"client_cert" json:"clientCert,omitempty"`
			ClientKey  *string   `tfsdk:"client_key" json:"clientKey,omitempty"`
			RootCas    *[]string `tfsdk:"root_cas" json:"rootCas,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		TlsFirst *bool   `tfsdk:"tls_first" json:"tlsFirst,omitempty"`
		Ttl      *string `tfsdk:"ttl" json:"ttl,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *JetstreamNatsIoObjectStoreV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_jetstream_nats_io_object_store_v1beta2_manifest"
}

func (r *JetstreamNatsIoObjectStoreV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
					"account": schema.StringAttribute{
						Description:         "Name of the account to which the Object Store belongs.",
						MarkdownDescription: "Name of the account to which the Object Store belongs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[^.*>]*$`), ""),
						},
					},

					"bucket": schema.StringAttribute{
						Description:         "A unique name for the Object Store.",
						MarkdownDescription: "A unique name for the Object Store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"compression": schema.BoolAttribute{
						Description:         "Object Store compression.",
						MarkdownDescription: "Object Store compression.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"creds": schema.StringAttribute{
						Description:         "NATS user credentials for connecting to servers. Please make sure your controller has mounted the creds on its path.",
						MarkdownDescription: "NATS user credentials for connecting to servers. Please make sure your controller has mounted the creds on its path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "The description of the Object Store.",
						MarkdownDescription: "The description of the Object Store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"js_domain": schema.StringAttribute{
						Description:         "The JetStream domain to use for the Object Store.",
						MarkdownDescription: "The JetStream domain to use for the Object Store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_bytes": schema.Int64Attribute{
						Description:         "The maximum size of the Store in bytes.",
						MarkdownDescription: "The maximum size of the Store in bytes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata": schema.MapAttribute{
						Description:         "Additional Object Store metadata.",
						MarkdownDescription: "Additional Object Store metadata.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nkey": schema.StringAttribute{
						Description:         "NATS user NKey for connecting to servers.",
						MarkdownDescription: "NATS user NKey for connecting to servers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"placement": schema.SingleNestedAttribute{
						Description:         "The Object Store placement via tags or cluster name.",
						MarkdownDescription: "The Object Store placement via tags or cluster name.",
						Attributes: map[string]schema.Attribute{
							"cluster": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.ListAttribute{
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

					"prevent_delete": schema.BoolAttribute{
						Description:         "When true, the managed Object Store will not be deleted when the resource is deleted.",
						MarkdownDescription: "When true, the managed Object Store will not be deleted when the resource is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prevent_update": schema.BoolAttribute{
						Description:         "When true, the managed Object Store will not be updated when the resource is updated.",
						MarkdownDescription: "When true, the managed Object Store will not be updated when the resource is updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "The number of replicas to keep for the Object Store in clustered JetStream.",
						MarkdownDescription: "The number of replicas to keep for the Object Store in clustered JetStream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(5),
						},
					},

					"servers": schema.ListAttribute{
						Description:         "A list of servers for creating the Object Store.",
						MarkdownDescription: "A list of servers for creating the Object Store.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage": schema.StringAttribute{
						Description:         "The storage backend to use for the Object Store.",
						MarkdownDescription: "The storage backend to use for the Object Store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("file", "memory"),
						},
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "A client's TLS certs and keys.",
						MarkdownDescription: "A client's TLS certs and keys.",
						Attributes: map[string]schema.Attribute{
							"client_cert": schema.StringAttribute{
								Description:         "A client's cert filepath. Should be mounted.",
								MarkdownDescription: "A client's cert filepath. Should be mounted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_key": schema.StringAttribute{
								Description:         "A client's key filepath. Should be mounted.",
								MarkdownDescription: "A client's key filepath. Should be mounted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"root_cas": schema.ListAttribute{
								Description:         "A list of filepaths to CAs. Should be mounted.",
								MarkdownDescription: "A list of filepaths to CAs. Should be mounted.",
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

					"tls_first": schema.BoolAttribute{
						Description:         "When true, the Object Store will initiate TLS before server INFO.",
						MarkdownDescription: "When true, the Object Store will initiate TLS before server INFO.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl": schema.StringAttribute{
						Description:         "The time expiry for keys.",
						MarkdownDescription: "The time expiry for keys.",
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

func (r *JetstreamNatsIoObjectStoreV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_jetstream_nats_io_object_store_v1beta2_manifest")

	var model JetstreamNatsIoObjectStoreV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("jetstream.nats.io/v1beta2")
	model.Kind = pointer.String("ObjectStore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
