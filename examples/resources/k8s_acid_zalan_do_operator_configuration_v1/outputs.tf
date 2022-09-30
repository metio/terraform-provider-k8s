output "resources" {
  value = {
    "big"   = k8s_acid_zalan_do_operator_configuration_v1.big.yaml
    "small" = k8s_acid_zalan_do_operator_configuration_v1.small.yaml
  }
}
