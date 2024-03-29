output "manifests" {
  value = {
    "example" = data.k8s_workload_codeflare_dev_app_wrapper_v1beta1_manifest.example.yaml
  }
}
