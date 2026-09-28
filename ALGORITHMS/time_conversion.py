def timeConversion(s):
    hour = s[0:2]
    mins = s[2:-2] #Python supports negative indexing so this still works.
    day = s[-2:]
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
        return s[-2] 
   
