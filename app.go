package codice

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/gocarina/gocsv"
	"github.com/saracen/go7z"
	"golang.org/x/net/html/charset"
)

const replacement = " "

var replacer = strings.NewReplacer(
	"\r\n", replacement,
	"\r", replacement,
	"\n", replacement,
	"\v", replacement,
	"\f", replacement,
	"\u0085", replacement,
	"\u2028", replacement,
	"\u2029", replacement,
)

type CodiceApp struct {
	feed *Feed
}

func NewCodiceApp() (*CodiceApp, error) {
	c := CodiceApp{}
	return &c, nil
}

// LoadZip downloads a zip file from a HTTP server and parses its content
func (c *CodiceApp) LoadWebZip(url string) (*[]Entry, error) {
	var r *bytes.Reader
	var size int64

	if strings.HasPrefix(url, "http") {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		i, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		bar := pb.Default.Start(i)
		bar.Set(pb.Bytes, true)
		bar.SetTemplate("Downloading... {{speed . }} (downloaded: {{counters .}})")
		bar.Start()
		reader := bar.NewProxyReader(resp.Body)
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		bar.Finish()
		r = bytes.NewReader(body)
		size = int64(len(body))
	} else {
		content, err := ioutil.ReadFile(url)
		if err != nil {
			log.Fatal(err)
		}
		r = bytes.NewReader(content)
		size = int64(len(content))
	}

	// Most of files are .zip, so we assume first we are dealing with one of those.
	parsedEntries, err := c.loadZipFromMemory(r, size)
	if err != nil {
		fmt.Println("Uh, this does not look like zip. Trying 7z...")
		parsedEntries, err = c.load7zFromMemory(r, size)
		if err != nil {
			fmt.Printf("Error parsing 7zip: %s", err)
		}
	}
	return parsedEntries, nil
}

// load7z loads a zip-formatted file already in memory
func (c *CodiceApp) loadZipFromMemory(r *bytes.Reader, size int64) (*[]Entry, error) {
	zipReader, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	parsedEntries := []Entry{}
	for _, zipFile := range zipReader.File {
		fmt.Println("Parsing file:", zipFile.Name)
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			return nil, err
		}
		entries, err := c.parseXML(unzippedFileBytes)
		if err != nil {
			return nil, err
		}
		parsedEntries = append(parsedEntries, *entries...)
	}
	return &parsedEntries, nil
}

// load7z loads a 7z-formatted file already in memory
func (c *CodiceApp) load7zFromMemory(r *bytes.Reader, size int64) (*[]Entry, error) {
	sevenR, err := go7z.NewReader(r, size)
	if err != nil {
		return nil, err
	}

	parsedEntries := []Entry{}
	for {
		fileInfo, err := sevenR.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}
		if fileInfo.IsEmptyStream && !fileInfo.IsEmptyFile {
			// If it is not an empty file nor empty stream then its a directory
			//fmt.Printf("%s is a directory. Ignoring...\n", fileInfo.Name)
			continue
		}
		fmt.Println("Parsing file:", fileInfo.Name)
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		entries, err := c.parseXML(data)
		if err != nil {
			return nil, err
		}
		parsedEntries = append(parsedEntries, *entries...)
	}
	return &parsedEntries, nil
}

