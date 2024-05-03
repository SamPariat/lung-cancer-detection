package predict_result

import (
	"errors"
	"fmt"
)

const numberErrorText = "where 1 means NO and 2 means YES, but provided value is : "

func (cliResult *CLIResult) ValidateCLIEntries() error {

	if cliResult.Gender != "M" && cliResult.Gender != "F" {
		return errors.New("gender must be either M or F")
	}

	if cliResult.LungCancer != "YES" && cliResult.LungCancer != "NO" {
		return errors.New("lung cancer must be either YES or NO")
	}

	if cliResult.Age < 0 || cliResult.Age > 100 {
		return fmt.Errorf("age must be between 0 and 100, but provided value is:%d", cliResult.Age)
	}

	if cliResult.Smoking < 1 || cliResult.Smoking > 2 {
		return fmt.Errorf("smoking must be either 1 or 2, %v%d", numberErrorText, cliResult.Smoking)
	}

	if cliResult.YellowFingers < 1 || cliResult.YellowFingers > 2 {
		return fmt.Errorf("yellow fingers must be either 1 or 2, %v%d", numberErrorText, cliResult.YellowFingers)
	}

	if cliResult.Anxiety < 1 || cliResult.Anxiety > 2 {
		return fmt.Errorf("anxiety must be either 1 or 2, %v%d", numberErrorText, cliResult.Anxiety)
	}

	if cliResult.PeerPressure < 1 || cliResult.PeerPressure > 2 {
		return fmt.Errorf("peer pressure must be either 1 or 2, %v%d", numberErrorText, cliResult.PeerPressure)
	}

	if cliResult.ChronicDisease < 1 || cliResult.ChronicDisease > 2 {
		return fmt.Errorf("chronic disease must be either 1 or 2, %v%d", numberErrorText, cliResult.ChronicDisease)
	}

	if cliResult.Fatigue < 1 || cliResult.Fatigue > 2 {
		return fmt.Errorf("fatigue must be either 1 or 2, %v%d", numberErrorText, cliResult.Fatigue)
	}

	if cliResult.Allergy < 1 || cliResult.Allergy > 2 {
		return fmt.Errorf("allergy must be either 1 or 2, %v%d", numberErrorText, cliResult.Allergy)
	}

	if cliResult.Wheezing < 1 || cliResult.Wheezing > 2 {
		return fmt.Errorf("wheezing must be either 1 or 2, %v%d", numberErrorText, cliResult.Wheezing)
	}

	if cliResult.AlcoholConsuming < 1 || cliResult.AlcoholConsuming > 2 {
		return fmt.Errorf("alcohol consuming must be either 1 or 2, %v%d", numberErrorText, cliResult.AlcoholConsuming)
	}

	if cliResult.Coughing < 1 || cliResult.Coughing > 2 {
		return fmt.Errorf("coughing must be either 1 or 2, %v%d", numberErrorText, cliResult.Coughing)
	}

	if cliResult.ShortnessOfBreath < 1 || cliResult.ShortnessOfBreath > 2 {
		return fmt.Errorf("shortness of breath must be either 1 or 2, %v%d", numberErrorText, cliResult.ShortnessOfBreath)
	}

	if cliResult.SwallowingDifficulty < 1 || cliResult.SwallowingDifficulty > 2 {
		return fmt.Errorf("swallowing difficulty must be either 1 or 2, %v%d", numberErrorText, cliResult.SwallowingDifficulty)
	}

	if cliResult.ChestPain < 1 || cliResult.ChestPain > 2 {
		return fmt.Errorf("chest pain must be either 1 or 2, %v%d", numberErrorText, cliResult.ChestPain)
	}

	return nil
}
