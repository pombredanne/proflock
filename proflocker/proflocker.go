package proflocker

import (
        "../iwscanner"
        "os"
        "encoding/json"
)

type APscore struct {
        essid string
        name string
        score float64
}

type Location struct {
        name string
        path string
        aps iwscanner.APs
}

type Locations []Location

func RecordLocation(location string, profile_dir string, device string) (error) {
        os.MkdirAll(profile_dir + "/" + location, 755)

        aps, err := iwscanner.GetAPs(device)
        if err != nil {
                return err
        }
        out, err := json.Marshal(aps)
        if err != nil {
                return err
        }

        f, err := os.OpenFile(profile_dir + "/" + location + "/data", os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
                return err
        }

        defer f.Close()

        if _, err = f.WriteString(string(out)); err != nil {
                return err
        }

        return nil
}

func ParseLocation(path string) (Location, error) {
        location := Location{path: path}
        return location, nil
}

func ParseLocationsDir(dir string) (Locations, error) {
        locations := Locations{}
        return locations, nil
}