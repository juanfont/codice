package codice

import "time"

type CodiceEntry struct {
	EntryID    string    `csv:"entry_id"`
	Updated    time.Time `csv:"updated"`
	FolderID   string    `csv:"folder_id"`
	Title      string    `csv:"title"`
	Summary    string    `csv:"summary"`
	StatusCode string    `csv:"status_code"`

	ContractingPartyIdentification         string `csv:"contracting_party_identification"`
	ContractingPartyName                   string `csv:"contracting_party_name"`
	ContractingPartyWebsite                string `csv:"contracting_party_website"`
	ContractingPartyType                   string `csv:"contracting_party_type"`
	ContractingPartyParentLocatedPartyName string `csv:"contracting_party_parent_name"`

	TypeCode          string `csv:"type_code"`
	SubTypeCode       string `csv:"type_sub_code"`
	CPVClassification string `csv:"cpv_classification_codes"`

	TechnicalInstructionsURL string `csv:"technical_instructions_url"`

	BudgetEstimatedOverallContractAmount string `csv:"budget_estimated_amount"`
	BudgetTotalAmount                    string `csv:"budget_total_amount"`
	BudgetTaxExclusiveAmount             string `csv:"budget_exc_tax_amount"`

	RealizedLocationCountrySubentity     string `csv:"realized_location_country_subentity"`
	RealizedLocationCountrySubentityCode string `csv:"realized_location_country_subentity_code"`
	RealizedLocationAddressCountry       string `csv:"realized_location_address_country"`
	RealizedLocationAddressCountryCode   string `csv:"realized_location_address_country_code"`
	RealizedLocationAddressCityName      string `csv:"realized_location_address_city"`
	RealizedLocationAddressPostalZone    string `csv:"realized_location_address_postal_zone"`
	RealizedLocationAddressAddressLine   string `csv:"realized_location_address_address_line"`

	PlannedPeriodDurationMeasure         string `csv:"planned_period_duration"`
	PlannedPeriodDurationMeasureUnitCode string `csv:"planned_period_duration_unit_code"`
	PlannedPeriodStartDate               string `csv:"planned_period_duration_start_date"`
	PlannedPeriodEndDate                 string `csv:"planned_period_duration_end_date"`

	ContractExtensionOptionsDescription        string `csv:"contract_extension_options_description"`
	ContractExtensionValidityPeriodDescription string `csv:"contract_extension_validity_period_description"`

	LotID                 string `csv:"lot_id"`
	LotName               string `csv:"lot_name"`
	LotTotalAmount        string `csv:"lot_total_amount"`
	LotTaxExclusiveAmount string `csv:"lot_exc_tax_amount"`
	LotCPVClassification  string `csv:"lot_cpv_classification_codes"`

	TenderResultCode         string `csv:"tender_result_code"`
	TenderDescription        string `csv:"tender_description"`
	TenderContractID         string `csv:"tender_contract_id"`
	TenderContractIssueDate  string `csv:"tender_contract_issue_date"`
	TenderWinningPartyID     string `csv:"tender_winning_party_id"`
	TenderWinningPartyName   string `csv:"tender_winning_party_name"`
	TenderWinningPartyScheme string `csv:"tender_winning_party_scheme"`
	TenderPayableAmount      string `csv:"tender_payable_amount"`
	TenderTaxExclusiveAmount string `csv:"tender_exc_tax_amount"`

	TenderStartDate              string `csv:"tender_start_date"`
	TenderAwardDate              string `csv:"tender_award_date"`
	TenderReceivedTenderQuantity string `csv:"tender_received_quantity"`
	TenderLowerLenderAmount      string `csv:"tender_lower_lender_amount"`
	TenderHigherTenderAmount     string `csv:"tender_higher_lender_amount"`

	TenderingTermsFundingProgram                  string `csv:"tendering_terms_funding_program"`
	TenderingTermsFundingProgramCode              string `csv:"tendering_terms_funding_program_code"`
	TenderingTermsLanguage                        string `csv:"tendering_terms_language"`
	TenderingTermsLRequiredCurriculaIndicator     string `csv:"tendering_terms_required_curricula_indicator"`
	TenderingTermsVariantConstraintIndicator      string `csv:"tendering_terms_variant_constraint_indicator"`
	TenderingTermsPriceRevisionFormulaDescription string `csv:"tendering_terms_price_revision_formula_description"`
	TenderingTermsSubcontractTermsRate            string `csv:"tendering_terms_subcontract_terms_rate"`
	TenderingTermsSubcontractTermsDescription     string `csv:"tendering_terms_subcontract_terms_description"`

	TenderingProcessProcedureCode         string `csv:"tendering_process_procedure_code"`
	TenderingProcessContractingSystemCode string `csv:"tendering_process_contracting_system_code"`
	TenderingProcessUrgencyCode           string `csv:"tendering_process_urgency_code"`
	TenderingProcessSubmissionMethodCode  string `csv:"tendering_process_submission_method_code"`
	TenderingProcessSubmissionEndDate     string `csv:"tendering_process_submission_end_date"`
	TenderingProcessSubmissionEndTime     string `csv:"tendering_process_submission_end_time"`

	TenderingProcessEconomicOperatorLimitationDescription string `csv:"tendering_process_economic_operator_limitation_description"`
	TenderingProcessEconomicOperatorExpectedQuantity      string `csv:"tendering_process_economic_operator_expected_quantity"`
	TenderingProcessEconomicOperatorMaximumQuantity       string `csv:"tendering_process_economic_operator_maximum_quantity"`
	TenderingProcessEconomicOperatorMinimunQuantity       string `csv:"tendering_process_economic_operator_minimum_quantity"`
}

type Modification struct {
	EntryID                                string `csv:"entry_id"`
	ID                                     string `csv:"id"`
	FolderID                               string `csv:"folder_id"`
	Note                                   string `csv:"note"`
	ContractModificationDuration           string `csv:"duration"`
	ContractModificationDurationUnit       string `csv:"duration_unit"`
	ContractModificationFinalDuration      string `csv:"final_duration"`
	ContractModificationFinalDurationUnit  string `csv:"final_duration_unit"`
	ContractModificationLegalMonetaryTotal string `csv:"monetary_total_tax_exclusive"`
	FinalLegalMonetaryTotal                string `csv:"final_legal_tax_exclusive"`
}

type FinancialGuarantee struct {
	EntryID                   string `csv:"entry_id"`
	FolderID                  string `csv:"folder_id"`
	GuaranteeTypeCode         string `csv:"guarantee_type_code"`
	AmountRate                string `csv:"amount_rate"`
	LiabilityAmount           string `csv:"liability_amount"`
	LiabilityAmountCurrencyID string `csv:"liability_amount_currency_id"`
}
