package config

import(
    "fmt"
)

type Config struct {
    User string
    Stacks string
}


func (c Config) Print() {
    fmt.Println("--------------")
    fmt.Println("Config")
    fmt.Println("User: ", c.User)
    fmt.Println("Stacks: ", c.Stacks)
    fmt.Println("--------------")
    fmt.Printf("\n")
}

func default_config() Config {
    return Config {
        User : "foo",
        Stacks : "bar",
    }
}
