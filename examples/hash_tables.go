package examples

//A hash function is something that returns the index of an entry in an array upon creation, for efficiency.
// A 'seperate-chained' hash table is similar to an array combined with a linked list. in case of entry collision, the next entry that tries to go
// in a populated slot, gets put into a linked list (Bucket).
// Time complexity of a hash table is constant 0(1).
// Downside: If all keys match, and get inserted into same bucket, time becomes 0(n) (Bad!)

const ArraySize = 7

//Hashtable structure
type HashTable struct {
	array [ArraySize]*bucket
}

//Bucket structure
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

//Hash function
