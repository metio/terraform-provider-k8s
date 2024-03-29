output "manifests" {
  value = {
    "example" = data.k8s_operator_victoriametrics_com_vm_single_v1beta1_manifest.example.yaml
  }
}
