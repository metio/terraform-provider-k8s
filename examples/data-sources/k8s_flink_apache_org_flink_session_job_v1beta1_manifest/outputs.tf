output "manifests" {
  value = {
    "example" = data.k8s_flink_apache_org_flink_session_job_v1beta1_manifest.example.yaml
  }
}
