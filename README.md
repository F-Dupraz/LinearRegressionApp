# Linear Regression App with Go

## Table of Contents

- [General Info](#general-info)
- [First Section](#first-section)
  - [Linear Regression Algorithm](#linear-regression-algorithm)
  - [Read the CSV](#read-the-csv)
- [Second Section](#second-section)
  - [Backend](#backend)
  - [Database](#database)
- [Some Things to Add](#some-things-to-add)

## General Info

This project exists because I felt like smashing ML, backend, and software development into one messy repository. The project is split into two main sections, and guess what, each of those is split again. Brilliant, right?

## First Section

### Linear Regression Algorithm

The logic for this algorithm came straight out of the book "Applied Regression Analysis: A Research Tool, Second Edition" by John O. Rawlings, Sastry G. Pantula, and David A. Dickey. I translated it into Golang code because, well, why the hell not? Everything you need to know is in the first chapter, so just get on with it.

### Read the CSV

The data I'm serving in this application is some stuff I downloaded from Kaggle. You can swap out the CSVs with whatever floats your boat. Reading these CSVs is dead simple—I skimmed the documentation and asked ChatGPT a couple of questions. Easy peasy.

## Second Section

### Backend

The backend is dead simple, too. There's a server document that kicks things off, and three handlers (only two are actually necessary) doing all the logic gymnastics with the data in the 'DB'. Check out the required JSON for the backend dance:

```json
/predict/candy

REQUEST
{
  "chocolate"        uint
  "fruity"           uint
  "caramel"          uint
  "peanutyalmondy"   uint
  "nougat"           uint
  "crispedricewafer" uint
  "hard"             uint
  "bar"              uint
  "pluribus"         uint
  "sugarpercent"     float64
  "pricepercent"     float64
}

RESPONSE (Predict)
{
  "winpercent" float64
}

/predict/heart

REQUEST
{
  "age"      uint
  "sex"      uint
  "cp"       uint
  "trestbps" uint
  "chol"     uint
  "fbs"      bool
  "restecg"  uint
  "thalach"  uint
  "exang"    uint
  "oldpeak"  float64
  "slope"    uint
  "ca"       uint
  "thal"     uint
}

RESPONSE (Predict)
{
  "target" uint
}

/predict/inurance

REQUEST
{
  "age"      uint
  "bmi"      float64
  "children" uint
}

RESPONSE (Predict)
{
  "charges" float64
}

/train/candy

REQUEST
{
  "chocolate"        uint
  "fruity"           uint
  "caramel"          uint
  "peanutyalmondy"   uint
  "nougat"           uint
  "crispedricewafer" uint
  "hard"             uint
  "bar"              uint
  "pluribus"         uint
  "sugarpercent"     float64
  "pricepercent"     float64
  "winpercent"       float64
}

/train/heart

REQUEST
{
  "age"      uint
  "sex"      uint
  "cp"       uint
  "trestbps" uint
  "chol"     uint
  "fbs"      bool
  "restecg"  uint
  "thalach"  uint
  "exang"    uint
  "oldpeak"  float64
  "slope"    uint
  "ca"       uint
  "thal"     uint
  "target"   uint
}

/train/inurance

REQUEST
{
  "age"      uint
  "bmi"      float64
  "children" uint
  "charges"  float64
}
```

### Database

The database is dirt poor—three CSV files with the Intercept taking the lead. It's always '[0]', followed by the coefficients. I did it because I'm all about keeping things simple and efficient.

## Some Things to Add

I reckon there are two main things to tack on here. Firstly, a specialized DB for ML would be nifty. Secondly, and perhaps more importantly, an authentication and authorization protocol. Picture a world where only the cool kids with authenticated passes can ask for predictions, and the authorized folks get to scribble in the 'DB'. I didn't add it, not because I couldn't, but simply because I couldn't be bothered.

## Execution Instructions

Here's a step-by-step guide to running this application:

### Prerequisites

Make sure you have Go installed on your system. You can get it [here](https://golang.org/doc/install).

### Installing Dependencies

1. Clone this repository:

   ```bash
   git clone https://github.com/F-Dupraz/LinearRegressionApp.git
   cd LinearRegressionApp
   ```

2. Install any necessary dependencies (if applicable):

   ```bash
    # Example: Installing Go packages
    github.com/gorilla/mux v1.8.1
	  github.com/joho/godotenv v1.5.1
	  github.com/rs/cors v1.10.1
   ```

4. Run the server:

   ```bash
   go run processHandler.go
   
   go run main.go
   ```

5. Done! The API is now available at [http://localhost:8000](http://localhost:8000).

## LICENSE

This project is under the MIT License. Check out [LICENSE.md](LICENSE.md) for the nitty-gritty details.

## Contributions and Known Issues

We welcome contributions! If you encounter any issues or have enhancements in mind, please open an issue or submit a pull request.

Enjoy using this application and have fun exploring the world of linear regression with Go! 😎🚀