package utils

import "log"

func Log_on_error(err error, err_msg string) {
	if err != nil {
		log.Println("Error condition ", err, err_msg)
	}
}
