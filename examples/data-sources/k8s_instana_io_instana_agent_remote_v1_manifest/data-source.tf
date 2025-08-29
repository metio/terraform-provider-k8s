data "k8s_instana_io_instana_agent_remote_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
