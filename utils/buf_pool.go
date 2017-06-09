package utils

var bufferSize = 1024 * 33
var freeList = make(chan []byte, 100)

type BufferPool struct {
	freeList chan []byte
}

func (pool *BufferPool) Get() (buffer []byte) {
	select {
	case buffer = <-pool.freeList:
	default:
		buffer = make([]byte, bufferSize)
	}
	return
}

func (pool *BufferPool) Put(buffer []byte) {
	select {
	case pool.freeList <- buffer:
	default:
	}
}

var BufPool = &BufferPool{freeList: make(chan []byte, 100)}