package config

import(
    "fmt"
    "errors"
    // "bytes"
    "log"
    "os"

    "gopkg.in/yaml.v2"

    "github.com/duclos-cavalcanti/go-org/cmd/org/util"
)

func readConfig(config_path string) Config {
    var conf Config

    data, err := util.ReadFile(config_path)
    if err != nil {
        log.Fatal(err)
    }

    err = yaml.Unmarshal(data, &conf)
    if err != nil {
        log.Fatal(err)
    }

    return conf
}

func writeConfig(config_path string, conf Config) {
    data, err := yaml.Marshal(&conf)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    err = util.WriteFile(config_path, data)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
}

func Setup(dir string) (Config, error) {
    // var buf bytes.Buffer
    // var logger = log.New(&buf, "SETUP: ", log.Ltime)
    var config_path = "config.yml"
    var conf Config

    if is_dir := util.IsDirectory(dir); is_dir {
        if err := os.Chdir(dir) ;err != nil {
            return conf,err
        } else {
            if util.ExistsFile(config_path) {
                conf = readConfig(config_path)
            } else {
                conf = default_config()
                writeConfig(config_path, conf)
            }

            // fmt.Print(&buf)
            return conf,nil
        }
    } else {
        return conf, errors.New(fmt.Sprintf("%s is not a valid config directory", dir))
    }
}
