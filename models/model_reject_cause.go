// SPDX-License-Identifier: Apache-2.0

/*
 * NSSF NS Selection
 *
 * Models manually added to support Rejected NSSAI (based on TS 29.531/29.571)
 *
 */
package models

type RejectCause string

const (
	RejectCause_S_NSSAI_NOT_AVAILABLE_IN_PLMN                             RejectCause = "S_NSSAI_NOT_AVAILABLE_IN_PLMN"
	RejectCause_S_NSSAI_NOT_AVAILABLE_IN_CURRENT_PLMN_OR_SNPN             RejectCause = "S_NSSAI_NOT_AVAILABLE_IN_CURRENT_PLMN_OR_SNPN"
	RejectCause_S_NSSAI_NOT_AVAILABLE_IN_TA                               RejectCause = "S_NSSAI_NOT_AVAILABLE_IN_TA"
	RejectCause_PLMN_NOT_ALLOWED                                          RejectCause = "PLMN_NOT_ALLOWED"
	RejectCause_REGISTRATION_AREA_MISMATCH                                RejectCause = "REGISTRATION_AREA_MISMATCH"
	RejectCause_S_NSSAI_NOT_AVAILABLE_DUE_TO_FAILED_OR_REVOKED_NSAA       RejectCause = "S_NSSAI_NOT_AVAILABLE_DUE_TO_FAILED_OR_REVOKED_NSAA"
	RejectCause_S_NSSAI_NOT_AVAILABLE_DUE_TO_MAXIMUM_NUMBER_OF_UE_REACHED RejectCause = "S_NSSAI_NOT_AVAILABLE_DUE_TO_MAXIMUM_NUMBER_OF_UE_REACHED"
)
