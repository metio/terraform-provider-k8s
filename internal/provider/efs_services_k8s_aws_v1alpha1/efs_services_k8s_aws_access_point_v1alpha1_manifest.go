/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package efs_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &EfsServicesK8SAwsAccessPointV1Alpha1Manifest{}
)

func NewEfsServicesK8SAwsAccessPointV1Alpha1Manifest() datasource.DataSource {
	return &EfsServicesK8SAwsAccessPointV1Alpha1Manifest{}
}

type EfsServicesK8SAwsAccessPointV1Alpha1Manifest struct{}

type EfsServicesK8SAwsAccessPointV1Alpha1ManifestData struct {
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
		FileSystemID  *string `tfsdk:"file_system_id" json:"fileSystemID,omitempty"`
		FileSystemRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"file_system_ref" json:"fileSystemRef,omitempty"`
		PosixUser *struct {
			Gid           *int64    `tfsdk:"gid" json:"gid,omitempty"`
			SecondaryGIDs *[]string `tfsdk:"secondary_gi_ds" json:"secondaryGIDs,omitempty"`
			Uid           *int64    `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"posix_user" json:"posixUser,omitempty"`
		RootDirectory *struct {
			CreationInfo *struct {
				OwnerGID    *int64  `tfsdk:"owner_gid" json:"ownerGID,omitempty"`
				OwnerUID    *int64  `tfsdk:"owner_uid" json:"ownerUID,omitempty"`
				Permissions *string `tfsdk:"permissions" json:"permissions,omitempty"`
			} `tfsdk:"creation_info" json:"creationInfo,omitempty"`
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"root_directory" json:"rootDirectory,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EfsServicesK8SAwsAccessPointV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_efs_services_k8s_aws_access_point_v1alpha1_manifest"
}

func (r *EfsServicesK8SAwsAccessPointV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AccessPoint is the Schema for the AccessPoints API",
		MarkdownDescription: "AccessPoint is the Schema for the AccessPoints API",
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
				Description:         "AccessPointSpec defines the desired state of AccessPoint.",
				MarkdownDescription: "AccessPointSpec defines the desired state of AccessPoint.",
				Attributes: map[string]schema.Attribute{
					"file_system_id": schema.StringAttribute{
						Description:         "The ID of the EFS file system that the access point provides access to.",
						MarkdownDescription: "The ID of the EFS file system that the access point provides access to.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"file_system_ref": schema.SingleNestedAttribute{
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

					"posix_user": schema.SingleNestedAttribute{
						Description:         "The operating system user and group applied to all file system requests madeusing the access point.",
						MarkdownDescription: "The operating system user and group applied to all file system requests madeusing the access point.",
						Attributes: map[string]schema.Attribute{
							"gid": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secondary_gi_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.Int64Attribute{
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

					"root_directory": schema.SingleNestedAttribute{
						Description:         "Specifies the directory on the EFS file system that the access point exposesas the root directory of your file system to NFS clients using the accesspoint. The clients using the access point can only access the root directoryand below. If the RootDirectory > Path specified does not exist, Amazon EFScreates it and applies the CreationInfo settings when a client connects toan access point. When specifying a RootDirectory, you must provide the Path,and the CreationInfo.Amazon EFS creates a root directory only if you have provided the CreationInfo:OwnUid, OwnGID, and permissions for the directory. If you do not providethis information, Amazon EFS does not create the root directory. If the rootdirectory does not exist, attempts to mount using the access point will fail.",
						MarkdownDescription: "Specifies the directory on the EFS file system that the access point exposesas the root directory of your file system to NFS clients using the accesspoint. The clients using the access point can only access the root directoryand below. If the RootDirectory > Path specified does not exist, Amazon EFScreates it and applies the CreationInfo settings when a client connects toan access point. When specifying a RootDirectory, you must provide the Path,and the CreationInfo.Amazon EFS creates a root directory only if you have provided the CreationInfo:OwnUid, OwnGID, and permissions for the directory. If you do not providethis information, Amazon EFS does not create the root directory. If the rootdirectory does not exist, attempts to mount using the access point will fail.",
						Attributes: map[string]schema.Attribute{
							"creation_info": schema.SingleNestedAttribute{
								Description:         "Required if the RootDirectory > Path specified does not exist. Specifiesthe POSIX IDs and permissions to apply to the access point's RootDirectory> Path. If the access point root directory does not exist, EFS creates itwith these settings when a client connects to the access point. When specifyingCreationInfo, you must include values for all properties.Amazon EFS creates a root directory only if you have provided the CreationInfo:OwnUid, OwnGID, and permissions for the directory. If you do not providethis information, Amazon EFS does not create the root directory. If the rootdirectory does not exist, attempts to mount using the access point will fail.If you do not provide CreationInfo and the specified RootDirectory does notexist, attempts to mount the file system using the access point will fail.",
								MarkdownDescription: "Required if the RootDirectory > Path specified does not exist. Specifiesthe POSIX IDs and permissions to apply to the access point's RootDirectory> Path. If the access point root directory does not exist, EFS creates itwith these settings when a client connects to the access point. When specifyingCreationInfo, you must include values for all properties.Amazon EFS creates a root directory only if you have provided the CreationInfo:OwnUid, OwnGID, and permissions for the directory. If you do not providethis information, Amazon EFS does not create the root directory. If the rootdirectory does not exist, attempts to mount using the access point will fail.If you do not provide CreationInfo and the specified RootDirectory does notexist, attempts to mount the file system using the access point will fail.",
								Attributes: map[string]schema.Attribute{
									"owner_gid": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"owner_uid": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"permissions": schema.StringAttribute{
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

							"path": schema.StringAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "Creates tags associated with the access point. Each tag is a key-value pair,each key must be unique. For more information, see Tagging Amazon Web Servicesresources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)in the Amazon Web Services General Reference Guide.",
						MarkdownDescription: "Creates tags associated with the access point. Each tag is a key-value pair,each key must be unique. For more information, see Tagging Amazon Web Servicesresources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)in the Amazon Web Services General Reference Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *EfsServicesK8SAwsAccessPointV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_efs_services_k8s_aws_access_point_v1alpha1_manifest")

	var model EfsServicesK8SAwsAccessPointV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("efs.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("AccessPoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
