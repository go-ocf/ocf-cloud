package service

import (
	"context"

	"github.com/plgd-dev/cloud/certificate-authority/pb"
	"github.com/plgd-dev/cloud/pkg/log"
	"github.com/plgd-dev/kit/security/signer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *RequestHandler) SignCertificate(ctx context.Context, req *pb.SignCertificateRequest) (*pb.SignCertificateResponse, error) {

	notBefore := r.ValidFrom()
	notAfter := notBefore.Add(r.ValidFor)
	signer := signer.NewBasicCertificateSigner(r.Certificate, r.PrivateKey, notBefore, notAfter)
	cert, err := signer.Sign(ctx, req.CertificateSigningRequest)
	if err != nil {
		return nil, log.LogAndReturnError(status.Errorf(codes.InvalidArgument, "cannot sign certificate: %v", err))
	}
	log.Debugf("RequestHandler.SignCertificate csr=%v crt=%v", string(req.CertificateSigningRequest), string(cert))

	return &pb.SignCertificateResponse{
		Certificate: cert,
	}, nil
}
