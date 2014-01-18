package tree_array

import "testing"

func Test_treearray(t *testing.T) {
  tree_array := TreeArray{}
  t.Log("Testing")
  if (tree_array.length != 0) {
    t.Error("New array should have length 0")
  }
}

func Test_put_get(t *testing.T) {
  tree_array := TreeArray{}
  tree_array.Put(100)
  t.Log(tree_array)
  if tree_array.length != 1 {
    t.Error("Array should have length 2 after 1 put")
  }
  v, e := tree_array.Get(1)
  if v != 100 {
    t.Error("Get 1 should return 100")
  }
  tree_array.Put(101)

  if tree_array.length != 2 {
    t.Error("Array should have length 2 after 2 put")
  }
  v, e = tree_array.Get(2)
  if v != 101 {
    t.Error("Get 2 should return 101")
  }
  tree_array.Put(102)
  tree_array.Put(103)
  tree_array.Put(104)

  v, e = tree_array.Get(5)
  if v != 104 {
    t.Error("Get 5 should return 104")
  }

  v, e = tree_array.Get(100)
  if e == nil {
    t.Log(e)
    t.Error("Should return error")
  }
}
