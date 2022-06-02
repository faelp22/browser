package certload

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
)

type CertLoadInterface interface {
	ConfigCertLoad(pathCertFile, pathKeyFile, pathCAFile string, SSLVerify bool)
	GetTLSConfig() *tls.Config
}

type certLoad struct {
	PathCertFile string
	PathKeyFile  string
	PathCAFile   string
	SSLVerify    bool
	tlsConfig    *tls.Config
}

func (c *certLoad) addCertFile(pathCertFile string) {
	_, err := os.Stat(pathCertFile)
	if err == nil {
		c.PathCertFile = pathCertFile
	} else {
		log.Fatal("Error to load [CertFile], check if path to file exist")
	}
}

func (c *certLoad) addKeyFile(pathKeyFile string) {
	_, err := os.Stat(pathKeyFile)
	if err == nil {
		c.PathKeyFile = pathKeyFile
	} else {
		log.Fatal("Error to load [KeyFile], check if path to file exist")
	}
}

func (c *certLoad) addCAFile(pathCAFile string) {
	_, err := os.Stat(pathCAFile)
	if err == nil {
		c.PathCAFile = pathCAFile
	} else {
		log.Fatal("Error to load [CAFile], check if path to file exist")
	}
}

func (c *certLoad) ConfigCertLoad(pathCertFile, pathKeyFile, pathCAFile string, SSLVerify bool) {

	c.addCertFile(pathCertFile)
	c.addKeyFile(pathKeyFile)
	c.addCAFile(pathCAFile)

	cert, err := tls.LoadX509KeyPair(c.PathCertFile, c.PathKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	caCert, err := ioutil.ReadFile(c.PathCAFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	c.tlsConfig = &tls.Config{
		InsecureSkipVerify: !SSLVerify,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
	}
	// c.tlsConfig.BuildNameToCertificate()
}

func (c *certLoad) GetTLSConfig() *tls.Config {
	return c.tlsConfig
}

func NewCertLoad(pathCertFile, pathKeyFile, pathCAFile string, SSLVerify bool) CertLoadInterface {
	cl := certLoad{}
	cl.ConfigCertLoad(pathCertFile, pathKeyFile, pathCAFile, SSLVerify)
	return &cl
}
