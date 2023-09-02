/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package pkg_crossplane_io_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &PkgCrossplaneIoLockV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &PkgCrossplaneIoLockV1Beta1DataSource{}
)

func NewPkgCrossplaneIoLockV1Beta1DataSource() datasource.DataSource {
	return &PkgCrossplaneIoLockV1Beta1DataSource{}
}

type PkgCrossplaneIoLockV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type PkgCrossplaneIoLockV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Packages *[]struct {
		Dependencies *[]struct {
			Constraints *string `tfsdk:"constraints" json:"constraints,omitempty"`
			Package     *string `tfsdk:"package" json:"package,omitempty"`
			Type        *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"dependencies" json:"dependencies,omitempty"`
		Name    *string `tfsdk:"name" json:"name,omitempty"`
		Source  *string `tfsdk:"source" json:"source,omitempty"`
		Type    *string `tfsdk:"type" json:"type,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"packages" json:"packages,omitempty"`
}

func (r *PkgCrossplaneIoLockV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pkg_crossplane_io_lock_v1beta1"
}

func (r *PkgCrossplaneIoLockV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Lock is the CRD type that tracks package dependencies.",
		MarkdownDescription: "Lock is the CRD type that tracks package dependencies.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"packages": schema.ListNestedAttribute{
				Description:         "",
				MarkdownDescription: "",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dependencies": schema.ListNestedAttribute{
							Description:         "Dependencies are the list of dependencies of this package. The order of the dependencies will dictate the order in which they are resolved.",
							MarkdownDescription: "Dependencies are the list of dependencies of this package. The order of the dependencies will dictate the order in which they are resolved.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"constraints": schema.StringAttribute{
										Description:         "Constraints is a valid semver range, which will be used to select a valid dependency version.",
										MarkdownDescription: "Constraints is a valid semver range, which will be used to select a valid dependency version.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"package": schema.StringAttribute{
										Description:         "Package is the OCI image name without a tag or digest.",
										MarkdownDescription: "Package is the OCI image name without a tag or digest.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the type of package. Can be either Configuration or Provider.",
										MarkdownDescription: "Type is the type of package. Can be either Configuration or Provider.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
							},
							Required: false,
							Optional: false,
							Computed: true,
						},

						"name": schema.StringAttribute{
							Description:         "Name corresponds to the name of the package revision for this package.",
							MarkdownDescription: "Name corresponds to the name of the package revision for this package.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"source": schema.StringAttribute{
							Description:         "Source is the OCI image name without a tag or digest.",
							MarkdownDescription: "Source is the OCI image name without a tag or digest.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"type": schema.StringAttribute{
							Description:         "Type is the type of package. Can be either Configuration or Provider.",
							MarkdownDescription: "Type is the type of package. Can be either Configuration or Provider.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"version": schema.StringAttribute{
							Description:         "Version is the tag or digest of the OCI image.",
							MarkdownDescription: "Version is the tag or digest of the OCI image.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *PkgCrossplaneIoLockV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *PkgCrossplaneIoLockV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_pkg_crossplane_io_lock_v1beta1")

	var data PkgCrossplaneIoLockV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1beta1", Resource: "Lock"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse PkgCrossplaneIoLockV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("pkg.crossplane.io/v1beta1")
	data.Kind = pointer.String("Lock")
	data.Metadata = readResponse.Metadata
	data.Packages = readResponse.Packages

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
