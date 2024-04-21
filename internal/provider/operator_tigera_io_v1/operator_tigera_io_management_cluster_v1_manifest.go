/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoManagementClusterV1Manifest{}
)

func NewOperatorTigeraIoManagementClusterV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoManagementClusterV1Manifest{}
}

type OperatorTigeraIoManagementClusterV1Manifest struct{}

type OperatorTigeraIoManagementClusterV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Address *string `tfsdk:"address" json:"address,omitempty"`
		Tls     *struct {
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoManagementClusterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_management_cluster_v1_manifest"
}

func (r *OperatorTigeraIoManagementClusterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The presence of ManagementCluster in your cluster, will configure it to be the management plane to which managedclusters can connect. At most one instance of this resource is supported. It must be named 'tigera-secure'.",
		MarkdownDescription: "The presence of ManagementCluster in your cluster, will configure it to be the management plane to which managedclusters can connect. At most one instance of this resource is supported. It must be named 'tigera-secure'.",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "ManagementClusterSpec defines the desired state of a ManagementCluster",
				MarkdownDescription: "ManagementClusterSpec defines the desired state of a ManagementCluster",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "This field specifies the externally reachable address to which your managed cluster will connect. When a managedcluster is added, this field is used to populate an easy-to-apply manifest that will connect both clusters.Valid examples are: '0.0.0.0:31000', 'example.com:32000', '[::1]:32500'",
						MarkdownDescription: "This field specifies the externally reachable address to which your managed cluster will connect. When a managedcluster is added, this field is used to populate an easy-to-apply manifest that will connect both clusters.Valid examples are: '0.0.0.0:31000', 'example.com:32000', '[::1]:32500'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS provides options for configuring how Managed Clusters can establish an mTLS connection with the Management Cluster.",
						MarkdownDescription: "TLS provides options for configuring how Managed Clusters can establish an mTLS connection with the Management Cluster.",
						Attributes: map[string]schema.Attribute{
							"secret_name": schema.StringAttribute{
								Description:         "SecretName indicates the name of the secret in the tigera-operator namespace that contains the private key and certificate that the management cluster uses when it listens for incoming connections.When set to tigera-management-cluster-connection voltron will use the same cert bundle which Guardian client certs are signed with.When set to manager-tls, voltron will use the same cert bundle which Manager UI is served with.This cert bundle must be a publicly signed cert created by the user.Note that Tigera Operator will generate a self-signed manager-tls cert if one does not exist,and use of that cert will result in Guardian being unable to verify Voltron's identity.If changed on a running cluster with connected managed clusters, all managed clusters will disconnect as they will no longer be able to verify Voltron's identity.To reconnect existing managed clusters, change the tls.ca of the  managed clusters' ManagementClusterConnection resource.One of: tigera-management-cluster-connection, manager-tlsDefault: tigera-management-cluster-connection",
								MarkdownDescription: "SecretName indicates the name of the secret in the tigera-operator namespace that contains the private key and certificate that the management cluster uses when it listens for incoming connections.When set to tigera-management-cluster-connection voltron will use the same cert bundle which Guardian client certs are signed with.When set to manager-tls, voltron will use the same cert bundle which Manager UI is served with.This cert bundle must be a publicly signed cert created by the user.Note that Tigera Operator will generate a self-signed manager-tls cert if one does not exist,and use of that cert will result in Guardian being unable to verify Voltron's identity.If changed on a running cluster with connected managed clusters, all managed clusters will disconnect as they will no longer be able to verify Voltron's identity.To reconnect existing managed clusters, change the tls.ca of the  managed clusters' ManagementClusterConnection resource.One of: tigera-management-cluster-connection, manager-tlsDefault: tigera-management-cluster-connection",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("tigera-management-cluster-connection", "manager-tls"),
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

func (r *OperatorTigeraIoManagementClusterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_management_cluster_v1_manifest")

	var model OperatorTigeraIoManagementClusterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("ManagementCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
