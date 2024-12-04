Vamos a profundizar en una implementación simple de un HashMap en Go. Nuestra implementación se centrará en la funcionalidad básica, incluyendo la inserción, recuperación y eliminación de pares clave-valor.

En nuestro escenario sólo tenemos nombres por simplicidad. En un HashMap real, tendríamos pares clave-valor. Y como puedes ver, cada nodo en un cubo incluirá una clave, un valor y un puntero al siguiente nodo en el mismo cubo. Con esta implementación tendríamos una lista enlazada en cada índice de la tabla hash.

```
type Node struct {
    Key   string
    Value string
    Next  *Node
}
```

Y aquí está nuestra tabla hash que simplemente incluye una matriz de punteros de nodo y un entero representa el tamaño de la tabla.


```
type HashMap struct {
    buckets []*Node
    size    int
}
```
Now we need a function to create a new HashMap. That’s pretty simple:

```
func NewHashMap(size int) *HashMap {
    return &HashMap{
        buckets: make([]*Node, size),
        size:    size,
    }
}
```

Ahora necesitamos una función hash para calcular el índice de cada par clave-valor que almacenaremos en nuestra tabla hash. Normalmente se debería utilizar un algoritmo hash adecuado para minimizar la colisión. Pero, de nuevo, en aras de la simplicidad nos limitaremos a obtener el resto de la longitud de nuestra clave dividido por el tamaño de nuestro HashMap como hash.

```
func hashFunction(key string, size int) uint {
	return uint(len(key) % size)
}
```

Aquí con el método de inserción de nuestro HashMap entenderemos mejor la estructura de los HashMaps.

```
// Insert method simply gets key and value as parameters.
func (hm *HashMap) Insert(key string, value string) {

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
```

Now, let’s implement a get method:

```
func (hm *HashMap) Get(key string) string {

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
```

Our last method will be the delete.

```
func (hm *HashMap) Delete(key string) {
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
```

Ejemplo de uso de nuestro HashMap

```
func main() {
    // Create a new hashmap with size 10
    myHashMap := NewHashMap(10)

    // Insert key-value pairs
    myHashMap.Insert("john", "doe")
    myHashMap.Insert("foo", "bar")

    // Get and print values
    value := myHashMap.Get("john")
    fmt.Println("Value for key john:", value)

    // Delete a key
    myHashMap.Delete("foo")
    /* If we try to get the value for key "foo" we will get an empty string. (You can return a
    proper error or a flag in your get method) */
}
```

Recapitulemos las características clave de los HashMaps

Almacenamiento clave-valor: Los Hashmaps almacenan datos en pares clave-valor, donde cada clave es única y está asociada a un valor específico.

Función hash: Una función hash asigna claves a índices en la tabla hash, facilitando el acceso rápido a los valores.

Gestión de colisiones: Las colisiones se producen cuando dos claves producen el mismo valor hash. Los HashMaps emplean técnicas de resolución de colisiones como el encadenamiento o el direccionamiento abierto para gestionar las colisiones con elegancia.

Tamaño dinámico: Muchas implementaciones de HashMap cambian dinámicamente de tamaño para mantener un factor de carga equilibrado, garantizando un rendimiento óptimo a medida que cambia el número de elementos.

Operaciones eficientes: Los HashMaps ofrecen un rendimiento medio en tiempo constante para operaciones comunes como la inserción, la eliminación y la búsqueda.