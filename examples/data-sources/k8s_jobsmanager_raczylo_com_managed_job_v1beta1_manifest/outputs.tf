output "manifests" {
  value = {
    "example" = data.k8s_jobsmanager_raczylo_com_managed_job_v1beta1_manifest.example.yaml
  }
}
