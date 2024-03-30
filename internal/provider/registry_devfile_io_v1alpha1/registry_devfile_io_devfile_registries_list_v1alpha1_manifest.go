/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package registry_devfile_io_v1alpha1

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
	_ datasource.DataSource = &RegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest{}
)

func NewRegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest() datasource.DataSource {
	return &RegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest{}
}

type RegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest struct{}

type RegistryDevfileIoDevfileRegistriesListV1Alpha1ManifestData struct {
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
		DevfileRegistries *[]struct {
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			SkipTLSVerify *bool   `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			Url           *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"devfile_registries" json:"devfileRegistries,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_registry_devfile_io_devfile_registries_list_v1alpha1_manifest"
}

func (r *RegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DevfileRegistriesList is a custom resource where namespace users can add a list of Devfile Registries to allow devfiles to be visible at the namespace level.  In order to be added to the list, the Devfile Registries must be reachable, supports the Devfile v2.0 spec and above, and is not using the default namespace.",
		MarkdownDescription: "DevfileRegistriesList is a custom resource where namespace users can add a list of Devfile Registries to allow devfiles to be visible at the namespace level.  In order to be added to the list, the Devfile Registries must be reachable, supports the Devfile v2.0 spec and above, and is not using the default namespace.",
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
				Description:         "DevfileRegistriesListSpec defines the desired state of DevfileRegistriesList",
				MarkdownDescription: "DevfileRegistriesListSpec defines the desired state of DevfileRegistriesList",
				Attributes: map[string]schema.Attribute{
					"devfile_registries": schema.ListNestedAttribute{
						Description:         "DevfileRegistries is a list of devfile registry services",
						MarkdownDescription: "DevfileRegistries is a list of devfile registry services",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the unique Name of the devfile registry.",
									MarkdownDescription: "Name is the unique Name of the devfile registry.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "SkipTLSVerify defaults to false.  Set to true in a non-production environment to bypass certificate checking",
									MarkdownDescription: "SkipTLSVerify defaults to false.  Set to true in a non-production environment to bypass certificate checking",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "URL is the unique URL of the devfile registry.",
									MarkdownDescription: "URL is the unique URL of the devfile registry.",
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

func (r *RegistryDevfileIoDevfileRegistriesListV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_registry_devfile_io_devfile_registries_list_v1alpha1_manifest")

	var model RegistryDevfileIoDevfileRegistriesListV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("registry.devfile.io/v1alpha1")
	model.Kind = pointer.String("DevfileRegistriesList")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
