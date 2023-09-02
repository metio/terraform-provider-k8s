output "manifests" {
  value = {
    "example" = data.k8s_keda_sh_scaled_job_v1alpha1_manifest.example.yaml
  }
}
