output "manifests" {
  value = {
    "example" = data.k8s_sriovnetwork_openshift_io_sriov_operator_config_v1_manifest.example.yaml
  }
}
