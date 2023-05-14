package models

type Server struct {
	NumUsers   int        `json:"numUsers"`
	Name       string     `json:"name"`
	Channels   []Channel  `json:"channels"`
	Owner      string     `json:"-"`
	AdminIDs   []string   `json:"-"`
	UserIDs    []string   `json:"-"`
}

func (s *Server) addUser(userID string) {
	s.UserIDs = append(s.UserIDs, userID)
	s.NumUsers++
}

func (s *Server) removeUser(userID string) {
	for i, id := range s.UserIDs {
		if id == userID {
			s.UserIDs = append(s.UserIDs[:i], s.UserIDs[i+1:]...)
			s.NumUsers--
			break
		}
	}
}

func (s *Server) addAdmin(adminID string) {
	s.AdminIDs = append(s.AdminIDs, adminID)
}

func (s *Server) removeAdmin(adminID string) {
	for i, id := range s.AdminIDs {
		if id == adminID {
			s.AdminIDs = append(s.AdminIDs[:i], s.AdminIDs[i+1:]...)
			break
		}
	}
}

func (s *Server) changeOwner(newOwner string) {
	s.Owner = newOwner
}

func (s *Server) addChannel(channel Channel) {
	s.Channels = append(s.Channels, channel)
}

func (s *Server) removeChannel(channelID string) {
	for i, channel := range s.Channels {
		if channel.ID == channelID {
			s.Channels = append(s.Channels[:i], s.Channels[i+1:]...)
			break
		}
	}
}

func (s *Server) editName(newName string) {
	s.Name = newName
}

type Channel struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	ItemIDs []string `json:"-"`
}

func (c *Channel) editChannel(newName, newType string) {
	c.Name = newName
	c.Type = newType
}

func (c *Channel) addItemByID(itemID string) {
	c.ItemIDs = append(c.ItemIDs, itemID)
}

func (c *Channel) editItemByID(itemID, newItemID string) {
	for i, id := range c.ItemIDs {
		if id == itemID {
			c.ItemIDs[i] = newItemID
			break
		}
	}
}

func (c *Channel) deleteItemByID(itemID string) {
	for i, id := range c.ItemIDs {
		if id == itemID {
			c.ItemIDs = append(c.ItemIDs[:i], c.ItemIDs[i+1:]...)
			break
		}
	}
}
