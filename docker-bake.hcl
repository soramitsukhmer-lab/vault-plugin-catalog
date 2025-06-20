target "default" {}

target "binaries" {
    output = [ "./out" ]
    platforms = [ "local" ]
    target = "binaries"
}
