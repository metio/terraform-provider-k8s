data "k8s_apisix_apache_org_apisix_route_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
