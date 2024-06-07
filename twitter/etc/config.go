package etc

// why we use config but not use env
// because the env need under the floder

var Conf = new(AppConf)

type AppConf struct {
}
