output "manifests" {
  value = {
    "example" = data.k8s_infinispan_org_infinispan_v1_manifest.example.yaml
  }
}
