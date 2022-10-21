resource "k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3" "example" {
  metadata = {
    name = "health-for-strangers"
  }
  spec = {
    matching_precedence = 1000
    priority_level_configuration = {
      name = "exempt"
    }
    rules = [
      {
        non_resource_rules = [
          {
            non_resource_urls = ["/healthz", "/livez", "/readyz"]
            verbs             = ["*"]
          }
        ]
        subjects = [
          {
            kind = "Group"
            group = {
              name = "system:unauthenticated"
            }
          }
        ]
      }
    ]
  }
}
