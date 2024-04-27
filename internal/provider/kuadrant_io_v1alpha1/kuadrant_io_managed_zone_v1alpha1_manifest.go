/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuadrant_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KuadrantIoManagedZoneV1Alpha1Manifest{}
)

func NewKuadrantIoManagedZoneV1Alpha1Manifest() datasource.DataSource {
	return &KuadrantIoManagedZoneV1Alpha1Manifest{}
}

type KuadrantIoManagedZoneV1Alpha1Manifest struct{}

type KuadrantIoManagedZoneV1Alpha1ManifestData struct {
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
		Description          *string `tfsdk:"description" json:"description,omitempty"`
		DnsProviderSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"dns_provider_secret_ref" json:"dnsProviderSecretRef,omitempty"`
		DomainName        *string `tfsdk:"domain_name" json:"domainName,omitempty"`
		Id                *string `tfsdk:"id" json:"id,omitempty"`
		ParentManagedZone *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"parent_managed_zone" json:"parentManagedZone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KuadrantIoManagedZoneV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuadrant_io_managed_zone_v1alpha1_manifest"
}

func (r *KuadrantIoManagedZoneV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ManagedZone is the Schema for the managedzones API",
		MarkdownDescription: "ManagedZone is the Schema for the managedzones API",
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
				Description:         "ManagedZoneSpec defines the desired state of ManagedZone",
				MarkdownDescription: "ManagedZoneSpec defines the desired state of ManagedZone",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "description for this ManagedZone",
						MarkdownDescription: "description for this ManagedZone",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"dns_provider_secret_ref": schema.SingleNestedAttribute{
						Description:         "dnsProviderSecretRef reference to a secret containing credentials to access a dns provider.",
						MarkdownDescription: "dnsProviderSecretRef reference to a secret containing credentials to access a dns provider.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"domain_name": schema.StringAttribute{
						Description:         "domainName of this ManagedZone",
						MarkdownDescription: "domainName of this ManagedZone",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
						},
					},

					"id": schema.StringAttribute{
						Description:         "id is the provider assigned id of this  zone (i.e. route53.HostedZone.ID).",
						MarkdownDescription: "id is the provider assigned id of this  zone (i.e. route53.HostedZone.ID).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parent_managed_zone": schema.SingleNestedAttribute{
						Description:         "parentManagedZone reference to another managed zone that this managed zone belongs to.",
						MarkdownDescription: "parentManagedZone reference to another managed zone that this managed zone belongs to.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "'name' is the name of the managed zone. Required",
								MarkdownDescription: "'name' is the name of the managed zone. Required",
								Required:            true,
								Optional:            false,
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

func (r *KuadrantIoManagedZoneV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuadrant_io_managed_zone_v1alpha1_manifest")

	var model KuadrantIoManagedZoneV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuadrant.io/v1alpha1")
	model.Kind = pointer.String("ManagedZone")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
