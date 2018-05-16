package rmail

// AnswerQuestion : send mail when added answer for the question
func (client *Client) AnswerQuestion(answerID int) {
	query := `
		mutation sendAnswerQuestionMail($answerID: Int!) {
			send_mail {
				customers_questions {
					answer_question(answer_id: $answerID)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"answerID": answerID,
	}
	client.sendMail(query, vars)
}