package server

type Child struct {
	Name    string
	Balance float32
}

type PocketMoneyCommand interface {
	commandType()
}

type GetChildPocketMoneyCommand struct {
	Name string
	Resp chan Child
}
type GetKidsPocketMoneyCommand struct {
	Resp chan []Child
}
type DepositPocketMoneyCommand struct{}
type WithdrawPocketMoneyCommand struct{}

func (GetChildPocketMoneyCommand) commandType() {}
func (GetKidsPocketMoneyCommand) commandType()  {}
func (DepositPocketMoneyCommand) commandType()  {}
func (WithdrawPocketMoneyCommand) commandType() {}

func getKids() []Child {
	kids := []Child{
		{Name: "Elizabeth", Balance: 10.0},
		{Name: "Matilda", Balance: 4.0},
		{Name: "Joseph", Balance: 4.0},
	}
	return kids
}

func PocketMoneyManager(commandChan chan PocketMoneyCommand, store ChildStore) {
	log.Info("pocketMoneyManager started")

	for cmd := range commandChan {
		switch v := cmd.(type) {
		case GetKidsPocketMoneyCommand:
			v.Resp <- store.GetAllChildren()

		case GetChildPocketMoneyCommand:
			v.Resp <- store.GetChild(v.Name)

		case DepositPocketMoneyCommand:

		case WithdrawPocketMoneyCommand:

		default:
			log.Info("unknown command", "command", v)

		}

	}

}
