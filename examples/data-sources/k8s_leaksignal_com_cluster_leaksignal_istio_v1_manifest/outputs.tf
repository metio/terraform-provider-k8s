output "manifests" {
  value = {
    "example" = data.k8s_leaksignal_com_cluster_leaksignal_istio_v1_manifest.example.yaml
  }
}
