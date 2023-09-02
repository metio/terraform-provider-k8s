output "manifests" {
  value = {
    "example" = data.k8s_infinispan_org_restore_v2alpha1_manifest.example.yaml
  }
}
