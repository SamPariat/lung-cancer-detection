package predict_result

type CLIResult struct {
	Gender               string `json:"gender"`
	Age                  int    `json:"age"`
	Smoking              int    `json:"smoking"`
	YellowFingers        int    `json:"yellowFingers"`
	Anxiety              int    `json:"anxiety"`
	PeerPressure         int    `json:"peerPressure"`
	ChronicDisease       int    `json:"chronicDisease"`
	Fatigue              int    `json:"fatigue"`
	Allergy              int    `json:"allergy"`
	Wheezing             int    `json:"wheezing"`
	AlcoholConsuming     int    `json:"alcoholConsuming"`
	Coughing             int    `json:"coughing"`
	ShortnessOfBreath    int    `json:"shortnessOfBreath"`
	SwallowingDifficulty int    `json:"swallowingDifficulty"`
	ChestPain            int    `json:"chestPain"`
	LungCancer           string `json:"lungCancer"`
}
