data "k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  webhooks = [
    {
      name = "test.terraform.io"

      admission_review_versions = ["v1", "v1beta1"]

      client_config = {
        service = {
          namespace = "example-namespace"
          name      = "example-service"
        }
      }

      rules = [
        {
          api_groups   = ["apps"]
          api_versions = ["v1"]
          operations   = ["CREATE"]
          resources    = ["deployments"]
          scope        = "Namespaced"
        },
      ]

      side_effects = "None"
    },
  ]
}
