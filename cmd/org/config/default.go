package config

import()

type config struct {
    User string
    Stacks string
}

func default_config() config {
    return config {
        User : "foo",
        Stacks : "bar",
    }
}
