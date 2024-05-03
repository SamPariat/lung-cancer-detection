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

with open('scalar_data.pkl', 'rb') as scalar_file:
    model_scaler = pickle.load(scalar_file)

with open('random_forest.pkl', 'rb') as random_forest_file:
    random_forest_classifier = pickle.load(random_forest_file)

with open('svm.pkl', 'rb') as svm_file:
    svm_classifier = pickle.load(svm_file)

with open('knn.pkl', 'rb') as knn_file:
    knn_classifier = pickle.load(knn_file)

with open('adaboost.pkl', 'rb') as adaboost_file:
    adaboost = pickle.load(adaboost_file)

@app.route('/prediction', methods=['POST'])
def predict_lung_cancer():
    print(request.json)

    data = request.json

    gender = 1 if data['gender'] == 'M' else 0
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
        smoking - 1,
        yellowFingers - 1,
        anxiety - 1,
        peerPressure - 1,
        chronicDisease - 1,
        fatigue - 1,
        allergy - 1,
        wheezing - 1,
        alcoholConsuming - 1,
        coughing - 1,
        shortnessOfBreath - 1,
        swallowingDifficulty - 1,
        chestPain - 1
    ])

    features = features.reshape(1, -1)

    print(features)

    random_forest_prediction = random_forest_classifier.predict(features)
    svm_prediction = svm_classifier.predict(features)
    knn_prediction = knn_classifier.predict(features)
    adaboost_prediction = adaboost.predict(features)

    print(random_forest_prediction)
    print(svm_prediction)
    print(knn_prediction)
    print(adaboost_prediction)

    result = {
        'randomForestPrediction': "YES" if int(random_forest_prediction[0]) == 1 else "NO",
        'svmPrediction': "YES" if int(svm_prediction[0]) == 1 else "NO",
        'knnPrediction': "YES" if int(knn_prediction[0]) == 1 else "NO",
        'adaboostPrediction': "YES" if int(adaboost_prediction[0]) == 1 else "NO",
    }

    return jsonify(result)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000, debug=True)
