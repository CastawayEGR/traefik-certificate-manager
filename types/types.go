package types

type CertificatesMap map[string]Certificates

type Certificates struct {
    Account      Account      `json:"Account"`
    Certificates []Certificate `json:"Certificates"`
}

type Account struct {
    Email        string       `json:"Email"`
    Registration Registration `json:"Registration"`
    PrivateKey   string       `json:"PrivateKey"`
    KeyType      string       `json:"KeyType"`
}

type Registration struct {
    Body Body   `json:"body"`
    Uri  string `json:"uri"`
}

type Body struct {
    Status  string   `json:"status"`
    Contact []string `json:"contact"`
}

type Certificate struct {
    Domain      Domain `json:"domain"`
    Certificate string `json:"certificate"`
    Key         string `json:"key"`
    Store       string `json:"Store"`
}

type Domain struct {
    Main string   `json:"main"`
    Sans []string `json:"sans,omitempty"`
}

type Options struct {
    File string `opts:"help=json file path"`
    Version bool `opts:"help=version info"`
}
