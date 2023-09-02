output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_host_endpoint_v1_manifest.example.yaml
  }
}
