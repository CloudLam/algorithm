/*!
 * tree.js v1.0.0
 * 2017-2018 Cloud Lam
 * Released under the MIT License.
 */

/*
 * Binary-Search-Tree
 */
class BinarySearchNode {
  constructor (value, left, right) {
    this.value = value || null;
    this.left = left || null;
    this.right = right || null;
  }
}

class BinarySearchTree {
  constructor (data) {
    this.root = null;
    if (data) {
      data.forEach(value => {
        this.insert(value)
      });
    }
  }
  search (value) {
    let current = this.root;
    while (current != null) {
      if (current.value == value) {
        return current;
      }
      if (current.value > value) {
        current = current.left;
      } else {
        current = current.right;
      }
    }
    return null;
  }
  getMax (node = this.root) {
    if (node.right == null) {
      return node;
    }
    return this.getMax(node.right);
  }
  getMin (node = this.root) {
    if (node.left == null) {
      return node;
    }
    return this.getMin(node.left);
  }
  insert (value) {
    let node = new BinarySearchNode(value);
    if (this.root === null) {
      this.root = node;
    } else {
      insertNode(this.root, node);
    }
    function insertNode (node, newNode) {
      if (newNode.value < node.value) {
        if (node.left === null) {
          node.left = newNode;
        } else {
          insertNode(node.left, newNode);
        }
      } else {
        if (node.right === null) {
          node.right = newNode;
        } else {
          insertNode(node.right, newNode);
        }
      }
    }
  }
  remove (value) {
    let _this = this;
    this.root = removeNode(this.root, value);
    function removeNode (node, value) {
      if (node == null) {
        return null;
      }
      if (node.value == value) {
        if (node.left == null && node.right == null) {
          return null;
        }
        if (node.left == null) {
          return node.right;
        }
        if (node.right == null) {
          return node.left;
        }
        var temp = _this.getMin(node.right);
        node.value = temp.value;
        node.right = removeNode(node.right, temp.value);
        return node;
      }
      if (node.value > value) {
        node.left = removeNode(node.left, value);
        return node;
      }
      if (node.value < value) {
        node.right = removeNode(node.right, value);
        return node;
      }
    }
  }
  preOrder () {
    let result = [];
    preOrderTraversal(this.root);
    return result;
    function preOrderTraversal (node) {
      if (node) {
        result.push(node.value);
      }
      if (node.left) {
        preOrderTraversal(node.left);
      }
      if (node.right) {
        preOrderTraversal(node.right);
      }
    }
  }
  inOrder () {
    let result = [];
    inOrderTraversal(this.root);
    return result;
    function inOrderTraversal (node) {
      if (node.left) {
        inOrderTraversal(node.left);
      }
      if (node) {
        result.push(node.value);
      }
      if (node.right) {
        inOrderTraversal(node.right);
      }
    }
  }
  postOrder () {
    let result = [];
    postOrderTraversal(this.root);
    return result;
    function postOrderTraversal (node) {
      if (node.left) {
        postOrderTraversal(node.left);
      }
      if (node.right) {
        postOrderTraversal(node.right);
      }
      if (node) {
        result.push(node.value);
      }
    }
  }
}

/*
 * B-Tree
 */
class BNode {
  constructor () {
    this.keys = [];
    this.parent = null;
    this.children = [];
  }
}

class BTree {
  constructor (order) {
    this.root = null;
    this.order = order > 2 ? order : 2;
  }
  search (value) {}
  insert (value) {}
  remove (value) {}
}

/*
 * B-Plus-Tree
 */
class BPlusNode extends BNode {
  constructor () {
    super ();
    this.isLeaf = false;
    this.next = null;
  }
}

class BPlusTree {
  constructor (order) {
    this.root = null;
    this.order = order > 2 ? order : 2;
  }
  search (value) {}
  insert (value) {}
  remove (value) {}
}
