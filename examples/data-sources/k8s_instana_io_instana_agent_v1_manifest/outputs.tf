output "manifests" {
  value = {
    "example" = data.k8s_instana_io_instana_agent_v1_manifest.example.yaml
  }
}
