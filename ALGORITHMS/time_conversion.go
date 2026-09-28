func timeConversion(s string) string {
    hour, err := strconv.Atoi(s[0:2])
    if err != nil {
        return fmt.Sprintln(err)
    }
    mins := s[2:len(s)-2]
    day := s[len(s) -2:]
    
    if hour == 12 && day == "AM"{
        return "00" + mins
    }else if hour < 12 && day == "PM"{
        hour += 12
        return fmt.Sprintf("%d%s", hour, mins)
    }

    return s[:len(s)-2] 
}
