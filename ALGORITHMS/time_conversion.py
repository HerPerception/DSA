def timeConversion(s):
    hour = s[0:2]
    mins = s[2:len(s)-2]
    day = s[len(s) -2:]
    try:
        hour = int(hour)
    except ValueError:
        return s
    if hour == 12 and day == 'AM':
        return '00'+ mins
    elif hour < 12 and day == 'PM':
        hour += 12
        return str(hour) + mins
    else:
        return s[:len(s)-2] 
   
