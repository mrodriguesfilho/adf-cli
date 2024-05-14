package internal

type ServiceData struct {
	Name        string
	Version     string
	DownloadUrl string
}

var StaticServiceDataArr = []ServiceData{
	{"adfweb", "0.0.1", "http://localhost:5000/downloadadf"},
	{"jvm", "12.0.0", "http://localhost:5000/downloadjvm"},
}
