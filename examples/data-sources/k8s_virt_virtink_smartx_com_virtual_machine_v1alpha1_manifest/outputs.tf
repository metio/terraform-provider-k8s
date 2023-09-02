output "manifests" {
  value = {
    "example" = data.k8s_virt_virtink_smartx_com_virtual_machine_v1alpha1_manifest.example.yaml
  }
}
