/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

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
	_ datasource.DataSource = &CiliumIoCiliumIdentityV2Manifest{}
)

func NewCiliumIoCiliumIdentityV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumIdentityV2Manifest{}
}

type CiliumIoCiliumIdentityV2Manifest struct{}

type CiliumIoCiliumIdentityV2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Security_labels *map[string]string `tfsdk:"security_labels" json:"security-labels,omitempty"`
}

func (r *CiliumIoCiliumIdentityV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_identity_v2_manifest"
}

func (r *CiliumIoCiliumIdentityV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumIdentity is a CRD that represents an identity managed by Cilium. It is intended as a backing store for identity allocation, acting as the global coordination backend, and can be used in place of a KVStore (such as etcd). The name of the CRD is the numeric identity and the labels on the CRD object are the kubernetes sourced labels seen by cilium. This is currently the only label source possible when running under kubernetes. Non-kubernetes labels are filtered but all labels, from all sources, are places in the SecurityLabels field. These also include the source and are used to define the identity. The labels under metav1.ObjectMeta can be used when searching for CiliumIdentity instances that include particular labels. This can be done with invocations such as: kubectl get ciliumid -l 'foo=bar'",
		MarkdownDescription: "CiliumIdentity is a CRD that represents an identity managed by Cilium. It is intended as a backing store for identity allocation, acting as the global coordination backend, and can be used in place of a KVStore (such as etcd). The name of the CRD is the numeric identity and the labels on the CRD object are the kubernetes sourced labels seen by cilium. This is currently the only label source possible when running under kubernetes. Non-kubernetes labels are filtered but all labels, from all sources, are places in the SecurityLabels field. These also include the source and are used to define the identity. The labels under metav1.ObjectMeta can be used when searching for CiliumIdentity instances that include particular labels. This can be done with invocations such as: kubectl get ciliumid -l 'foo=bar'",
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

			"security_labels": schema.MapAttribute{
				Description:         "SecurityLabels is the source-of-truth set of labels for this identity.",
				MarkdownDescription: "SecurityLabels is the source-of-truth set of labels for this identity.",
				ElementType:         types.StringType,
				Required:            true,
				Optional:            false,
				Computed:            false,
			},
		},
	}
}

func (r *CiliumIoCiliumIdentityV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_identity_v2_manifest")

	var model CiliumIoCiliumIdentityV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumIdentity")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
