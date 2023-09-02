output "manifests" {
  value = {
    "example" = data.k8s_sematext_com_sematext_agent_v1_manifest.example.yaml
  }
}
