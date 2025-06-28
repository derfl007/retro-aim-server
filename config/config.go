package config

//go:generate go run github.com/mk6i/retro-aim-server/cmd/config_generator unix settings.env
type Config struct {
	ApiHost      string `envconfig:"API_HOST" require:"true" val:"127.0.0.1" description:"Specifies the IP address or hostname that the management API binds to for incoming connections (127.0.0.1 restricts to same machine only)."`
	ApiPort      string `envconfig:"API_PORT" required:"true" val:"8080" description:"The port that the management API service binds to."`
	KerberosPort string `envconfig:"KERBEROS_PORT" required:"true" val:"1088" description:"The port that the Kerberos server binds to."`
	AlertPort    string `envconfig:"ALERT_PORT" required:"true" val:"5194" description:"The port that the Alert service binds to."`
	AuthPort     string `envconfig:"AUTH_PORT" required:"true" val:"5190" description:"The port that the auth service binds to."`
	BARTPort     string `envconfig:"BART_PORT" required:"true" val:"5195" description:"The port that the BART service binds to."`
	BOSPort      string `envconfig:"BOS_PORT" required:"true" val:"5191" description:"The port that the BOS service binds to."`
	ChatNavPort  string `envconfig:"CHAT_NAV_PORT" required:"true" val:"5193" description:"The port that the chat nav service binds to."`
	ChatPort     string `envconfig:"CHAT_PORT" required:"true" val:"5192" description:"The port that the chat service binds to."`
	AdminPort    string `envconfig:"ADMIN_PORT" required:"true" val:"5196" description:"The port that the admin service binds to."`
	ODirPort     string `envconfig:"ODIR_PORT" required:"true" val:"5197" description:"The port that the ODir service binds to."`
	DBPath       string `envconfig:"DB_PATH" required:"true" val:"oscar.sqlite" description:"The path to the SQLite database file. The file and DB schema are auto-created if they doesn't exist."`
	DisableAuth  bool   `envconfig:"DISABLE_AUTH" required:"true" val:"true" description:"Disable password check and auto-create new users at login time. Useful for quickly creating new accounts during development without having to register new users via the management API."`
	LogLevel     string `envconfig:"LOG_LEVEL" required:"true" val:"info" description:"Set logging granularity. Possible values: 'trace', 'debug', 'info', 'warn', 'error'."`
	OSCARHost    string `envconfig:"OSCAR_HOST" required:"true" val:"127.0.0.1" description:"The hostname that AIM clients connect to in order to reach OSCAR services (auth, BOS, BUCP, etc). Make sure the hostname is reachable by all clients. For local development, the default loopback address should work provided the server and AIM client(s) are running on the same machine. For LAN-only clients, a private IP address (e.g. 192.168..) or hostname should suffice. For clients connecting over the Internet, specify your public IP address and ensure that TCP ports 5190-5197 are open on your firewall."`
	TOCHost      string `envconfig:"TOC_HOST" require:"true" val:"0.0.0.0" description:"Specifies the IP address or hostname that the TOC service binds to for incoming connections (0.0.0.0 listens on all interfaces)."`
	TOCPort      string `envconfig:"TOC_PORT" required:"true" val:"9898" description:"The port that the TOC service binds to."`
}

type Build struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}
