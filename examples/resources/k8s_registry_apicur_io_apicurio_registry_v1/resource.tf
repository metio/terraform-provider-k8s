resource "k8s_registry_apicur_io_apicurio_registry_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_registry_apicur_io_apicurio_registry_v1" "example" {
  metadata = {
    name = "example-apicurioregistry-kafkasql"
  }
  spec = {
    configuration = {
      kafkasql = {
        bootstrap_servers = "<service name>.<namespace>.svc:9092"
      }
      persistence = "kafkasql"
    }
  }
}
