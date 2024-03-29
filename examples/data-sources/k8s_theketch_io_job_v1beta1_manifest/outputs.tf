output "manifests" {
  value = {
    "example" = data.k8s_theketch_io_job_v1beta1_manifest.example.yaml
  }
}
