package ransomwatch

import (
	"strings"
)

// Retourne une map de correspondances entre les organisations
// et les posts et groups 'ransomwatch' correspondants.
func MatchOrgsName(exact bool, orgs ...string) (matchs map[string]Match, err error) {

	matchs = make(map[string]Match)

	posts, err := GetPosts()
	if err != nil {
		return
	}

	groups, err := GetGroups()
	if err != nil {
		return
	}

	groupsByName := make(map[string]Group)
	for _, group := range groups {
		groupsByName[group.Name] = group
	}

	for _, org := range orgs {
		for _, post := range posts {
			if exact && post.Title == org {
				matchs[org] = Match{
					Post:  post,
					Group: groupsByName[post.GroupName],
				}
			} else if lowerOrg, lowerTitle := strings.ToLower(org), strings.ToLower(post.Title); !exact &&
				strings.Contains(lowerTitle, lowerOrg) {
				matchs[org] = Match{
					Post:  post,
					Group: groupsByName[post.GroupName],
				}
			}
		}
	}

	return
}
