output "manifests" {
  value = {
    "example" = data.k8s_infrastructure_cluster_x_k8s_io_kubevirt_machine_v1alpha1_manifest.example.yaml
  }
}
