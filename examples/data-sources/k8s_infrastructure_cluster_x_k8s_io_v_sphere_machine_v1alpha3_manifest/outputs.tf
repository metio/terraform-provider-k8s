output "manifests" {
  value = {
    "example" = data.k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha3_manifest.example.yaml
  }
}