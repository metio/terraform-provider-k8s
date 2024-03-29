output "manifests" {
  value = {
    "example" = data.k8s_multicluster_crd_antrea_io_cluster_claim_v1alpha2_manifest.example.yaml
  }
}
