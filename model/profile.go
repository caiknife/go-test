package model

type Profile struct {
	Team   string `json:"team" xml:"team"`
	Member string `json:"member" xml:"member"`
}

type Person struct {
	Profile
}
