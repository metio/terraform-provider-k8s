output "manifests" {
  value = {
    "example" = data.k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta1_manifest.example.yaml
  }
}
