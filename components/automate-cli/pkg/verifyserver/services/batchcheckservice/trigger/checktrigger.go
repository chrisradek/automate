package trigger

import "github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"

type CheckTrigger struct {
	HardwareResourceCountCheck    ICheck
	SshUserAccessCheck            ICheck
	CertificateCheck              ICheck
	ExternalOpensearchCheck       ICheck
	ExternalPostgresCheck         ICheck
	FirewallCheck                 ICheck
	FqdnCheck                     ICheck
	NfsBackupConfigCheck          ICheck
	OpensearchS3BucketAccessCheck ICheck
	S3BackupConfigCheck           ICheck
	SoftwareVersionCheck          ICheck
	SystemResourceCheck           ICheck
	SystemUserCheck               ICheck
}
type ICheck interface {
	Run(config models.Config) []models.CheckTriggerResponse
}

func NewCheckTrigger(hrc, sshC, cert, eop, epc, fc, fqdn, nfs, os3, s3b, svc, src, suc ICheck) CheckTrigger {
	return CheckTrigger{
		HardwareResourceCountCheck:    hrc,
		SshUserAccessCheck:            sshC,
		CertificateCheck:              cert,
		ExternalOpensearchCheck:       eop,
		ExternalPostgresCheck:         epc,
		FirewallCheck:                 fc,
		FqdnCheck:                     fqdn,
		NfsBackupConfigCheck:          nfs,
		OpensearchS3BucketAccessCheck: os3,
		S3BackupConfigCheck:           s3b,
		SoftwareVersionCheck:          svc,
		SystemResourceCheck:           src,
		SystemUserCheck:               suc,
	}
}
