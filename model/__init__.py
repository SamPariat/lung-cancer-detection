import os
from flask import Flask, request, jsonify
import numpy as np
import pickle
from sklearn.preprocessing import LabelEncoder, StandardScaler

app = Flask(__name__)

@app.route('/', methods=['GET'])
def health_check():
    return jsonify({
        'message': 'Server is up',
        'status': 200,
    })

with open('model.pkl', 'rb') as model_file:
    model = pickle.load(model_file)

with open('scalar_data.pkl', 'rb') as scalar_file:
    model_scaler = pickle.load(scalar_file)

# Fit LabelEncoders for categorical features
label_encoders = {}
for feature in ['alcoholConsuming', 'allergy', 'anxiety', 'chestPain',
                'chronicDisease', 'coughing', 'fatigue', 'lungCancer',
                'peerPressure', 'shortnessOfBreath', 'smoking',
                'swallowingDifficulty', 'wheezing', 'yellowFingers']:
    label_encoders[feature] = LabelEncoder()
    label_encoders[feature].fit([1, 2])  # Assuming values are 1 and 2

@app.route('/prediction', methods=['POST'])
def predict_lung_cancer():
    data = request.json

    gender = data['gender']
    age = data['age']
    smoking = data['smoking']
    yellowFingers = data['yellowFingers']
    anxiety = data['anxiety']
    peerPressure = data['peerPressure']
    chronicDisease = data['chronicDisease']
    fatigue = data['fatigue']
    allergy = data['allergy']
    wheezing = data['wheezing']
    alcoholConsuming = data['alcoholConsuming']
    coughing = data['coughing']
    shortnessOfBreath = data['shortnessOfBreath']
    swallowingDifficulty = data['swallowingDifficulty']
    chestPain = data['chestPain']

    scalar = StandardScaler()
    scalar.mean_ = model_scaler['mean']
    scalar.scale_ = model_scaler['scale']

    age_scaled = scalar.transform([[age]])[0][0]

    features = np.array([
        gender,
        age_scaled,
        smoking,
        yellowFingers,
        anxiety,
        peerPressure,
        chronicDisease,
        fatigue,
        allergy,
        wheezing,
        alcoholConsuming,
        coughing,
        shortnessOfBreath,
        swallowingDifficulty,
        chestPain
    ])

    # Make predictions using the model
    prediction = model.predict(features.reshape(1, -1))

    # Prepare the prediction result in JSON format
    result = {'prediction': prediction[0]}

    # Return the prediction result as JSON
    return jsonify(result)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=1234, debug=True)




    {
  "age": 2,
  "alcoholConsuming": 2,
  "allergy": 2,
  "anxiety": 2,
  "chestPain": 2,
  "chronicDisease": 2,
  "coughing": 2,
  "fatigue": 2,
  "gender": "M",
  "lungCancer": "YES",
  "peerPressure": 2,
  "shortnessOfBreath": 2,
  "smoking": 2,
  "swallowingDifficulty": 2,
  "wheezing": 2,
  "yellowFingers": 2
}