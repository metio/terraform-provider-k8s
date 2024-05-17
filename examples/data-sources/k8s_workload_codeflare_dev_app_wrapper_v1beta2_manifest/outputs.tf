output "manifests" {
  value = {
    "example" = data.k8s_workload_codeflare_dev_app_wrapper_v1beta2_manifest.example.yaml
  }
}
