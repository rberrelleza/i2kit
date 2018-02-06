package cf

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	log "github.com/sirupsen/logrus"
)

//Create creates a AWS Cloud Formation stack
func Create(name, template string, config *aws.Config) (string, error) {
	log.Infof("Creating stack '%s'...", name)
	stack := &cloudformation.CreateStackInput{
		Capabilities: []*string{aws.String("CAPABILITY_IAM")},
		StackName:    aws.String(name),
		TemplateBody: aws.String(template),
		Tags: []*cloudformation.Tag{
			&cloudformation.Tag{
				Key:   aws.String("i2kit"),
				Value: aws.String("alpha"),
			},
		},
	}
	svc := cloudformation.New(session.New(), config)
	response, err := svc.CreateStack(stack)
	if err != nil {
		return "", err
	}
	return *response.StackId, nil
}