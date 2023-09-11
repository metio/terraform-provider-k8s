/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &ChaosMeshOrgPodIochaosV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgPodIochaosV1Alpha1DataSource{}
)

func NewChaosMeshOrgPodIochaosV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgPodIochaosV1Alpha1DataSource{}
}

type ChaosMeshOrgPodIochaosV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgPodIochaosV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Actions *[]struct {
			Atime *struct {
				Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
				Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
			} `tfsdk:"atime" json:"atime,omitempty"`
			Blocks *int64 `tfsdk:"blocks" json:"blocks,omitempty"`
			Ctime  *struct {
				Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
				Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
			} `tfsdk:"ctime" json:"ctime,omitempty"`
			Faults *[]struct {
				Errno  *int64 `tfsdk:"errno" json:"errno,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"faults" json:"faults,omitempty"`
			Gid     *int64    `tfsdk:"gid" json:"gid,omitempty"`
			Ino     *int64    `tfsdk:"ino" json:"ino,omitempty"`
			Kind    *string   `tfsdk:"kind" json:"kind,omitempty"`
			Latency *string   `tfsdk:"latency" json:"latency,omitempty"`
			Methods *[]string `tfsdk:"methods" json:"methods,omitempty"`
			Mistake *struct {
				Filling        *string `tfsdk:"filling" json:"filling,omitempty"`
				MaxLength      *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
				MaxOccurrences *int64  `tfsdk:"max_occurrences" json:"maxOccurrences,omitempty"`
			} `tfsdk:"mistake" json:"mistake,omitempty"`
			Mtime *struct {
				Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
				Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
			} `tfsdk:"mtime" json:"mtime,omitempty"`
			Nlink   *int64  `tfsdk:"nlink" json:"nlink,omitempty"`
			Path    *string `tfsdk:"path" json:"path,omitempty"`
			Percent *int64  `tfsdk:"percent" json:"percent,omitempty"`
			Perm    *int64  `tfsdk:"perm" json:"perm,omitempty"`
			Rdev    *int64  `tfsdk:"rdev" json:"rdev,omitempty"`
			Size    *int64  `tfsdk:"size" json:"size,omitempty"`
			Source  *string `tfsdk:"source" json:"source,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
			Uid     *int64  `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"actions" json:"actions,omitempty"`
		Container       *string `tfsdk:"container" json:"container,omitempty"`
		VolumeMountPath *string `tfsdk:"volume_mount_path" json:"volumeMountPath,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgPodIochaosV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_pod_io_chaos_v1alpha1"
}

func (r *ChaosMeshOrgPodIochaosV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PodIOChaos is the Schema for the podiochaos API",
		MarkdownDescription: "PodIOChaos is the Schema for the podiochaos API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "PodIOChaosSpec defines the desired state of IOChaos",
				MarkdownDescription: "PodIOChaosSpec defines the desired state of IOChaos",
				Attributes: map[string]schema.Attribute{
					"actions": schema.ListNestedAttribute{
						Description:         "Actions are a list of IOChaos actions",
						MarkdownDescription: "Actions are a list of IOChaos actions",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"atime": schema.SingleNestedAttribute{
									Description:         "Timespec represents a time",
									MarkdownDescription: "Timespec represents a time",
									Attributes: map[string]schema.Attribute{
										"nsec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"blocks": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ctime": schema.SingleNestedAttribute{
									Description:         "Timespec represents a time",
									MarkdownDescription: "Timespec represents a time",
									Attributes: map[string]schema.Attribute{
										"nsec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"faults": schema.ListNestedAttribute{
									Description:         "Faults represents the fault to inject",
									MarkdownDescription: "Faults represents the fault to inject",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"errno": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"weight": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
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

								"gid": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ino": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "FileType represents type of file",
									MarkdownDescription: "FileType represents type of file",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"latency": schema.StringAttribute{
									Description:         "Latency represents the latency to inject",
									MarkdownDescription: "Latency represents the latency to inject",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"methods": schema.ListAttribute{
									Description:         "Methods represents the method that the action will inject in",
									MarkdownDescription: "Methods represents the method that the action will inject in",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"mistake": schema.SingleNestedAttribute{
									Description:         "MistakeSpec represents the mistake to inject",
									MarkdownDescription: "MistakeSpec represents the mistake to inject",
									Attributes: map[string]schema.Attribute{
										"filling": schema.StringAttribute{
											Description:         "Filling determines what is filled in the mistake data.",
											MarkdownDescription: "Filling determines what is filled in the mistake data.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"max_length": schema.Int64Attribute{
											Description:         "Max length of each wrong data segment in bytes",
											MarkdownDescription: "Max length of each wrong data segment in bytes",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"max_occurrences": schema.Int64Attribute{
											Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
											MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"mtime": schema.SingleNestedAttribute{
									Description:         "Timespec represents a time",
									MarkdownDescription: "Timespec represents a time",
									Attributes: map[string]schema.Attribute{
										"nsec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"nlink": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "Path represents a glob of injecting path",
									MarkdownDescription: "Path represents a glob of injecting path",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"percent": schema.Int64Attribute{
									Description:         "Percent represents the percent probability of injecting this action",
									MarkdownDescription: "Percent represents the percent probability of injecting this action",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"perm": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rdev": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"size": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"source": schema.StringAttribute{
									Description:         "Source represents the source of current rules",
									MarkdownDescription: "Source represents the source of current rules",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "IOChaosType represents the type of IOChaos Action",
									MarkdownDescription: "IOChaosType represents the type of IOChaos Action",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"uid": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
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

					"container": schema.StringAttribute{
						Description:         "TODO: support multiple different container to inject in one pod",
						MarkdownDescription: "TODO: support multiple different container to inject in one pod",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"volume_mount_path": schema.StringAttribute{
						Description:         "VolumeMountPath represents the target mount path It must be a root of mount path now. TODO: search the mount parent of any path automatically. TODO: support multiple different volume mount path in one pod",
						MarkdownDescription: "VolumeMountPath represents the target mount path It must be a root of mount path now. TODO: search the mount parent of any path automatically. TODO: support multiple different volume mount path in one pod",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ChaosMeshOrgPodIochaosV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ChaosMeshOrgPodIochaosV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_pod_io_chaos_v1alpha1")

	var data ChaosMeshOrgPodIochaosV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "podiochaos"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse ChaosMeshOrgPodIochaosV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("PodIOChaos")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
