package outbound

type RickMortyInter interface {
	GetCharactersCount() (uint64, error)
	GetPageCount() (uint64, error)
	GetCharactersOfPages(uint) ([]CharacterInter, error)
}

type CharacterInter interface {
	GetID() uint64
	GetName() string
	GetAvatar() string
}

