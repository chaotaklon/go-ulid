package ulid

import (
	"math/rand"
	"time"
	"sync"
)


const (
	ENCODING = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
)

var (
	mux sync.Mutex
)


func Init() {
	rand.Seed(time.Now().UnixNano())
}


func NewCommon(threadSafe bool) []byte {
	
	unixTimeStampMs := time.Now().UnixNano() / int64(time.Millisecond)  // int64 type unix timestamp in milliseconds
	timeStampPart := make([]byte, 6)
	for i := 5; i >= 0; i-- {  // convert int64 timestamp to [6]byte in big endian way
		timeStampPart[i] = byte(unixTimeStampMs & 0x000000FF)
		unixTimeStampMs >>= 8
	}
	
	randPart := make([]byte, 10)
	if threadSafe {
		mux.Lock()
	}
	rand.Read(randPart)  // math/rand library is not thread safe
	if threadSafe {
		mux.Unlock()
	}
	

	ulidBase32 := make([]byte, 26)  // Crockford's Base32 https://www.crockford.com/base32.html
	
	// 10 byte timestamp
	ulidBase32[0] = ENCODING[(timeStampPart[0]&224)>>5]
	ulidBase32[1] = ENCODING[timeStampPart[0]&31]
	ulidBase32[2] = ENCODING[(timeStampPart[1]&248)>>3]
	ulidBase32[3] = ENCODING[((timeStampPart[1]&7)<<2)|((timeStampPart[2]&192)>>6)]
	ulidBase32[4] = ENCODING[(timeStampPart[2]&62)>>1]
	ulidBase32[5] = ENCODING[((timeStampPart[2]&1)<<4)|((timeStampPart[3]&240)>>4)]
	ulidBase32[6] = ENCODING[((timeStampPart[3]&15)<<1)|((timeStampPart[4]&128)>>7)]
	ulidBase32[7] = ENCODING[(timeStampPart[4]&124)>>2]
	ulidBase32[8] = ENCODING[((timeStampPart[4]&3)<<3)|((timeStampPart[5]&224)>>5)]
	ulidBase32[9] = ENCODING[timeStampPart[5]&31]

	// 16 bytes of entropy
	ulidBase32[10] = ENCODING[(randPart[0]&248)>>3]
	ulidBase32[11] = ENCODING[((randPart[0]&7)<<2)|((randPart[1]&192)>>6)]
	ulidBase32[12] = ENCODING[(randPart[1]&62)>>1]
	ulidBase32[13] = ENCODING[((randPart[1]&1)<<4)|((randPart[2]&240)>>4)]
	ulidBase32[14] = ENCODING[((randPart[2]&15)<<1)|((randPart[3]&128)>>7)]
	ulidBase32[15] = ENCODING[(randPart[3]&124)>>2]
	ulidBase32[16] = ENCODING[((randPart[3]&3)<<3)|((randPart[4]&224)>>5)]
	ulidBase32[17] = ENCODING[randPart[4]&31]
	ulidBase32[18] = ENCODING[(randPart[5]&248)>>3]
	ulidBase32[19] = ENCODING[((randPart[5]&7)<<2)|((randPart[6]&192)>>6)]
	ulidBase32[20] = ENCODING[(randPart[6]&62)>>1]
	ulidBase32[21] = ENCODING[((randPart[6]&1)<<4)|((randPart[7]&240)>>4)]
	ulidBase32[22] = ENCODING[((randPart[7]&15)<<1)|((randPart[8]&128)>>7)]
	ulidBase32[23] = ENCODING[(randPart[8]&124)>>2]
	ulidBase32[24] = ENCODING[((randPart[8]&3)<<3)|((randPart[9]&224)>>5)]
	ulidBase32[25] = ENCODING[randPart[9]&31]

	return ulidBase32
}


// return string type, thread safe
func NewString() string {
	return string(NewCommon(true))
}

// return byte array type, thread safe
func NewByteArray() []byte {
	return NewCommon(true)
}

// return string type, no mutex lock during math/rand, for single thread application requiring highest performance
func NewStringThreadUnsafe() string {
	return string(NewCommon(false))
}

// return byte array type, no mutex lock during math/rand, for single thread application requiring highest performance
func NewByteArrayThreadUnsafe() []byte {
	return NewCommon(false)
}


// the most common use case: thread safe, return string
func New() string {
	return NewString()
}
