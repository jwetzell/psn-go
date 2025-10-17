package chunks

type ChunkHeader struct {
	Id           uint16
	DataLen      uint16
	HasSubchunks bool
}

type Chunk struct {
	ChunkData []byte
	Header    ChunkHeader
}
