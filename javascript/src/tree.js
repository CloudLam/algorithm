/*!
 * tree.js v1.0.0
 * 2017-2018 Cloud Lam
 * Released under the MIT License.
 */

class Node {
  constructor (value, left, right) {
    this.value = value || null;
    this.left = left || null;
    this.right = right || null;
  }
}

class BinaryTree {
  constructor (data) {
    this.root = null;
    if (data) {
      data.forEach(value => {
        this.insert(value)
      });
    }
  }
  insert (value) {
    let node = new Node(value);
    if (this.root === null) {
      this.root = node;
    } else {
      insertNode(this.root, node);
    }
    function insertNode(node, newNode) {
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
    inOrderTraversal(this.root);
    return result;
    function inOrderTraversal (node) {
      if (node.left) {
        inOrderTraversal(node.left);
      }
      if (node.right) {
        inOrderTraversal(node.right);
      }
      if (node) {
        result.push(node.value);
      }
    }
  }
}
