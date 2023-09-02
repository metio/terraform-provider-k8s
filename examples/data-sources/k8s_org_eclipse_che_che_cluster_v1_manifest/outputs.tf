output "manifests" {
  value = {
    "example" = data.k8s_org_eclipse_che_che_cluster_v1_manifest.example.yaml
  }
}
