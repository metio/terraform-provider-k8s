output "manifests" {
  value = {
    "example" = data.k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha4_manifest.example.yaml
  }
}
