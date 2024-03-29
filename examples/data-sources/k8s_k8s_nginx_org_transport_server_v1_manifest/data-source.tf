data "k8s_k8s_nginx_org_transport_server_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
