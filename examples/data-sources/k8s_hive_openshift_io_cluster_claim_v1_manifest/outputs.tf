output "manifests" {
  value = {
    "example" = data.k8s_hive_openshift_io_cluster_claim_v1_manifest.example.yaml
  }
}
