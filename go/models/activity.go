package models

type Activity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

var Activities = []Activity{
	{ID: 1, Title: "Yoga Class", Type: "Fitness"},
	{ID: 2, Title: "Cooking Workshop", Type: "Lifestyle"},
	{ID: 3, Title: "Painting Season", Type: "Art"},
}

func FindActivityByID(id int) (Activity, bool) {
	for _, activity := range Activities {
		if activity.ID == id {
			return activity, true
		}
	}
	return Activity{}, false
}
