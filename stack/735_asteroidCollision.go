func asteroidCollision(asteroids []int) []int {
    var stack []int
    if len(asteroids)<1 {
        return stack
    }
    //1 + + push
    //2 - - push
    //3 - + push
    //4 + - collision
    for i:=range asteroids {
        var isMinusBig bool
        if len(stack)<1 {
            //avoid middle stack is empty like 1 -1 -2 -2
            stack=append(stack,asteroids[i])
        }else {
            head:=stack[len(stack)-1]
            if head>0&&asteroids[i]<0 {
                //'left + and right - 'must collision
                for head>0&&asteroids[i]<0 {
                    if head>abs(asteroids[i]) {
                        isMinusBig=false//reset no push
                        break
                    }else {
                        if len(stack)>0 {
                            stack=stack[0:len(stack)-1]//pop
                            if head==abs(asteroids[i]) {
                                isMinusBig=false//reset no push
                                break
                            }
                            isMinusBig=true
                            if len(stack)==0 {
                                isMinusBig=false//reset no push
                                stack=append(stack,asteroids[i])
                            }
                            head=stack[len(stack)-1]
                        }
                    }
                }
                if isMinusBig {
                    stack=append(stack,asteroids[i])
                }
            }else {
                //same direction or left - and right +
                stack=append(stack,asteroids[i])
            }
        }
    }
    return stack
}

func abs(a int)int {
    if a<0 {
        return -a
    }else {
        return a
    }
}
