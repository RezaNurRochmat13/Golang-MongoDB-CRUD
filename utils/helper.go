package utils

import "log"

func GlobalErrorException(message error) bool {
	if message != nil {
		log.Printf("Caused error exceptions : %s", message.Error())
	}

	return message == nil
}

func GlobalErrorDatabaseException(message error) bool {
	if message != nil {
		log.Printf("Caused error database exceptions : %s", message.Error())
	}

	return message == nil
}
