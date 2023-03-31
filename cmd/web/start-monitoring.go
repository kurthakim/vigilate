package main

import "log"

type job struct {
	HostServiceID int
}

func (j job) Run() {
	repo.ScheduledCheck(j.HostServiceID)
}

func startMonitoring() {
	if preferenceMap["monitoring_live"] == "1" {
		log.Println("************************ starting monitor")

		// trigger a message to broadcast to all clients that app is starting to monitor
		data := make(map[string]string)
		data["message"] = "Monitoring is starting..."
		err := app.WsClient.Trigger("public-channel", "app-starting", data)
		if err != nil {
			log.Println(err)
		}

		// get all of the services that we want to monitor
		servicesToMonitor, err := repo.DB.GetServicesToMonitor()
		if err != nil {
			log.Println(err)
		}

		log.Println("Length of servicesToMonitor is", len(servicesToMonitor))

		for _, x := range servicesToMonitor {
			log.Println("******Service to monitor on", x.HostName, "is", x.Service.ServiceName)
		}

		// range through the services

		// get the schedule unit and number

		// create a job

		// save the id of the job, so we can start/stop it

		// broadcast over websocket the fact that the service is scheduled

		// end range

	}
}
