output "manifests" {
  value = {
    "example" = data.k8s_operator_victoriametrics_com_vm_scrape_config_v1beta1_manifest.example.yaml
  }
}
