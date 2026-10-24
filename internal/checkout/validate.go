package checkout

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	postalUS = regexp.MustCompile(`^\d{5}(-\d{4})?$`)
	postalCA = regexp.MustCompile(`^[A-Z]\d[A-Z] ?\d[A-Z]\d$`)
	postalGB = regexp.MustCompile(`^[A-Z]{1,2}\d[A-Z\d]? ?\d[A-Z]{2}$`)
	postalDE = regexp.MustCompile(`^\d{5}$`)
	postalFR = regexp.MustCompile(`^\d{5}$`)
	postalES = regexp.MustCompile(`^\d{5}$`)
	postalIT = regexp.MustCompile(`^\d{5}$`)
	postalNL = regexp.MustCompile(`^\d{4} ?[A-Z]{2}$`)
	postalBE = regexp.MustCompile(`^\d{4}$`)
	postalCH = regexp.MustCompile(`^\d{4}$`)
	postalAT = regexp.MustCompile(`^\d{4}$`)
	postalSE = regexp.MustCompile(`^\d{3} ?\d{2}$`)
	postalNO = regexp.MustCompile(`^\d{4}$`)
	postalDK = regexp.MustCompile(`^\d{4}$`)
	postalFI = regexp.MustCompile(`^\d{5}$`)
	postalPL = regexp.MustCompile(`^\d{2}-\d{3}$`)
	postalPT = regexp.MustCompile(`^\d{4}-\d{3}$`)
	postalIE = regexp.MustCompile(`^[A-Z]\d{2} ?[A-Z\d]{4}$`)
	postalAU = regexp.MustCompile(`^\d{4}$`)
	postalJP = regexp.MustCompile(`^\d{3}-\d{4}$`)
)

// ValidateAddress checks that a shipping address is complete and that its
// postal code has the right shape for its country.
func ValidateAddress(a Address) error {
	if strings.TrimSpace(a.Line1) == "" {
		return errors.New("checkout: address line 1 is required")
	}
	if strings.TrimSpace(a.City) == "" {
		return errors.New("checkout: city is required")
	}
	postal := strings.ToUpper(strings.TrimSpace(a.Postal))
	switch strings.ToUpper(a.Country) {
	case "US":
		return validateUS(postal)
	case "CA":
		return validateCA(postal)
	case "GB":
		return validateGB(postal)
	case "DE":
		return validateDE(postal)
	case "FR":
		return validateFR(postal)
	case "ES":
		return validateES(postal)
	case "IT":
		return validateIT(postal)
	case "NL":
		return validateNL(postal)
	case "BE":
		return validateBE(postal)
	case "CH":
		return validateCH(postal)
	case "AT":
		return validateAT(postal)
	case "SE":
		return validateSE(postal)
	case "NO":
		return validateNO(postal)
	case "DK":
		return validateDK(postal)
	case "FI":
		return validateFI(postal)
	case "PL":
		return validatePL(postal)
	case "PT":
		return validatePT(postal)
	case "IE":
		return validateIE(postal)
	case "AU":
		return validateAU(postal)
	case "JP":
		return validateJP(postal)
	default:
		return fmt.Errorf("checkout: unsupported country %q", a.Country)
	}
}

// validateUS checks a US postal code.
func validateUS(postal string) error {
	if !postalUS.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid US postal code", postal)
	}
	return nil
}

// validateCA checks a CA postal code.
func validateCA(postal string) error {
	if !postalCA.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid CA postal code", postal)
	}
	return nil
}

// validateGB checks a GB postal code.
func validateGB(postal string) error {
	if !postalGB.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid GB postal code", postal)
	}
	return nil
}

// validateDE checks a DE postal code.
func validateDE(postal string) error {
	if !postalDE.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid DE postal code", postal)
	}
	return nil
}

// validateFR checks a FR postal code.
func validateFR(postal string) error {
	if !postalFR.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid FR postal code", postal)
	}
	return nil
}

// validateES checks a ES postal code.
func validateES(postal string) error {
	if !postalES.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid ES postal code", postal)
	}
	return nil
}

// validateIT checks a IT postal code.
func validateIT(postal string) error {
	if !postalIT.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid IT postal code", postal)
	}
	return nil
}

// validateNL checks a NL postal code.
func validateNL(postal string) error {
	if !postalNL.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid NL postal code", postal)
	}
	return nil
}

// validateBE checks a BE postal code.
func validateBE(postal string) error {
	if !postalBE.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid BE postal code", postal)
	}
	return nil
}

// validateCH checks a CH postal code.
func validateCH(postal string) error {
	if !postalCH.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid CH postal code", postal)
	}
	return nil
}

// validateAT checks a AT postal code.
func validateAT(postal string) error {
	if !postalAT.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid AT postal code", postal)
	}
	return nil
}

// validateSE checks a SE postal code.
func validateSE(postal string) error {
	if !postalSE.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid SE postal code", postal)
	}
	return nil
}

// validateNO checks a NO postal code.
func validateNO(postal string) error {
	if !postalNO.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid NO postal code", postal)
	}
	return nil
}

// validateDK checks a DK postal code.
func validateDK(postal string) error {
	if !postalDK.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid DK postal code", postal)
	}
	return nil
}

// validateFI checks a FI postal code.
func validateFI(postal string) error {
	if !postalFI.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid FI postal code", postal)
	}
	return nil
}

// validatePL checks a PL postal code.
func validatePL(postal string) error {
	if !postalPL.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid PL postal code", postal)
	}
	return nil
}

// validatePT checks a PT postal code.
func validatePT(postal string) error {
	if !postalPT.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid PT postal code", postal)
	}
	return nil
}

// validateIE checks a IE postal code.
func validateIE(postal string) error {
	if !postalIE.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid IE postal code", postal)
	}
	return nil
}

// validateAU checks a AU postal code.
func validateAU(postal string) error {
	if !postalAU.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid AU postal code", postal)
	}
	return nil
}

// validateJP checks a JP postal code.
func validateJP(postal string) error {
	if !postalJP.MatchString(postal) {
		return fmt.Errorf("checkout: %q is not a valid JP postal code", postal)
	}
	return nil
}
