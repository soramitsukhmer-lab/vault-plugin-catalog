variable "VERSION" {
    default = ""
}

target "default" {
    tags = [ "localhost:5000/soramitsukhmer-lab/vault-plugin-catalog:dev" ]
}

target "release" {
    args = {
      "VERSION" = "${VERSION}"
    }
    platforms = [
        "linux/amd64",
        "linux/arm64",
    ]
    tags = [ "harbor.sorakh.io/soramitsukhmer-lab/vault-plugin-catalog:${VERSION}" ]
}

target "binaries" {
    output = [ "./binaries" ]
    platforms = [ "local" ]
    target = "binaries"
}
