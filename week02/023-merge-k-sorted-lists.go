/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
        return nil
    }
    return merge_sort(lists,0,len(lists)-1)
}

func merge_sort(lists []*ListNode,low,high int) *ListNode{
    if low == high{
        return lists[low]
    }
    if low > high{
        return nil   
    }
    mid := low + (high-low)/2
    return merge(merge_sort(lists,low,mid),merge_sort(lists,mid+1,high))
}

func merge(list1,list2 *ListNode) *ListNode{
    if list1 == nil{
        return list2
    }
    if list2 == nil{
        return list1
    }
    head := &ListNode{}
    p := head
    for list1 != nil&&list2 != nil{
        if list1.Val < list2.Val{
            p.Next =list1
            list1 = list1.Next
        }else{
            p.Next = list2
            list2 = list2.Next
        }
        p = p.Next
    }
    if list1 != nil{
        p.Next = list1
    }
    if list2 != nil{
        p.Next = list2
    }
    return head.Next
}
