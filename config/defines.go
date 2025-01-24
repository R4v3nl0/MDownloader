package config

type Prefixes struct {
	Cover string `yaml:"cover"`
	Video string `yaml:"video"`
}

type Suffixes struct {
	Playlist string `yaml:"playlist"`
}

type Requests struct {
	Timeout int               `yaml:"timeout"`
	Retry   int               `yaml:"retry"`
	Delay   int               `yaml:"delay"`
	Proxy   string            `yaml:"proxy"`
	Headers map[string]string `yaml:"headers"`
}

type RegHref struct {
	MovieCollection string `yaml:"movieCollection"`
	PublicPlaylist  string `yaml:"publicPlaylist"`
	NextPage        string `yaml:"nextPage"`
}

type RegData struct {
	Uuid       string `yaml:"uuid"`
	Title      string `yaml:"title"`
	Resolution string `yaml:"resolution"`
	VideoSlice string `yaml:"videoSlice"`
}

type Regexes struct {
	Href RegHref `yaml:"href"`
	Data RegData `yaml:"data"`
}

type Config struct {
	Prefixes Prefixes `yaml:"prefix"`
	Suffixes Suffixes `yaml:"suffix"`
	Requests Requests `yaml:"requests"`
	Regexes  Regexes  `yaml:"regex"`

	Ffmpeg  string `yaml:"ffmpeg"`
	Quality int    `yaml:"quality"`

	SavePath string `yaml:"savePath"`
}
