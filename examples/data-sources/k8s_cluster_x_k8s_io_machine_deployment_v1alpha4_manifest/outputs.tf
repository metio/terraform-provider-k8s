output "manifests" {
  value = {
    "example" = data.k8s_cluster_x_k8s_io_machine_deployment_v1alpha4_manifest.example.yaml
  }
}
