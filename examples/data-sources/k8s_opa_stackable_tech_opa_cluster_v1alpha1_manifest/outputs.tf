output "manifests" {
  value = {
    "example" = data.k8s_opa_stackable_tech_opa_cluster_v1alpha1_manifest.example.yaml
  }
}
