output "manifests" {
  value = {
    "example" = data.k8s_quota_codeflare_dev_quota_subtree_v1alpha1_manifest.example.yaml
  }
}
