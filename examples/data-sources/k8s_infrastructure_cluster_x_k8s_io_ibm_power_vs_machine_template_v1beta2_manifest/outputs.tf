output "manifests" {
  value = {
    "example" = data.k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_template_v1beta2_manifest.example.yaml
  }
}
