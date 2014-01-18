package tree_array
import "errors"
type ArrayNode struct {
  parent *ArrayNode
  left *ArrayNode
  right *ArrayNode
  value int
  pos   int
}

type TreeArray struct {
  length int
  root   *ArrayNode
}


func (ta *TreeArray) Put(val int) {
  if ta.length > 0 {
    node := ta.root
    key  := ta.length + 1
    for node != nil {
      if node.pos == key {
        node.value = val
        break
      }
      if node.pos % 2 == 0 {
        if node.left == nil {
          node.left = &ArrayNode{value: val, pos: key}
          break
        } else {
          node = node.left
        }
      } else {
        if node.right == nil {
          node.right = &ArrayNode{value: val, pos: key}
          break
        } else {
          node = node.right
        }
      }
    }
  } else {
    ta.root = &ArrayNode{ value: val, pos: 1 }
  }
  ta.length += 1
}

func (ta TreeArray) Get(n int) (int, error) {
  node := ta.root
  if n > ta.length {
    return 0, errors.New("Out of bounds")
  }
  for n > 1 {
    if node.pos == n {
      break
    }
    if node.pos % 2 == 0 {
      node = node.left
    } else {
      node = node.right
    }
  }
  return node.value, nil
}
