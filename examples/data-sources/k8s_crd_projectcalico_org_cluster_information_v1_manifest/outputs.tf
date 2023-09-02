output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_cluster_information_v1_manifest.example.yaml
  }
}
