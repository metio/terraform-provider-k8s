resource "k8s_binding_operators_coreos_com_service_binding_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    application = {
      group   = "apps"
      version = "v1"
    }
    services = []
  }
}

resource "k8s_binding_operators_coreos_com_service_binding_v1alpha1" "example" {
  metadata = {
    name = "example-servicebinding"
  }
  spec = {
    application = {
      group    = "apps"
      name     = "nodejs-rest-http-crud"
      resource = "deployments"
      version  = "v1"
    }
    services = [
      {
        group   = "postgresql.example.dev"
        kind    = "Database"
        name    = "pg-instance"
        version = "v1alpha1"
      }
    ]
  }
}
