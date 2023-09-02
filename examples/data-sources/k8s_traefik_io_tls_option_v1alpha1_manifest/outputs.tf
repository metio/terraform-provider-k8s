output "manifests" {
  value = {
    "example" = data.k8s_traefik_io_tls_option_v1alpha1_manifest.example.yaml
  }
}
