output "manifests" {
  value = {
    "example" = data.k8s_ray_io_ray_job_v1alpha1_manifest.example.yaml
  }
}
