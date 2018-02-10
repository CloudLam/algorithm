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
  compare (a, b) {
    return a - b;
  }
  split (order) {
    let temp = this.keys.splice(order / 2, 1);

    let left = new BNode();
    left.keys = this.keys.splice(0, order / 2);

    let right = new BNode();
    right.keys = this.keys;

    if (this.children.length > 0) {
      left.children = this.children.splice(0, order / 2 + 1);
      left.children.forEach(element => {
        element.parent = left;
      });
      right.children = this.children;
      right.children.forEach(element => {
        element.parent = right;
      });
    }

    if (this.parent) {
      if (temp[0] < this.parent.keys[0]) {
        this.parent.keys.splice(0, 0, temp[0]);
        this.parent.children.splice(0, 1, right);
        this.parent.children.splice(0, 0, left);
      } else {
        this.parent.keys.push(temp[0]);
        this.parent.children.pop();
        this.parent.children.push(left);
        this.parent.children.push(right);
      }
      left.parent = this.parent;
      right.parent = this.parent;
      if (this.parent.keys.length >= order) {
        this.parent.split(order);
      }
    } else {
      this.keys = temp;
      this.children = [];
      this.children.push(left);
      this.children.push(right);
      left.parent = this;
      right.parent = this;
    }
  }
  insert (value, order) {
    if (this.children.length == 0) {
      this.keys.push(value);
      this.keys.sort(this.compare);
      if (this.keys.length < order) {
        return;
      } else {
        this.split(order);
      }
    } else {
      for (let i = 0; i < this.keys.length; i++) {
        if (value < this.keys[i]) {
          this.children[i].insert(value, order);
          return;
        }
      }
      this.children[this.keys.length].insert(value, order);
    }
  }
  remove (value, order) {
    for (let i = 0; i < this.keys.length; i++) {
      if (value == this.keys[i]) {
        return;
      }
      if (value < this.keys[i]) {
        this.children[i].remove(value, order);
        return;
      }
    }
    this.children[this.keys.length].remove(value, order);
  }
}

class BTree {
  constructor (order) {
    this.order = order > 2 ? order : 2;
    this.root = new BNode();
  }
  search (value) {}
  insert (value) {
    this.root.insert(value, this.order);
  }
  remove (value) {
    this.root.remove(value, this.order);
  }
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
    this.order = order > 2 ? order : 2;
    this.root = new BPlusNode();
  }
  search (value) {}
  insert (value) {}
  remove (value) {}
}
