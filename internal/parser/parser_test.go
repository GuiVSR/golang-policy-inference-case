package parser

import (
	"fmt"
	"lab/internal/models"
	"testing"
)

func TestParsePolicy_Valid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		check func(t *testing.T, g *models.Graph)
	}{
		{
			name:  "ExampleDot",
			input: `digraph Policy { start [result=""]; approved [ result="approved=true,segment=prime" ]; rejected [ result="approved=false" ]; review [ result="approved=false,segment=manual" ]; start -> approved [cond="age>=18 && score>700"]; start -> review [cond="age>=18 && score<=700"]; start -> rejected [cond="age<18"]; }`,
			check: func(t *testing.T, g *models.Graph) {
				output := g.Edges
				fmt.Println(output)
			},
		},
		{
			name:  "ExampleDot2",
			input: `digraph Policy { start [result=""]; age_check [result="age_checked=true"]; start -> age_check [cond="age!=null"]; adult [result="age_group=adult"]; minor [result="age_group=minor"]; senior [result="age_group=senior"]; age_check -> adult [cond="age>=18 && age<65"]; age_check -> minor [cond="age<18"]; age_check -> senior [cond="age>=65"]; credit_check [result="credit_checked=true"]; adult -> credit_check [cond="true"]; prime [result="credit_tier=prime,score_factor=1.2"]; near_prime [result="credit_tier=near_prime,score_factor=1.0"]; subprime [result="credit_tier=subprime,score_factor=0.8"]; credit_check -> prime [cond="score>700"]; credit_check -> near_prime [cond="score>600 && score<=700"]; credit_check -> subprime [cond="score<=600"]; income_check [result="income_checked=true"]; adult -> income_check [cond="true"]; senior -> income_check [cond="true"]; high_income [result="income_level=high,debt_ratio_max=0.5"]; medium_income [result="income_level=medium,debt_ratio_max=0.4"]; low_income [result="income_level=low,debt_ratio_max=0.3"]; income_check -> high_income [cond="income>100000"]; income_check -> medium_income [cond="income>50000 && income<=100000"]; income_check -> low_income [cond="income<=50000"]; employment_check [result="employment_checked=true"]; adult -> employment_check [cond="true"]; employed [result="employment_status=employed,job_stability=high"]; self_employed [result="employment_status=self_employed,job_stability=medium"]; retired [result="employment_status=retired,job_stability=high"]; employment_check -> employed [cond="employment_type=='full_time' || employment_type=='part_time'"]; employment_check -> self_employed [cond="employment_type=='self_employed'"]; senior -> retired [cond="true"]; dti_check [result="dti_checked=true"]; high_income -> dti_check [cond="true"]; medium_income -> dti_check [cond="true"]; low_income -> dti_check [cond="true"]; dti_good [result="dti_status=good,risk_level=low"]; dti_warning [result="dti_status=warning,risk_level=medium"]; dti_bad [result="dti_status=bad,risk_level=high"]; dti_check -> dti_good [cond="existing_debt/income <= 0.3"]; dti_check -> dti_warning [cond="existing_debt/income > 0.3 && existing_debt/income <= 0.5"]; dti_check -> dti_bad [cond="existing_debt/income > 0.5"]; guardian_check [result="guardian_checked=true"]; minor -> guardian_check [cond="true"]; guardian_approved [result="minor_status=approved_with_guardian,approved=false"]; guardian_denied [result="minor_status=denied_no_guardian,approved=false"]; guardian_check -> guardian_approved [cond="has_guardian==true && guardian_consent==true"]; guardian_check -> guardian_denied [cond="has_guardian==false || guardian_consent==false"]; final_approved [result="approved=true,final_segment=prime,final_limit=high"]; final_approved_medium [result="approved=true,final_segment=standard,final_limit=medium"]; final_approved_low [result="approved=true,final_segment=basic,final_limit=low"]; final_review [result="approved=false,final_segment=manual_review"]; final_denied [result="approved=false,final_segment=denied"]; prime -> final_approved [cond="dti_status=='good' && employment_status!=''"]; prime -> final_approved_medium [cond="dti_status=='warning'"]; prime -> final_review [cond="dti_status=='bad'"]; near_prime -> final_approved_medium [cond="dti_status=='good' || dti_status=='warning'"]; near_prime -> final_review [cond="dti_status=='bad'"]; subprime -> final_approved_low [cond="dti_status=='good'"]; subprime -> final_review [cond="dti_status=='warning'"]; subprime -> final_denied [cond="dti_status=='bad'"]; senior -> final_approved_medium [cond="dti_status!='bad'"]; senior -> final_review [cond="dti_status=='bad'"]; guardian_approved -> final_review; guardian_denied -> final_denied; }`,
			check: func(t *testing.T, g *models.Graph) {
				output := g.Edges
				fmt.Println(output)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g, err := ParsePolicy(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if g == nil {
				t.Fatalf("expected graph, got nil")
			}
			tc.check(t, g)
		})
	}
}

func TestParsePolicy_Invalid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "IncompleteDot",
			input: "digraph { start -> ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParsePolicy(tc.input)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
