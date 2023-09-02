output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_bgp_filter_v1_manifest.example.yaml
  }
}
