output "manifests" {
  value = {
    "example" = data.k8s_ingress_operator_openshift_io_dns_record_v1_manifest.example.yaml
  }
}
