output "manifests" {
  value = {
    "example" = data.k8s_jobset_x_k8s_io_job_set_v1alpha2_manifest.example.yaml
  }
}
