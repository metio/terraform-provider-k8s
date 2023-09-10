data "k8s_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    allow = {
      "@self" = {
        tcp_socket = ["listen"]
      }
      http_cache_port_t = {
        tcp_socket = ["name_bind"]
      }
      node_t = {
        tcp_socket = ["node_bind"]
      }
    }
    inherit = [
      {
        kind = "System"
        name = "container"
      }
    ]
  }
}
