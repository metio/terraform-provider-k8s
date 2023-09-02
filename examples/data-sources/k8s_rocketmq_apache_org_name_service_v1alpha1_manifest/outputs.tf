output "manifests" {
  value = {
    "example" = data.k8s_rocketmq_apache_org_name_service_v1alpha1_manifest.example.yaml
  }
}
