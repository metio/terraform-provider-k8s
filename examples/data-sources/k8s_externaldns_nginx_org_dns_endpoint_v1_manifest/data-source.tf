data "k8s_externaldns_nginx_org_dns_endpoint_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
