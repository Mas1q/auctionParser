package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func print(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("START!")
	realm := RealmInfo{}
	parcels, err := realm.FetchRealmInfo("19555")
	if err != nil {
		print(err)
	}
	for parcel := range parcels.Parcels {
		//parcelID := parcels.Parcels[parcel].Id
		fmt.Println(parcels.Parcels[parcel].Id)
		fmt.Println(parcels.Parcels[parcel].LastChanneledAlchemica)
		print(len(parcels.Parcels[parcel].EquippedInstallations))
		for installation := range parcels.Parcels[parcel].EquippedInstallations {
			fmt.Print(string(parcels.Parcels[parcel].EquippedInstallations[installation].Id) + ", ")
		}
	}
}
