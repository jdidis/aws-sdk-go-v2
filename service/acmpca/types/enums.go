// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionType string

// Enum values for ActionType
const (
	ActionTypeIssuecertificate ActionType = "IssueCertificate"
	ActionTypeGetcertificate   ActionType = "GetCertificate"
	ActionTypeListpermissions  ActionType = "ListPermissions"
)

type AuditReportResponseFormat string

// Enum values for AuditReportResponseFormat
const (
	AuditReportResponseFormatJson AuditReportResponseFormat = "JSON"
	AuditReportResponseFormatCsv  AuditReportResponseFormat = "CSV"
)

type AuditReportStatus string

// Enum values for AuditReportStatus
const (
	AuditReportStatusCreating AuditReportStatus = "CREATING"
	AuditReportStatusSuccess  AuditReportStatus = "SUCCESS"
	AuditReportStatusFailed   AuditReportStatus = "FAILED"
)

type CertificateAuthorityStatus string

// Enum values for CertificateAuthorityStatus
const (
	CertificateAuthorityStatusCreating            CertificateAuthorityStatus = "CREATING"
	CertificateAuthorityStatusPending_certificate CertificateAuthorityStatus = "PENDING_CERTIFICATE"
	CertificateAuthorityStatusActive              CertificateAuthorityStatus = "ACTIVE"
	CertificateAuthorityStatusDeleted             CertificateAuthorityStatus = "DELETED"
	CertificateAuthorityStatusDisabled            CertificateAuthorityStatus = "DISABLED"
	CertificateAuthorityStatusExpired             CertificateAuthorityStatus = "EXPIRED"
	CertificateAuthorityStatusFailed              CertificateAuthorityStatus = "FAILED"
)

type CertificateAuthorityType string

// Enum values for CertificateAuthorityType
const (
	CertificateAuthorityTypeRoot        CertificateAuthorityType = "ROOT"
	CertificateAuthorityTypeSubordinate CertificateAuthorityType = "SUBORDINATE"
)

type FailureReason string

// Enum values for FailureReason
const (
	FailureReasonRequest_timed_out     FailureReason = "REQUEST_TIMED_OUT"
	FailureReasonUnsupported_algorithm FailureReason = "UNSUPPORTED_ALGORITHM"
	FailureReasonOther                 FailureReason = "OTHER"
)

type KeyAlgorithm string

// Enum values for KeyAlgorithm
const (
	KeyAlgorithmRsa_2048      KeyAlgorithm = "RSA_2048"
	KeyAlgorithmRsa_4096      KeyAlgorithm = "RSA_4096"
	KeyAlgorithmEc_prime256v1 KeyAlgorithm = "EC_prime256v1"
	KeyAlgorithmEc_secp384r1  KeyAlgorithm = "EC_secp384r1"
)

type RevocationReason string

// Enum values for RevocationReason
const (
	RevocationReasonUnspecified                      RevocationReason = "UNSPECIFIED"
	RevocationReasonKey_compromise                   RevocationReason = "KEY_COMPROMISE"
	RevocationReasonCertificate_authority_compromise RevocationReason = "CERTIFICATE_AUTHORITY_COMPROMISE"
	RevocationReasonAffiliation_changed              RevocationReason = "AFFILIATION_CHANGED"
	RevocationReasonSuperseded                       RevocationReason = "SUPERSEDED"
	RevocationReasonCessation_of_operation           RevocationReason = "CESSATION_OF_OPERATION"
	RevocationReasonPrivilege_withdrawn              RevocationReason = "PRIVILEGE_WITHDRAWN"
	RevocationReasonA_a_compromise                   RevocationReason = "A_A_COMPROMISE"
)

type SigningAlgorithm string

// Enum values for SigningAlgorithm
const (
	SigningAlgorithmSha256withecdsa SigningAlgorithm = "SHA256WITHECDSA"
	SigningAlgorithmSha384withecdsa SigningAlgorithm = "SHA384WITHECDSA"
	SigningAlgorithmSha512withecdsa SigningAlgorithm = "SHA512WITHECDSA"
	SigningAlgorithmSha256withrsa   SigningAlgorithm = "SHA256WITHRSA"
	SigningAlgorithmSha384withrsa   SigningAlgorithm = "SHA384WITHRSA"
	SigningAlgorithmSha512withrsa   SigningAlgorithm = "SHA512WITHRSA"
)

type ValidityPeriodType string

// Enum values for ValidityPeriodType
const (
	ValidityPeriodTypeEnd_date ValidityPeriodType = "END_DATE"
	ValidityPeriodTypeAbsolute ValidityPeriodType = "ABSOLUTE"
	ValidityPeriodTypeDays     ValidityPeriodType = "DAYS"
	ValidityPeriodTypeMonths   ValidityPeriodType = "MONTHS"
	ValidityPeriodTypeYears    ValidityPeriodType = "YEARS"
)