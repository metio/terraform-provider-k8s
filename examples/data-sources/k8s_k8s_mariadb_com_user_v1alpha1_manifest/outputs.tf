output "manifests" {
  value = {
    "example" = data.k8s_k8s_mariadb_com_user_v1alpha1_manifest.example.yaml
  }
}
