output "manifests" {
  value = {
    "example" = data.k8s_keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest.example.yaml
  }
}
