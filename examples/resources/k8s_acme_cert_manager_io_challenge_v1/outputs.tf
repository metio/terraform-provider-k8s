output "resources" {
  value = {
    "big"   = k8s_acme_cert_manager_io_challenge_v1.big.yaml
    "small" = k8s_acme_cert_manager_io_challenge_v1.small.yaml
  }
}
