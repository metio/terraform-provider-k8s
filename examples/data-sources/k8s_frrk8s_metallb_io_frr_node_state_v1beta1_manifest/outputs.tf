output "manifests" {
  value = {
    "example" = data.k8s_frrk8s_metallb_io_frr_node_state_v1beta1_manifest.example.yaml
  }
}
