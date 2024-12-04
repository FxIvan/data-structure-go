package hashmap

type Node struct {
	Key   int
	Value string
	Next  *Node
}

type HashMap struct {
	buckets []*Node
	size    int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func hashFunction(key int, size int) uint {
	return uint(key % size)
}

// Insert method simply gets key and value as parameters.
func (hm *HashMap) Insert(key int, value string) {

	// Here we calculate the index that we will put our key-value pair (node) using the hashFunction.
	index := hashFunction(key, hm.size)

	// Here we create our node with key and value
	node := &Node{Key: key, Value: value}

	/* If there are no key-value pairs in the generated index, we simply put our new node to
	   that index. */
	if hm.buckets[index] == nil {
		hm.buckets[index] = node
	} else {

		/* If there is a node in the generated index, that means we have a collision. No panic,
		   that's why we have Next in our node. We should continuously set a current node to the
		   next node until we have the next node "nil". That means there are no nodes after that
		   current one. */
		current := hm.buckets[index]
		for current.Next != nil {
			current = current.Next
		}
		/* Now we can set the next node of the current node to our newly generated node with
		   peace in mind. */
		current.Next = node
	}
}

func (hm *HashMap) Get(key int) string {

	// Calculate the index with hashFunction to know where we should look.
	index := hashFunction(key, hm.size)

	// We get the first node in this index and assign it to a variable.
	current := hm.buckets[index]

	for current != nil {
		/* We can not simply compare our key with the key of the current node. Because it may
		   not be the node that we are looking for, but what about the next node of it, the next
		   node of the next node of it... Thus, we should continuously set the current node to the
		   next node until the key of the current node matches with our key. Then we return the
		   value of that node. */
		if current.Key == key {
			return current.Value
		}
		current = current.Next
	}

	/* If the code reaches here, it means that there is not a node that matches with our key.
	   We simply return an empty string here. ( In the real world we would handle this situation
	   differently.) */
	return ""
}

func (hm *HashMap) Delete(key int) {
	// Calculate the index of the key using the hash function
	index := hashFunction(key, hm.size)
	// Get the first node at the calculated index
	current := hm.buckets[index]
	// Initialize a pointer to keep track of the previous node
	var prev *Node

	// Traverse the linked list at the calculated index
	for current != nil {
		// If the current node's key matches the key to be deleted
		if current.Key == key {
			// If the previous pointer is nil, it means the node to delete is the first node
			if prev == nil {
				/* Update the head of the linked list to skip the current node. It means we no
				   longer have the current node. */
				hm.buckets[index] = current.Next
			} else {
				/* Skip the current node by updating the previous node's next pointer. When we
				   set the next node of the previous node to the next node of the current node, we
				   simply no longer have the current node. It looks complicated but give it a
				   shoot, you will understand. */
				prev.Next = current.Next
			}
			// Exit the method after deletion
			return
		}
		/* If the key of current node not matches with our key, we move to the next node in the
		   linked list */
		prev = current
		current = current.Next
	}
}
