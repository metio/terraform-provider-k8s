output "manifests" {
  value = {
    "example" = data.k8s_documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest.example.yaml
  }
}
