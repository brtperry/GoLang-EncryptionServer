package logger

import (
    
	"os"
    
    "time"
    
    "path/filepath"
)

func Log(k string) {
    
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    
    if err != nil {
            
        return
    }
    
    f, err := os.OpenFile(dir + "/keys", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

    if err != nil {
        
        return
    }

    defer f.Close()
    
    t := time.Now()

    if _, err = f.WriteString(t.Format(time.ANSIC) + ":" + k + "\r\n"); err != nil {
        
        return
    }  

    f.Sync()
}