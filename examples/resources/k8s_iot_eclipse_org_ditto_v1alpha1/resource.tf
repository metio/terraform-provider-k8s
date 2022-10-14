resource "k8s_iot_eclipse_org_ditto_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}

resource "k8s_iot_eclipse_org_ditto_v1alpha1" "example" {
  metadata = {
    name = "example-ditto"
  }
  spec = {
    mongo_db = {
      host = "mongodb"
    }
  }
}
