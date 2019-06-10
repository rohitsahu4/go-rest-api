package thread

//easyjson:json
type ThreadMap map[int]Thread
//easyjson:json
type ThreadSlice [] Thread

//easyjson:json
type Thread struct{
	Id int 
	Discription string
	Title string 
	TimeCreated string
}
