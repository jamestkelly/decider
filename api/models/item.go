package models

import "time"

/*
Item
Interface to represent Items.
*/
type Item struct {
	ID        string    `json:"id"`
	Type 	  string    `json:"type"`
	User      User      `json:"user"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Media     []string  `json:"media"`
	Comments  []Comment `json:"comments"`
	Votes     []Vote 	`json:"votes"`
}

func (i *Item) addVote(vote Vote) {
	i.Votes = append(i.Votes, vote)
}

func (i *Item) editVote(voteID string, newVote Vote) {
	for index, vote := range i.Votes {
		if vote.ID == voteID {
			i.Votes[index] = newVote
			break
		}
	}
}

func (i *Item) removeVote(voteID string) {
	for index, vote := range i.Votes {
		if vote.ID == voteID {
			i.Votes = append(i.Votes[:index], i.Votes[index+1:]...)
			break
		}
	}
}

func (i *Item) calculateVotes() int {
	return len(i.Votes)
}

func (i *Item) addComment(comment Comment) {
	i.Comments = append(i.Comments, comment)
}

func (i *Item) editComment(commentID string, newComment Comment) {
	for index, comment := range i.Comments {
		if comment.ID == commentID {
			i.Comments[index] = newComment
			break
		}
	}
}

func (i *Item) removeComment(commentID string) {
	for index, comment := range i.Comments {
		if comment.ID == commentID {
			i.Comments = append(i.Comments[:index], i.Comments[index+1:]...)
			break
		}
	}
}

/*
Comment
Interface to represent Comments of Items.
*/
type Comment struct {
	ID        string    `json:"id"`
	UserID 	  string    `json:"userID"`
	Timestamp time.Time `json:"timestamp"`
	Votes     []Vote    `json:"Votes"`
}

/*
Vote
Interface to represent Votes of Comments or Items.
*/
type Vote struct {
	ID        string      `json:"id"`
	Vote 	  bool        `json:"vote"`
	Timestamp time.Time   `json:"timestamp"`
}