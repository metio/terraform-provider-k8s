/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComBundlesV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComBundlesV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComBundlesV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComBundlesV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComBundlesV1Alpha1ManifestData struct {
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
		CliMaxVersion   *string `tfsdk:"cli_max_version" json:"cliMaxVersion,omitempty"`
		CliMinVersion   *string `tfsdk:"cli_min_version" json:"cliMinVersion,omitempty"`
		Number          *int64  `tfsdk:"number" json:"number,omitempty"`
		VersionsBundles *[]struct {
			Aws *struct {
				ClusterTemplate *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_template" json:"clusterTemplate,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Bootstrap *struct {
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"bootstrap" json:"bootstrap,omitempty"`
			BottlerocketHostContainers *struct {
				Admin *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"admin" json:"admin,omitempty"`
				Control *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"control" json:"control,omitempty"`
				KubeadmBootstrap *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kubeadm_bootstrap" json:"kubeadmBootstrap,omitempty"`
			} `tfsdk:"bottlerocket_host_containers" json:"bottlerocketHostContainers,omitempty"`
			CertManager *struct {
				Acmesolver *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"acmesolver" json:"acmesolver,omitempty"`
				Cainjector *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cainjector" json:"cainjector,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				Ctl *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"ctl" json:"ctl,omitempty"`
				Manifest *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"manifest" json:"manifest,omitempty"`
				Startupapicheck *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"startupapicheck" json:"startupapicheck,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
				Webhook *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"webhook" json:"webhook,omitempty"`
			} `tfsdk:"cert_manager" json:"certManager,omitempty"`
			Cilium *struct {
				Cilium *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cilium" json:"cilium,omitempty"`
				HelmChart *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"helm_chart" json:"helmChart,omitempty"`
				Manifest *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"manifest" json:"manifest,omitempty"`
				Operator *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"operator" json:"operator,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"cilium" json:"cilium,omitempty"`
			CloudStack *struct {
				ClusterAPIController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_api_controller" json:"clusterAPIController,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				KubeRbacProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_rbac_proxy" json:"kubeRbacProxy,omitempty"`
				KubeVip *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_vip" json:"kubeVip,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"cloud_stack" json:"cloudStack,omitempty"`
			ClusterAPI *struct {
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"cluster_api" json:"clusterAPI,omitempty"`
			ControlPlane *struct {
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"control_plane" json:"controlPlane,omitempty"`
			Docker *struct {
				ClusterTemplate *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_template" json:"clusterTemplate,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Manager *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"manager" json:"manager,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"docker" json:"docker,omitempty"`
			EksD *struct {
				Ami *struct {
					Bottlerocket *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
						Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"bottlerocket" json:"bottlerocket,omitempty"`
				} `tfsdk:"ami" json:"ami,omitempty"`
				Channel    *string `tfsdk:"channel" json:"channel,omitempty"`
				Components *string `tfsdk:"components" json:"components,omitempty"`
				Containerd *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
					Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"containerd" json:"containerd,omitempty"`
				Crictl *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
					Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"crictl" json:"crictl,omitempty"`
				Etcdadm *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
					Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"etcdadm" json:"etcdadm,omitempty"`
				GitCommit    *string `tfsdk:"git_commit" json:"gitCommit,omitempty"`
				Imagebuilder *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
					Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"imagebuilder" json:"imagebuilder,omitempty"`
				KindNode *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kind_node" json:"kindNode,omitempty"`
				KubeVersion *string `tfsdk:"kube_version" json:"kubeVersion,omitempty"`
				ManifestUrl *string `tfsdk:"manifest_url" json:"manifestUrl,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Ova         *struct {
					Bottlerocket *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
						Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"bottlerocket" json:"bottlerocket,omitempty"`
				} `tfsdk:"ova" json:"ova,omitempty"`
				Raw *struct {
					Bottlerocket *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
						Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"bottlerocket" json:"bottlerocket,omitempty"`
				} `tfsdk:"raw" json:"raw,omitempty"`
			} `tfsdk:"eks_d" json:"eksD,omitempty"`
			Eksa *struct {
				CliTools *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cli_tools" json:"cliTools,omitempty"`
				ClusterController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_controller" json:"clusterController,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				DiagnosticCollector *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"diagnostic_collector" json:"diagnosticCollector,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"eksa" json:"eksa,omitempty"`
			EtcdadmBootstrap *struct {
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"etcdadm_bootstrap" json:"etcdadmBootstrap,omitempty"`
			EtcdadmController *struct {
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Controller *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"controller" json:"controller,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"etcdadm_controller" json:"etcdadmController,omitempty"`
			Flux *struct {
				HelmController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"helm_controller" json:"helmController,omitempty"`
				KustomizeController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kustomize_controller" json:"kustomizeController,omitempty"`
				NotificationController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"notification_controller" json:"notificationController,omitempty"`
				SourceController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"source_controller" json:"sourceController,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"flux" json:"flux,omitempty"`
			Haproxy *struct {
				Image *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
			} `tfsdk:"haproxy" json:"haproxy,omitempty"`
			Kindnetd *struct {
				Manifest *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"manifest" json:"manifest,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"kindnetd" json:"kindnetd,omitempty"`
			KubeVersion *string `tfsdk:"kube_version" json:"kubeVersion,omitempty"`
			Nutanix     *struct {
				CloudProvider *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cloud_provider" json:"cloudProvider,omitempty"`
				ClusterAPIController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_api_controller" json:"clusterAPIController,omitempty"`
				ClusterTemplate *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_template" json:"clusterTemplate,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				KubeVip *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_vip" json:"kubeVip,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"nutanix" json:"nutanix,omitempty"`
			PackageController *struct {
				CredentialProviderPackage *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"credential_provider_package" json:"credentialProviderPackage,omitempty"`
				HelmChart *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"helm_chart" json:"helmChart,omitempty"`
				PackageController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"package_controller" json:"packageController,omitempty"`
				TokenRefresher *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"token_refresher" json:"tokenRefresher,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"package_controller" json:"packageController,omitempty"`
			Snow *struct {
				BottlerocketBootstrapSnow *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"bottlerocket_bootstrap_snow" json:"bottlerocketBootstrapSnow,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				KubeVip *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_vip" json:"kubeVip,omitempty"`
				Manager *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"manager" json:"manager,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"snow" json:"snow,omitempty"`
			Tinkerbell *struct {
				ClusterAPIController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_api_controller" json:"clusterAPIController,omitempty"`
				ClusterTemplate *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_template" json:"clusterTemplate,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Envoy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"envoy" json:"envoy,omitempty"`
				KubeVip *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_vip" json:"kubeVip,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				TinkerbellStack *struct {
					Actions *struct {
						Cexec *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"cexec" json:"cexec,omitempty"`
						ImageToDisk *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"image_to_disk" json:"imageToDisk,omitempty"`
						Kexec *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"kexec" json:"kexec,omitempty"`
						OciToDisk *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"oci_to_disk" json:"ociToDisk,omitempty"`
						Reboot *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"reboot" json:"reboot,omitempty"`
						WriteFile *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"write_file" json:"writeFile,omitempty"`
					} `tfsdk:"actions" json:"actions,omitempty"`
					Boots *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"boots" json:"boots,omitempty"`
					Hegel *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"hegel" json:"hegel,omitempty"`
					Hook *struct {
						Bootkit *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"bootkit" json:"bootkit,omitempty"`
						Docker *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"docker" json:"docker,omitempty"`
						Initramfs *struct {
							Amd *struct {
								Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
								Description *string   `tfsdk:"description" json:"description,omitempty"`
								Name        *string   `tfsdk:"name" json:"name,omitempty"`
								Os          *string   `tfsdk:"os" json:"os,omitempty"`
								OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
								Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
								Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
								Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
							} `tfsdk:"amd" json:"amd,omitempty"`
							Arm *struct {
								Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
								Description *string   `tfsdk:"description" json:"description,omitempty"`
								Name        *string   `tfsdk:"name" json:"name,omitempty"`
								Os          *string   `tfsdk:"os" json:"os,omitempty"`
								OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
								Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
								Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
								Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
							} `tfsdk:"arm" json:"arm,omitempty"`
						} `tfsdk:"initramfs" json:"initramfs,omitempty"`
						Kernel *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"kernel" json:"kernel,omitempty"`
						Vmlinuz *struct {
							Amd *struct {
								Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
								Description *string   `tfsdk:"description" json:"description,omitempty"`
								Name        *string   `tfsdk:"name" json:"name,omitempty"`
								Os          *string   `tfsdk:"os" json:"os,omitempty"`
								OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
								Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
								Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
								Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
							} `tfsdk:"amd" json:"amd,omitempty"`
							Arm *struct {
								Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
								Description *string   `tfsdk:"description" json:"description,omitempty"`
								Name        *string   `tfsdk:"name" json:"name,omitempty"`
								Os          *string   `tfsdk:"os" json:"os,omitempty"`
								OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
								Sha256      *string   `tfsdk:"sha256" json:"sha256,omitempty"`
								Sha512      *string   `tfsdk:"sha512" json:"sha512,omitempty"`
								Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
							} `tfsdk:"arm" json:"arm,omitempty"`
						} `tfsdk:"vmlinuz" json:"vmlinuz,omitempty"`
					} `tfsdk:"hook" json:"hook,omitempty"`
					Rufio *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"rufio" json:"rufio,omitempty"`
					Stack *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"stack" json:"stack,omitempty"`
					Tink *struct {
						Nginx *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"nginx" json:"nginx,omitempty"`
						TinkController *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"tink_controller" json:"tinkController,omitempty"`
						TinkRelay *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"tink_relay" json:"tinkRelay,omitempty"`
						TinkRelayInit *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"tink_relay_init" json:"tinkRelayInit,omitempty"`
						TinkServer *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"tink_server" json:"tinkServer,omitempty"`
						TinkWorker *struct {
							Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
							Description *string   `tfsdk:"description" json:"description,omitempty"`
							ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
							Name        *string   `tfsdk:"name" json:"name,omitempty"`
							Os          *string   `tfsdk:"os" json:"os,omitempty"`
							OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
							Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"tink_worker" json:"tinkWorker,omitempty"`
					} `tfsdk:"tink" json:"tink,omitempty"`
					TinkerbellChart *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"tinkerbell_chart" json:"tinkerbellChart,omitempty"`
					TinkerbellCrds *struct {
						Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
						Description *string   `tfsdk:"description" json:"description,omitempty"`
						ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
						Name        *string   `tfsdk:"name" json:"name,omitempty"`
						Os          *string   `tfsdk:"os" json:"os,omitempty"`
						OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
						Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"tinkerbell_crds" json:"tinkerbellCrds,omitempty"`
				} `tfsdk:"tinkerbell_stack" json:"tinkerbellStack,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"tinkerbell" json:"tinkerbell,omitempty"`
			Upgrader *struct {
				Upgrader *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"upgrader" json:"upgrader,omitempty"`
			} `tfsdk:"upgrader" json:"upgrader,omitempty"`
			VSphere *struct {
				ClusterAPIController *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_api_controller" json:"clusterAPIController,omitempty"`
				ClusterTemplate *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"cluster_template" json:"clusterTemplate,omitempty"`
				Components *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Driver *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"driver" json:"driver,omitempty"`
				KubeProxy *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_proxy" json:"kubeProxy,omitempty"`
				KubeVip *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kube_vip" json:"kubeVip,omitempty"`
				Manager *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"manager" json:"manager,omitempty"`
				Metadata *struct {
					Uri *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Syncer *struct {
					Arch        *[]string `tfsdk:"arch" json:"arch,omitempty"`
					Description *string   `tfsdk:"description" json:"description,omitempty"`
					ImageDigest *string   `tfsdk:"image_digest" json:"imageDigest,omitempty"`
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Os          *string   `tfsdk:"os" json:"os,omitempty"`
					OsName      *string   `tfsdk:"os_name" json:"osName,omitempty"`
					Uri         *string   `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"syncer" json:"syncer,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"v_sphere" json:"vSphere,omitempty"`
		} `tfsdk:"versions_bundles" json:"versionsBundles,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComBundlesV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComBundlesV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Bundles is the Schema for the bundles API.",
		MarkdownDescription: "Bundles is the Schema for the bundles API.",
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
				Description:         "BundlesSpec defines the desired state of Bundles.",
				MarkdownDescription: "BundlesSpec defines the desired state of Bundles.",
				Attributes: map[string]schema.Attribute{
					"cli_max_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cli_min_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"number": schema.Int64Attribute{
						Description:         "Monotonically increasing release number",
						MarkdownDescription: "Monotonically increasing release number",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"versions_bundles": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws": schema.SingleNestedAttribute{
									Description:         "This field has been deprecated",
									MarkdownDescription: "This field has been deprecated",
									Attributes: map[string]schema.Attribute{
										"cluster_template": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"bootstrap": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

								"bottlerocket_host_containers": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"admin": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"control": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kubeadm_bootstrap": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"cert_manager": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"acmesolver": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"cainjector": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"ctl": schema.SingleNestedAttribute{
											Description:         "This field has been deprecated",
											MarkdownDescription: "This field has been deprecated",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"manifest": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"startupapicheck": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"webhook": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"cilium": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cilium": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"helm_chart": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"manifest": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"operator": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"cloud_stack": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cluster_api_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_rbac_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_vip": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cluster_api": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

								"control_plane": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

								"docker": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cluster_template": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"manager": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

								"eks_d": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"ami": schema.SingleNestedAttribute{
											Description:         "Ami points to a collection of AMIs built with this eks-d version",
											MarkdownDescription: "Ami points to a collection of AMIs built with this eks-d version",
											Attributes: map[string]schema.Attribute{
												"bottlerocket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sha256": schema.StringAttribute{
															Description:         "The sha256 of the asset, only applies for 'file' store",
															MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sha512": schema.StringAttribute{
															Description:         "The sha512 of the asset, only applies for 'file' store",
															MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The URI where the asset is located",
															MarkdownDescription: "The URI where the asset is located",
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

										"channel": schema.StringAttribute{
											Description:         "Release branch of the EKS-D release like 1-19, 1-20",
											MarkdownDescription: "Release branch of the EKS-D release like 1-19, 1-20",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"components": schema.StringAttribute{
											Description:         "Components refers to the url that points to the EKS-D release CRD",
											MarkdownDescription: "Components refers to the url that points to the EKS-D release CRD",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"containerd": schema.SingleNestedAttribute{
											Description:         "Containerd points to the containerd binary baked into this eks-D based node image",
											MarkdownDescription: "Containerd points to the containerd binary baked into this eks-D based node image",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha256": schema.StringAttribute{
													Description:         "The sha256 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha512": schema.StringAttribute{
													Description:         "The sha512 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The URI where the asset is located",
													MarkdownDescription: "The URI where the asset is located",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"crictl": schema.SingleNestedAttribute{
											Description:         "Crictl points to the crictl binary/tarball built for this eks-d kube version",
											MarkdownDescription: "Crictl points to the crictl binary/tarball built for this eks-d kube version",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha256": schema.StringAttribute{
													Description:         "The sha256 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha512": schema.StringAttribute{
													Description:         "The sha512 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The URI where the asset is located",
													MarkdownDescription: "The URI where the asset is located",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"etcdadm": schema.SingleNestedAttribute{
											Description:         "Etcdadm points to the etcdadm binary/tarball built for this eks-d kube version",
											MarkdownDescription: "Etcdadm points to the etcdadm binary/tarball built for this eks-d kube version",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha256": schema.StringAttribute{
													Description:         "The sha256 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha512": schema.StringAttribute{
													Description:         "The sha512 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The URI where the asset is located",
													MarkdownDescription: "The URI where the asset is located",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"git_commit": schema.StringAttribute{
											Description:         "Git commit the component is built from, before any patches",
											MarkdownDescription: "Git commit the component is built from, before any patches",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"imagebuilder": schema.SingleNestedAttribute{
											Description:         "ImageBuilder points to the image-builder binary used to build eks-D based node images",
											MarkdownDescription: "ImageBuilder points to the image-builder binary used to build eks-D based node images",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha256": schema.StringAttribute{
													Description:         "The sha256 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sha512": schema.StringAttribute{
													Description:         "The sha512 of the asset, only applies for 'file' store",
													MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The URI where the asset is located",
													MarkdownDescription: "The URI where the asset is located",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kind_node": schema.SingleNestedAttribute{
											Description:         "KindNode points to a kind image built with this eks-d version",
											MarkdownDescription: "KindNode points to a kind image built with this eks-d version",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kube_version": schema.StringAttribute{
											Description:         "Release number of EKS-D release",
											MarkdownDescription: "Release number of EKS-D release",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"manifest_url": schema.StringAttribute{
											Description:         "Url pointing to the EKS-D release manifest using which assets where created",
											MarkdownDescription: "Url pointing to the EKS-D release manifest using which assets where created",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ova": schema.SingleNestedAttribute{
											Description:         "Ova points to a collection of OVAs built with this eks-d version",
											MarkdownDescription: "Ova points to a collection of OVAs built with this eks-d version",
											Attributes: map[string]schema.Attribute{
												"bottlerocket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sha256": schema.StringAttribute{
															Description:         "The sha256 of the asset, only applies for 'file' store",
															MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sha512": schema.StringAttribute{
															Description:         "The sha512 of the asset, only applies for 'file' store",
															MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The URI where the asset is located",
															MarkdownDescription: "The URI where the asset is located",
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

										"raw": schema.SingleNestedAttribute{
											Description:         "Raw points to a collection of Raw images built with this eks-d version",
											MarkdownDescription: "Raw points to a collection of Raw images built with this eks-d version",
											Attributes: map[string]schema.Attribute{
												"bottlerocket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sha256": schema.StringAttribute{
															Description:         "The sha256 of the asset, only applies for 'file' store",
															MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sha512": schema.StringAttribute{
															Description:         "The sha512 of the asset, only applies for 'file' store",
															MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The URI where the asset is located",
															MarkdownDescription: "The URI where the asset is located",
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
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"eksa": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cli_tools": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"cluster_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"diagnostic_collector": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"etcdadm_bootstrap": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

								"etcdadm_controller": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

								"flux": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"helm_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kustomize_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"notification_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"source_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"haproxy": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"image": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
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

								"kindnetd": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"manifest": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"kube_version": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"nutanix": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cloud_provider": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"cluster_api_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"cluster_template": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_vip": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"package_controller": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"credential_provider_package": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"helm_chart": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"package_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"token_refresher": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"snow": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"bottlerocket_bootstrap_snow": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_vip": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"manager": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tinkerbell": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cluster_api_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"cluster_template": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"envoy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_vip": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"tinkerbell_stack": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"actions": schema.SingleNestedAttribute{
													Description:         "Tinkerbell Template Actions.",
													MarkdownDescription: "Tinkerbell Template Actions.",
													Attributes: map[string]schema.Attribute{
														"cexec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"image_to_disk": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"kexec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"oci_to_disk": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"reboot": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"write_file": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"boots": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_digest": schema.StringAttribute{
															Description:         "The SHA256 digest of the image manifest",
															MarkdownDescription: "The SHA256 digest of the image manifest",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The image repository, name, and tag",
															MarkdownDescription: "The image repository, name, and tag",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"hegel": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_digest": schema.StringAttribute{
															Description:         "The SHA256 digest of the image manifest",
															MarkdownDescription: "The SHA256 digest of the image manifest",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The image repository, name, and tag",
															MarkdownDescription: "The image repository, name, and tag",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"hook": schema.SingleNestedAttribute{
													Description:         "Tinkerbell hook OS.",
													MarkdownDescription: "Tinkerbell hook OS.",
													Attributes: map[string]schema.Attribute{
														"bootkit": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"docker": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"initramfs": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"amd": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"arch": schema.ListAttribute{
																			Description:         "Architectures of the asset",
																			MarkdownDescription: "Architectures of the asset",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"description": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "The asset name",
																			MarkdownDescription: "The asset name",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"os": schema.StringAttribute{
																			Description:         "Operating system of the asset",
																			MarkdownDescription: "Operating system of the asset",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("linux", "darwin", "windows"),
																			},
																		},

																		"os_name": schema.StringAttribute{
																			Description:         "Name of the OS like ubuntu, bottlerocket",
																			MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha256": schema.StringAttribute{
																			Description:         "The sha256 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha512": schema.StringAttribute{
																			Description:         "The sha512 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uri": schema.StringAttribute{
																			Description:         "The URI where the asset is located",
																			MarkdownDescription: "The URI where the asset is located",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"arm": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"arch": schema.ListAttribute{
																			Description:         "Architectures of the asset",
																			MarkdownDescription: "Architectures of the asset",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"description": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "The asset name",
																			MarkdownDescription: "The asset name",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"os": schema.StringAttribute{
																			Description:         "Operating system of the asset",
																			MarkdownDescription: "Operating system of the asset",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("linux", "darwin", "windows"),
																			},
																		},

																		"os_name": schema.StringAttribute{
																			Description:         "Name of the OS like ubuntu, bottlerocket",
																			MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha256": schema.StringAttribute{
																			Description:         "The sha256 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha512": schema.StringAttribute{
																			Description:         "The sha512 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uri": schema.StringAttribute{
																			Description:         "The URI where the asset is located",
																			MarkdownDescription: "The URI where the asset is located",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"kernel": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"vmlinuz": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"amd": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"arch": schema.ListAttribute{
																			Description:         "Architectures of the asset",
																			MarkdownDescription: "Architectures of the asset",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"description": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "The asset name",
																			MarkdownDescription: "The asset name",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"os": schema.StringAttribute{
																			Description:         "Operating system of the asset",
																			MarkdownDescription: "Operating system of the asset",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("linux", "darwin", "windows"),
																			},
																		},

																		"os_name": schema.StringAttribute{
																			Description:         "Name of the OS like ubuntu, bottlerocket",
																			MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha256": schema.StringAttribute{
																			Description:         "The sha256 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha512": schema.StringAttribute{
																			Description:         "The sha512 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uri": schema.StringAttribute{
																			Description:         "The URI where the asset is located",
																			MarkdownDescription: "The URI where the asset is located",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"arm": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"arch": schema.ListAttribute{
																			Description:         "Architectures of the asset",
																			MarkdownDescription: "Architectures of the asset",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"description": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "The asset name",
																			MarkdownDescription: "The asset name",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"os": schema.StringAttribute{
																			Description:         "Operating system of the asset",
																			MarkdownDescription: "Operating system of the asset",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("linux", "darwin", "windows"),
																			},
																		},

																		"os_name": schema.StringAttribute{
																			Description:         "Name of the OS like ubuntu, bottlerocket",
																			MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha256": schema.StringAttribute{
																			Description:         "The sha256 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha256 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sha512": schema.StringAttribute{
																			Description:         "The sha512 of the asset, only applies for 'file' store",
																			MarkdownDescription: "The sha512 of the asset, only applies for 'file' store",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uri": schema.StringAttribute{
																			Description:         "The URI where the asset is located",
																			MarkdownDescription: "The URI where the asset is located",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"rufio": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_digest": schema.StringAttribute{
															Description:         "The SHA256 digest of the image manifest",
															MarkdownDescription: "The SHA256 digest of the image manifest",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The image repository, name, and tag",
															MarkdownDescription: "The image repository, name, and tag",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"stack": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_digest": schema.StringAttribute{
															Description:         "The SHA256 digest of the image manifest",
															MarkdownDescription: "The SHA256 digest of the image manifest",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The image repository, name, and tag",
															MarkdownDescription: "The image repository, name, and tag",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"tink": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"nginx": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"tink_controller": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"tink_relay": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"tink_relay_init": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"tink_server": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"tink_worker": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"arch": schema.ListAttribute{
																	Description:         "Architectures of the asset",
																	MarkdownDescription: "Architectures of the asset",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"description": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_digest": schema.StringAttribute{
																	Description:         "The SHA256 digest of the image manifest",
																	MarkdownDescription: "The SHA256 digest of the image manifest",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The asset name",
																	MarkdownDescription: "The asset name",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"os": schema.StringAttribute{
																	Description:         "Operating system of the asset",
																	MarkdownDescription: "Operating system of the asset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("linux", "darwin", "windows"),
																	},
																},

																"os_name": schema.StringAttribute{
																	Description:         "Name of the OS like ubuntu, bottlerocket",
																	MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uri": schema.StringAttribute{
																	Description:         "The image repository, name, and tag",
																	MarkdownDescription: "The image repository, name, and tag",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"tinkerbell_chart": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_digest": schema.StringAttribute{
															Description:         "The SHA256 digest of the image manifest",
															MarkdownDescription: "The SHA256 digest of the image manifest",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The image repository, name, and tag",
															MarkdownDescription: "The image repository, name, and tag",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"tinkerbell_crds": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"arch": schema.ListAttribute{
															Description:         "Architectures of the asset",
															MarkdownDescription: "Architectures of the asset",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_digest": schema.StringAttribute{
															Description:         "The SHA256 digest of the image manifest",
															MarkdownDescription: "The SHA256 digest of the image manifest",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The asset name",
															MarkdownDescription: "The asset name",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"os": schema.StringAttribute{
															Description:         "Operating system of the asset",
															MarkdownDescription: "Operating system of the asset",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("linux", "darwin", "windows"),
															},
														},

														"os_name": schema.StringAttribute{
															Description:         "Name of the OS like ubuntu, bottlerocket",
															MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uri": schema.StringAttribute{
															Description:         "The image repository, name, and tag",
															MarkdownDescription: "The image repository, name, and tag",
															Required:            false,
															Optional:            true,
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

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"upgrader": schema.SingleNestedAttribute{
									Description:         "UpgraderBundle is a In-place Kubernetes version upgrader bundle.",
									MarkdownDescription: "UpgraderBundle is a In-place Kubernetes version upgrader bundle.",
									Attributes: map[string]schema.Attribute{
										"upgrader": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
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

								"v_sphere": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cluster_api_controller": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"cluster_template": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"driver": schema.SingleNestedAttribute{
											Description:         "This field has been deprecated",
											MarkdownDescription: "This field has been deprecated",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kube_proxy": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"kube_vip": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"manager": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"uri": schema.StringAttribute{
													Description:         "URI points to the manifest yaml file",
													MarkdownDescription: "URI points to the manifest yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"syncer": schema.SingleNestedAttribute{
											Description:         "This field has been deprecated",
											MarkdownDescription: "This field has been deprecated",
											Attributes: map[string]schema.Attribute{
												"arch": schema.ListAttribute{
													Description:         "Architectures of the asset",
													MarkdownDescription: "Architectures of the asset",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_digest": schema.StringAttribute{
													Description:         "The SHA256 digest of the image manifest",
													MarkdownDescription: "The SHA256 digest of the image manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The asset name",
													MarkdownDescription: "The asset name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.StringAttribute{
													Description:         "Operating system of the asset",
													MarkdownDescription: "Operating system of the asset",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("linux", "darwin", "windows"),
													},
												},

												"os_name": schema.StringAttribute{
													Description:         "Name of the OS like ubuntu, bottlerocket",
													MarkdownDescription: "Name of the OS like ubuntu, bottlerocket",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "The image repository, name, and tag",
													MarkdownDescription: "The image repository, name, and tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"version": schema.StringAttribute{
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

func (r *AnywhereEksAmazonawsComBundlesV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_bundles_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComBundlesV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("Bundles")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
