output "manifests" {
  value = {
    "example" = data.k8s_workload_codeflare_dev_scheduling_spec_v1beta1_manifest.example.yaml
  }
}
