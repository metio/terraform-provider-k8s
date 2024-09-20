output "manifests" {
  value = {
    "example" = data.k8s_sriovnetwork_openshift_io_sriov_network_v1_manifest.example.yaml
  }
}
