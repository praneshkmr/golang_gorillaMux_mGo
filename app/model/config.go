package model

type Configuration struct {
	MongoDB mongoConf `json:"mongoDB"`
}

type mongoConf struct {
	Host     string `json:"host"`
	Database string `json:"database"`
}
