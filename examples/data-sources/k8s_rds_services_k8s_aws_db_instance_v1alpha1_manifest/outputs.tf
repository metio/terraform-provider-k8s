output "manifests" {
  value = {
    "example" = data.k8s_rds_services_k8s_aws_db_instance_v1alpha1_manifest.example.yaml
  }
}
