class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return build(preorder, inorder, 0, 0, inorder.length - 1);
    }

    private TreeNode build(int[] preorder, int[] inorder, int preIndex, int inStart, int inEnd) {
        if (inStart > inEnd || preIndex >= preorder.length) {
            return null;
        }
        TreeNode root = new TreeNode(preorder[preIndex]);
        int inIndex = 0;
        for (int i = inStart; i <= inEnd; i++) {
            if (root.val == inorder[i]) {
                inIndex = i;
                break;
            }
        }
        root.left = build(preorder, inorder, preIndex + 1, inStart, inIndex - 1);
        root.right = build(preorder, inorder, preIndex + inIndex - inStart + 1, inIndex + 1, inEnd);
        return root;
    }
}
