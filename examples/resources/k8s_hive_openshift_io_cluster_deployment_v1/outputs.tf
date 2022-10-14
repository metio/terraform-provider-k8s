output "resources" {
  value = {
    "minimal" = k8s_hive_openshift_io_cluster_deployment_v1.minimal.yaml
  }
}