// LoadXMLFromFs loads a XML from contrataciondelestado.es from the file system
func (c *CodiceApp) LoadXMLFromFs(path string) (*[]Entry, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	xmlData, err := ioutil.ReadAll(xmlFile)
	xmlFile.Close()
	if err != nil {
		return nil, err
	}
	entries, err := c.parseXML(xmlData)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (c *CodiceApp) parseXML(xmlData []byte) (*[]Entry, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	feed := Feed{}

	if err := decoder.Decode(&feed); err != nil {
		return nil, err
	}
	return &feed.Entry, nil
}

// FlattenToCsv takes a slice of Entry and flattens it into a series of csv files.
func (c *CodiceApp) FlattenToCsv(entries *[]Entry, outputPrefix string) error {
	lines := []CodiceEntry{}
	modificationLines := []Modification{}
	guaranteeLines := []FinancialGuarantee{{}}
	for _, e := range *entries {
		cntr := e.ContractFolderStatus

		cpvs := []string{}
		for _, c := range cntr.ProcurementProject.RequiredCommodityClassification {
			cpvs = append(cpvs, c.ItemClassificationCode.Text)
		}

		l := CodiceEntry{
			EntryID:    e.ID,
			Updated:    e.Updated,
			FolderID:   sanitize(cntr.ContractFolderID),
			Title:      sanitize(cntr.ProcurementProject.Name),
			Summary:    sanitize(e.Summary.Text),
			StatusCode: sanitize(cntr.ContractFolderStatusCode.Text),

			ContractingPartyIdentification: sanitize(cntr.LocatedContractingParty.Party.PartyIdentification.ID.Text),
			ContractingPartyName:           sanitize(cntr.LocatedContractingParty.Party.PartyName.Name),
			ContractingPartyWebsite:        sanitize(cntr.LocatedContractingParty.Party.WebsiteURI),
			ContractingPartyType:           sanitize(cntr.LocatedContractingParty.ContractingPartyTypeCode.Text),

			TypeCode:          sanitize(cntr.ProcurementProject.TypeCode.Text),
			SubTypeCode:       sanitize(cntr.ProcurementProject.SubTypeCode.Text),
			CPVClassification: strings.Join(cpvs, ";"),

			TechnicalInstructionsURL: cntr.TechnicalDocumentReference.Attachment.ExternalReference.URI,

			BudgetEstimatedOverallContractAmount: cntr.ProcurementProject.BudgetAmount.EstimatedOverallContractAmount.Text,
			BudgetTotalAmount:                    cntr.ProcurementProject.BudgetAmount.TotalAmount.Text,
			BudgetTaxExclusiveAmount:             cntr.ProcurementProject.BudgetAmount.TaxExclusiveAmount.Text,

			RealizedLocationCountrySubentity:     cntr.ProcurementProject.RealizedLocation.CountrySubentity,
			RealizedLocationCountrySubentityCode: cntr.ProcurementProject.SubTypeCode.Text,
			RealizedLocationAddressCountry:       cntr.ProcurementProject.RealizedLocation.Address.Country.Name,
			RealizedLocationAddressCountryCode:   cntr.ProcurementProject.RealizedLocation.Address.Country.IdentificationCode.Text,
			RealizedLocationAddressCityName:      cntr.ProcurementProject.RealizedLocation.Address.CityName,
			RealizedLocationAddressPostalZone:    cntr.ProcurementProject.RealizedLocation.Address.PostalZone,
			RealizedLocationAddressAddressLine:   cntr.ProcurementProject.RealizedLocation.Address.AddressLine.Line,

			PlannedPeriodDurationMeasure:         cntr.ProcurementProject.PlannedPeriod.DurationMeasure.Text,
			PlannedPeriodDurationMeasureUnitCode: cntr.ProcurementProject.PlannedPeriod.DurationMeasure.UnitCode,
			PlannedPeriodStartDate:               cntr.ProcurementProject.PlannedPeriod.StartDate,
			PlannedPeriodEndDate:                 cntr.ProcurementProject.PlannedPeriod.EndDate,

			ContractExtensionOptionsDescription:        cntr.ProcurementProject.ContractExtension.OptionsDescription,
			ContractExtensionValidityPeriodDescription: cntr.ProcurementProject.ContractExtension.OptionValidityPeriod.Description,

			TenderingTermsFundingProgram:                  sanitize(cntr.TenderingTerms.FundingProgram),
			TenderingTermsFundingProgramCode:              sanitize(cntr.TenderingTerms.FundingProgramCode.Text),
			TenderingTermsLanguage:                        sanitize(cntr.TenderingTerms.Language.ID),
			TenderingTermsLRequiredCurriculaIndicator:     sanitize(cntr.TenderingTerms.RequiredCurriculaIndicator),
			TenderingTermsVariantConstraintIndicator:      sanitize(cntr.TenderingTerms.VariantConstraintIndicator),
			TenderingTermsPriceRevisionFormulaDescription: sanitize(cntr.TenderingTerms.PriceRevisionFormulaDescription),

			TenderingTermsSubcontractTermsRate:        sanitize(cntr.TenderingTerms.AllowedSubcontractTerms.Rate),
			TenderingTermsSubcontractTermsDescription: sanitize(cntr.TenderingTerms.AllowedSubcontractTerms.Description),

			TenderingProcessProcedureCode:         sanitize(cntr.TenderingProcess.ProcedureCode.Text),
			TenderingProcessContractingSystemCode: sanitize(cntr.TenderingProcess.ContractingSystemCode.Text),
			TenderingProcessSubmissionMethodCode:  sanitize(cntr.TenderingProcess.SubmissionMethodCode.Text),
			TenderingProcessUrgencyCode:           sanitize(cntr.TenderingProcess.UrgencyCode.Text),
			TenderingProcessSubmissionEndDate:     sanitize(cntr.TenderingProcess.TenderSubmissionDeadlinePeriod.EndDate),
			TenderingProcessSubmissionEndTime:     sanitize(cntr.TenderingProcess.TenderSubmissionDeadlinePeriod.EndTime),

			TenderingProcessEconomicOperatorLimitationDescription: sanitize(cntr.TenderingProcess.EconomicOperatorShortList.LimitationDescription),
			TenderingProcessEconomicOperatorExpectedQuantity:      sanitize(cntr.TenderingProcess.EconomicOperatorShortList.ExpectedQuantity),
			TenderingProcessEconomicOperatorMaximumQuantity:       sanitize(cntr.TenderingProcess.EconomicOperatorShortList.MaximumQuantity),
			TenderingProcessEconomicOperatorMinimunQuantity:       sanitize(cntr.TenderingProcess.EconomicOperatorShortList.MinimumQuantity),
		}

		if len(cntr.ProcurementProjectLot) > 0 {
			for _, lot := range cntr.ProcurementProjectLot {
				l.LotID = sanitize(lot.ID.Text)
				l.LotName = sanitize(lot.ProcurementProject.Name)
				l.LotTotalAmount = sanitize(lot.ProcurementProject.BudgetAmount.TotalAmount.Text)
				l.LotTaxExclusiveAmount = sanitize(lot.ProcurementProject.BudgetAmount.TaxExclusiveAmount.Text)
				lcpvs := []string{}
				for _, c := range lot.ProcurementProject.RequiredCommodityClassification {
					lcpvs = append(lcpvs, c.ItemClassificationCode.Text)
				}
				l.LotCPVClassification = strings.Join(lcpvs, ";")

				for _, r := range cntr.TenderResult {
					if r.AwardedTenderedProject.ProcurementProjectLotID == lot.ID.Text {
						l.TenderResultCode = sanitize(r.ResultCode.Text)
						l.TenderDescription = sanitize(r.Description)
						l.TenderContractID = sanitize(r.Contract.ID)
						l.TenderContractIssueDate = sanitize(r.Contract.IssueDate)
						l.TenderWinningPartyID = sanitize(r.WinningParty.PartyIdentification.ID.Text)
						l.TenderWinningPartyName = sanitize(r.WinningParty.PartyName.Name)
						l.TenderWinningPartyScheme = sanitize(r.WinningParty.PartyIdentification.ID.SchemeName)
						l.TenderPayableAmount = sanitize(r.AwardedTenderedProject.LegalMonetaryTotal.PayableAmount.Text)
						l.TenderTaxExclusiveAmount = sanitize(r.AwardedTenderedProject.LegalMonetaryTotal.TaxExclusiveAmount.Text)
						l.TenderStartDate = sanitize(r.StartDate)
						l.TenderAwardDate = sanitize(r.AwardDate)
						l.TenderReceivedTenderQuantity = sanitize(r.ReceivedTenderQuantity)
						l.TenderLowerLenderAmount = sanitize(r.LowerTenderAmount.Text)
						l.TenderHigherTenderAmount = sanitize(r.HigherTenderAmount.Text)
					}
				}
				lines = append(lines, l)

			}
		} else {
			if len(cntr.TenderResult) == 1 {
				r := cntr.TenderResult[0]
				l.TenderResultCode = sanitize(r.ResultCode.Text)
				l.TenderDescription = sanitize(r.Description)
				l.TenderStartDate = sanitize(r.StartDate)
				l.TenderContractID = sanitize(r.Contract.ID)
				l.TenderContractIssueDate = sanitize(r.Contract.IssueDate)
				l.TenderWinningPartyID = sanitize(r.WinningParty.PartyIdentification.ID.Text)
				l.TenderWinningPartyName = sanitize(r.WinningParty.PartyName.Name)
				l.TenderWinningPartyScheme = sanitize(r.WinningParty.PartyIdentification.ID.SchemeName)
				l.TenderPayableAmount = sanitize(r.AwardedTenderedProject.LegalMonetaryTotal.PayableAmount.Text)
				l.TenderTaxExclusiveAmount = sanitize(r.AwardedTenderedProject.LegalMonetaryTotal.TaxExclusiveAmount.Text)
				l.TenderStartDate = sanitize(r.StartDate)
				l.TenderAwardDate = sanitize(r.AwardDate)
				l.TenderReceivedTenderQuantity = sanitize(r.ReceivedTenderQuantity)
				l.TenderLowerLenderAmount = sanitize(r.LowerTenderAmount.Text)
				l.TenderHigherTenderAmount = sanitize(r.HigherTenderAmount.Text)
			}

			lines = append(lines, l)
		}

		for _, m := range cntr.ContractModification {
			modification := Modification{
				EntryID:                                sanitize(e.ID),
				ID:                                     sanitize(m.ID),
				FolderID:                               sanitize(m.ContractID),
				Note:                                   sanitize(m.Note),
				ContractModificationDuration:           m.ContractModificationDurationMeasure.Text,
				ContractModificationDurationUnit:       m.ContractModificationDurationMeasure.UnitCode,
				ContractModificationFinalDuration:      m.FinalDurationMeasure.Text,
				ContractModificationFinalDurationUnit:  m.FinalDurationMeasure.UnitCode,
				ContractModificationLegalMonetaryTotal: m.ContractModificationLegalMonetaryTotal.TaxExclusiveAmount.Text,
				FinalLegalMonetaryTotal:                m.FinalLegalMonetaryTotal.TaxExclusiveAmount.Text,
			}
			modificationLines = append(modificationLines, modification)
		}

		for _, fg := range cntr.TenderingTerms.RequiredFinancialGuarantee {
			guarantee := FinancialGuarantee{
				EntryID:                   sanitize(e.ID),
				FolderID:                  sanitize(cntr.ContractFolderID),
				GuaranteeTypeCode:         fg.GuaranteeTypeCode.Text,
				AmountRate:                fg.AmountRate,
				LiabilityAmount:           fg.LiabilityAmount.Text,
				LiabilityAmountCurrencyID: fg.LiabilityAmount.CurrencyID,
			}
			guaranteeLines = append(guaranteeLines, guarantee)
		}
	}

	entriesPath := fmt.Sprintf("%s_entries.csv", outputPrefix)
	outputEntriesFile, err := os.Create(entriesPath)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(lines, outputEntriesFile)
	outputEntriesFile.Close()
	fmt.Printf("Entries rows written to %s\n", entriesPath)

	modificationsPath := fmt.Sprintf("%s_modifications.csv", outputPrefix)
	outputModificationsFile, err := os.Create(modificationsPath)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(modificationLines, outputModificationsFile)
	outputModificationsFile.Close()
	fmt.Printf("Modifications rows written to %s\n", modificationsPath)

	guaranteesPath := fmt.Sprintf("%s_financial_guarantees.csv", outputPrefix)
	outputGuaranteesFile, err := os.Create(guaranteesPath)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(guaranteeLines, outputGuaranteesFile)
	outputGuaranteesFile.Close()
	fmt.Printf("Financial Guarantee rows written to %s\n", guaranteesPath)

	return nil
}

func sanitize(str string) string {
	str = strings.TrimSpace(str)
	str = replacer.Replace(str)
	return str
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
