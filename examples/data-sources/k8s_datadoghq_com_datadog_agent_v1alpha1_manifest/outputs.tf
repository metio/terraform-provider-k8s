output "manifests" {
  value = {
    "example" = data.k8s_datadoghq_com_datadog_agent_v1alpha1_manifest.example.yaml
  }
}
