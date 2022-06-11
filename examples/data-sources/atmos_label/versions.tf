terraform {
  required_providers {
    atmos = {
      source = "cloudposse/atmos"
      # For local development,
      # install the provider on local computer by running `make install` from the root of the repo,
      # and uncomment the version below
      version = "9999.99.99"
    }
  }
}
