output "resources" {
  value = {
    "minimal" = k8s_app_lightbend_com_akka_cluster_v1alpha1.minimal.yaml
    "example" = k8s_app_lightbend_com_akka_cluster_v1alpha1.example.yaml
  }
}
