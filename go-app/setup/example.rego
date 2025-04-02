package example
import rego.v1

default allow = false

allow if{
    input.method == "GET"
    input.path == "/api/user"
}

allow if{
    input.method == "GET"
    input.path == "/api/policy"
}
