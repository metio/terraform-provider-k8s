output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_ip_pool_v1_manifest.example.yaml
  }
}
