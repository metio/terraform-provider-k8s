/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sharedresource_openshift_io_v1alpha1

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
	_ datasource.DataSource = &SharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest{}
)

func NewSharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest() datasource.DataSource {
	return &SharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest{}
}

type SharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest struct{}

type SharedresourceOpenshiftIoSharedSecretV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Description *string `tfsdk:"description" json:"description,omitempty"`
		SecretRef   *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sharedresource_openshift_io_shared_secret_v1alpha1_manifest"
}

func (r *SharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SharedSecret allows a Secret to be shared across namespaces. Pods can mount the shared Secret by adding a CSI volume to the pod specification using the 'csi.sharedresource.openshift.io' CSI driver and a reference to the SharedSecret in the volume attributes:  spec: volumes: - name: shared-secret csi: driver: csi.sharedresource.openshift.io volumeAttributes: sharedSecret: my-share  For the mount to be successful, the pod's service account must be granted permission to 'use' the named SharedSecret object within its namespace with an appropriate Role and RoleBinding. For compactness, here are example 'oc' invocations for creating such Role and RoleBinding objects.  'oc create role shared-resource-my-share --verb=use --resource=sharedsecrets.sharedresource.openshift.io --resource-name=my-share' 'oc create rolebinding shared-resource-my-share --role=shared-resource-my-share --serviceaccount=my-namespace:default'  Shared resource objects, in this case Secrets, have default permissions of list, get, and watch for system authenticated users.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.",
		MarkdownDescription: "SharedSecret allows a Secret to be shared across namespaces. Pods can mount the shared Secret by adding a CSI volume to the pod specification using the 'csi.sharedresource.openshift.io' CSI driver and a reference to the SharedSecret in the volume attributes:  spec: volumes: - name: shared-secret csi: driver: csi.sharedresource.openshift.io volumeAttributes: sharedSecret: my-share  For the mount to be successful, the pod's service account must be granted permission to 'use' the named SharedSecret object within its namespace with an appropriate Role and RoleBinding. For compactness, here are example 'oc' invocations for creating such Role and RoleBinding objects.  'oc create role shared-resource-my-share --verb=use --resource=sharedsecrets.sharedresource.openshift.io --resource-name=my-share' 'oc create rolebinding shared-resource-my-share --role=shared-resource-my-share --serviceaccount=my-namespace:default'  Shared resource objects, in this case Secrets, have default permissions of list, get, and watch for system authenticated users.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "spec is the specification of the desired shared secret",
				MarkdownDescription: "spec is the specification of the desired shared secret",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "description is a user readable explanation of what the backing resource provides.",
						MarkdownDescription: "description is a user readable explanation of what the backing resource provides.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "secretRef is a reference to the Secret to share",
						MarkdownDescription: "secretRef is a reference to the Secret to share",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name represents the name of the Secret that is being referenced.",
								MarkdownDescription: "name represents the name of the Secret that is being referenced.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace represents the namespace where the referenced Secret is located.",
								MarkdownDescription: "namespace represents the namespace where the referenced Secret is located.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *SharedresourceOpenshiftIoSharedSecretV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sharedresource_openshift_io_shared_secret_v1alpha1_manifest")

	var model SharedresourceOpenshiftIoSharedSecretV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("sharedresource.openshift.io/v1alpha1")
	model.Kind = pointer.String("SharedSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
