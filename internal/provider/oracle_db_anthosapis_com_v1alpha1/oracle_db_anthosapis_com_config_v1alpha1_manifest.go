/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package oracle_db_anthosapis_com_v1alpha1

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
	_ datasource.DataSource = &OracleDbAnthosapisComConfigV1Alpha1Manifest{}
)

func NewOracleDbAnthosapisComConfigV1Alpha1Manifest() datasource.DataSource {
	return &OracleDbAnthosapisComConfigV1Alpha1Manifest{}
}

type OracleDbAnthosapisComConfigV1Alpha1Manifest struct{}

type OracleDbAnthosapisComConfigV1Alpha1ManifestData struct {
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
		Disks *[]struct {
			AccessModes *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Selector    *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Size         *string `tfsdk:"size" json:"size,omitempty"`
			StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
			VolumeName   *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
		} `tfsdk:"disks" json:"disks,omitempty"`
		HostAntiAffinityNamespaces *[]string          `tfsdk:"host_anti_affinity_namespaces" json:"hostAntiAffinityNamespaces,omitempty"`
		Images                     *map[string]string `tfsdk:"images" json:"images,omitempty"`
		LogLevel                   *map[string]string `tfsdk:"log_level" json:"logLevel,omitempty"`
		Platform                   *string            `tfsdk:"platform" json:"platform,omitempty"`
		StorageClass               *string            `tfsdk:"storage_class" json:"storageClass,omitempty"`
		VolumeSnapshotClass        *string            `tfsdk:"volume_snapshot_class" json:"volumeSnapshotClass,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OracleDbAnthosapisComConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_oracle_db_anthosapis_com_config_v1alpha1_manifest"
}

func (r *OracleDbAnthosapisComConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Config is the Schema for the configs API.",
		MarkdownDescription: "Config is the Schema for the configs API.",
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
				Description:         "ConfigSpec defines the desired state of Config.",
				MarkdownDescription: "ConfigSpec defines the desired state of Config.",
				Attributes: map[string]schema.Attribute{
					"disks": schema.ListNestedAttribute{
						Description:         "Disks slice describes at minimum two disks: data and log (archive log), and optionally a backup disk.",
						MarkdownDescription: "Disks slice describes at minimum two disks: data and log (archive log), and optionally a backup disk.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_modes": schema.ListAttribute{
									Description:         "AccessModes contains the desired access modes the volume should have.",
									MarkdownDescription: "AccessModes contains the desired access modes the volume should have.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"annotations": schema.MapAttribute{
									Description:         "A map of string keys and values to be stored in the annotations of the PVC. These can be read and write by external tools through Kubernetes.",
									MarkdownDescription: "A map of string keys and values to be stored in the annotations of the PVC. These can be read and write by external tools through Kubernetes.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of a disk.",
									MarkdownDescription: "Name of a disk.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"selector": schema.SingleNestedAttribute{
									Description:         "A label query over volumes to consider for binding.",
									MarkdownDescription: "A label query over volumes to consider for binding.",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
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

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"size": schema.StringAttribute{
									Description:         "Disk size. If not specified, the defaults are: DataDisk:'100Gi', LogDisk:'150Gi',BackupDisk:'100Gi'",
									MarkdownDescription: "Disk size. If not specified, the defaults are: DataDisk:'100Gi', LogDisk:'150Gi',BackupDisk:'100Gi'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"storage_class": schema.StringAttribute{
									Description:         "StorageClass points to a particular CSI driver and is used for disk provisioning.",
									MarkdownDescription: "StorageClass points to a particular CSI driver and is used for disk provisioning.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"volume_name": schema.StringAttribute{
									Description:         "VolumeName is the binding reference to the PersistentVolume tied to this disk.",
									MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume tied to this disk.",
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

					"host_anti_affinity_namespaces": schema.ListAttribute{
						Description:         "HostAntiAffinityNamespaces is an optional list of namespaces that need to be included in anti-affinity by hostname rule. The effect of the rule is forbidding scheduling a database pod in the current namespace on a host that already runs a database pod in any of the listed namespaces.",
						MarkdownDescription: "HostAntiAffinityNamespaces is an optional list of namespaces that need to be included in anti-affinity by hostname rule. The effect of the rule is forbidding scheduling a database pod in the current namespace on a host that already runs a database pod in any of the listed namespaces.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"images": schema.MapAttribute{
						Description:         "Service agent and other data plane agent images. This is an optional map that allows a customer to specify agent images different from those chosen/provided by the operator by default.",
						MarkdownDescription: "Service agent and other data plane agent images. This is an optional map that allows a customer to specify agent images different from those chosen/provided by the operator by default.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.MapAttribute{
						Description:         "Log Levels for the various components. This is an optional map for component -> log level",
						MarkdownDescription: "Log Levels for the various components. This is an optional map for component -> log level",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"platform": schema.StringAttribute{
						Description:         "Deployment platform. Presently supported values are: GCP (default), BareMetal, Minikube and Kind.",
						MarkdownDescription: "Deployment platform. Presently supported values are: GCP (default), BareMetal, Minikube and Kind.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("GCP", "BareMetal", "Minikube", "Kind"),
						},
					},

					"storage_class": schema.StringAttribute{
						Description:         "Storage class to use for dynamic provisioning. This value varies depending on a platform. For GCP (the default), it is 'standard-rwo'.",
						MarkdownDescription: "Storage class to use for dynamic provisioning. This value varies depending on a platform. For GCP (the default), it is 'standard-rwo'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_snapshot_class": schema.StringAttribute{
						Description:         "Volume Snapshot class to use for storage snapshots. This value varies from platform to platform. For GCP (the default), it is 'csi-gce-pd-snapshot-class'.",
						MarkdownDescription: "Volume Snapshot class to use for storage snapshots. This value varies from platform to platform. For GCP (the default), it is 'csi-gce-pd-snapshot-class'.",
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

func (r *OracleDbAnthosapisComConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_oracle_db_anthosapis_com_config_v1alpha1_manifest")

	var model OracleDbAnthosapisComConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("oracle.db.anthosapis.com/v1alpha1")
	model.Kind = pointer.String("Config")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
