output "manifests" {
  value = {
    "example" = data.k8s_awx_ansible_com_awx_backup_v1beta1_manifest.example.yaml
  }
}
