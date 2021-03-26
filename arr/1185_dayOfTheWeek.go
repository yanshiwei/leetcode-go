func dayOfTheWeek(day int, month int, year int) string {
    //1971年1月1日为星期五 如2021年03月26日
    var Week=[]string{"Friday", "Saturday","Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
    var days int
    //1971-2020所有按年计算的天数
    for y:=1971;y<year;y++{
        if isLeapYear(y){
            days+=366
        }else{
            days+=365
        }
    }
    //2021年1月开始到02月，所有按月计算的天数
    for m:=1;m<month;m++{
        if m==2 {
            //闰月
            if isLeapYear(year){
                days+=29
            }else{
                days+=28
            }
        }else if (m==1||m==3||m==5||m==7||m==8||m==10||m==12){
            days+=31
        }else{
            days+=30
        }
    }
    //2021年3月01日开始到26日，所有按日计算的天数
    days+=day
    return Week[(days-1)%7]
}

func isLeapYear(year int)bool{
    if year%400==0{
        return true
    }
    if year%4==0&&year%100!=0{
        return true
    }
    return false
}
