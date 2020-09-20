package codice

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	XMLName     xml.Name `xml:"feed"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	CbcPlaceExt string   `xml:"cbc-place-ext,attr"`
	CacPlaceExt string   `xml:"cac-place-ext,attr"`
	Cbc         string   `xml:"cbc,attr"`
	Cac         string   `xml:"cac,attr"`
	Ns1         string   `xml:"ns1,attr"`
	Author      struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name"`
		URI   string `xml:"uri"`
		Email string `xml:"email"`
	} `xml:"author"`
	ID   string `xml:"id"`
	Link []struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
	} `xml:"link"`
	Title   string  `xml:"title"`
	Updated string  `xml:"updated"`
	Entry   []Entry `xml:"entry"`
}

type Entry struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id"`
	Link struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Summary struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"summary"`
	Title                string    `xml:"title"`
	Updated              time.Time `xml:"updated"`
	ContractFolderStatus struct {
		Text                     string `xml:",chardata"`
		ContractFolderID         string `xml:"ContractFolderID"`
		ContractFolderStatusCode struct {
			Text       string `xml:",chardata"`
			LanguageID string `xml:"languageID,attr"`
			ListURI    string `xml:"listURI,attr"`
		} `xml:"ContractFolderStatusCode"`
		LocatedContractingParty struct {
			Text                     string `xml:",chardata"`
			ContractingPartyTypeCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"ContractingPartyTypeCode"`
			Party struct {
				Text                string `xml:",chardata"`
				WebsiteURI          string `xml:"WebsiteURI"`
				PartyIdentification struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text       string `xml:",chardata"`
						SchemeName string `xml:"schemeName,attr"`
					} `xml:"ID"`
				} `xml:"PartyIdentification"`
				PartyName struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name"`
				} `xml:"PartyName"`
				PostalAddress struct {
					Text        string `xml:",chardata"`
					CityName    string `xml:"CityName"`
					PostalZone  string `xml:"PostalZone"`
					AddressLine struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode struct {
							Text    string `xml:",chardata"`
							ListURI string `xml:"listURI,attr"`
						} `xml:"IdentificationCode"`
						Name string `xml:"Name"`
					} `xml:"Country"`
				} `xml:"PostalAddress"`
				Contact struct {
					Text           string `xml:",chardata"`
					Name           string `xml:"Name"`
					Telephone      string `xml:"Telephone"`
					Telefax        string `xml:"Telefax"`
					ElectronicMail string `xml:"ElectronicMail"`
				} `xml:"Contact"`
			} `xml:"Party"`
			ParentLocatedParty struct {
				Text      string `xml:",chardata"`
				PartyName struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name"`
				} `xml:"PartyName"`
				ParentLocatedParty struct {
					Text      string `xml:",chardata"`
					PartyName struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name"`
					} `xml:"PartyName"`
					ParentLocatedParty struct {
						Text      string `xml:",chardata"`
						PartyName struct {
							Text string `xml:",chardata"`
							Name string `xml:"Name"`
						} `xml:"PartyName"`
						ParentLocatedParty struct {
							Text      string `xml:",chardata"`
							PartyName struct {
								Text string `xml:",chardata"`
								Name string `xml:"Name"`
							} `xml:"PartyName"`
							ParentLocatedParty struct {
								Text      string `xml:",chardata"`
								PartyName struct {
									Text string `xml:",chardata"`
									Name string `xml:"Name"`
								} `xml:"PartyName"`
								ParentLocatedParty struct {
									Text      string `xml:",chardata"`
									PartyName struct {
										Text string `xml:",chardata"`
										Name string `xml:"Name"`
									} `xml:"PartyName"`
								} `xml:"ParentLocatedParty"`
							} `xml:"ParentLocatedParty"`
						} `xml:"ParentLocatedParty"`
					} `xml:"ParentLocatedParty"`
				} `xml:"ParentLocatedParty"`
			} `xml:"ParentLocatedParty"`
		} `xml:"LocatedContractingParty"`
		ProcurementProject struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"Name"`
			TypeCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"TypeCode"`
			SubTypeCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"SubTypeCode"`
			BudgetAmount struct {
				Text                           string `xml:",chardata"`
				EstimatedOverallContractAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"EstimatedOverallContractAmount"`
				TotalAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TotalAmount"`
				TaxExclusiveAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxExclusiveAmount"`
			} `xml:"BudgetAmount"`
			RequiredCommodityClassification []struct {
				Text                   string `xml:",chardata"`
				ItemClassificationCode struct {
					Text    string `xml:",chardata"`
					ListURI string `xml:"listURI,attr"`
				} `xml:"ItemClassificationCode"`
			} `xml:"RequiredCommodityClassification"`
			RealizedLocation struct {
				Text                 string `xml:",chardata"`
				CountrySubentity     string `xml:"CountrySubentity"`
				CountrySubentityCode struct {
					Text    string `xml:",chardata"`
					ListURI string `xml:"listURI,attr"`
				} `xml:"CountrySubentityCode"`
				Address struct {
					Text    string `xml:",chardata"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode struct {
							Text    string `xml:",chardata"`
							ListURI string `xml:"listURI,attr"`
						} `xml:"IdentificationCode"`
						Name string `xml:"Name"`
					} `xml:"Country"`
					CityName    string `xml:"CityName"`
					PostalZone  string `xml:"PostalZone"`
					AddressLine struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
				} `xml:"Address"`
			} `xml:"RealizedLocation"`
			PlannedPeriod struct {
				Text            string `xml:",chardata"`
				DurationMeasure struct {
					Text     string `xml:",chardata"`
					UnitCode string `xml:"unitCode,attr"`
				} `xml:"DurationMeasure"`
				StartDate string `xml:"StartDate"`
				EndDate   string `xml:"EndDate"`
			} `xml:"PlannedPeriod"`
			ContractExtension struct {
				Text                 string `xml:",chardata"`
				OptionsDescription   string `xml:"OptionsDescription"`
				OptionValidityPeriod struct {
					Text        string `xml:",chardata"`
					Description string `xml:"Description"`
				} `xml:"OptionValidityPeriod"`
			} `xml:"ContractExtension"`
		} `xml:"ProcurementProject"`
		TenderResult   []TenderResult `xml:"TenderResult"`
		TenderingTerms struct {
			Text                       string `xml:",chardata"`
			VariantConstraintIndicator string `xml:"VariantConstraintIndicator"`
			Language                   struct {
				Text string `xml:",chardata"`
				ID   string `xml:"ID"`
			} `xml:"Language"`
			RequiredFinancialGuarantee []struct {
				Text              string `xml:",chardata"`
				GuaranteeTypeCode struct {
					Text    string `xml:",chardata"`
					ListURI string `xml:"listURI,attr"`
				} `xml:"GuaranteeTypeCode"`
				AmountRate      string `xml:"AmountRate"`
				LiabilityAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"LiabilityAmount"`
			} `xml:"RequiredFinancialGuarantee"`
			TendererQualificationRequest struct {
				Text                        string `xml:",chardata"`
				TechnicalEvaluationCriteria []struct {
					Text                       string `xml:",chardata"`
					EvaluationCriteriaTypeCode struct {
						Text    string `xml:",chardata"`
						ListURI string `xml:"listURI,attr"`
					} `xml:"EvaluationCriteriaTypeCode"`
					Description string `xml:"Description"`
				} `xml:"TechnicalEvaluationCriteria"`
				FinancialEvaluationCriteria []struct {
					Text                       string `xml:",chardata"`
					EvaluationCriteriaTypeCode struct {
						Text    string `xml:",chardata"`
						ListURI string `xml:"listURI,attr"`
					} `xml:"EvaluationCriteriaTypeCode"`
					Description string `xml:"Description"`
				} `xml:"FinancialEvaluationCriteria"`
				SpecificTendererRequirement []struct {
					Text                string `xml:",chardata"`
					RequirementTypeCode struct {
						Text    string `xml:",chardata"`
						ListURI string `xml:"listURI,attr"`
					} `xml:"RequirementTypeCode"`
				} `xml:"SpecificTendererRequirement"`
				Description                          string `xml:"Description"`
				RequiredBusinessClassificationScheme struct {
					Text                   string `xml:",chardata"`
					ID                     string `xml:"ID"`
					ClassificationCategory []struct {
						Text      string `xml:",chardata"`
						CodeValue string `xml:"CodeValue"`
					} `xml:"ClassificationCategory"`
				} `xml:"RequiredBusinessClassificationScheme"`
				PersonalSituation string `xml:"PersonalSituation"`
			} `xml:"TendererQualificationRequest"`
			FundingProgramCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"FundingProgramCode"`
			AllowedSubcontractTerms struct {
				Text        string `xml:",chardata"`
				Rate        string `xml:"Rate"`
				Description string `xml:"Description"`
			} `xml:"AllowedSubcontractTerms"`
			PriceRevisionFormulaDescription string `xml:"PriceRevisionFormulaDescription"`
			FundingProgram                  string `xml:"FundingProgram"`
			RequiredCurriculaIndicator      string `xml:"RequiredCurriculaIndicator"`
		} `xml:"TenderingTerms"`
		TenderingProcess struct {
			Text          string `xml:",chardata"`
			ProcedureCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"ProcedureCode"`
			UrgencyCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"UrgencyCode"`
			SubmissionMethodCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"SubmissionMethodCode"`
			TenderSubmissionDeadlinePeriod struct {
				Text        string `xml:",chardata"`
				EndDate     string `xml:"EndDate"`
				EndTime     string `xml:"EndTime"`
				Description string `xml:"Description"`
			} `xml:"TenderSubmissionDeadlinePeriod"`
			DocumentAvailabilityPeriod struct {
				Text    string `xml:",chardata"`
				EndDate string `xml:"EndDate"`
				EndTime string `xml:"EndTime"`
			} `xml:"DocumentAvailabilityPeriod"`
			ContractingSystemCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"ContractingSystemCode"`
			EconomicOperatorShortList struct {
				Text                  string `xml:",chardata"`
				ExpectedQuantity      string `xml:"ExpectedQuantity"`
				MaximumQuantity       string `xml:"MaximumQuantity"`
				MinimumQuantity       string `xml:"MinimumQuantity"`
				LimitationDescription string `xml:"LimitationDescription"`
			} `xml:"EconomicOperatorShortList"`
		} `xml:"TenderingProcess"`
		LegalDocumentReference struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"ID"`
			Attachment struct {
				Text              string `xml:",chardata"`
				ExternalReference struct {
					Text         string `xml:",chardata"`
					URI          string `xml:"URI"`
					DocumentHash string `xml:"DocumentHash"`
				} `xml:"ExternalReference"`
			} `xml:"Attachment"`
		} `xml:"LegalDocumentReference"`
		TechnicalDocumentReference struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"ID"`
			Attachment struct {
				Text              string `xml:",chardata"`
				ExternalReference struct {
					Text         string `xml:",chardata"`
					URI          string `xml:"URI"`
					DocumentHash string `xml:"DocumentHash"`
				} `xml:"ExternalReference"`
			} `xml:"Attachment"`
		} `xml:"TechnicalDocumentReference"`
		ValidNoticeInfo []struct {
			Text           string `xml:",chardata"`
			NoticeTypeCode struct {
				Text    string `xml:",chardata"`
				ListURI string `xml:"listURI,attr"`
			} `xml:"NoticeTypeCode"`
			AdditionalPublicationStatus []struct {
				Text                                   string `xml:",chardata"`
				PublicationMediaName                   string `xml:"PublicationMediaName"`
				AdditionalPublicationDocumentReference []struct {
					Text      string `xml:",chardata"`
					IssueDate string `xml:"IssueDate"`
				} `xml:"AdditionalPublicationDocumentReference"`
			} `xml:"AdditionalPublicationStatus"`
		} `xml:"ValidNoticeInfo"`
		AdditionalDocumentReference []struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"ID"`
			Attachment struct {
				Text              string `xml:",chardata"`
				ExternalReference struct {
					Text         string `xml:",chardata"`
					URI          string `xml:"URI"`
					DocumentHash string `xml:"DocumentHash"`
				} `xml:"ExternalReference"`
			} `xml:"Attachment"`
		} `xml:"AdditionalDocumentReference"`
		ProcurementProjectLot []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text       string `xml:",chardata"`
				SchemeName string `xml:"schemeName,attr"`
			} `xml:"ID"`
			ProcurementProject struct {
				Text         string `xml:",chardata"`
				Name         string `xml:"Name"`
				BudgetAmount struct {
					Text        string `xml:",chardata"`
					TotalAmount struct {
						Text       string `xml:",chardata"`
						CurrencyID string `xml:"currencyID,attr"`
					} `xml:"TotalAmount"`
					TaxExclusiveAmount struct {
						Text       string `xml:",chardata"`
						CurrencyID string `xml:"currencyID,attr"`
					} `xml:"TaxExclusiveAmount"`
				} `xml:"BudgetAmount"`
				RequiredCommodityClassification []struct {
					Text                   string `xml:",chardata"`
					ItemClassificationCode struct {
						Text    string `xml:",chardata"`
						ListURI string `xml:"listURI,attr"`
					} `xml:"ItemClassificationCode"`
				} `xml:"RequiredCommodityClassification"`
			} `xml:"ProcurementProject"`
		} `xml:"ProcurementProjectLot"`
		ContractModification []struct {
			Text                                string `xml:",chardata"`
			ID                                  string `xml:"ID"`
			Note                                string `xml:"Note"`
			ContractModificationDurationMeasure struct {
				Text     string `xml:",chardata"`
				UnitCode string `xml:"unitCode,attr"`
			} `xml:"ContractModificationDurationMeasure"`
			FinalDurationMeasure struct {
				Text     string `xml:",chardata"`
				UnitCode string `xml:"unitCode,attr"`
			} `xml:"FinalDurationMeasure"`
			ContractID                             string `xml:"ContractID"`
			ContractModificationLegalMonetaryTotal struct {
				Text               string `xml:",chardata"`
				TaxExclusiveAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxExclusiveAmount"`
			} `xml:"ContractModificationLegalMonetaryTotal"`
			FinalLegalMonetaryTotal struct {
				Text               string `xml:",chardata"`
				TaxExclusiveAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxExclusiveAmount"`
			} `xml:"FinalLegalMonetaryTotal"`
		} `xml:"ContractModification"`
	} `xml:"ContractFolderStatus"`
}

