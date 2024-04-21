output "manifests" {
  value = {
    "example" = data.k8s_k8s_mariadb_com_max_scale_v1alpha1_manifest.example.yaml
  }
}
