output "manifests" {
  value = {
    "example" = data.k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest.example.yaml
  }
}
