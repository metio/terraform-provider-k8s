output "manifests" {
  value = {
    "example" = data.k8s_logging_banzaicloud_io_node_agent_v1beta1_manifest.example.yaml
  }
}
