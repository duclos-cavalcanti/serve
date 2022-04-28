package config

import(
    "fmt"
    "errors"
    "log"
    "os"

    "gopkg.in/yaml.v2"
)

func readConfig() {
    config_path := "config.yml"
    if existsFile(config_path) {
        fmt.Println("Config exists...")
        // do things
    } else {
        fmt.Println("Creating config...")
        info := struct {
            User string
            Stacks string
        } {
            User: "foo",
            Stacks: "bar",
        }
        data, err := yaml.Marshal(&info)
        if err != nil {
            log.Fatalf("error: %v", err)
        }

        err = writeFile(config_path, data)
        if err != nil {
            log.Fatalf("error: %v", err)
        }
    }

}

func Setup(config string) error {
    is_dir, err := isDirectory(config)
    if err != nil {
        return err
    } else {
        if is_dir {
            err = os.Chdir(config)
            if err != nil {
                return err
            } else {
                readConfig()
                return nil
            }
        } else {
            return errors.New(fmt.Sprintf("%s is not a valid config directory", config))
        }
    }
}
