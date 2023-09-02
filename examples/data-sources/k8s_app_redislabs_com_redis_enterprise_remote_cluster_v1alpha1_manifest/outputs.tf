output "manifests" {
  value = {
    "example" = data.k8s_app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest.example.yaml
  }
}
