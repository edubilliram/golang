type UpdateTicketReq struct {
	OrderType                  string `json:"jobType"`
	Summary                    string
	DispatchTechSupervisor     string
	LighthouseWorkOrderID      string
	RemedyLoginID              string
	FirstName                  string `json:"dispatchTechFirstName"`
	LastName                   string `json:"dispatchTechLastName"`
	Region                     string
	Ma                         string
	Hub                        string
	Node                       string
	AssignedGroup              string
	Assignee                   string `json:"rocOwner"`
	IncidentStartTime          string
	status                     string
	StatusReason               string
	ProductCategorizationTier1 string
	ProductCategorizationTier2 string
	ProductCategorizationTier3 string
	Resolution                 string
	RequiredResolutionDateTime string
	ProblemCode                string
	ResolutionCode             string
	IncidentNumber             string
	notes                      string
}


"request": {
                
	"incidentNumber": "INC000000332556",
	"updateNotes": "testing update func",
	"dispatchTechFirstName": "Patrick",
	"dispatchTechLastName": "Nice",
	"dispatchTechSupervisor": "Yama Safi",
	"lighthouseWorkOrderId": "180",
	"incidentStartTime": "2018-02-02T13:24:17-05:00",
	"rocOwner": "Eric Burch"
}

"requests": [
        {
            "routing_key": "scope-service-remedy-devr.update",
            "request": {
                
                "incidentNumber": "INC000000332556",
                "updateNotes": "testing update func",
                "dispatchTechFirstName": "Patrick",
                "dispatchTechLastName": "Nice",
                "dispatchTechSupervisor": "Yama Safi",
                "lighthouseWorkOrderId": "180",
                "incidentStartTime": "2018-02-02T13:24:17-05:00",
                "status" : "Resolved",
                "statusReason" : "No Further Action Required",
                "productCategorizationTier2" : "Distribution Passive",
                "productCategorizationTier3" : "Inline EQ",
                "resolution" : "TEST",
                "RequiredResolutionDateTime" : "2018-03-16T14:59:35-06:00",
                "problemCode" : "Third Party Damage",
                "resolutionCode" : "Resolved",
                "selfInflicted" : "No",
                "rocOwner": "Eric Burch"
            }
        }
	]
	
	