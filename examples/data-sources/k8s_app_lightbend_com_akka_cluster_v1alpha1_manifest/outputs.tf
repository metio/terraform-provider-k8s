output "manifests" {
  value = {
    "example" = data.k8s_app_lightbend_com_akka_cluster_v1alpha1_manifest.example.yaml
  }
}
