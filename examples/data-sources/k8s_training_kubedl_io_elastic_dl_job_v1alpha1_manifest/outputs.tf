output "manifests" {
  value = {
    "example" = data.k8s_training_kubedl_io_elastic_dl_job_v1alpha1_manifest.example.yaml
  }
}
