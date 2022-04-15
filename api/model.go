package main

type SiteStatistics struct {
	Visits int `json:"Visits"`
}

// TODO: figure out the pattern to return a refreshed copy of itself

func (s *SiteStatistics) getVisits() (int, error) {
	var err error

	// TODO: read the visit count from the database here

	return s.Visits, err
}

func (s *SiteStatistics) addVisit() error {
	var err error

	// TODO: update the count in the database here
	s.Visits++

	return err
}
