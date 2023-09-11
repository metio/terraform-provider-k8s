/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strconv"
)

type K8sProvider struct {
	client *dynamic.Interface
}

type K8sProviderModel struct {
	Kubeconfig     types.String `tfsdk:"kubeconfig"`
	Context        types.String `tfsdk:"context"`
	FieldManager   types.String `tfsdk:"field_manager"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts"`
	Timeout        types.Int64  `tfsdk:"timeout"`
	Offline        types.Bool   `tfsdk:"offline"`
}

var _ provider.Provider = &K8sProvider{}

func New() provider.Provider {
	return &K8sProvider{}
}

func NewWithClient(client dynamic.Interface) provider.Provider {
	return &K8sProvider{client: &client}
}

func (p *K8sProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "k8s"
}

func (p *K8sProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Provider for Kubernetes resources using server-side apply. Requires Terraform 1.0 or later.",
		MarkdownDescription: "Provider for [Kubernetes](https://kubernetes.io/) resources using [server-side apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/). Requires Terraform 1.0 or later.",
		Attributes: map[string]schema.Attribute{
			"kubeconfig": schema.StringAttribute{
				Description:         "An explicit path to a kubeconfig file. Can be specified with the 'TF_K8S_CONFIG' environment variable. Uses Kubernetes defaults if not specified ('KUBECONFIG', or your home directory).",
				MarkdownDescription: "An explicit path to a kubeconfig file. Can be specified with the `TF_K8S_CONFIG` environment variable. Uses Kubernetes defaults if not specified (`KUBECONFIG`, or your home directory).",
				Required:            false,
				Optional:            true,
				Sensitive:           false,
			},
			"context": schema.StringAttribute{
				Description:         "The context to use from your kubeconfig. Can be specified with the 'TF_K8S_CONTEXT' environment variable. Defaults to the current context in your config.",
				MarkdownDescription: "The context to use from your kubeconfig. Can be specified with the `TF_K8S_CONTEXT` environment variable. Defaults to the current context in your config.",
				Required:            false,
				Optional:            true,
				Sensitive:           false,
			},
			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. Can be specified with the 'TF_K8S_FIELD_MANAGER' environment variable. Defaults to 'terraform-provider-k8s'.",
				MarkdownDescription: "The name of the manager used to track field ownership. Can be specified with the `TF_K8S_FIELD_MANAGER` environment variable. Defaults to `terraform-provider-k8s`.",
				Required:            false,
				Optional:            true,
				Sensitive:           false,
			},
			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. Can be specified with the 'TF_K8S_FORCE_CONFLICTS' environment variable. Defaults to 'true'.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. Can be specified with the `TF_K8S_FORCE_CONFLICTS` environment variable. Defaults to `true`.",
				Required:            false,
				Optional:            true,
				Sensitive:           false,
			},
			"timeout": schema.Int64Attribute{
				Description:         "The timeout to apply for HTTP requests in seconds. Can be specified with the 'TF_K8S_TIMEOUT' environment variable. Defaults to '32'.",
				MarkdownDescription: "The timeout to apply for HTTP requests in seconds. Can be specified with the `TF_K8S_TIMEOUT` environment variable. Defaults to `32`.",
				Required:            false,
				Optional:            true,
				Sensitive:           false,
			},
			"offline": schema.BoolAttribute{
				Description:         "Enable offline mode for this provider. In offline mode, no connection to a kubernetes cluster will be performed, therefore no resource or data source can be created except manifest data sources (those ending with _manifest). Can be specified with the 'TF_K8S_OFFLINE' environment variable. Defaults to 'false'.",
				MarkdownDescription: "Enable offline mode for this provider. In offline mode, no connection to a kubernetes cluster will be performed, therefore no resource or data source can be created except manifest data sources (those ending with _manifest). Can be specified with the `TF_K8S_OFFLINE` environment variable. Defaults to `false`.",
				Required:            false,
				Optional:            true,
				Sensitive:           false,
			},
		},
	}
}

