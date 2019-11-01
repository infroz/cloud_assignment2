package commits

/*  Created this struct incase I need to get more data later from
 *  the projects request, for now this will be used to count
 *  commits from each project via /project/id/repository/commits
 */
type project struct {
	ID int `json:"id"`
}
