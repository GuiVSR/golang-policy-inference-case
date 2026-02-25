package mocks

const InputSimpleApproved = `{"age": 30, "score": 750}`
const InputSimpleReview = `{"age": 30, "score": 650}`
const InputSimpleRejected = `{"age": 16, "score": 750}`
const InputComplexPrime = `{
  "age": 35,
  "income": 90000,
  "documentation": "verified",
  "credit_score": 720,
  "no_defaults": true,
  "employment_years": 5,
  "employment_status": "permanent",
  "debt_to_income_ratio": 0.25,
  "savings_amount": 25000,
  "emergency_fund": "sufficient",
  "monthly_expenses": 2000,
  "home_ownership": true,
  "education_level": "master",
  "marital_status": "married",
  "dependents": 2
}`
const InputComplexSubprime = `{
	"age": 42,
	"income": 95000,
	"documentation": "verified",
	"credit_score": 680,
	"no_defaults": true,
	"employment_years": 3,
	"employment_status": "permanent",
	"debt_to_income_ratio": 0.43,
	"savings_amount": 25000,
	"emergency_fund": "sufficient",
	"monthly_expenses": 2800,
	"home_ownership": true,
	"education_level": "associate",
	"marital_status": "married",
	"dependents": 1
}`
const InputComplexManualReview = `{
  "age": 65,
  "income": 40000,
  "documentation": "verified",
  "credit_score": 680,
  "no_defaults": false,
  "defaults_age": 30,
  "employment_years": 15,
  "employment_status": "permanent",
  "debt_to_income_ratio": 0.38,
  "savings_amount": 8000,
  "emergency_fund": "partial",
  "monthly_expenses": 1500,
  "home_ownership": true,
  "education_level": "highschool",
  "marital_status": "married",
  "dependents": 0
}`
const InputComplexConditionalApproval = `{
  "age": 28,
  "income": 110000,
  "documentation": "pending",
  "credit_score": 710,
  "no_defaults": true,
  "employment_years": 3,
  "employment_status": "contract",
  "contract_length": 8,
  "debt_to_income_ratio": 0.32,
  "savings_amount": 15000,
  "emergency_fund": "sufficient",
  "monthly_expenses": 1200,
  "home_ownership": false,
  "education_level": "bachelor",
  "marital_status": "single",
  "dependents": 0
}`
const InputComplexConditionalApprovalExperiencedWorker = `{
  "age": 31,
  "income": 55000,
  "documentation": "verified",
  "credit_score": 690,
  "no_defaults": true,
  "employment_years": 4,
  "employment_status": "contract",
  "contract_length": 18,
  "debt_to_income_ratio": 0.3,
  "savings_amount": 12000,
  "emergency_fund": "sufficient",
  "monthly_expenses": 1600,
  "home_ownership": true,
  "education_level": "bachelor",
  "marital_status": "married",
  "dependents": 2
}`
const InputComplexConditionalApprovalHighIncome = `{
  "age": 39,
  "income": 125000,
  "documentation": "verified",
  "credit_score": 700,
  "no_defaults": true,
  "employment_years": 7,
  "employment_status": "permanent",
  "debt_to_income_ratio": 0.28,
  "savings_amount": 55000,
  "emergency_fund": "sufficient",
  "monthly_expenses": 1800,
  "home_ownership": false,
  "education_level": "master",
  "marital_status": "single",
  "dependents": 0
}`
const InputComplexNearPrime = `{
  "age": 45,
  "income": 160000,
  "documentation": "verified",
  "credit_score": 700,
  "no_defaults": true,
  "employment_years": 12,
  "employment_status": "permanent",
  "debt_to_income_ratio": 0.35,
  "savings_amount": 40000,
  "emergency_fund": "partial",
  "monthly_expenses": 4500,
  "home_ownership": true,
  "education_level": "master",
  "marital_status": "married",
  "dependents": 4
}`
const InputComplexRejected = `{
  "age": 25,
  "income": 28000,
  "documentation": "verified",
  "credit_score": 620,
  "no_defaults": false,
  "defaults_age": 12,
  "employment_years": 1,
  "employment_status": "contract",
  "contract_length": 6,
  "debt_to_income_ratio": 0.55,
  "savings_amount": 3000,
  "emergency_fund": "none",
  "monthly_expenses": 2500,
  "home_ownership": false,
  "education_level": "highschool",
  "marital_status": "single",
  "dependents": 0
}`

const InvalidInput1 = `{"age": "thirty", "income": "fifty thousand"}`
const MissingFieldInput = `{"age": 30}`
