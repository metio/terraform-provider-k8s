output "manifests" {
  value = {
    "example" = data.k8s_slinky_slurm_net_cluster_v1alpha1_manifest.example.yaml
  }
}
