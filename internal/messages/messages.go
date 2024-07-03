package messages

import (
	"github.com/jmoiron/sqlx"
	"music-app/pkg/system/messages"
)

type Model struct {
	db *sqlx.DB
}

func NewMsgs(db *sqlx.DB) Model {
	return Model{
		db: db,
	}
}

func (m *Model) GetByCode(code int) (int, string, string) {

	//db := dbx.GetConnection()
	repoMsg := messages.FactoryStorage(m.db, nil, "")
	srvMsg := messages.NewMessageService(repoMsg, nil, "")
	msg, _, err := srvMsg.GetMessageByID(code)
	if err != nil {
		return 70, "", ""
	}

	if msg == nil {
		return code, "", ""
	}

	return msg.ID, msg.TypeMessage, msg.Spa

}
