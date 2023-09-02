output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest.example.yaml
  }
}
