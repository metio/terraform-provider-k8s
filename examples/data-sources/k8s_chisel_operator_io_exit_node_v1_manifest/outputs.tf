output "manifests" {
  value = {
    "example" = data.k8s_chisel_operator_io_exit_node_v1_manifest.example.yaml
  }
}
