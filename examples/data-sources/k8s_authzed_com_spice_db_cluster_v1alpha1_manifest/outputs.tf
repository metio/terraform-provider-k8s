output "manifests" {
  value = {
    "example" = data.k8s_authzed_com_spice_db_cluster_v1alpha1_manifest.example.yaml
  }
}