type TenderResult struct {
	Text       string `xml:",chardata"`
	ResultCode struct {
		Text    string `xml:",chardata"`
		ListURI string `xml:"listURI,attr"`
	} `xml:"ResultCode"`
	Description string `xml:"Description"`
	Contract    struct {
		Text      string `xml:",chardata"`
		IssueDate string `xml:"IssueDate"`
		ID        string `xml:"ID"`
	} `xml:"Contract"`
	WinningParty struct {
		Text                string `xml:",chardata"`
		PartyIdentification struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text       string `xml:",chardata"`
				SchemeName string `xml:"schemeName,attr"`
			} `xml:"ID"`
		} `xml:"PartyIdentification"`
		PartyName struct {
			Text string `xml:",chardata"`
			Name string `xml:"Name"`
		} `xml:"PartyName"`
	} `xml:"WinningParty"`
	AwardedTenderedProject struct {
		Text               string `xml:",chardata"`
		LegalMonetaryTotal struct {
			Text               string `xml:",chardata"`
			TaxExclusiveAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxExclusiveAmount"`
			PayableAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"PayableAmount"`
		} `xml:"LegalMonetaryTotal"`
		ProcurementProjectLotID string `xml:"ProcurementProjectLotID"`
	} `xml:"AwardedTenderedProject"`
	StartDate              string `xml:"StartDate"`
	AwardDate              string `xml:"AwardDate"`
	ReceivedTenderQuantity string `xml:"ReceivedTenderQuantity"`
	LowerTenderAmount      struct {
		Text       string `xml:",chardata"`
		CurrencyID string `xml:"currencyID,attr"`
	} `xml:"LowerTenderAmount"`
	HigherTenderAmount struct {
		Text       string `xml:",chardata"`
		CurrencyID string `xml:"currencyID,attr"`
	} `xml:"HigherTenderAmount"`
}
