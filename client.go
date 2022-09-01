package libkitsune

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/lusory/libkitsune/proto/kitsune/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// clientCache is a global cache for NewOrCachedKitsuneClient.
var clientCache = make(map[string]*KitsuneClient)

// KitsuneClient is an implementation of KitsuneClientInterface and a collection of a grpc.ClientConnInterface
// and the services provided over that connection.
type KitsuneClient struct {
	Cc            *grpc.ClientConn
	ImageRegistry v1.ImageRegistryServiceClient
	VmRegistry    v1.VirtualMachineRegistryServiceClient
}

// KitsuneClientInterface is the abstract version of KitsuneClient.
type KitsuneClientInterface interface {
	Close() error
}

// Close closes the grpc.ClientConn of the KitsuneClient and removes it from cache if present.
func (c *KitsuneClient) Close() error {
	delete(clientCache, c.Cc.Target())
	return c.Cc.Close()
}

// NewKitsuneClient dials the supplied target and creates a new KitsuneClient with the resulting grpc.ClientConn.
func NewKitsuneClient(target string, ssl bool) (*KitsuneClient, error) {
	creds := insecure.NewCredentials()
	if ssl {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return &KitsuneClient{}, err
		}

		creds = credentials.NewTLS(&tls.Config{RootCAs: certPool})
	}

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return &KitsuneClient{}, err
	}

	return &KitsuneClient{
		conn,
		v1.NewImageRegistryServiceClient(conn),
		v1.NewVirtualMachineRegistryServiceClient(conn),
	}, nil
}

// NewOrCachedKitsuneClient attempts to fetch a KitsuneClient corresponding to the supplied target from cache,
// creating a new KitsuneClient and inserting it into the cache if not successful.
func NewOrCachedKitsuneClient(target string, ssl bool) (*KitsuneClient, error) {
	if client, ok := clientCache[target]; ok {
		return &client, nil
	}

	client, err := NewKitsuneClient(target, ssl)
	if err != nil {
		return &KitsuneClient{}, err
	}

	clientCache[target] = client
	return client, nil
}
