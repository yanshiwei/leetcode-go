func countTriplets(arr []int) int {
    /*
    //因为a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1] b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k] ；  a == b 那么 a ^ b == 0 即 arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1] ^ arr[j] ^ arr[j + 1] ^ ... ^ arr[k] == 0 在这种情况下 j 可以为(i , k] (不包含i，包含k)的任意值， count = count + （j - i）
    */
    var last int
    var count int
    for i:=0;i<len(arr)-1;i++{
        //j must be bigger than i, i取值[0,len)
        last=arr[i]
        for j:=i+1;j<len(arr);j++{
            //j can be k，k取值(0,len)
            last=last^arr[j]
            if last==0{
                //fmt.Println(j,i)
                count+=(j-i)
            }
        }
    }
    return count
}
