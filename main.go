package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type Case struct {
	MemberStatus  string
	Reason        string
	Problem       string
	Temperature   string
	OtherSymptoms string
	ValidID       string
	SupportLevel  string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Create a sample case
	myCase := &Case{}

	// Collect user input sequentially
	fmt.Print("Are you a member? (yes/no): ")
	memberStatus, _ := reader.ReadString('\n')
	myCase.MemberStatus = strings.TrimSpace(memberStatus)

	fmt.Print("Do you have a valid ID? (yes/no): ")
	validID, _ := reader.ReadString('\n')
	myCase.ValidID = strings.TrimSpace(validID)

	fmt.Print("What is the reason for your visit? (new_case/follow_up_case/information_other): ")
	reason, _ := reader.ReadString('\n')
	myCase.Reason = strings.TrimSpace(reason)

	fmt.Print("What is your temperature? (normal/abnormal/unknown): ")
	temperature, _ := reader.ReadString('\n')
	myCase.Temperature = strings.TrimSpace(temperature)

	fmt.Print("Do you have other symptoms? (yes/no): ")
	otherSymptoms, _ := reader.ReadString('\n')
	myCase.OtherSymptoms = strings.TrimSpace(otherSymptoms)

	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MF", myCase)
	if err != nil {
		panic(err)
	}

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	fileRes := pkg.NewFileResource("rules.grl")
	err = ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
	if err != nil {
		panic(err)
	}

	knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
	if err != nil {
		panic(err)
	}

	engine := engine.NewGruleEngine()

	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}

	fmt.Println("Support Level:", myCase.SupportLevel)
}
