package outbound

type RickMortyInter interface {
	GetCharactersCount() (uint64, error)
	GetPageCount() (uint64, error)
	GetCharactersOfPages(uint) ([]CharacterInputInter, error)
}

type CharacterInputInter interface {
	GetID() uint64
	CharacterOutputInter
}

type CharacterOutputInter interface {
	GetName() string
	GetAvatar() string
}
