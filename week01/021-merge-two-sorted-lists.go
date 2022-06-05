/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    sentinel := &ListNode{0, nil}
    head := sentinel
    for list1 != nil || list2 != nil {
        if list2 == nil || (list1!= nil && list1.Val <= list2.Val) {
            head.Next = list1
            list1 = list1.Next
        } else {
            head.Next = list2
            list2 = list2.Next
        }
        head = head.Next
    }
    return sentinel.Next
}