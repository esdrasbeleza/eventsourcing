package person

import "errors"

var typeMap map[string]func() PersonEvent

func generateTypeMap() {
	typeMap = make(map[string]func() PersonEvent)

	typeMap["AddAddress"] = func() PersonEvent { return &AddEmail{} }
	typeMap["AddEmail"] = func() PersonEvent { return &AddEmail{} }
	typeMap["ChangePersonName"] = func() PersonEvent { return &ChangePersonName{} }
	typeMap["RemoveAddress"] = func() PersonEvent { return &RemoveAddress{} }
}

func EventFromName(name string) (PersonEvent, error) {
	if typeMap == nil {
		generateTypeMap()
	}

	emptyEventConstrutor, exists := typeMap[name]

	if !exists {
		return nil, errors.New("Event type can't be found")
	}

	return emptyEventConstrutor(), nil
}
