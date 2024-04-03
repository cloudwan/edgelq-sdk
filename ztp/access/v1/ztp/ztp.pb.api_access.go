// Code generated by protoc-gen-goten-access
// Service: Ztp
// DO NOT EDIT!!!

package ztp_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	edgelq_instance_access "github.com/cloudwan/edgelq-sdk/ztp/access/v1/edgelq_instance"
	hardware_access "github.com/cloudwan/edgelq-sdk/ztp/access/v1/hardware"
	project_access "github.com/cloudwan/edgelq-sdk/ztp/access/v1/project"
	tpm_attestation_cert_access "github.com/cloudwan/edgelq-sdk/ztp/access/v1/tpm_attestation_cert"
	ztp_client "github.com/cloudwan/edgelq-sdk/ztp/client/v1/ztp"
	edgelq_instance "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/edgelq_instance"
	hardware "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/hardware"
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
	tpm_attestation_cert "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/tpm_attestation_cert"
)

type ZtpApiAccess interface {
	gotenresource.Access

	edgelq_instance.EdgelqInstanceAccess
	hardware.HardwareAccess
	project.ProjectAccess
	tpm_attestation_cert.TpmAttestationCertAccess
}

type apiZtpAccess struct {
	gotenresource.Access

	edgelq_instance.EdgelqInstanceAccess
	hardware.HardwareAccess
	project.ProjectAccess
	tpm_attestation_cert.TpmAttestationCertAccess
}

func NewApiAccess(client ztp_client.ZtpClient) ZtpApiAccess {

	edgelqInstanceAccess := edgelq_instance_access.NewApiEdgelqInstanceAccess(client)
	hardwareAccess := hardware_access.NewApiHardwareAccess(client)
	projectAccess := project_access.NewApiProjectAccess(client)
	tpmAttestationCertAccess := tpm_attestation_cert_access.NewApiTpmAttestationCertAccess(client)

	return &apiZtpAccess{
		Access: gotenresource.NewCompositeAccess(

			edgelq_instance.AsAnyCastAccess(edgelqInstanceAccess),
			hardware.AsAnyCastAccess(hardwareAccess),
			project.AsAnyCastAccess(projectAccess),
			tpm_attestation_cert.AsAnyCastAccess(tpmAttestationCertAccess),
		),

		EdgelqInstanceAccess:     edgelqInstanceAccess,
		HardwareAccess:           hardwareAccess,
		ProjectAccess:            projectAccess,
		TpmAttestationCertAccess: tpmAttestationCertAccess,
	}
}