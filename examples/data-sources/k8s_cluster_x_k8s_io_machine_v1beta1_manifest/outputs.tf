output "manifests" {
  value = {
    "example" = data.k8s_cluster_x_k8s_io_machine_v1beta1_manifest.example.yaml
  }
}
