output "manifests" {
  value = {
    "example" = data.k8s_sriovnetwork_openshift_io_sriov_network_pool_config_v1_manifest.example.yaml
  }
}
