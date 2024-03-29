output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_insights_operator_v1_manifest.example.yaml
  }
}
