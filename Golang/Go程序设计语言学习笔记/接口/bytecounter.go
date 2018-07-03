package bytecounter



type ByteCounter int

func (c *ByteCounter) Write(p []byte)(int,error){
	*c += ByteCouter(len(p))
	return len(p),nil
}