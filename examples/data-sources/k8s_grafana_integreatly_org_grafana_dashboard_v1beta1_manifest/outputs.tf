output "manifests" {
  value = {
    "example" = data.k8s_grafana_integreatly_org_grafana_dashboard_v1beta1_manifest.example.yaml
  }
}
