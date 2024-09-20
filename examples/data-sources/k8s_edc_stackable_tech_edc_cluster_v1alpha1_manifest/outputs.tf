output "manifests" {
  value = {
    "example" = data.k8s_edc_stackable_tech_edc_cluster_v1alpha1_manifest.example.yaml
  }
}
