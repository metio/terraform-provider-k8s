data "k8s_spark_stackable_tech_spark_application_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    main_application_file = "some-file"
    mode                  = "cluster"
    spark_image           = {}
  }
}
