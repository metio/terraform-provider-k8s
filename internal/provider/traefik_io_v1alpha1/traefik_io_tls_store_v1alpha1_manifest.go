/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

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
	_ datasource.DataSource = &TraefikIoTlsstoreV1Alpha1Manifest{}
)

func NewTraefikIoTlsstoreV1Alpha1Manifest() datasource.DataSource {
	return &TraefikIoTlsstoreV1Alpha1Manifest{}
}

type TraefikIoTlsstoreV1Alpha1Manifest struct{}

type TraefikIoTlsstoreV1Alpha1ManifestData struct {
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
		Certificates *[]struct {
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"certificates" json:"certificates,omitempty"`
		DefaultCertificate *struct {
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"default_certificate" json:"defaultCertificate,omitempty"`
		DefaultGeneratedCert *struct {
			Domain *struct {
				Main *string   `tfsdk:"main" json:"main,omitempty"`
				Sans *[]string `tfsdk:"sans" json:"sans,omitempty"`
			} `tfsdk:"domain" json:"domain,omitempty"`
			Resolver *string `tfsdk:"resolver" json:"resolver,omitempty"`
		} `tfsdk:"default_generated_cert" json:"defaultGeneratedCert,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoTlsstoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_tls_store_v1alpha1_manifest"
}

func (r *TraefikIoTlsstoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TLSStore is the CRD implementation of a Traefik TLS Store. For the time being, only the TLSStore named default is supported. This means that you cannot have two stores that are named default in different Kubernetes namespaces. More info: https://doc.traefik.io/traefik/v3.0/https/tls/#certificates-stores",
		MarkdownDescription: "TLSStore is the CRD implementation of a Traefik TLS Store. For the time being, only the TLSStore named default is supported. This means that you cannot have two stores that are named default in different Kubernetes namespaces. More info: https://doc.traefik.io/traefik/v3.0/https/tls/#certificates-stores",
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
				Description:         "TLSStoreSpec defines the desired state of a TLSStore.",
				MarkdownDescription: "TLSStoreSpec defines the desired state of a TLSStore.",
				Attributes: map[string]schema.Attribute{
					"certificates": schema.ListNestedAttribute{
						Description:         "Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.",
						MarkdownDescription: "Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"secret_name": schema.StringAttribute{
									Description:         "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
									MarkdownDescription: "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
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

					"default_certificate": schema.SingleNestedAttribute{
						Description:         "DefaultCertificate defines the default certificate configuration.",
						MarkdownDescription: "DefaultCertificate defines the default certificate configuration.",
						Attributes: map[string]schema.Attribute{
							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								MarkdownDescription: "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_generated_cert": schema.SingleNestedAttribute{
						Description:         "DefaultGeneratedCert defines the default generated certificate configuration.",
						MarkdownDescription: "DefaultGeneratedCert defines the default generated certificate configuration.",
						Attributes: map[string]schema.Attribute{
							"domain": schema.SingleNestedAttribute{
								Description:         "Domain is the domain definition for the DefaultCertificate.",
								MarkdownDescription: "Domain is the domain definition for the DefaultCertificate.",
								Attributes: map[string]schema.Attribute{
									"main": schema.StringAttribute{
										Description:         "Main defines the main domain name.",
										MarkdownDescription: "Main defines the main domain name.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sans": schema.ListAttribute{
										Description:         "SANs defines the subject alternative domain names.",
										MarkdownDescription: "SANs defines the subject alternative domain names.",
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

							"resolver": schema.StringAttribute{
								Description:         "Resolver is the name of the resolver that will be used to issue the DefaultCertificate.",
								MarkdownDescription: "Resolver is the name of the resolver that will be used to issue the DefaultCertificate.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TraefikIoTlsstoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_io_tls_store_v1alpha1_manifest")

	var model TraefikIoTlsstoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("traefik.io/v1alpha1")
	model.Kind = pointer.String("TLSStore")

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
