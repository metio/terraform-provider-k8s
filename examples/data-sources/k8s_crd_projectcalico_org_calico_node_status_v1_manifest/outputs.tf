output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_calico_node_status_v1_manifest.example.yaml
  }
}
