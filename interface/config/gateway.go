package config

import (
	"crypto/x509"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path"
	"time"
)

type PeerConfig struct {
	MSPID               string
	CertPath            string
	KeyPath             string
	TLSCertPath         string
	PeerEndpoint        string
	GatewayPeer         string
	Network             *client.Network
	TransactionContract *client.Contract
	InspectionContract  *client.Contract
}

const (
	CHANNEL_NAME              = "vehicles"
	PROJECT_PATH              = "/home/jeho/blockchain-repo/cloud_chain/organizations/peerOrganizations"
	INSPECTION_CONTRACT_NAME  = "inspection"
	TRANSACTION_CONTRACT_NAME = "transaction"
)

var (
	SellerConfig = PeerConfig{
		MSPID:        "sellerMSP",
		CertPath:     PROJECT_PATH + "/seller.pnu.cse" + "/users/User1@seller.pnu.cse/msp/signcerts/cert.pem",
		KeyPath:      PROJECT_PATH + "/seller.pnu.cse" + "/users/User1@seller.pnu.cse/msp/keystore/",
		TLSCertPath:  PROJECT_PATH + "/seller.pnu.cse" + "/peers/peer0.seller.pnu.cse/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.seller.pnu.cse",
	}

	BuyerConfig = PeerConfig{
		MSPID:        "buyerMSP",
		CertPath:     PROJECT_PATH + "/buyer.pnu.cse" + "/users/User1@buyer.pnu.cse/msp/signcerts/cert.pem",
		KeyPath:      PROJECT_PATH + "/buyer.pnu.cse" + "/users/User1@buyer.pnu.cse/msp/keystore/",
		TLSCertPath:  PROJECT_PATH + "/buyer.pnu.cse" + "/peers/peer0.buyer.pnu.cse/tls/ca.crt",
		PeerEndpoint: "localhost:9051", // Buyer 피어의 gRPC 엔드포인트
		GatewayPeer:  "peer0.buyer.pnu.cse",
	}

	InspectorConfig = PeerConfig{
		MSPID:        "inspectorMSP",
		CertPath:     PROJECT_PATH + "/inspector.pnu.cse" + "/users/User1@inspector.pnu.cse/msp/signcerts/cert.pem",
		KeyPath:      PROJECT_PATH + "/inspector.pnu.cse" + "/users/User1@inspector.pnu.cse/msp/keystore/",
		TLSCertPath:  PROJECT_PATH + "/inspector.pnu.cse" + "/peers/peer0.inspector.pnu.cse/tls/ca.crt",
		PeerEndpoint: "localhost:11051", // Inspector 피어의 gRPC 엔드포인트
		GatewayPeer:  "peer0.inspector.pnu.cse",
	}
)

// NewPeerConfig 생성자 함수
func NewPeerConfig(MSPID, CertPath, KeyPath, TLSCertPath, PeerEndpoint, GatewayPeer string) *PeerConfig {
	return &PeerConfig{
		MSPID:        MSPID,
		CertPath:     CertPath,
		KeyPath:      KeyPath,
		TLSCertPath:  TLSCertPath,
		PeerEndpoint: PeerEndpoint,
		GatewayPeer:  GatewayPeer,
	}
}

// Connect 함수
func (pc *PeerConfig) Connect() {
	clientConnection := pc.newGrpcConnection()

	id := pc.newIdentity()
	sign := pc.newSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	pc.Network = gateway.GetNetwork(CHANNEL_NAME)
	pc.TransactionContract = pc.Network.GetContract(TRANSACTION_CONTRACT_NAME)
	pc.InspectionContract = pc.Network.GetContract(INSPECTION_CONTRACT_NAME)
}

func (pc *PeerConfig) newGrpcConnection() *grpc.ClientConn {
	certificate, err := pc.loadCertificate(pc.TLSCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, pc.GatewayPeer)

	connection, err := grpc.Dial(pc.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func (pc *PeerConfig) newIdentity() *identity.X509Identity {
	certificate, err := pc.loadCertificate(pc.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(pc.MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func (pc *PeerConfig) newSign() identity.Sign {
	files, err := os.ReadDir(pc.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(pc.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func (pc *PeerConfig) loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
