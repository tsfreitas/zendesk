package handlers

func GenerateSQSHandler() SQSHandlerPort {
	return newSQSHandler()
}
