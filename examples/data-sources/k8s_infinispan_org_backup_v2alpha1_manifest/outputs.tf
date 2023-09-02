output "manifests" {
  value = {
    "example" = data.k8s_infinispan_org_backup_v2alpha1_manifest.example.yaml
  }
}
