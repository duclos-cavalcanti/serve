package config

import(
    "fmt"
    "errors"
    "log"
    "os"

    "gopkg.in/yaml.v2"

    "github.com/duclos-cavalcanti/go-org/cmd/org/util"
)

func readConfig(config_path string) {

}

func writeConfig(config_path string) {
    conf := default_config()
    data, err := yaml.Marshal(&conf)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    err = util.WriteFile(config_path, data)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
}

func Setup(config string) error {
    is_dir, err := util.IsDirectory(config)
    if err != nil {
        return err
    } else {
        if is_dir {
            err = os.Chdir(config)
            if err != nil {
                return err
            } else {
                config_path := "config.yml"
                if util.ExistsFile(config_path) {
                    fmt.Println("Config exists...")
                    readConfig(config_path)
                } else {
                    fmt.Println("Creating config...")
                    writeConfig(config_path)
                }
                return nil
            }
        } else {
            return errors.New(fmt.Sprintf("%s is not a valid config directory", config))
        }
    }
}