func (p *K8sProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Kubernetes client")

	var config K8sProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Kubeconfig.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("kubeconfig"),
			"Unknown kubeconfig",
			"The provider cannot create a Kubernetes client as there is an unknown configuration value for the kubeconfig option. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_K8S_CONFIG environment variable.",
		)
	}

	if config.Context.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("context"),
			"Unknown context",
			"The provider cannot create a Kubernetes client as there is an unknown configuration value for the context option. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_K8S_CONTEXT environment variable.",
		)
	}

	if config.FieldManager.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("field_manager"),
			"Unknown field_manager",
			"The provider cannot create a Kubernetes client as there is an unknown configuration value for the field_manager option. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_K8S_FIELD_MANAGER environment variable.",
		)
	}

	if config.ForceConflicts.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("force_conflicts"),
			"Unknown force_conflicts",
			"The provider cannot create a Kubernetes client as there is an unknown configuration value for the force_conflicts option. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_K8S_FORCE_CONFLICTS environment variable.",
		)
	}

	if config.Timeout.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("timeout"),
			"Unknown timeout",
			"The provider cannot create a Kubernetes client as there is an unknown configuration value for the timeout option. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_K8S_TIMEOUT environment variable.",
		)
	}

	if config.Offline.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("offline"),
			"Unknown offline",
			"The provider cannot create a Kubernetes client as there is an unknown configuration value for the offline option. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_K8S_OFFLINE environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	kubeconfig := os.Getenv("TF_K8S_CONFIG")
	clientContext := os.Getenv("TF_K8S_CONTEXT")
	fieldManager := os.Getenv("TF_K8S_FIELD_MANAGER")
	forceConflicts := os.Getenv("TF_K8S_FORCE_CONFLICTS")
	timeout := os.Getenv("TF_K8S_TIMEOUT")
	offline := os.Getenv("TF_K8S_OFFLINE")

	if !config.Kubeconfig.IsNull() {
		kubeconfig = config.Kubeconfig.ValueString()
	}

	if !config.Context.IsNull() {
		clientContext = config.Context.ValueString()
	}

	if !config.FieldManager.IsNull() {
		fieldManager = config.FieldManager.ValueString()
	}

	if !config.ForceConflicts.IsNull() {
		forceConflicts = strconv.FormatBool(config.ForceConflicts.ValueBool())
	}

	if !config.Timeout.IsNull() {
		timeout = strconv.FormatInt(config.Timeout.ValueInt64(), 10)
	}

	if !config.Offline.IsNull() {
		offline = strconv.FormatBool(config.Offline.ValueBool())
	}

	if fieldManager == "" {
		fieldManager = "terraform-provider-k8s"
	}

	if forceConflicts == "" {
		forceConflicts = "true"
	}

	if timeout == "" {
		timeout = "32"
	}

	if offline == "" {
		offline = "false"
	}

	conflicts, err := strconv.ParseBool(forceConflicts)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("force_conflicts"),
			"Invalid force_conflicts value",
			"The supplied force_conflicts value cannot be parsed into a bool: "+err.Error(),
		)
	}

	offlineMode, err := strconv.ParseBool(offline)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("offline"),
			"Invalid offline value",
			"The supplied offline value cannot be parsed into a bool: "+err.Error(),
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "kubeconfig", kubeconfig)
	ctx = tflog.SetField(ctx, "context", clientContext)
	ctx = tflog.SetField(ctx, "field_manager", fieldManager)
	ctx = tflog.SetField(ctx, "force_conflicts", forceConflicts)
	ctx = tflog.SetField(ctx, "timeout", timeout)
	ctx = tflog.SetField(ctx, "offline", offline)

	if offlineMode {
		resp.DataSourceData = &utilities.DataSourceData{
			Offline: offlineMode,
		}
		resp.ResourceData = &utilities.ResourceData{
			Offline: offlineMode,
		}
	} else {
		tflog.Debug(ctx, "Creating Kubernetes client")

		var client dynamic.Interface
		if p.client == nil {
			loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
			if kubeconfig != "" {
				loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}
			}

			configOverrides := &clientcmd.ConfigOverrides{
				Timeout: timeout,
			}

			if clientContext != "" {
				configOverrides.CurrentContext = clientContext
			}

			kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

			rawConfig, err := kubeConfig.RawConfig()
			if err != nil {
				resp.Diagnostics.AddError(
					"Unable to create Kubernetes client",
					fmt.Sprintf("An unexpected error occurred when creating the Kubernetes client. "+
						"If the error is not clear, please contact the provider developers.\n\n"+
						"Kubernetes client error (%T): %s", err, err.Error()),
				)
				return
			}
			err = clientcmd.ConfirmUsable(rawConfig, rawConfig.CurrentContext)
			if err != nil {
				resp.Diagnostics.AddError(
					"Unable to create Kubernetes client",
					fmt.Sprintf("An unexpected error occurred when creating the Kubernetes client. "+
						"If the error is not clear, please contact the provider developers.\n\n"+
						"Kubernetes client error (%T): %s", err, err.Error()),
				)
				return
			}

			clientConfig, err := kubeConfig.ClientConfig()
			if err != nil {
				resp.Diagnostics.AddError(
					"Unable to create Kubernetes client",
					fmt.Sprintf("An unexpected error occurred when creating the Kubernetes client. "+
						"If the error is not clear, please contact the provider developers.\n\n"+
						"Kubernetes client error (%T): %s", err, err.Error()),
				)
				return
			}

			client, err = dynamic.NewForConfig(clientConfig)
			if err != nil {
				resp.Diagnostics.AddError(
					"Unable to create Kubernetes client",
					fmt.Sprintf("An unexpected error occurred when creating the Kubernetes client. "+
						"If the error is not clear, please contact the provider developers.\n\n"+
						"Kubernetes client error (%T): %s", err, err.Error()),
				)
				return
			}
		} else {
			client = *p.client
		}

		resp.DataSourceData = &utilities.DataSourceData{
			Client:  client,
			Offline: offlineMode,
		}
		resp.ResourceData = &utilities.ResourceData{
			Client:         client,
			FieldManager:   fieldManager,
			ForceConflicts: conflicts,
			Offline:        offlineMode,
		}

		tflog.Info(ctx, "Configured Kubernetes client")
	}
}

func (p *K8sProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return allDataSources()
}

func (p *K8sProvider) Resources(_ context.Context) []func() resource.Resource {
	return allResources()
}
