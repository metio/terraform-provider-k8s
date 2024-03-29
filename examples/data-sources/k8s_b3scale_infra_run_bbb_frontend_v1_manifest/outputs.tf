output "manifests" {
  value = {
    "example" = data.k8s_b3scale_infra_run_bbb_frontend_v1_manifest.example.yaml
  }
}
