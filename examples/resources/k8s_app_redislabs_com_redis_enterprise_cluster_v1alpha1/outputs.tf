output "resources" {
  value = {
    "minimal" = k8s_app_redislabs_com_redis_enterprise_cluster_v1alpha1.minimal.yaml
    "example" = k8s_app_redislabs_com_redis_enterprise_cluster_v1alpha1.example.yaml
  }
}
