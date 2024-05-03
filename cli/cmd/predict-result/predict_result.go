/*
Copyright © 2024 NAME HERE apunthelegend@outlook.com
*/
package predict_result

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/lung-cancer-detection/constants"
)

var cliResult CLIResult

func RunCmd(cmd *cobra.Command, args []string) {
	err := cliResult.ValidateCLIEntries()
	if err != nil {
		log.Fatal(err)
	}

	postBodyBytes, err := json.Marshal(&cliResult)
	if err != nil {
		log.Fatal(err)
	}

	postBodyBuffer := bytes.NewBuffer(postBodyBytes)

	httpResponse, err := http.Post(constants.BaseUrl, constants.ApplicationJson, postBodyBuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	stringBody := string(body)

	log.Printf("%s", stringBody)
}

var PredictResultCmd = &cobra.Command{
	Use:   "predict-result",
	Short: "The command used for predicting lung cancer detection result",
	Long: `Run the predict-result command to predict whether the patient has lung cancer or not.

Example of usage: predict-result
`,
	Run: RunCmd,
}

func init() {
	PredictResultCmd.PersistentFlags().StringVarP(
		&cliResult.Gender,
		constants.GenderFlag,
		constants.GenderFlagShort,
		"",
		"Gender",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Age,
		constants.AgeFlag,
		constants.AgeFlagShort,
		0,
		"Age",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Smoking,
		constants.SmokingFlag,
		constants.SmokingFlagShort,
		0,
		"Shortness of breath",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.YellowFingers,
		constants.YellowFingersFlag,
		constants.YellowFingersFlagShort,
		0,
		"Yellow fingers",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Anxiety,
		constants.AnxietyFlag,
		constants.AnxietyFlagShort,
		0,
		"Anxiety",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.PeerPressure,
		constants.PeerPressureFlag,
		constants.PeerPressureFlagShort,
		0,
		"Peer pressure",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.ChronicDisease,
		constants.ChronicDisease,
		constants.ChronicDiseaseFlagShort,
		0,
		"Chronic disease",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Fatigue,
		constants.FatigueFlag,
		constants.FatigueFlagShort,
		0,
		"Fatigue",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Allergy,
		constants.AllergyFlag,
		constants.AllergyFlagShort,
		0,
		"Allergy",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Wheezing,
		constants.WheezingFlag,
		constants.WheezingFlagShort,
		0,
		"Wheezing",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.AlcoholConsuming,
		constants.AlcoholConsumingFlag,
		constants.AlcoholConsumingFlagShort,
		0,
		"Alcohol consumption",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.Coughing,
		constants.CoughingFlag,
		constants.CoughingFlagShort,
		0,
		"Coughing",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.ShortnessOfBreath,
		constants.ShortnessOfBreathFlag,
		constants.ShortnessOfBreathFlagShort,
		0,
		"Shortness of breath",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.SwallowingDifficulty,
		constants.SwallowingDifficultyFlag,
		constants.SwallowingDifficultyFlagShort,
		0,
		"Swallowing difficulty",
	)
	PredictResultCmd.PersistentFlags().IntVarP(
		&cliResult.ChestPain,
		constants.ChestPainFlag,
		constants.ChestPainFlagShort,
		0,
		"Chest pain",
	)
	PredictResultCmd.PersistentFlags().StringVarP(
		&cliResult.LungCancer,
		constants.LungCancerFlag,
		constants.LungCancerFlagShort,
		"",
		"Lung cancer",
	)

	PredictResultCmd.MarkPersistentFlagRequired(constants.GenderFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.AgeFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.SmokingFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.YellowFingersFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.AnxietyFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.PeerPressureFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.ChronicDisease)
	PredictResultCmd.MarkPersistentFlagRequired(constants.FatigueFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.AllergyFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.WheezingFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.AlcoholConsumingFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.CoughingFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.ShortnessOfBreathFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.SwallowingDifficultyFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.ChestPainFlag)
	PredictResultCmd.MarkPersistentFlagRequired(constants.LungCancerFlag)
}
