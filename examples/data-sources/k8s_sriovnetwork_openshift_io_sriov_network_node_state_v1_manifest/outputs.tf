output "manifests" {
  value = {
    "example" = data.k8s_sriovnetwork_openshift_io_sriov_network_node_state_v1_manifest.example.yaml
  }
}
