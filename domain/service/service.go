package service

import "github.com/dulguundd/logError-lib/errs"

type Service struct {
	id                  string   `json:"id"`
	createdAt           int      `json:"created_at"`
	updatedAt           int      `json:"updated_at"`
	name                string   `json:"name"`
	retries             int      `json:"retries"`
	protocol            string   `json:"protocol"`
	host                string   `json:"host"`
	port                string   `json:"port"`
	path                string   `json:"path"`
	connectTimeout      int      `json:"connect_timeout"`
	writeTimeout        int      `json:"write_timeout"`
	readTimeout         int      `json:"read_timeout"`
	tags                []string `json:"tags"`
	clientCertificateId string   `json:"client_certificate_id"`
	tlsVerify           string   `json:"tls_verify"`
	tlsVerifyDepth      int      `json:"tls_verify_depth"`
	caCertificates      string   `json:"ca_certificates"`
	wsId                string   `json:"ws_id"`
}

type ServiceRepository interface {
	GetServices() ([]Service, *errs.AppError)
	//GetService(id string) (Service, *errs.AppError)
}