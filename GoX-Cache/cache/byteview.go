package cache

/*
ByteView has only one data member, which stores the actual cached value.
The choice of the byte type is to support the storage of arbitrary data types, such as strings, images, etc.
The Len() method is implemented, as in the implementation of lru.
Cache, where cached objects must implement the Value interface, which includes the Len() method to return the memory size they occupy.
data is read-only, and the ByteSlice() method returns a copy of it to prevent the cached value from being modified by external programs.
*/

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	data []byte
}

// Len returns the view's length.
func (v ByteView) Len() int {
	return len(v.data)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.data)
}

// String returns the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.data)
}

func cloneBytes(data []byte) []byte {
	c := make([]byte, len(data))
	copy(c, data)
	return c
}
