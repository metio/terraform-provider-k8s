output "manifests" {
  value = {
    "example" = data.k8s_sonataflow_org_sonata_flow_v1alpha08_manifest.example.yaml
  }
}
