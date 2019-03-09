package players

/*	An implementation of the Player interface that interacts with an
 *	instance of MongoDB
 */
type MongoPlayer struct {
	playerID  int
	firstName string
	lastName  string
	nickname  string
	//totalRecord TotalRecord
}

func NewPlayer() *MongoPlayer {
	mongoPlayer := new(MongoPlayer)
	//mongoPlayer.playerID = generateID()
	return mongoPlayer
}

func (player *MongoPlayer) createPlayer() {

}

func (player *MongoPlayer) getPlayer() {

}

func (player *MongoPlayer) updatePlayer() {

}

func (player *MongoPlayer) deletePlayer() {

}
