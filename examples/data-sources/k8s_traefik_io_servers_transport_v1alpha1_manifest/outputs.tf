output "manifests" {
  value = {
    "example" = data.k8s_traefik_io_servers_transport_v1alpha1_manifest.example.yaml
  }
}
