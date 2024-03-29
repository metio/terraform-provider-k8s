output "manifests" {
  value = {
    "example" = data.k8s_flux_framework_org_mini_cluster_v1alpha1_manifest.example.yaml
  }
}
