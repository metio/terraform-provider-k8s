output "manifests" {
  value = {
    "example" = data.k8s_insights_openshift_io_data_gather_v1alpha1_manifest.example.yaml
  }
}
