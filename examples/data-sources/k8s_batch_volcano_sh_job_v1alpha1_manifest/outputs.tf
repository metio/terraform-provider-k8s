output "manifests" {
  value = {
    "example" = data.k8s_batch_volcano_sh_job_v1alpha1_manifest.example.yaml
  }
}
