package filecontent

//LoadContent is a var
var LoadContent = []byte(
	`
	package auto
	
	import (
		"automatepi/database"
		"automatepi/models"
		//"automatepi/utils/console"
		"log"
	)
	//Load is
	func Load()  {
		db, err := database.Connect()
		if err != nil {
			log.Fatal("this is an error :", err)
		}
		defer db.Close()
	
		// err = db.Debug().DropTableIfExists(&models.User{}).Error
		// if err != nil {
		// 	log.Fatal("this is an error :", err)
		// }
	
		err = db.Debug().AutoMigrate(&models.User{}).Error
		if err != nil {
			 log.Fatal("error occured : ", err )
		 }
	
	
	}`)